package ratelimit

import (
	"errors"
	"fmt"
)

// Sentinel errors for rate limiting operations.
//
// These errors are part of the public API and can be used by custom Store
// implementations or for error type checking. The built-in methods (Allow,
// Wait, Take) use boolean returns for simplicity, but custom implementations
// may return these sentinel errors for more detailed error handling.
//
// All errors support errors.Is() for type checking:
//
//	if errors.Is(err, ratelimit.ErrStoreClosed) { ... }
//
// To wrap errors with additional context while preserving Is() behavior:
//
//	return fmt.Errorf("redis error: %w", ratelimit.ErrStorageUnavailable)
var (
	// ErrLimitExceeded indicates the rate limit was exceeded.
	// This can be used by custom Store implementations or error wrappers.
	// Note: The standard Allow() method returns false instead of this error.
	ErrLimitExceeded = errors.New("ratelimit: limit exceeded")

	// ErrStorageUnavailable indicates the storage backend is unreachable.
	// Custom Store implementations should return this when the backend fails.
	ErrStorageUnavailable = errors.New("ratelimit: storage unavailable")

	// ErrInvalidTokenCount indicates zero or negative tokens were requested.
	// Custom implementations may return this; Take() returns false for simplicity.
	ErrInvalidTokenCount = errors.New("ratelimit: invalid token count")

	// ErrExcessiveTokens indicates the token request exceeds MaxTokensPerRequest.
	// Custom implementations may return this; Take() returns false for simplicity.
	ErrExcessiveTokens = errors.New("ratelimit: token request exceeds maximum")

	// ErrStoreClosed is returned when operations are attempted on a closed store.
	ErrStoreClosed = errors.New("ratelimit: store is closed")

	// ErrKeyLimitExceeded is returned when the maximum number of keys is reached.
	ErrKeyLimitExceeded = errors.New("ratelimit: key limit exceeded")

	// ErrKeyTooLong is returned when a key exceeds the maximum length.
	ErrKeyTooLong = errors.New("ratelimit: key exceeds maximum length")

	// ErrWaitTimeout is returned when Wait() exceeds maximum iterations or time limit.
	ErrWaitTimeout = errors.New("ratelimit: wait timeout exceeded")

	// ErrRateLimiterClosed is returned when operations are attempted on a closed rate limiter.
	ErrRateLimiterClosed = errors.New("ratelimit: rate limiter is closed")
)

// WrapStorageError wraps an underlying error with ErrStorageUnavailable context.
// This allows callers to use errors.Is(err, ErrStorageUnavailable) while also
// accessing the underlying cause via errors.Unwrap().
//
// Example:
//
//	return ratelimit.WrapStorageError(redisErr)
//	// Later: errors.Is(err, ratelimit.ErrStorageUnavailable) returns true
//	// And: errors.Unwrap(err) returns redisErr
func WrapStorageError(cause error) error {
	if cause == nil {
		return ErrStorageUnavailable
	}
	return fmt.Errorf("%w: %w", ErrStorageUnavailable, cause)
}
