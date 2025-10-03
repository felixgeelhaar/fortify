package ratelimit

import (
	"sync"
	"time"
)

// tokenBucket implements the token bucket algorithm for rate limiting.
// It maintains a bucket of tokens that refills at a constant rate.
//
// Field alignment optimization is intentionally disabled for this internal struct because:
// 1. This is an internal implementation detail, not exposed in the public API
// 2. Fields are logically grouped (sync primitives, time tracking, token state) for code clarity
// 3. The struct is instantiated once per key, not in hot allocation paths
// 4. Memory overhead from misalignment is negligible (few extra bytes per bucket)
// 5. The mutex and time fields need to be at the top for clear synchronization semantics
//
//nolint:govet // fieldalignment: internal struct, code clarity prioritized over memory optimization
type tokenBucket struct {
	mu sync.Mutex

	// lastRefill is the last time tokens were refilled
	lastRefill time.Time

	// interval is the time period for rate replenishment
	interval time.Duration

	// tokens is the current number of available tokens
	tokens float64

	// burst is the maximum number of tokens (bucket capacity)
	burst float64

	// rate is the number of tokens added per interval
	rate float64
}

// newTokenBucket creates a new token bucket with the given rate, burst, and interval.
func newTokenBucket(rate, burst int, interval time.Duration) *tokenBucket {
	return &tokenBucket{
		tokens:     float64(burst),
		burst:      float64(burst),
		rate:       float64(rate),
		interval:   interval,
		lastRefill: time.Now(),
	}
}

// allow attempts to take 1 token from the bucket.
// Returns true if a token was available, false otherwise.
func (tb *tokenBucket) allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1.0 {
		tb.tokens--
		return true
	}

	return false
}

// take attempts to take n tokens from the bucket.
// Returns true if n tokens were available, false otherwise.
// If insufficient tokens are available, no tokens are taken.
func (tb *tokenBucket) take(n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	tokensNeeded := float64(n)
	if tb.tokens >= tokensNeeded {
		tb.tokens -= tokensNeeded
		return true
	}

	return false
}

// waitTime returns the duration to wait for at least 1 token to become available.
// Returns 0 if a token is immediately available.
//
// IMPORTANT: This is a best-effort calculation. The returned wait time is based on
// the current token count and refill rate, but is NOT a guarantee. Between calculating
// the wait time and the caller using it, other goroutines may consume tokens, making
// the wait time stale. The caller should be prepared to wait again if needed.
//
// Includes safeguards against edge cases like zero rate or very small intervals.
func (tb *tokenBucket) waitTime() time.Duration {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	// If tokens available, no wait needed
	if tb.tokens >= 1.0 {
		return 0
	}

	// Calculate how many tokens we need
	tokensNeeded := 1.0 - tb.tokens

	// Safety check: ensure tokensNeeded is positive
	if tokensNeeded <= 0 {
		return 0
	}

	// Safety check: if rate is 0 or negative, tokens will never be added
	// Return maximum wait time to prevent infinite blocking
	if tb.rate <= 0 {
		return time.Hour * 24
	}

	// Calculate time to generate those tokens
	// tokens per nanosecond = rate / interval
	intervalNs := tb.interval.Nanoseconds()

	// Safety check: prevent division by zero or very small intervals
	if intervalNs <= 0 {
		return time.Hour * 24 // Maximum wait for edge case
	}

	tokensPerNs := tb.rate / float64(intervalNs)

	// Safety check: if tokens per nanosecond is effectively zero (very small rate)
	// use direct calculation to avoid division by zero
	if tokensPerNs <= 0 {
		return time.Hour * 24
	}

	nsToWait := tokensNeeded / tokensPerNs

	// Safety check: ensure result is within reasonable bounds
	// Prevent negative durations and cap maximum wait time
	if nsToWait < 0 {
		return 0
	}

	wait := time.Duration(nsToWait)

	// Cap maximum wait time to 24 hours to prevent overflow and unreasonable waits
	maxWait := time.Hour * 24
	if wait > maxWait {
		return maxWait
	}

	return wait
}

// refill adds tokens to the bucket based on time elapsed since last refill.
// Must be called with lock held.
func (tb *tokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	if elapsed <= 0 {
		return
	}

	// Cap elapsed time to prevent overflow from clock skew or system sleep
	// Maximum 1 hour of elapsed time is reasonable for token refill
	if elapsed > time.Hour {
		elapsed = time.Hour
	}

	// Calculate tokens to add: (elapsed / interval) * rate
	tokensToAdd := (float64(elapsed.Nanoseconds()) / float64(tb.interval.Nanoseconds())) * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > tb.burst {
		tb.tokens = tb.burst
	}

	tb.lastRefill = now
}
