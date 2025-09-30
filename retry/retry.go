// Package retry provides automatic retry logic with intelligent backoff strategies
// for handling transient failures in distributed systems.
//
// The retry package supports multiple backoff policies (exponential, linear, constant),
// error classification for determining retryability, and context-aware cancellation.
//
// Example usage:
//
//	r := retry.New[*User](&retry.Config{
//	    MaxAttempts:   3,
//	    InitialDelay:  100 * time.Millisecond,
//	    Multiplier:    2.0,
//	    BackoffPolicy: retry.BackoffExponential,
//	    Jitter:        true,
//	})
//
//	user, err := r.Do(ctx, func(ctx context.Context) (*User, error) {
//	    return fetchUser(ctx, userID)
//	})
package retry

import (
	"context"
	"errors"
	"log/slog"
	"time"

	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

// Retry is a generic interface for retry pattern implementation.
// It automatically retries failed operations with configurable backoff strategies.
type Retry[T any] interface {
	// Do executes the given function with automatic retries on failure.
	// It respects context cancellation and stops retrying if the context is cancelled.
	// Returns the result and error from the last attempt.
	Do(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
}

// retry is the concrete implementation of Retry.
type retry[T any] struct {
	config Config
}

// New creates a new Retry instance with the given configuration.
func New[T any](config *Config) Retry[T] {
	if config == nil {
		config = &Config{}
	}
	config.setDefaults()
	return &retry[T]{
		config: *config,
	}
}

// Do implements the Retry interface.
func (r *retry[T]) Do(ctx context.Context, fn func(context.Context) (T, error)) (T, error) {
	var result T
	var err error

	for attempt := 1; attempt <= r.config.MaxAttempts; attempt++ {
		// Check context before attempting
		if err := ctx.Err(); err != nil {
			return result, err
		}

		// Execute the function
		result, err = fn(ctx)

		// Success - return immediately
		if err == nil {
			return result, nil
		}

		// Check if error is retryable
		if !r.isRetryable(err) {
			r.logAttempt(ctx, attempt, err, false)
			return result, err
		}

		// Last attempt - don't wait
		if attempt == r.config.MaxAttempts {
			r.logAttempt(ctx, attempt, err, false)
			return result, err
		}

		// Log retry attempt
		r.logAttempt(ctx, attempt, err, true)

		// Call OnRetry callback
		if r.config.OnRetry != nil {
			r.config.OnRetry(attempt+1, err)
		}

		// Calculate backoff delay
		delay := calculateBackoff(
			r.config.BackoffPolicy,
			attempt,
			r.config.InitialDelay,
			r.config.MaxDelay,
			r.config.Multiplier,
			r.config.Jitter,
		)

		// Wait before retry with context cancellation support
		select {
		case <-time.After(delay):
			// Continue to next attempt
		case <-ctx.Done():
			return result, ctx.Err()
		}
	}

	return result, err
}

// isRetryable determines if an error should trigger a retry.
func (r *retry[T]) isRetryable(err error) bool {
	if err == nil {
		return false
	}

	// Custom predicate takes precedence
	if r.config.IsRetryable != nil {
		return r.config.IsRetryable(err)
	}

	// Check non-retryable errors first
	for _, nonRetryable := range r.config.NonRetryableErrors {
		if errors.Is(err, nonRetryable) {
			return false
		}
	}

	// Check retryable errors
	if len(r.config.RetryableErrors) > 0 {
		for _, retryable := range r.config.RetryableErrors {
			if errors.Is(err, retryable) {
				return true
			}
		}
		// If RetryableErrors is specified but error doesn't match, don't retry
		return false
	}

	// Check if error implements RetryableError interface
	if fortifyerrors.IsRetryable(err) {
		return true
	}

	// Default: retry all errors if no classification is configured
	return true
}

// logAttempt logs retry attempts using structured logging.
func (r *retry[T]) logAttempt(ctx context.Context, attempt int, err error, willRetry bool) {
	if r.config.Logger == nil {
		return
	}

	if willRetry {
		r.config.Logger.WarnContext(ctx, "retry attempt failed, retrying",
			slog.Int("attempt", attempt),
			slog.Int("max_attempts", r.config.MaxAttempts),
			slog.String("error", err.Error()),
		)
	} else {
		r.config.Logger.ErrorContext(ctx, "retry exhausted",
			slog.Int("attempt", attempt),
			slog.Int("max_attempts", r.config.MaxAttempts),
			slog.String("error", err.Error()),
		)
	}
}
