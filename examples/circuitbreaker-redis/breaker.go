// Package redisbreaker implements a distributed circuit breaker backed by
// Redis. State (Closed / Open / HalfOpen), counts, and timing all live in a
// single Redis Hash; transitions happen atomically via Lua scripts.
//
// Trade-offs vs. the in-process fortify/circuitbreaker:
//
//   - Coordination cost: every Execute incurs at least one Redis round-trip
//     for the admission check. Fortify's in-process CB is sub-microsecond;
//     this implementation is bound by your Redis network latency.
//   - Single point of failure: if Redis is unreachable, the breaker fails
//     open by default (the operation runs unguarded). Configurable via
//     FailMode.
//   - Predicate restriction: the in-process CB lets you supply an arbitrary
//     ReadyToTrip closure. This implementation only supports a
//     consecutive-failure threshold, since the predicate must execute in Lua
//     server-side.
//   - State change callbacks are not delivered cross-process. The
//     OnStateChange callback fires only on the instance that observes a
//     transition; sibling instances learn about it on their next admission
//     check.
//
// Use this when multiple replicas of your service must agree on whether a
// shared downstream is healthy. For per-replica protection, use the
// in-process fortify/circuitbreaker — it is faster and simpler.
package redisbreaker

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/redis/go-redis/v9"
)

// FailMode controls behavior when Redis is unreachable.
type FailMode int

const (
	// FailOpen runs the operation unguarded if Redis errors. Favors
	// availability of the downstream over CB enforcement.
	FailOpen FailMode = iota
	// FailClosed rejects with ferrors.ErrCircuitOpen if Redis errors.
	// Favors enforcement over availability.
	FailClosed
)

// Config holds the breaker configuration.
type Config struct {
	// Client is the Redis client. Required.
	Client redis.UniversalClient

	// Key is the Redis key for this breaker's state Hash. Required.
	// Use distinct keys for distinct downstreams.
	Key string

	// FailureThreshold trips the breaker Open after this many consecutive
	// failures. Defaults to 5.
	FailureThreshold int

	// OpenTimeout is how long the breaker stays Open before HalfOpen
	// trials. Defaults to 30s.
	OpenTimeout time.Duration

	// MaxRequests is the number of HalfOpen trials admitted before deciding
	// the verdict. Defaults to 1.
	MaxRequests int

	// FailMode controls behavior on Redis errors. Defaults to FailOpen.
	FailMode FailMode

	// OnStateChange (optional) fires on state transitions observed by THIS
	// instance. Sibling instances do not deliver this callback.
	OnStateChange func(from, to circuitbreaker.State)
}

func (c *Config) setDefaults() {
	if c.FailureThreshold <= 0 {
		c.FailureThreshold = 5
	}
	if c.OpenTimeout <= 0 {
		c.OpenTimeout = 30 * time.Second
	}
	if c.MaxRequests <= 0 {
		c.MaxRequests = 1
	}
}

// Breaker is a distributed circuit breaker. Implements
// circuitbreaker.CircuitBreaker[T].
type Breaker[T any] struct {
	cfg       Config
	mu        sync.Mutex
	lastState circuitbreaker.State
}

// New constructs a Breaker from the given configuration.
func New[T any](cfg Config) (*Breaker[T], error) {
	if cfg.Client == nil {
		return nil, errors.New("redisbreaker: Client is required")
	}
	if cfg.Key == "" {
		return nil, errors.New("redisbreaker: Key is required")
	}
	cfg.setDefaults()
	return &Breaker[T]{cfg: cfg, lastState: circuitbreaker.StateClosed}, nil
}

// admitScript handles the admission check + state transition.
//
// KEYS[1] = state hash key
// ARGV[1] = failure threshold
// ARGV[2] = open timeout (ms)
// ARGV[3] = max half-open trials
// ARGV[4] = now (unix ms)
//
// Returns: { state ("closed"/"open"/"half-open"), generation,
//   allowed (0|1), retry_after_ms }
const admitScriptSrc = `
local key       = KEYS[1]
local threshold = tonumber(ARGV[1])
local openMs    = tonumber(ARGV[2])
local maxTrials = tonumber(ARGV[3])
local now       = tonumber(ARGV[4])

local h = redis.call('HMGET', key, 'state', 'generation', 'expiry', 'consecutive_failures', 'half_open_admitted')
local state           = h[1] or 'closed'
local generation      = tonumber(h[2] or '0')
local expiry          = tonumber(h[3] or '0')
local consecFails     = tonumber(h[4] or '0')
local halfOpenAdmits  = tonumber(h[5] or '0')

local function setHash(s, exp)
  redis.call('HMSET', key,
    'state', s,
    'generation', generation,
    'expiry', exp,
    'consecutive_failures', consecFails,
    'half_open_admitted', halfOpenAdmits)
end

-- Auto-transition Open -> HalfOpen on timeout.
if state == 'open' and now >= expiry then
  state = 'half-open'
  generation = generation + 1
  halfOpenAdmits = 0
  setHash(state, 0)
end

local allowed = 0
local retryAfter = 0

if state == 'closed' then
  allowed = 1
elseif state == 'open' then
  retryAfter = expiry - now
  if retryAfter < 0 then retryAfter = 0 end
elseif state == 'half-open' then
  if halfOpenAdmits < maxTrials then
    halfOpenAdmits = halfOpenAdmits + 1
    setHash(state, expiry)
    allowed = 1
  end
end

return { state, tostring(generation), allowed, tostring(retryAfter) }
`

// recordScript records the outcome of a previously-admitted call.
//
// KEYS[1] = state hash key
// ARGV[1] = generation observed at admission
// ARGV[2] = success (1) or failure (0)
// ARGV[3] = failure threshold
// ARGV[4] = open timeout (ms)
// ARGV[5] = now (unix ms)
//
// Returns: { state, generation } after applying the result.
const recordScriptSrc = `
local key       = KEYS[1]
local admitGen  = tonumber(ARGV[1])
local success   = tonumber(ARGV[2])
local threshold = tonumber(ARGV[3])
local openMs    = tonumber(ARGV[4])
local now       = tonumber(ARGV[5])

local h = redis.call('HMGET', key, 'state', 'generation', 'expiry', 'consecutive_failures', 'half_open_admitted')
local state          = h[1] or 'closed'
local generation     = tonumber(h[2] or '0')
local expiry         = tonumber(h[3] or '0')
local consecFails    = tonumber(h[4] or '0')
local halfOpenAdmits = tonumber(h[5] or '0')

-- Discard stale results from a previous generation (state has rotated under us).
if admitGen ~= generation then
  return { state, tostring(generation) }
end

if state == 'closed' then
  if success == 1 then
    consecFails = 0
  else
    consecFails = consecFails + 1
    if consecFails >= threshold then
      state = 'open'
      generation = generation + 1
      expiry = now + openMs
      consecFails = 0
    end
  end
elseif state == 'half-open' then
  if success == 1 then
    halfOpenAdmits = halfOpenAdmits - 1
    -- All half-open trials succeeded: close the breaker.
    if halfOpenAdmits <= 0 then
      state = 'closed'
      generation = generation + 1
      consecFails = 0
      halfOpenAdmits = 0
    end
  else
    state = 'open'
    generation = generation + 1
    expiry = now + openMs
    consecFails = 0
    halfOpenAdmits = 0
  end
end

redis.call('HMSET', key,
  'state', state,
  'generation', generation,
  'expiry', expiry,
  'consecutive_failures', consecFails,
  'half_open_admitted', halfOpenAdmits)

return { state, tostring(generation) }
`

var (
	admitScript  = redis.NewScript(admitScriptSrc)
	recordScript = redis.NewScript(recordScriptSrc)
)

// Execute implements circuitbreaker.CircuitBreaker[T].
func (b *Breaker[T]) Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
	var zero T

	if err := ctx.Err(); err != nil {
		return zero, err
	}

	state, generation, allowed, retryAfter, err := b.admit(ctx)
	if err != nil {
		if b.cfg.FailMode == FailClosed {
			return zero, &ferrors.CircuitOpenError{State: "unknown"}
		}
		// FailOpen: run unguarded.
		return fn(ctx)
	}

	b.fireStateChangeIfChanged(state)

	if !allowed {
		return zero, &ferrors.CircuitOpenError{
			State:      state.String(),
			RetryAfter: retryAfter,
		}
	}

	result, runErr := fn(ctx)
	success := runErr == nil

	if newState, _, recErr := b.record(ctx, generation, success); recErr == nil {
		b.fireStateChangeIfChanged(newState)
	}
	return result, runErr
}

// State implements circuitbreaker.CircuitBreaker[T]. Performs a Redis read
// to get the latest cross-instance state.
func (b *Breaker[T]) State() circuitbreaker.State {
	state, _, _, _, err := b.admit(context.Background())
	if err != nil {
		b.mu.Lock()
		defer b.mu.Unlock()
		return b.lastState
	}
	return state
}

// Reset implements circuitbreaker.CircuitBreaker[T]. Clears all breaker
// state in Redis.
func (b *Breaker[T]) Reset() {
	_ = b.cfg.Client.Del(context.Background(), b.cfg.Key).Err()
	b.mu.Lock()
	defer b.mu.Unlock()
	b.lastState = circuitbreaker.StateClosed
}

// Close implements circuitbreaker.CircuitBreaker[T]. No-op; the caller owns
// the Redis client lifecycle.
func (b *Breaker[T]) Close() error { return nil }

// admit runs the admission Lua and parses results.
func (b *Breaker[T]) admit(ctx context.Context) (circuitbreaker.State, int64, bool, time.Duration, error) {
	now := time.Now().UnixMilli()
	res, err := admitScript.Run(ctx, b.cfg.Client,
		[]string{b.cfg.Key},
		b.cfg.FailureThreshold,
		int64(b.cfg.OpenTimeout/time.Millisecond),
		b.cfg.MaxRequests,
		now,
	).Result()
	if err != nil {
		return circuitbreaker.StateClosed, 0, false, 0, fmt.Errorf("redisbreaker: admit: %w", err)
	}
	arr, ok := res.([]interface{})
	if !ok || len(arr) != 4 {
		return circuitbreaker.StateClosed, 0, false, 0, errors.New("redisbreaker: unexpected admit shape")
	}
	state := parseState(toString(arr[0]))
	generation, _ := strconv.ParseInt(toString(arr[1]), 10, 64)
	allowedI, _ := strconv.ParseInt(toString(arr[2]), 10, 64)
	retryMs, _ := strconv.ParseInt(toString(arr[3]), 10, 64)
	return state, generation, allowedI == 1, time.Duration(retryMs) * time.Millisecond, nil
}

// record runs the recording Lua and parses results.
func (b *Breaker[T]) record(ctx context.Context, generation int64, success bool) (circuitbreaker.State, int64, error) {
	successFlag := 0
	if success {
		successFlag = 1
	}
	now := time.Now().UnixMilli()
	res, err := recordScript.Run(ctx, b.cfg.Client,
		[]string{b.cfg.Key},
		generation,
		successFlag,
		b.cfg.FailureThreshold,
		int64(b.cfg.OpenTimeout/time.Millisecond),
		now,
	).Result()
	if err != nil {
		return circuitbreaker.StateClosed, 0, fmt.Errorf("redisbreaker: record: %w", err)
	}
	arr, ok := res.([]interface{})
	if !ok || len(arr) != 2 {
		return circuitbreaker.StateClosed, 0, errors.New("redisbreaker: unexpected record shape")
	}
	state := parseState(toString(arr[0]))
	gen, _ := strconv.ParseInt(toString(arr[1]), 10, 64)
	return state, gen, nil
}

// fireStateChangeIfChanged invokes the user callback on local-observed transitions.
func (b *Breaker[T]) fireStateChangeIfChanged(newState circuitbreaker.State) {
	b.mu.Lock()
	prev := b.lastState
	b.lastState = newState
	b.mu.Unlock()
	if prev == newState {
		return
	}
	if b.cfg.OnStateChange != nil {
		// Run synchronously; user can wrap with goroutine if desired.
		defer func() { _ = recover() }()
		b.cfg.OnStateChange(prev, newState)
	}
}

func parseState(s string) circuitbreaker.State {
	switch s {
	case "open":
		return circuitbreaker.StateOpen
	case "half-open":
		return circuitbreaker.StateHalfOpen
	default:
		return circuitbreaker.StateClosed
	}
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

// Compile-time interface check.
var _ circuitbreaker.CircuitBreaker[int] = (*Breaker[int])(nil)
