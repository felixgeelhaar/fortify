// Package redis provides a Redis-backed rate limiter for distributed systems.
//
// This package implements the same RateLimiter interface as the core fortify/ratelimit
// package but stores state in Redis, enabling rate limiting across multiple application
// instances in a distributed system.
//
// The implementation uses Lua scripts for atomic operations and follows production-grade
// best practices including proper error handling, observability integration, and support
// for Redis Cluster and Sentinel.
//
// Example usage:
//
//	import (
//	    "github.com/redis/go-redis/v9"
//	    redisrl "github.com/felixgeelhaar/fortify/backends/redis"
//	)
//
//	// Create Redis client
//	client := redis.NewClient(&redis.Options{
//	    Addr: "localhost:6379",
//	})
//
//	// Create distributed rate limiter
//	limiter := redisrl.New(redisrl.Config{
//	    Client:   client,
//	    Rate:     100,
//	    Burst:    200,
//	    Interval: time.Second,
//	})
//
//	// Use same interface as in-memory limiter
//	if limiter.Allow(ctx, "user-123") {
//	    // Process request
//	} else {
//	    // Return 429 Too Many Requests
//	}
package redis

import (
	"context"
	_ "embed"
	"errors"
	"log/slog"
	"time"

	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/redis/go-redis/v9"
)

//go:embed lua/token_bucket.lua
var tokenBucketScript string

// Sentinel errors for configuration validation.
var (
	ErrNoRedisClient  = errors.New("redis client is required")
	ErrInvalidRate    = errors.New("rate must be positive")
	ErrInvalidBurst   = errors.New("burst must be positive")
	ErrInvalidInterval = errors.New("interval must be positive")
)

// redisRateLimiter is the Redis-backed implementation of RateLimiter.
type redisRateLimiter struct {
	config Config
	script *redis.Script
}

// New creates a new Redis-backed RateLimiter with the given configuration.
// The returned RateLimiter implements the same interface as the in-memory version,
// allowing for drop-in replacement in distributed systems.
func New(config Config) (ratelimit.RateLimiter, error) {
	config.setDefaults()

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &redisRateLimiter{
		config: config,
		script: redis.NewScript(tokenBucketScript),
	}, nil
}

// Allow implements the RateLimiter interface.
// It performs a non-blocking check against Redis to determine if the request
// should be allowed based on the token bucket state.
func (rl *redisRateLimiter) Allow(ctx context.Context, key string) bool {
	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)
	redisKey := rl.config.KeyPrefix + resolvedKey

	// Prepare script arguments
	result, err := rl.executeScript(ctx, redisKey, 1)
	if err != nil {
		return rl.handleError(ctx, resolvedKey, err)
	}

	allowed := result == 1

	// Log and callback if rate limited
	if !allowed {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.OnLimit != nil {
			rl.safeCallback(func() {
				rl.config.OnLimit(resolvedKey)
			})
		}
	}

	return allowed
}

// Wait implements the RateLimiter interface.
// It blocks until a token is available or the context is cancelled.
// This is implemented as a retry loop with exponential backoff to avoid
// overwhelming Redis with polling requests.
func (rl *redisRateLimiter) Wait(ctx context.Context, key string) error {
	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)
	redisKey := rl.config.KeyPrefix + resolvedKey

	// Initial backoff: start with 10ms, max 100ms
	backoff := 10 * time.Millisecond
	maxBackoff := 100 * time.Millisecond

	for {
		// Check context first
		if err := ctx.Err(); err != nil {
			return err
		}

		// Try to take a token
		result, err := rl.executeScript(ctx, redisKey, 1)
		if err != nil {
			// Handle Redis errors
			if rl.config.FallbackOnError {
				return nil // Allow on error if configured
			}
			return err
		}

		// If allowed, return success
		if result == 1 {
			return nil
		}

		// Wait before retrying with exponential backoff
		select {
		case <-time.After(backoff):
			// Increase backoff for next iteration
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Take implements the RateLimiter interface.
// It attempts to take n tokens from the bucket atomically.
func (rl *redisRateLimiter) Take(ctx context.Context, key string, tokens int) bool {
	if tokens <= 0 {
		return false
	}

	// Resolve key
	resolvedKey := rl.resolveKey(ctx, key)
	redisKey := rl.config.KeyPrefix + resolvedKey

	// Execute script with specified token count
	result, err := rl.executeScript(ctx, redisKey, tokens)
	if err != nil {
		return rl.handleError(ctx, resolvedKey, err)
	}

	taken := result == 1

	// Log and callback if rate limited
	if !taken {
		rl.logRateLimit(ctx, resolvedKey)
		if rl.config.OnLimit != nil {
			rl.safeCallback(func() {
				rl.config.OnLimit(resolvedKey)
			})
		}
	}

	return taken
}

// executeScript runs the token bucket Lua script in Redis.
// This ensures atomic refill + consume operations.
func (rl *redisRateLimiter) executeScript(ctx context.Context, redisKey string, tokensToTake int) (int64, error) {
	now := time.Now().UnixNano()

	result, err := rl.script.Run(ctx, rl.config.Client, []string{redisKey},
		rl.config.Rate,                        // ARGV[1]: rate
		rl.config.Burst,                       // ARGV[2]: burst
		rl.config.Interval.Nanoseconds(),      // ARGV[3]: interval_ns
		tokensToTake,                          // ARGV[4]: tokens_to_take
		now,                                   // ARGV[5]: current_time_ns
		int64(rl.config.BucketTTL.Seconds()), // ARGV[6]: ttl_seconds
	).Result()

	if err != nil {
		return 0, err
	}

	// Result should be 1 (allowed) or 0 (denied)
	allowed, ok := result.(int64)
	if !ok {
		return 0, errors.New("unexpected script result type")
	}

	return allowed, nil
}

// resolveKey determines the rate limiting key using KeyFunc if configured.
func (rl *redisRateLimiter) resolveKey(ctx context.Context, key string) string {
	if rl.config.KeyFunc != nil {
		return rl.config.KeyFunc(ctx)
	}
	return key
}

// handleError processes Redis errors according to configuration.
func (rl *redisRateLimiter) handleError(ctx context.Context, key string, err error) bool {
	// Log the error
	if rl.config.Logger != nil {
		rl.config.Logger.ErrorContext(ctx, "redis rate limiter error",
			slog.String("key", key),
			slog.String("error", err.Error()),
		)
	}

	// Return based on fallback configuration
	// FallbackOnError = true: allow on error (availability over consistency)
	// FallbackOnError = false: deny on error (consistency over availability)
	return rl.config.FallbackOnError
}

// logRateLimit logs rate limiting events using structured logging.
func (rl *redisRateLimiter) logRateLimit(ctx context.Context, key string) {
	if rl.config.Logger != nil {
		rl.config.Logger.WarnContext(ctx, "rate limit exceeded",
			slog.String("backend", "redis"),
			slog.String("key", key),
			slog.Int("rate", rl.config.Rate),
			slog.Int("burst", rl.config.Burst),
		)
	}
}

// safeCallback executes a callback with panic recovery.
func (rl *redisRateLimiter) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if rl.config.Logger != nil {
				rl.config.Logger.Error("redis rate limiter callback panic",
					slog.String("backend", "redis"),
					slog.Any("panic", r),
				)
			}
		}
	}()
	fn()
}
