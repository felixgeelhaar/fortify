// Package ferrors provides standard error types and utilities for the fortify resilience package.
//
// This package defines sentinel errors used across all resilience patterns and provides
// a RetryableError interface for classifying errors that should trigger retry logic.
package ferrors

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Sentinel errors used by resilience patterns.
var (
	// ErrCircuitOpen indicates the circuit breaker is in open state and rejecting requests.
	ErrCircuitOpen = errors.New("circuit breaker is open")

	// ErrRateLimitExceeded indicates the rate limit has been exceeded for the given key.
	ErrRateLimitExceeded = errors.New("rate limit exceeded")

	// ErrBulkheadFull indicates the bulkhead is at capacity and cannot accept more requests.
	ErrBulkheadFull = errors.New("bulkhead at capacity")

	// ErrTimeout indicates an operation has exceeded its timeout duration.
	ErrTimeout = errors.New("operation timeout")

	// ErrMaxAttemptsReached indicates the maximum number of retry attempts has been reached.
	ErrMaxAttemptsReached = errors.New("max retry attempts reached")
)

// RetryableError is an interface for errors that indicate the operation should be retried.
// Implementations of this interface signal to retry logic that the underlying error
// represents a transient failure that may succeed on retry.
type RetryableError interface {
	error
	// Retryable returns true if the error represents a retryable failure.
	Retryable() bool
}

// retryableError wraps an error and marks it as retryable.
type retryableError struct {
	err error
}

// Error implements the error interface, returning the wrapped error's message.
func (e *retryableError) Error() string {
	if e.err == nil {
		return "retryable: <nil>"
	}
	return e.err.Error()
}

// Retryable always returns true for retryableError.
func (e *retryableError) Retryable() bool {
	return true
}

// Unwrap returns the underlying error, supporting error unwrapping with errors.Is and errors.As.
func (e *retryableError) Unwrap() error {
	return e.err
}

// AsRetryable wraps an error to mark it as retryable.
// If err is nil, returns a retryableError wrapping nil.
//
// Example:
//
//	if err := callExternalAPI(); err != nil {
//	    return AsRetryable(err) // Mark as retryable
//	}
func AsRetryable(err error) error {
	return &retryableError{err: err}
}

// IsRetryable checks if an error is marked as retryable.
// Returns true if the error or any error in its chain implements RetryableError
// and returns true from its Retryable() method.
//
// Example:
//
//	if IsRetryable(err) {
//	    // Retry the operation
//	}
func IsRetryable(err error) bool {
	if err == nil {
		return false
	}
	var re RetryableError
	return errors.As(err, &re) && re.Retryable()
}

// CircuitOpenError is returned when a circuit breaker rejects a request because
// it is in the Open state (or HalfOpen with no remaining trial capacity).
// Use errors.As to extract the structured fields:
//
//	var coe *ferrors.CircuitOpenError
//	if errors.As(err, &coe) {
//	    log.Printf("CB %q is %s; retry after %s", coe.Name, coe.State, coe.RetryAfter)
//	}
//
// errors.Is(err, ErrCircuitOpen) continues to match.
type CircuitOpenError struct {
	// Name is an optional identifier for the circuit breaker (empty if unset).
	Name string
	// State is the breaker state at the time of rejection.
	State string
	// RetryAfter is the duration until the breaker is expected to allow trial requests.
	// Zero if unknown.
	RetryAfter time.Duration
	// Counts captures the request counters that drove the rejection (informational).
	TotalRequests        uint32
	TotalFailures        uint32
	ConsecutiveFailures  uint32
	ConsecutiveSuccesses uint32
}

// Error implements error.
func (e *CircuitOpenError) Error() string {
	if e.RetryAfter > 0 {
		return fmt.Sprintf("circuit breaker is open (state=%s, retry_after=%s)", e.State, e.RetryAfter)
	}
	if e.State != "" {
		return fmt.Sprintf("circuit breaker is open (state=%s)", e.State)
	}
	return ErrCircuitOpen.Error()
}

// Unwrap allows errors.Is(err, ErrCircuitOpen) to keep matching.
func (e *CircuitOpenError) Unwrap() error { return ErrCircuitOpen }

// RateLimitError is returned when a rate limiter denies a request.
// Use errors.As to extract Key and RetryAfter; errors.Is(err, ErrRateLimitExceeded) keeps matching.
type RateLimitError struct {
	// Key is the rate-limit key that was denied.
	Key string
	// RetryAfter is the estimated time until a token becomes available.
	// Zero if unknown.
	RetryAfter time.Duration
}

// Error implements error.
func (e *RateLimitError) Error() string {
	if e.RetryAfter > 0 {
		return fmt.Sprintf("rate limit exceeded (retry_after=%s)", e.RetryAfter)
	}
	return ErrRateLimitExceeded.Error()
}

// Unwrap allows errors.Is(err, ErrRateLimitExceeded) to keep matching.
func (e *RateLimitError) Unwrap() error { return ErrRateLimitExceeded }

// TimeoutError is returned when an operation exceeds its configured deadline.
// Use errors.As to read Timeout; errors.Is(err, ErrTimeout) keeps matching.
type TimeoutError struct {
	// Timeout is the configured deadline that was exceeded.
	Timeout time.Duration
}

// Error implements error.
func (e *TimeoutError) Error() string {
	if e.Timeout > 0 {
		return fmt.Sprintf("operation timeout after %s", e.Timeout)
	}
	return ErrTimeout.Error()
}

// Unwrap exposes both ErrTimeout and context.DeadlineExceeded so that
// errors.Is matches either sentinel. Existing callers using
// errors.Is(err, context.DeadlineExceeded) continue to work.
func (e *TimeoutError) Unwrap() []error {
	return []error{ErrTimeout, context.DeadlineExceeded}
}
