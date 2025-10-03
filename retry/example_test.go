package retry_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/retry"
)

// Example demonstrates basic retry usage with exponential backoff.
func Example() {
	// Create retry with default configuration (3 attempts, exponential backoff)
	r := retry.New[string](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 100,
	})

	attempt := 0
	result, err := r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		if attempt < 3 {
			return "", errors.New("temporary failure")
		}
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s after %d attempts\n", result, attempt)
	// Output: Result: success after 3 attempts
}

// Example_exponentialBackoff demonstrates exponential backoff with custom multiplier.
func Example_exponentialBackoff() {
	r := retry.New[int](retry.Config{
		MaxAttempts:   4,
		InitialDelay:  time.Millisecond * 100,
		MaxDelay:      time.Second * 5,
		Multiplier:    2.0,
		BackoffPolicy: retry.BackoffExponential,
		Jitter:        false, // Disable jitter for predictable output
	})

	attempt := 0
	_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		fmt.Printf("Attempt %d\n", attempt)
		return 0, errors.New("still failing")
	})

	fmt.Printf("Failed after %d attempts: %v\n", attempt, err)
	// Output:
	// Attempt 1
	// Attempt 2
	// Attempt 3
	// Attempt 4
	// Failed after 4 attempts: still failing
}

// Example_linearBackoff demonstrates linear backoff strategy.
func Example_linearBackoff() {
	r := retry.New[string](retry.Config{
		MaxAttempts:   3,
		InitialDelay:  time.Millisecond * 200,
		BackoffPolicy: retry.BackoffLinear,
		OnRetry: func(attempt int, err error) {
			fmt.Printf("Retry attempt %d after error: %v\n", attempt, err)
		},
	})

	attempt := 0
	//nolint:errcheck // intentionally ignoring error in example
	r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		return "", fmt.Errorf("error %d", attempt)
	})
	// Output:
	// Retry attempt 2 after error: error 1
	// Retry attempt 3 after error: error 2
}

// Example_constantBackoff demonstrates constant delay between retries.
func Example_constantBackoff() {
	r := retry.New[int](retry.Config{
		MaxAttempts:   4,
		InitialDelay:  time.Millisecond * 500,
		BackoffPolicy: retry.BackoffConstant,
	})

	attempt := 0
	_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		if attempt < 4 {
			return 0, errors.New("not yet")
		}
		return attempt, nil
	})

	if err == nil {
		fmt.Printf("Succeeded on attempt %d\n", attempt)
	}
	// Output: Succeeded on attempt 4
}

// Example_selectiveRetry demonstrates retrying only specific errors.
func Example_selectiveRetry() {
	var ErrTemporary = errors.New("temporary error")
	var ErrPermanent = errors.New("permanent error")

	r := retry.New[string](retry.Config{
		MaxAttempts:  5,
		InitialDelay: time.Millisecond * 100,
		IsRetryable: func(err error) bool {
			// Only retry temporary errors
			return errors.Is(err, ErrTemporary)
		},
	})

	// This will be retried
	attempt := 0
	_, err := r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		if attempt < 3 {
			return "", ErrTemporary
		}
		return "success", nil
	})

	fmt.Printf("Temporary error - attempts: %d, error: %v\n", attempt, err)

	// This will NOT be retried
	attempt = 0
	_, err = r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		return "", ErrPermanent
	})

	fmt.Printf("Permanent error - attempts: %d, error: %v\n", attempt, err)
	// Output:
	// Temporary error - attempts: 3, error: <nil>
	// Permanent error - attempts: 1, error: permanent error
}

// Example_retryableErrors demonstrates using a list of retryable errors.
func Example_retryableErrors() {
	var ErrNetworkTimeout = errors.New("network timeout")
	var ErrServiceUnavailable = errors.New("service unavailable")
	var ErrUnauthorized = errors.New("unauthorized")

	r := retry.New[int](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 100,
		RetryableErrors: []error{
			ErrNetworkTimeout,
			ErrServiceUnavailable,
		},
	})

	// Network timeout - will be retried
	attempt := 0
	_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		if attempt < 2 {
			return 0, ErrNetworkTimeout
		}
		return 42, nil
	})

	fmt.Printf("Network timeout - attempts: %d, success: %v\n", attempt, err == nil)

	// Unauthorized error - will NOT be retried
	attempt = 0
	_, err = r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		return 0, ErrUnauthorized
	})

	fmt.Printf("Unauthorized - attempts: %d, error: %v\n", attempt, err)
	// Output:
	// Network timeout - attempts: 2, success: true
	// Unauthorized - attempts: 1, error: unauthorized
}

// Example_nonRetryableErrors demonstrates excluding specific errors from retry.
func Example_nonRetryableErrors() {
	var ErrBadRequest = errors.New("bad request")
	var ErrInternalError = errors.New("internal error")

	r := retry.New[string](retry.Config{
		MaxAttempts:  4,
		InitialDelay: time.Millisecond * 100,
		NonRetryableErrors: []error{
			ErrBadRequest, // Don't retry client errors
		},
	})

	// Bad request - will NOT be retried (client error)
	attempt := 0
	_, err := r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		return "", ErrBadRequest
	})

	fmt.Printf("Bad request - attempts: %d, error present: %v\n", attempt, err != nil)

	// Internal error - will be retried (server error)
	attempt = 0
	_, err = r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		if attempt < 3 {
			return "", ErrInternalError
		}
		return "recovered", nil
	})

	fmt.Printf("Internal error - attempts: %d, error: %v\n", attempt, err)
	// Output:
	// Bad request - attempts: 1, error present: true
	// Internal error - attempts: 3, error: <nil>
}

// Example_contextCancellation demonstrates respecting context cancellation.
func Example_contextCancellation() {
	r := retry.New[int](retry.Config{
		MaxAttempts:  10,
		InitialDelay: time.Millisecond * 200,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	attempt := 0
	_, err := r.Do(ctx, func(ctx context.Context) (int, error) {
		attempt++
		return 0, errors.New("persistent failure")
	})

	// Context will be cancelled before all retries complete
	fmt.Printf("Attempts: %d, Error: %v\n", attempt, err)
	// Output: Attempts: 2, Error: context deadline exceeded
}

// Example_jitter demonstrates adding random jitter to prevent thundering herd.
func Example_jitter() {
	r := retry.New[string](retry.Config{
		MaxAttempts:   3,
		InitialDelay:  time.Millisecond * 100,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		Jitter:        true, // Add 0-10% random jitter
	})

	attempt := 0
	//nolint:errcheck // intentionally ignoring error in example
	r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		if attempt < 3 {
			return "", errors.New("retry me")
		}
		return "success", nil
	})

	fmt.Printf("Succeeded after %d attempts with jitter\n", attempt)
	// Output: Succeeded after 3 attempts with jitter
}

// Example_httpClient demonstrates retry with HTTP client pattern.
func Example_httpClient() {
	r := retry.New[int](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 500,
		MaxDelay:     time.Second * 5,
		IsRetryable: func(err error) bool {
			// Retry on network errors and 5xx server errors
			// Don't retry on 4xx client errors
			return err != nil // Simplified for example
		},
		OnRetry: func(attempt int, err error) {
			fmt.Printf("Retrying HTTP request (attempt %d): %v\n", attempt, err)
		},
	})

	statusCode, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		// Simulate HTTP call
		// resp, err := httpClient.Get("https://api.example.com")
		// if err != nil {
		//     return 0, err
		// }
		// return resp.StatusCode, nil
		return 200, nil
	})

	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}

	fmt.Printf("HTTP Status: %d\n", statusCode)
	// Output: HTTP Status: 200
}
