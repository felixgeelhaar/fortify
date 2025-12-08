package ratelimit

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// maxElapsed is the maximum elapsed time considered for refill calculations.
	// This prevents overflow from clock skew or system sleep.
	maxElapsed = time.Hour

	// defaultMinWaitDuration is the minimum wait duration between retry attempts in Wait().
	defaultMinWaitDuration = 10 * time.Millisecond

	// defaultMaxWaitDuration is the maximum wait duration between retry attempts in Wait().
	defaultMaxWaitDuration = 100 * time.Millisecond

	// floatEpsilon is the threshold for float comparisons in token calculations.
	// This value (0.01% of a token) prevents adding negligible fractional tokens
	// that could accumulate due to float64 precision limits, while still allowing
	// accurate token accounting for practical rate limiting scenarios.
	floatEpsilon = 0.0001

	// maxWaitTimeLimit is the maximum total time Wait() can be blocked if rate is zero.
	maxWaitTimeLimit = 24 * time.Hour

	// defaultMaxWaitIterations is the maximum number of retry iterations in Wait().
	// This prevents infinite loops in edge cases.
	defaultMaxWaitIterations = 10000

	// defaultMaxTotalWaitTime is the maximum total time Wait() will block.
	// After this duration, Wait() returns ErrWaitTimeout.
	defaultMaxTotalWaitTime = 5 * time.Minute
)

// RateLimiter controls the rate of operations per key using the token bucket algorithm.
// It maintains separate token buckets for each key, allowing independent rate limiting.
type RateLimiter interface {
	// Allow checks if a request should be allowed based on rate limits.
	// Returns true if the request is allowed, false if it should be rejected.
	// Does not block - returns immediately.
	Allow(ctx context.Context, key string) bool

	// Wait blocks until a token is available or the context is cancelled.
	// Returns an error if the context is cancelled before a token becomes available.
	Wait(ctx context.Context, key string) error

	// Take attempts to take n tokens from the bucket.
	// Returns true if n tokens were available, false otherwise.
	// This is useful for operations that consume multiple tokens.
	//
	// Returns false if tokens exceeds MaxTokensPerRequest or is <= 0.
	Take(ctx context.Context, key string, tokens int) bool

	// Execute runs the provided operation if rate limiting allows.
	// It first checks if a token is available for the given key, and if so,
	// executes the operation and returns its result.
	//
	// Returns ErrLimitExceeded if the rate limit is exceeded.
	// Returns ErrRateLimiterClosed if the rate limiter has been closed.
	// Returns any error from the operation itself if it fails.
	//
	// Panic Handling: If operation panics, the panic propagates to the caller.
	// The consumed token is not refunded. Callers should implement panic
	// recovery if needed.
	//
	// This is a convenience method that combines Allow() with operation execution,
	// ensuring the operation only runs when rate limiting permits.
	Execute(ctx context.Context, key string, operation func() error) error

	// ExecuteN runs the provided operation if n tokens are available.
	// Similar to Execute but consumes multiple tokens per operation.
	//
	// Returns ErrLimitExceeded if the rate limit is exceeded.
	// Returns ErrRateLimiterClosed if the rate limiter has been closed.
	// Returns ErrInvalidTokenCount if tokens is <= 0.
	// Returns ErrExcessiveTokens if tokens exceeds MaxTokensPerRequest.
	// Returns any error from the operation itself if it fails.
	//
	// Panic Handling: If operation panics, the panic propagates to the caller.
	// The consumed tokens are not refunded. Callers should implement panic
	// recovery if needed.
	ExecuteN(ctx context.Context, key string, tokens int, operation func() error) error

	// HealthCheck verifies the rate limiter and its underlying store are operational.
	// Returns nil if healthy, or an error describing the issue.
	// This is useful for health checks in distributed systems.
	HealthCheck(ctx context.Context) error

	// Reset clears all rate limiting state.
	// This removes all buckets from the underlying store.
	// Useful for testing or administrative purposes.
	//
	// SECURITY WARNING: Reset() has no built-in authorization. Applications MUST
	// implement authorization checks before calling Reset() to prevent unauthorized
	// rate limiting bypass. Only call Reset() from administrative interfaces with
	// proper access controls, or in testing environments.
	//
	// Note: Reset is not atomic with respect to concurrent operations.
	// For deterministic clearing, ensure no concurrent operations are in progress.
	Reset(ctx context.Context) error

	// BucketCount returns the current number of active buckets in the store.
	// Returns -1 if:
	//   - The underlying store does not support bucket counting
	//   - The rate limiter has been closed
	// This is useful for monitoring and debugging.
	BucketCount() int64

	// Close releases any resources held by the rate limiter.
	// This includes closing the underlying Store.
	// After Close is called, the rate limiter should not be used.
	Close() error
}

// rateLimiter is the concrete implementation of RateLimiter.
type rateLimiter struct {
	config              Config
	closed              atomic.Bool
	healthCheckWarnOnce sync.Once

	// Cached values for hot path optimization (calculated once at construction)
	intervalNs int64   // config.Interval.Nanoseconds()
	burstFloat float64 // float64(config.Burst)
	rateFloat  float64 // float64(config.Rate)
}

// New creates a new RateLimiter with the given configuration.
func New(config *Config) RateLimiter {
	if config == nil {
		config = &Config{}
	}
	config.setDefaults()

	return &rateLimiter{
		config: *config,
		// Pre-calculate hot path values to avoid repeated conversions
		intervalNs: config.Interval.Nanoseconds(),
		burstFloat: float64(config.Burst),
		rateFloat:  float64(config.Rate),
	}
}

// Allow implements the RateLimiter interface.
func (rl *rateLimiter) Allow(ctx context.Context, key string) bool {
	if rl.closed.Load() {
		return false
	}
	resolvedKey := rl.resolveKey(ctx, key)

	// Only measure time when metrics are enabled to avoid overhead
	var start time.Time
	if rl.config.Metrics != nil {
		start = time.Now()
	}

	allowed, err := rl.tryConsume(ctx, resolvedKey, 1)

	// Record latency metrics (only if metrics enabled)
	if rl.config.Metrics != nil {
		rl.config.Metrics.OnStoreLatency(ctx, "allow", time.Since(start))
	}

	if err != nil {
		failOpen := rl.handleError(ctx, resolvedKey, err)
		if failOpen && rl.config.Metrics != nil {
			rl.config.Metrics.OnAllow(ctx, resolvedKey)
		}
		return failOpen
	}

	if allowed {
		if rl.config.Metrics != nil {
			rl.config.Metrics.OnAllow(ctx, resolvedKey)
		}
	} else {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.Metrics != nil {
			rl.config.Metrics.OnDeny(ctx, resolvedKey)
		}
		if rl.config.OnLimit != nil {
			rl.safeCallback(func() {
				rl.config.OnLimit(ctx, resolvedKey)
			})
		}
	}

	return allowed
}

// Wait implements the RateLimiter interface.
//
//nolint:gocyclo // complexity is inherent in wait logic with context/timing
func (rl *rateLimiter) Wait(ctx context.Context, key string) error {
	if rl.closed.Load() {
		return ErrRateLimiterClosed
	}
	resolvedKey := rl.resolveKey(ctx, key)
	start := time.Now()
	iterations := 0

	// Create timer once and reuse with Reset() to reduce allocations.
	// Start with a stopped timer - we'll reset it when needed.
	timer := time.NewTimer(0)
	if !timer.Stop() {
		// Drain the channel if timer already fired.
		// No default case needed: if Stop() returned false, the timer fired
		// and the channel is guaranteed to have a value.
		<-timer.C
	}
	defer timer.Stop()

	for {
		iterations++

		// Check iteration limit to prevent infinite loops
		if iterations > defaultMaxWaitIterations {
			return ErrWaitTimeout
		}

		// Check total wait time limit
		if time.Since(start) > defaultMaxTotalWaitTime {
			return ErrWaitTimeout
		}
		// Check context first
		if err := ctx.Err(); err != nil {
			return err
		}

		// Try to take a token
		allowed, err := rl.tryConsume(ctx, resolvedKey, 1)
		if err != nil {
			if rl.handleError(ctx, resolvedKey, err) {
				// FailOpen: allow the request
				if rl.config.Metrics != nil {
					rl.config.Metrics.OnStoreLatency(ctx, "wait", time.Since(start))
					rl.config.Metrics.OnAllow(ctx, resolvedKey)
				}
				return nil
			}
			// FailClosed: continue retrying with backoff
		}

		if allowed {
			if rl.config.Metrics != nil {
				rl.config.Metrics.OnStoreLatency(ctx, "wait", time.Since(start))
				rl.config.Metrics.OnAllow(ctx, resolvedKey)
			}
			return nil
		}

		// Calculate wait time based on current state (read-only)
		waitDuration := rl.calculateWaitTime(ctx, resolvedKey)
		if waitDuration <= 0 {
			waitDuration = defaultMinWaitDuration
		}

		// Cap wait time to prevent long blocking
		if waitDuration > defaultMaxWaitDuration {
			waitDuration = defaultMaxWaitDuration
		}

		// Reset and reuse timer (reduces allocations in loop)
		timer.Reset(waitDuration)
		select {
		case <-timer.C:
			// Timer fired normally, continue to next iteration
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Take implements the RateLimiter interface.
func (rl *rateLimiter) Take(ctx context.Context, key string, tokens int) bool {
	if rl.closed.Load() {
		return false
	}
	// Validate token count
	if tokens <= 0 {
		return false
	}

	// Check upper bound to prevent DoS and overflow
	if tokens > rl.config.MaxTokensPerRequest {
		if rl.config.Logger != nil {
			rl.config.Logger.WarnContext(ctx, "excessive token request rejected",
				slog.String("key", key),
				slog.Int("requested", tokens),
				slog.Int("max_allowed", rl.config.MaxTokensPerRequest),
			)
		}
		return false
	}

	resolvedKey := rl.resolveKey(ctx, key)

	// Only measure time when metrics are enabled to avoid overhead
	var start time.Time
	if rl.config.Metrics != nil {
		start = time.Now()
	}

	taken, err := rl.tryConsume(ctx, resolvedKey, tokens)

	// Record latency metrics (only if metrics enabled)
	if rl.config.Metrics != nil {
		rl.config.Metrics.OnStoreLatency(ctx, "take", time.Since(start))
	}

	if err != nil {
		failOpen := rl.handleError(ctx, resolvedKey, err)
		if failOpen && rl.config.Metrics != nil {
			rl.config.Metrics.OnAllow(ctx, resolvedKey)
		}
		return failOpen
	}

	if taken {
		if rl.config.Metrics != nil {
			rl.config.Metrics.OnAllow(ctx, resolvedKey)
		}
	} else {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.Metrics != nil {
			rl.config.Metrics.OnDeny(ctx, resolvedKey)
		}
		if rl.config.OnLimit != nil {
			rl.safeCallback(func() {
				rl.config.OnLimit(ctx, resolvedKey)
			})
		}
	}

	return taken
}

// Execute implements the RateLimiter interface.
func (rl *rateLimiter) Execute(ctx context.Context, key string, operation func() error) error {
	if rl.closed.Load() {
		return ErrRateLimiterClosed
	}

	if !rl.Allow(ctx, key) {
		return ErrLimitExceeded
	}

	return operation()
}

// ExecuteN implements the RateLimiter interface.
func (rl *rateLimiter) ExecuteN(ctx context.Context, key string, tokens int, operation func() error) error {
	if rl.closed.Load() {
		return ErrRateLimiterClosed
	}

	// Validate token count
	if tokens <= 0 {
		return ErrInvalidTokenCount
	}

	// Check upper bound
	if tokens > rl.config.MaxTokensPerRequest {
		return ErrExcessiveTokens
	}

	if !rl.Take(ctx, key, tokens) {
		return ErrLimitExceeded
	}

	return operation()
}

// Reset implements the RateLimiter interface.
func (rl *rateLimiter) Reset(ctx context.Context) error {
	if rl.closed.Load() {
		return ErrRateLimiterClosed
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	// Check if the store implements Resetter
	if r, ok := rl.config.Store.(Resetter); ok {
		return r.Reset(ctx)
	}

	// Warn if store doesn't implement Resetter
	if rl.config.Logger != nil {
		rl.config.Logger.WarnContext(ctx,
			"Store does not implement Resetter - reset not supported",
			slog.String("store_type", fmt.Sprintf("%T", rl.config.Store)),
		)
	}

	return nil
}

// HealthCheck implements the RateLimiter interface.
func (rl *rateLimiter) HealthCheck(ctx context.Context) error {
	if rl.closed.Load() {
		return ErrRateLimiterClosed
	}
	// Check if the store implements HealthChecker
	if hc, ok := rl.config.Store.(HealthChecker); ok {
		return hc.HealthCheck(ctx)
	}
	// Warn once if store doesn't implement HealthChecker
	// This helps catch missing health check implementations in custom stores
	rl.healthCheckWarnOnce.Do(func() {
		if rl.config.Logger != nil {
			rl.config.Logger.WarnContext(ctx,
				"Store does not implement HealthChecker - health checks always pass",
				slog.String("store_type", fmt.Sprintf("%T", rl.config.Store)),
			)
		}
	})
	return nil
}

// BucketCount implements the RateLimiter interface.
func (rl *rateLimiter) BucketCount() int64 {
	if rl.closed.Load() {
		return -1
	}

	// Check if the store implements BucketCounter
	if bc, ok := rl.config.Store.(BucketCounter); ok {
		return bc.BucketCount()
	}

	// Store doesn't support bucket counting
	return -1
}

// Close implements the RateLimiter interface.
func (rl *rateLimiter) Close() error {
	if rl.closed.Swap(true) {
		return nil // Already closed
	}
	return rl.config.Store.Close()
}

// tryConsume attempts to consume tokens from the bucket using the Store.
// Returns (allowed, error) where allowed indicates if tokens were consumed.
func (rl *rateLimiter) tryConsume(ctx context.Context, key string, tokensNeeded int) (bool, error) {
	var allowed bool

	_, err := rl.config.Store.AtomicUpdate(ctx, key, func(state *BucketState) *BucketState {
		now := time.Now()

		// Initialize new bucket if needed
		if state == nil {
			state = &BucketState{
				Tokens:     float64(rl.config.Burst),
				LastRefill: now,
			}
		}

		// Refill tokens based on elapsed time
		state = rl.refill(state, now)

		// Check if we have enough tokens
		tokensRequired := float64(tokensNeeded)
		if state.Tokens >= tokensRequired {
			// Consume tokens - create new state only when changing
			allowed = true
			return &BucketState{
				Tokens:     state.Tokens - tokensRequired,
				LastRefill: state.LastRefill,
			}
		}

		// Not enough tokens - return existing state unchanged
		allowed = false
		return state
	})

	return allowed, err
}

// refill calculates the new token count based on elapsed time.
// Returns the same state pointer if no refill is needed (optimization),
// or a new BucketState if tokens were added.
func (rl *rateLimiter) refill(state *BucketState, now time.Time) *BucketState {
	elapsed := now.Sub(state.LastRefill)

	// Don't refill if no time has passed or clock went backwards
	if elapsed <= 0 {
		return state
	}

	// Cap elapsed time to prevent overflow from clock skew or system sleep
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	// Use cached intervalNs (calculated once at construction)
	intervalNs := rl.intervalNs
	if intervalNs <= 0 {
		intervalNs = time.Second.Nanoseconds() // Safety fallback
	}

	// Calculate tokens to add using cached rateFloat
	tokensToAdd := (float64(elapsed.Nanoseconds()) / float64(intervalNs)) * rl.rateFloat

	// If no tokens to add, return original state (avoid allocation)
	if tokensToAdd < floatEpsilon {
		return state
	}

	newTokens := state.Tokens + tokensToAdd

	// Use cached burstFloat for comparisons
	burstFloat := rl.burstFloat

	// Cap at burst limit
	if newTokens > burstFloat {
		newTokens = burstFloat
	}

	// If already at burst limit, just update LastRefill timestamp
	if state.Tokens >= burstFloat && newTokens >= burstFloat {
		return &BucketState{
			Tokens:     burstFloat,
			LastRefill: now,
		}
	}

	return &BucketState{
		Tokens:     newTokens,
		LastRefill: now,
	}
}

// calculateWaitTime estimates how long to wait for a token to become available.
// This is a read-only operation that does not modify state.
func (rl *rateLimiter) calculateWaitTime(ctx context.Context, key string) time.Duration {
	// Use Get() for read-only access instead of AtomicUpdate
	state, err := rl.config.Store.Get(ctx, key)
	if err != nil || state == nil {
		return 0
	}

	now := time.Now()

	// Calculate current tokens after refill (without storing)
	elapsed := now.Sub(state.LastRefill)
	if elapsed <= 0 {
		elapsed = 0
	}
	if elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	intervalNs := rl.config.Interval.Nanoseconds()
	if intervalNs <= 0 {
		intervalNs = time.Second.Nanoseconds()
	}

	tokensToAdd := (float64(elapsed.Nanoseconds()) / float64(intervalNs)) * float64(rl.config.Rate)
	currentTokens := state.Tokens + tokensToAdd

	burstFloat := float64(rl.config.Burst)
	if currentTokens > burstFloat {
		currentTokens = burstFloat
	}

	// If tokens available, no wait needed
	if currentTokens >= 1.0 {
		return 0
	}

	// Calculate how many tokens we need
	tokensNeeded := 1.0 - currentTokens
	if tokensNeeded <= 0 {
		return 0
	}

	// Safety check: if rate is 0 or negative, tokens will never be added
	if rl.config.Rate <= 0 {
		return maxWaitTimeLimit
	}

	// Calculate time to generate those tokens
	tokensPerNs := float64(rl.config.Rate) / float64(intervalNs)
	if tokensPerNs <= 0 {
		return maxWaitTimeLimit
	}

	nsToWait := tokensNeeded / tokensPerNs
	if nsToWait < 0 {
		return 0
	}

	wait := time.Duration(nsToWait)

	// Cap maximum wait time
	if wait > maxWaitTimeLimit {
		wait = maxWaitTimeLimit
	}

	return wait
}

// handleError handles storage errors based on FailOpen configuration.
// Returns true if the request should be allowed (FailOpen), false otherwise.
func (rl *rateLimiter) handleError(ctx context.Context, key string, err error) bool {
	if rl.config.Logger != nil {
		rl.config.Logger.ErrorContext(ctx, "rate limiter storage error",
			slog.String("key", key),
			slog.String("error", err.Error()),
			slog.Bool("fail_open", rl.config.FailOpen),
		)
	}

	if rl.config.Metrics != nil {
		rl.config.Metrics.OnError(ctx, key, err)
	}

	return rl.config.FailOpen
}

// resolveKey determines the rate limiting key using KeyFunc if configured.
func (rl *rateLimiter) resolveKey(ctx context.Context, key string) string {
	if rl.config.KeyFunc != nil {
		return rl.config.KeyFunc(ctx)
	}
	return key
}

// logRateLimit logs rate limiting events using structured logging.
func (rl *rateLimiter) logRateLimit(ctx context.Context, key string) {
	if rl.config.Logger != nil {
		rl.config.Logger.WarnContext(ctx, "rate limit exceeded",
			slog.String("key", key),
			slog.Int("rate", rl.config.Rate),
			slog.Int("burst", rl.config.Burst),
		)
	}
}

// safeCallback executes a callback with panic recovery.
func (rl *rateLimiter) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if rl.config.Logger != nil {
				rl.config.Logger.Error("rate limiter callback panic",
					slog.String("pattern", "rate_limit"),
					slog.Any("panic", r),
					slog.String("stack", string(debug.Stack())),
				)
			}
		}
	}()
	fn()
}

// Compile-time interface check.
var _ RateLimiter = (*rateLimiter)(nil)
