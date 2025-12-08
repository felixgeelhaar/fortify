package ratelimit

import (
	"context"
	"log/slog"
	"time"
)

// Config holds the configuration for a RateLimiter.
type Config struct {
	// KeyFunc extracts the rate limiting key from the context.
	// If nil, the key parameter passed to Allow/Wait/Take is used.
	// This is useful for extracting keys from request context (user ID, IP, tenant, etc.)
	//
	// IMPORTANT: When KeyFunc is set, the key parameter in Allow/Wait/Take is IGNORED.
	// KeyFunc takes precedence and the returned value is used as the rate limiting key.
	// If you need to use the key parameter, either don't set KeyFunc or incorporate
	// the key into your context before calling the rate limiter methods.
	//
	// Key length constraints: The returned key must respect the Store's maximum key
	// length (e.g., DefaultMaxKeyLength=1024 for MemoryStore). Keys exceeding the
	// limit will cause ErrKeyTooLong errors from the Store. Keep keys reasonably
	// sized (e.g., user IDs, tenant names, IP addresses).
	KeyFunc func(ctx context.Context) string

	// OnLimit is called when a request is rate limited.
	// It receives the context and key that was rate limited.
	//
	// By default, the callback is executed synchronously, which may block the hot path.
	// For non-blocking behavior, either:
	//   1. Keep the callback lightweight (logging, counter increment)
	//   2. Dispatch to a goroutine or channel within your callback
	//   3. Use a bounded worker pool for heavy operations
	//
	// The callback is wrapped with panic recovery - panics are logged but don't crash.
	OnLimit func(ctx context.Context, key string)

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// Metrics is used for metrics collection. If nil, no metrics are recorded.
	// Implement the Metrics interface to integrate with your metrics system.
	Metrics Metrics

	// Store is the storage backend for rate limiter state.
	// If nil, an in-memory store is created automatically using NewMemoryStore().
	// For distributed rate limiting across multiple instances, provide a custom
	// Store implementation backed by Redis, DynamoDB, or another distributed backend.
	Store Store

	// Interval is the time period over which Rate tokens are added.
	// Defaults to 1 second if zero.
	Interval time.Duration

	// Rate is the number of tokens added to the bucket per Interval.
	// Must be positive. Defaults to 100 if zero or negative.
	//
	// Example: Rate=10 with Interval=time.Second allows 10 requests per second.
	Rate int

	// Burst is the maximum number of tokens in the bucket (bucket capacity).
	// This allows short bursts of requests up to this limit.
	// Must be positive. Defaults to Rate if zero or negative.
	//
	// Setting Burst > Rate allows temporary bursts above the sustained rate.
	// Setting Burst == Rate enforces strict rate limiting with no burst.
	Burst int

	// MaxTokensPerRequest limits the maximum tokens that can be requested
	// in a single Take() call. This prevents integer overflow and DoS attacks.
	// Defaults to Burst * 10 if zero or negative.
	MaxTokensPerRequest int

	// FailOpen determines behavior when storage operations fail.
	// If true, allows requests when storage is unavailable (favors availability).
	// If false (default), denies requests when storage fails (favors consistency).
	// This only applies when using a custom Store that can fail (e.g., Redis, DynamoDB).
	//
	// Security note: FailOpen=true may allow more requests than intended during
	// storage outages. Use with caution in security-critical applications.
	FailOpen bool
}

const (
	// DefaultRate is the default rate limit.
	DefaultRate = 100

	// MaxRate is the maximum allowed rate limit.
	MaxRate = 1000000

	// MaxBurst is the maximum allowed burst size.
	MaxBurst = 1000000

	// MaxInterval is the maximum allowed interval.
	MaxInterval = 24 * time.Hour

	// DefaultMaxTokensMultiplier is the default multiplier for MaxTokensPerRequest.
	DefaultMaxTokensMultiplier = 10

	// MaxMaxTokensPerRequest is the maximum allowed value for MaxTokensPerRequest.
	// This prevents integer overflow in token calculations.
	MaxMaxTokensPerRequest = 10000000
)

// setDefaults applies default values to unset configuration fields.
func (c *Config) setDefaults() {
	if c.Rate <= 0 {
		c.Rate = DefaultRate
	}
	// Cap rate to prevent overflow
	if c.Rate > MaxRate {
		c.Rate = MaxRate
	}

	if c.Burst <= 0 {
		c.Burst = c.Rate
	}
	// Cap burst to prevent overflow
	if c.Burst > MaxBurst {
		c.Burst = MaxBurst
	}

	if c.Interval <= 0 {
		c.Interval = time.Second
	}
	// Cap interval to reasonable maximum
	if c.Interval > MaxInterval {
		c.Interval = MaxInterval
	}

	if c.MaxTokensPerRequest <= 0 {
		c.MaxTokensPerRequest = c.Burst * DefaultMaxTokensMultiplier
	}
	// Cap MaxTokensPerRequest to prevent overflow
	if c.MaxTokensPerRequest > MaxMaxTokensPerRequest {
		c.MaxTokensPerRequest = MaxMaxTokensPerRequest
	}

	if c.Store == nil {
		c.Store = NewMemoryStore()
	}
}
