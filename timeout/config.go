package timeout

import (
	"log/slog"
	"time"
)

// Config holds the configuration for a Timeout instance.
type Config struct {
	// OnTimeout is called when an operation times out.
	// This can be used for cleanup, logging, or metrics.
	OnTimeout func()

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// DefaultTimeout is the timeout duration used when Execute is called with 0 duration.
	// If DefaultTimeout is 0, defaults to 30 seconds.
	DefaultTimeout time.Duration
}

// setDefaults applies default values to unset configuration fields.
func (c *Config) setDefaults() {
	if c.DefaultTimeout == 0 {
		c.DefaultTimeout = 30 * time.Second
	}
}
