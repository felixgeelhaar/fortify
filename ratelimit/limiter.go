// Package ratelimit provides token bucket based rate limiting for controlling
// request rates and preventing resource exhaustion.
//
// The rate limiter uses a token bucket algorithm where tokens are added at a
// constant rate up to a maximum burst capacity. Each request consumes one or more
// tokens. When the bucket is empty, requests are either rejected (Allow) or wait
// for tokens to become available (Wait).
//
// Example usage:
//
//	limiter := ratelimit.New(ratelimit.Config{
//	    Rate:     100,
//	    Burst:    150,
//	    Interval: time.Second,
//	    KeyFunc: func(ctx context.Context) string {
//	        return ctx.Value("user_id").(string)
//	    },
//	})
//
//	if limiter.Allow(ctx, "") {
//	    // Process request
//	} else {
//	    // Return 429 Too Many Requests
//	}
package ratelimit

import (
	"context"
	"log/slog"
	"sync"
	"time"
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
	Take(ctx context.Context, key string, tokens int) bool
}

// rateLimiter is the concrete implementation of RateLimiter.
type rateLimiter struct {
	buckets sync.Map // map[string]*tokenBucket
	config  Config
}

// New creates a new RateLimiter with the given configuration.
func New(config Config) RateLimiter {
	config.setDefaults()
	return &rateLimiter{
		config: config,
	}
}

// Allow implements the RateLimiter interface.
func (rl *rateLimiter) Allow(ctx context.Context, key string) bool {
	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)

	// Get or create bucket for this key
	bucket := rl.getBucket(resolvedKey)

	// Try to take a token
	allowed := bucket.allow()

	// Log and callback if rate limited
	if !allowed {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.OnLimit != nil {
			rl.config.OnLimit(resolvedKey)
		}
	}

	return allowed
}

// Wait implements the RateLimiter interface.
func (rl *rateLimiter) Wait(ctx context.Context, key string) error {
	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)

	// Get or create bucket for this key
	bucket := rl.getBucket(resolvedKey)

	for {
		// Check context first
		if err := ctx.Err(); err != nil {
			return err
		}

		// Try to take a token
		if bucket.allow() {
			return nil
		}

		// Calculate wait time
		waitDuration := bucket.waitTime()

		// Wait with context cancellation support
		select {
		case <-time.After(waitDuration):
			// Try again after wait
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Take implements the RateLimiter interface.
func (rl *rateLimiter) Take(ctx context.Context, key string, tokens int) bool {
	if tokens <= 0 {
		return false
	}

	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)

	// Get or create bucket for this key
	bucket := rl.getBucket(resolvedKey)

	// Try to take tokens
	taken := bucket.take(tokens)

	// Log and callback if rate limited
	if !taken {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.OnLimit != nil {
			rl.config.OnLimit(resolvedKey)
		}
	}

	return taken
}

// getBucket retrieves or creates a token bucket for the given key.
func (rl *rateLimiter) getBucket(key string) *tokenBucket {
	// Try to load existing bucket
	if bucket, ok := rl.buckets.Load(key); ok {
		//nolint:errcheck // type assertion safe here
		tb, _ := bucket.(*tokenBucket)
		return tb
	}

	// Create new bucket
	newBucket := newTokenBucket(rl.config.Rate, rl.config.Burst, rl.config.Interval)

	// Store and return (LoadOrStore handles race conditions)
	actual, _ := rl.buckets.LoadOrStore(key, newBucket)
	//nolint:errcheck // type assertion safe here
	tb, _ := actual.(*tokenBucket)
	return tb
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
