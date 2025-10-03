// Package fallback provides a fallback pattern for handling operation failures gracefully.
//
// The fallback pattern executes a primary operation and, if it fails, executes a fallback
// operation to provide a default value or alternative behavior. This is useful for
// graceful degradation in distributed systems.
//
// Example usage:
//
//	fb := fallback.New[*UserData](fallback.Config{
//	    Fallback: func(ctx context.Context, err error) (*UserData, error) {
//	        // Return cached data when primary fails
//	        return getCachedUserData(ctx)
//	    },
//	})
//
//	userData, err := fb.Execute(ctx, func(ctx context.Context) (*UserData, error) {
//	    return fetchUserDataFromAPI(ctx)
//	})
package fallback

import (
	"context"
	"log/slog"
)

// Fallback is the interface for fallback pattern implementation.
type Fallback[T any] interface {
	// Execute runs the primary function and falls back to the fallback function on failure.
	// Returns the result from either the primary or fallback function.
	Execute(ctx context.Context, primary func(context.Context) (T, error)) (T, error)
}

// fallback is the concrete implementation of Fallback.
type fallback[T any] struct {
	config Config[T]
}

// New creates a new Fallback with the given configuration.
func New[T any](config Config[T]) Fallback[T] {
	config.setDefaults()

	return &fallback[T]{
		config: config,
	}
}

// Execute implements the Fallback interface.
func (f *fallback[T]) Execute(ctx context.Context, primary func(context.Context) (T, error)) (T, error) {
	var zero T

	// Check context first
	if err := ctx.Err(); err != nil {
		return zero, err
	}

	// Execute primary function
	result, err := primary(ctx)

	// If successful, return immediately
	if err == nil {
		if f.config.OnSuccess != nil {
			f.safeCallback(f.config.OnSuccess)
		}
		f.logEvent("primary_success", slog.Attr{})
		return result, nil
	}

	// Check if we should use fallback for this error
	if f.config.ShouldFallback != nil && !f.config.ShouldFallback(err) {
		f.logEvent("no_fallback", slog.String("reason", "should_fallback_returned_false"))
		return zero, err
	}

	// Fallback function is required
	if f.config.Fallback == nil {
		f.logEvent("no_fallback", slog.String("reason", "fallback_not_configured"))
		return zero, err
	}

	// Call OnFallback callback
	if f.config.OnFallback != nil {
		f.safeCallback(func() {
			f.config.OnFallback(err)
		})
	}

	f.logEvent("fallback_triggered", slog.String("primary_error", err.Error()))

	// Execute fallback function
	fallbackResult, fallbackErr := f.config.Fallback(ctx, err)

	if fallbackErr != nil {
		f.logEvent("fallback_failed", slog.String("error", fallbackErr.Error()))
		// Return original error, not fallback error
		return zero, err
	}

	f.logEvent("fallback_success", slog.Attr{})
	return fallbackResult, nil
}

// logEvent logs fallback events using structured logging.
func (f *fallback[T]) logEvent(event string, attr slog.Attr) {
	if f.config.Logger == nil {
		return
	}

	args := []any{
		slog.String("pattern", "fallback"),
		slog.String("event", event),
	}

	if attr.Key != "" {
		args = append(args, attr)
	}

	f.config.Logger.Info("fallback event", args...)
}

// safeCallback executes a callback with panic recovery.
func (f *fallback[T]) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if f.config.Logger != nil {
				f.config.Logger.Error("fallback callback panic",
					slog.String("pattern", "fallback"),
					slog.Any("panic", r),
				)
			}
		}
	}()
	fn()
}
