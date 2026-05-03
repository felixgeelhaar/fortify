package hedge

import (
	"log/slog"
	"time"
)

// defaultMaxAttempts is the default upper bound on parallel attempts when
// MaxAttempts is unset.
const defaultMaxAttempts = 2

// maxAttemptsCap bounds MaxAttempts to prevent fan-out storms from a
// pathologically slow primary; sixteen parallel attempts is already extreme.
const maxAttemptsCap = 16

// defaultHedgeDelay is the default wait before firing the first hedge attempt.
const defaultHedgeDelay = 100 * time.Millisecond

// Config holds the configuration for a Hedge instance.
type Config struct {
	// OnHedge is called each time a hedge attempt is fired (i.e., not the
	// initial primary). Receives the 1-based attempt index of the just-fired
	// hedge (so the second attempt overall is OnHedge(2)).
	//
	// The callback is wrapped with panic recovery and runs synchronously on
	// the caller's goroutine.
	OnHedge func(attempt int)

	// Logger is used for structured logging. If nil, no logging is performed.
	Logger *slog.Logger

	// HedgeDelay is the time to wait before firing the next hedge attempt
	// after the previous one. Defaults to 100ms.
	//
	// A short HedgeDelay reduces tail latency at the cost of extra work
	// (more parallel calls). A long HedgeDelay approaches the no-hedge case.
	HedgeDelay time.Duration

	// MaxAttempts is the total number of parallel attempts INCLUDING the
	// primary. MaxAttempts == 1 disables hedging. Defaults to 2.
	// Capped at 16.
	MaxAttempts int
}

// setDefaults applies default values to unset configuration fields and
// clamps invalid values to safe defaults.
func (c *Config) setDefaults() {
	if c.MaxAttempts <= 0 {
		c.MaxAttempts = defaultMaxAttempts
	}
	if c.MaxAttempts > maxAttemptsCap {
		c.MaxAttempts = maxAttemptsCap
	}
	if c.HedgeDelay <= 0 {
		c.HedgeDelay = defaultHedgeDelay
	}
}
