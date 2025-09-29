package bulkhead

import (
	"log/slog"
	"time"
)

// Config holds the configuration for a Bulkhead instance.
type Config struct {
	// QueueTimeout is the maximum time a request can wait in the queue.
	// If 0, no timeout is applied (requests can wait indefinitely).
	QueueTimeout time.Duration

	// OnRejected is called when a request is rejected due to bulkhead being full.
	// This can be used for metrics or alerting.
	OnRejected func()

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// MaxConcurrent is the maximum number of concurrent executions allowed.
	// If MaxConcurrent is 0, defaults to 10.
	MaxConcurrent int

	// MaxQueue is the maximum number of requests that can be queued
	// when the bulkhead is at capacity. If 0, no queueing is allowed.
	MaxQueue int
}

// setDefaults applies default values to unset configuration fields.
func (c *Config) setDefaults() {
	if c.MaxConcurrent == 0 {
		c.MaxConcurrent = 10
	}
}
