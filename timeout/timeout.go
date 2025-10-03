// Package timeout provides context-based timeout enforcement for operations,
// allowing graceful degradation when operations exceed time limits.
//
// The timeout package wraps operations with context deadlines, automatically
// cancelling them when the timeout is exceeded. It supports configurable
// default timeouts and cleanup callbacks.
//
// Example usage:
//
//	t := timeout.New[[]byte](timeout.Config{
//	    DefaultTimeout: 5 * time.Second,
//	    OnTimeout: func() {
//	        log.Println("Operation timed out")
//	    },
//	})
//
//	data, err := t.Execute(ctx, 2*time.Second, func(ctx context.Context) ([]byte, error) {
//	    return fetchData(ctx)
//	})
package timeout

import (
	"context"
	"log/slog"
	"time"
)

// Timeout is a generic interface for enforcing operation time limits.
// It wraps operations with context deadlines and provides timeout callbacks.
type Timeout[T any] interface {
	// Execute runs the given function with a timeout.
	// If timeout is 0, uses the DefaultTimeout from Config.
	// Returns the function result or a context.DeadlineExceeded error if the timeout is exceeded.
	Execute(ctx context.Context, timeout time.Duration, fn func(context.Context) (T, error)) (T, error)
}

// timeout is the concrete implementation of Timeout.
type timeout[T any] struct {
	config Config
}

// New creates a new Timeout instance with the given configuration.
func New[T any](config Config) Timeout[T] {
	config.setDefaults()
	return &timeout[T]{
		config: config,
	}
}

// Execute implements the Timeout interface.
//
// IMPORTANT: The provided function MUST respect context cancellation.
// If the function does not check ctx.Err() or respond to context.Done(),
// the timeout will not be enforced and the function will run to completion.
func (t *timeout[T]) Execute(ctx context.Context, timeout time.Duration, fn func(context.Context) (T, error)) (T, error) {
	var zero T

	// Use default timeout if not specified
	if timeout == 0 {
		timeout = t.config.DefaultTimeout
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Execute function directly - context handles cancellation
	result, err := fn(ctx)

	// Check context errors after execution
	ctxErr := ctx.Err()
	if ctxErr != nil {
		// Timeout occurred
		if ctxErr == context.DeadlineExceeded {
			t.logTimeout(timeout)
			if t.config.OnTimeout != nil {
				// Run callback synchronously - it's panic-safe via safeCallback
				// Synchronous execution prevents goroutine leaks and ensures
				// callbacks complete before timeout error is returned
				t.safeCallback(t.config.OnTimeout)
			}
		}
		// Return context error (DeadlineExceeded or Canceled from parent)
		return zero, ctxErr
	}

	return result, err
}

// safeCallback executes a callback with panic recovery.
func (t *timeout[T]) safeCallback(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			if t.config.Logger != nil {
				t.config.Logger.Error("timeout callback panic",
					slog.String("pattern", "timeout"),
					slog.Any("panic", r),
				)
			}
		}
	}()
	fn()
}

// logTimeout logs timeout events using structured logging.
func (t *timeout[T]) logTimeout(duration time.Duration) {
	if t.config.Logger != nil {
		t.config.Logger.Warn("operation timeout",
			slog.Duration("timeout", duration),
		)
	}
}
