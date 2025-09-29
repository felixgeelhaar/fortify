package ratelimit

import (
	"context"
	"log/slog"
	"time"
)

// Config holds the configuration for a RateLimiter.
type Config struct {
	// Interval is the time period over which Rate tokens are added.
	// If Interval is 0, defaults to 1 second.
	Interval time.Duration

	// KeyFunc extracts the rate limiting key from the context.
	// If nil, the key parameter passed to Allow/Wait/Take is used.
	// This is useful for extracting keys from request context (user ID, IP, tenant, etc.)
	KeyFunc func(ctx context.Context) string

	// OnLimit is called when a request is rate limited.
	// It receives the key that was rate limited.
	OnLimit func(key string)

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// Rate is the number of tokens added to the bucket per Interval.
	// If Rate is 0, defaults to 100.
	Rate int

	// Burst is the maximum number of tokens in the bucket (bucket capacity).
	// This allows short bursts of requests up to this limit.
	// If Burst is 0, defaults to Rate.
	Burst int
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
}
