package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config holds the configuration for a Redis-backed RateLimiter.
type Config struct {
	// Client is the Redis client to use for rate limiting.
	// This can be a redis.Client, redis.ClusterClient, or redis.SentinelClient.
	// Required field - no default.
	Client redis.UniversalClient

	// KeyPrefix is prepended to all Redis keys to namespace rate limit data.
	// If empty, defaults to "fortify:ratelimit:".
	KeyPrefix string

	// KeyFunc extracts the rate limiting key from the context.
	// If nil, the key parameter passed to Allow/Wait/Take is used.
	// This is useful for extracting keys from request context (user ID, IP, tenant, etc.)
	KeyFunc func(ctx context.Context) string

	// OnLimit is called when a request is rate limited.
	// It receives the key that was rate limited.
	OnLimit func(key string)

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// Interval is the time period over which Rate tokens are added.
	// If Interval is 0, defaults to 1 second.
	Interval time.Duration

	// Rate is the number of tokens added to the bucket per Interval.
	// If Rate is 0, defaults to 100.
	Rate int

	// Burst is the maximum number of tokens in the bucket (bucket capacity).
	// This allows short bursts of requests up to this limit.
	// If Burst is 0, defaults to Rate.
	Burst int

	// BucketTTL is the time-to-live for idle rate limit buckets in Redis.
	// Unused buckets are automatically cleaned up after this duration.
	// If BucketTTL is 0, defaults to 1 hour.
	BucketTTL time.Duration

	// FallbackOnError determines behavior when Redis is unavailable.
	// If true, allows all requests when Redis operations fail.
	// If false (default), denies requests when Redis is unavailable.
	// This provides flexibility between availability and rate limit enforcement.
	FallbackOnError bool
}

// setDefaults applies default values to unset configuration fields.
func (c *Config) setDefaults() {
	if c.Rate <= 0 {
		c.Rate = 100
	}

	if c.Burst <= 0 {
		c.Burst = c.Rate
	}

	if c.Interval == 0 {
		c.Interval = time.Second
	}

	if c.KeyPrefix == "" {
		c.KeyPrefix = "fortify:ratelimit:"
	}

	if c.BucketTTL == 0 {
		c.BucketTTL = time.Hour
	}
}

// validate checks that required configuration is present and valid.
func (c *Config) validate() error {
	if c.Client == nil {
		return ErrNoRedisClient
	}

	if c.Rate <= 0 {
		return ErrInvalidRate
	}

	if c.Burst <= 0 {
		return ErrInvalidBurst
	}

	if c.Interval <= 0 {
		return ErrInvalidInterval
	}

	return nil
}
