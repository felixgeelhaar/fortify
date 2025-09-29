// Package errors provides standard error types and utilities for the fortify resilience package.
//
// This package defines sentinel errors used across all resilience patterns and provides
// a RetryableError interface for classifying errors that should trigger retry logic.
package errors

import "errors"

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