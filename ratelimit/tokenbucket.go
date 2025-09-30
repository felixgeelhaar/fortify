package ratelimit

import (
	"sync"
	"time"
)

// tokenBucket implements the token bucket algorithm for rate limiting.
// It maintains a bucket of tokens that refills at a constant rate.
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
func (tb *tokenBucket) waitTime() time.Duration {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1.0 {
		return 0
	}

	// Calculate how many tokens we need
	tokensNeeded := 1.0 - tb.tokens

	// Calculate time to generate those tokens
	// tokens per nanosecond = rate / interval
	tokensPerNs := tb.rate / float64(tb.interval.Nanoseconds())
	nsToWait := tokensNeeded / tokensPerNs

	return time.Duration(nsToWait)
}

// refill adds tokens to the bucket based on time elapsed since last refill.
// Must be called with lock held.
func (tb *tokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	if elapsed <= 0 {
		return
	}

	// Calculate tokens to add: (elapsed / interval) * rate
	tokensToAdd := (float64(elapsed.Nanoseconds()) / float64(tb.interval.Nanoseconds())) * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > tb.burst {
		tb.tokens = tb.burst
	}

	tb.lastRefill = now
}
