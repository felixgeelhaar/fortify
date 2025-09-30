package fallback

import (
	"context"
	"log/slog"
)

// Config configures the fallback pattern behavior.
type Config[T any] struct {
	// Fallback is the fallback function to execute when the primary function fails.
	// It receives the context and the error from the primary function.
	// Required.
	Fallback func(context.Context, error) (T, error)

	// ShouldFallback determines whether to execute the fallback function for a given error.
	// If nil, fallback is always executed on primary function failure.
	// Optional.
	ShouldFallback func(error) bool

	// OnFallback is called when the fallback function is triggered.
	// It receives the error from the primary function.
	// Optional.
	OnFallback func(error)

	// OnSuccess is called when the primary function succeeds.
	// Optional.
	OnSuccess func()

	// Logger is used for structured logging of fallback events.
	// If nil, no logging is performed.
	// Optional.
	Logger *slog.Logger
}

// setDefaults sets default values for optional configuration fields.
func (c *Config[T]) setDefaults() {
	// No defaults needed - all fields are optional except Fallback
}
