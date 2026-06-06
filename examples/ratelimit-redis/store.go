// Package redisstore implements ratelimit.Store backed by Redis using a Lua
// script for atomic read-modify-write of token-bucket state.
//
// This is a reference implementation. It is intentionally a separate Go
// module so that fortify's core module does not depend on go-redis.
//
// Caveats and tuning:
//
//   - Time source: the Lua script uses Redis server time (TIME command), not
//     client time. This avoids clock skew between application instances at
//     the cost of one extra round-trip per update.
//   - TTL: each bucket key receives a TTL on every update so dormant keys
//     expire automatically. Default 1h; configurable via WithTTL.
//   - Failures: returns a wrapped ErrStorageUnavailable on Redis errors,
//     suitable for the rate limiter's FailOpen behavior.
//   - Key hygiene: caller-supplied keys are passed through verbatim. The
//     RateLimiter's resolveKey already sanitizes; the Store assumes the key
//     is safe.
package redisstore

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"go.klarlabs.de/fortify/ratelimit"
	"github.com/redis/go-redis/v9"
)

// Store implements ratelimit.Store backed by Redis.
type Store struct {
	client redis.UniversalClient
	prefix string
	ttl    time.Duration
}

// Option configures a Store.
type Option func(*Store)

// WithPrefix sets the key prefix applied to every bucket key. Defaults to "rl:".
func WithPrefix(prefix string) Option {
	return func(s *Store) {
		s.prefix = prefix
	}
}

// WithTTL sets the bucket-key TTL. Defaults to 1 hour.
func WithTTL(ttl time.Duration) Option {
	return func(s *Store) {
		if ttl > 0 {
			s.ttl = ttl
		}
	}
}

// New creates a Redis-backed Store.
func New(client redis.UniversalClient, opts ...Option) *Store {
	s := &Store{
		client: client,
		prefix: "rl:",
		ttl:    1 * time.Hour,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// atomicScript implements read-modify-write of token-bucket state.
//
// KEYS[1] = bucket key
// ARGV[1] = tokens to consume
// ARGV[2] = burst (max tokens)
// ARGV[3] = rate (tokens per intervalNs)
// ARGV[4] = intervalNs
// ARGV[5] = ttl seconds
//
// Returns: { tokens_remaining_x_1e6, last_refill_unix_nanos, allowed (0|1) }
//
// Tokens are stored as scaled integers (× 1e6) to avoid Redis Lua's lack of
// reliable float storage. We track up to ~9.2e12 effective tokens, more
// than enough for any reasonable burst value.
const atomicScriptSrc = `
local key      = KEYS[1]
local tokens   = tonumber(ARGV[1])
local burst    = tonumber(ARGV[2])
local rate     = tonumber(ARGV[3])
local interval = tonumber(ARGV[4])
local ttl      = tonumber(ARGV[5])

local now_arr   = redis.call('TIME')
local now_secs  = tonumber(now_arr[1])
local now_micro = tonumber(now_arr[2])
local now_ns    = now_secs * 1000000000 + now_micro * 1000

local existing = redis.call('HMGET', key, 'tokens', 'last')
local cur_tokens
local last_refill
if existing[1] == false then
  cur_tokens  = burst * 1000000
  last_refill = now_ns
else
  cur_tokens  = tonumber(existing[1])
  last_refill = tonumber(existing[2])
end

local elapsed = now_ns - last_refill
if elapsed < 0 then elapsed = 0 end
if elapsed > 3600000000000 then elapsed = 3600000000000 end

local refill = (elapsed / interval) * rate * 1000000
cur_tokens = cur_tokens + refill
local cap = burst * 1000000
if cur_tokens > cap then cur_tokens = cap end

local need = tokens * 1000000
local allowed = 0
if cur_tokens >= need then
  cur_tokens = cur_tokens - need
  allowed = 1
end

redis.call('HMSET', key, 'tokens', cur_tokens, 'last', now_ns)
redis.call('PEXPIRE', key, ttl * 1000)

return { tostring(cur_tokens), tostring(now_ns), allowed }
`

var atomicScript = redis.NewScript(atomicScriptSrc)

// AtomicUpdate implements ratelimit.Store.
//
// The Lua script performs the read-refill-consume cycle atomically server-side.
// The updateFn closure is NOT invoked: Redis-backed buckets cannot run
// arbitrary Go closures inside a transaction. Instead, we encode the
// token-bucket protocol directly in Lua. This means a Redis Store works only
// for the standard token-bucket semantics; if you need custom update logic
// (e.g., a different replenishment curve), implement your own Store.
func (s *Store) AtomicUpdate(ctx context.Context, key string, updateFn func(*ratelimit.BucketState) *ratelimit.BucketState) (*ratelimit.BucketState, error) {
	// We need the rate-limit parameters to run the Lua script. Since
	// AtomicUpdate doesn't receive them directly, we infer by running
	// updateFn against a probe state and inspecting how it transformed
	// the tokens. For a typical RateLimiter the closure encodes:
	//   * Burst = capacity
	//   * desired tokens to consume
	//
	// The simpler and more correct approach is to require the caller to
	// pass parameters via context or a wrapper Store. For this reference
	// implementation we expose a typed helper and document that AtomicUpdate
	// only supports the standard fortify rate limiter contract.
	probeState := updateFn(nil)
	if probeState == nil {
		return nil, errors.New("redisstore: updateFn returned nil for probe state; not a token-bucket update")
	}
	// Burst inferred from probe initial token count.
	burst := int64(probeState.Tokens)
	if burst <= 0 {
		burst = 1
	}
	// Tokens consumed inferred by running probe twice (no-elapsed-time mock).
	probeAfter := updateFn(probeState)
	consumed := int64(probeState.Tokens - probeAfter.Tokens)
	if consumed <= 0 {
		consumed = 1
	}

	// Refill rate and interval are not directly inferrable from updateFn;
	// for a reference implementation we treat the bucket as already-warm and
	// approximate refill at 1 token per "rate seconds" using context lookup.
	rate, interval := paramsFromContext(ctx)
	if rate <= 0 || interval <= 0 {
		rate = burst
		interval = int64(time.Second)
	}

	res, err := atomicScript.Run(ctx, s.client, []string{s.prefix + key},
		consumed, burst, rate, interval, int64(s.ttl/time.Second),
	).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ratelimit.ErrStorageUnavailable, err)
	}

	arr, ok := res.([]interface{})
	if !ok || len(arr) != 3 {
		return nil, fmt.Errorf("%w: unexpected response shape", ratelimit.ErrStorageUnavailable)
	}
	tokensScaled, _ := strconv.ParseFloat(toString(arr[0]), 64)
	lastNs, _ := strconv.ParseInt(toString(arr[1]), 10, 64)

	return &ratelimit.BucketState{
		Tokens:     tokensScaled / 1_000_000,
		LastRefill: time.Unix(0, lastNs),
	}, nil
}

// Get implements ratelimit.Store. Read-only access without consuming tokens.
func (s *Store) Get(ctx context.Context, key string) (*ratelimit.BucketState, error) {
	vals, err := s.client.HMGet(ctx, s.prefix+key, "tokens", "last").Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ratelimit.ErrStorageUnavailable, err)
	}
	if vals[0] == nil {
		return nil, nil //nolint:nilnil // signaling "not found" per Store contract
	}
	tokensScaled, _ := strconv.ParseFloat(toString(vals[0]), 64)
	lastNs, _ := strconv.ParseInt(toString(vals[1]), 10, 64)
	return &ratelimit.BucketState{
		Tokens:     tokensScaled / 1_000_000,
		LastRefill: time.Unix(0, lastNs),
	}, nil
}

// Delete implements ratelimit.Store.
func (s *Store) Delete(ctx context.Context, key string) error {
	if err := s.client.Del(ctx, s.prefix+key).Err(); err != nil {
		return fmt.Errorf("%w: %w", ratelimit.ErrStorageUnavailable, err)
	}
	return nil
}

// Close implements ratelimit.Store. Does not close the underlying client;
// the caller owns the client lifecycle.
func (s *Store) Close() error { return nil }

// HealthCheck pings Redis.
func (s *Store) HealthCheck(ctx context.Context) error {
	if err := s.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("%w: %w", ratelimit.ErrStorageUnavailable, err)
	}
	return nil
}

// paramsContextKey carries rate/interval parameters for AtomicUpdate.
type paramsContextKey struct{}

// WithBucketParams attaches the rate-limiter's Rate and Interval to the
// context so AtomicUpdate can construct the Lua args. The Fortify rate
// limiter does this transparently when using the Store; if you call
// AtomicUpdate directly, attach via:
//
//	ctx = redisstore.WithBucketParams(ctx, rate, interval)
func WithBucketParams(ctx context.Context, rate int, interval time.Duration) context.Context {
	return context.WithValue(ctx, paramsContextKey{}, struct {
		rate     int
		interval time.Duration
	}{rate: rate, interval: interval})
}

func paramsFromContext(ctx context.Context) (rate, intervalNs int64) {
	v, ok := ctx.Value(paramsContextKey{}).(struct {
		rate     int
		interval time.Duration
	})
	if !ok {
		return 0, 0
	}
	return int64(v.rate), v.interval.Nanoseconds()
}

func toString(v interface{}) string {
	switch x := v.(type) {
	case string:
		return x
	case []byte:
		return string(x)
	default:
		return fmt.Sprintf("%v", x)
	}
}
