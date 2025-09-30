package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

func TestRetryDo(t *testing.T) {
	t.Run("succeeds on first attempt", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts: 3,
		})

		attempts := 0
		result, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want %v", result, 42)
		}
		if attempts != 1 {
			t.Errorf("attempts = %v, want 1", attempts)
		}
	})

	t.Run("retries on failure and succeeds", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:  3,
			InitialDelay: 10 * time.Millisecond,
		})

		attempts := 0
		result, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			if attempts < 3 {
				return 0, errors.New("temporary failure")
			}
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want %v", result, 42)
		}
		if attempts != 3 {
			t.Errorf("attempts = %v, want 3", attempts)
		}
	})

	t.Run("returns error after max attempts", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:  3,
			InitialDelay: 10 * time.Millisecond,
		})

		attempts := 0
		expectedErr := errors.New("persistent failure")
		_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, expectedErr
		})

		if !errors.Is(err, expectedErr) {
			t.Errorf("error = %v, want %v", err, expectedErr)
		}
		if attempts != 3 {
			t.Errorf("attempts = %v, want 3", attempts)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:  5,
			InitialDelay: 100 * time.Millisecond,
		})

		ctx, cancel := context.WithCancel(context.Background())

		attempts := 0
		_, err := r.Do(ctx, func(ctx context.Context) (int, error) {
			attempts++
			if attempts == 2 {
				cancel()
			}
			return 0, errors.New("failure")
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("error = %v, want %v", err, context.Canceled)
		}
		if attempts > 3 {
			t.Errorf("attempts = %v, should stop after context cancel", attempts)
		}
	})

	t.Run("does not retry non-retryable errors", func(t *testing.T) {
		nonRetryableErr := errors.New("non-retryable")

		r := New[int](Config{
			MaxAttempts:        5,
			InitialDelay:       10 * time.Millisecond,
			NonRetryableErrors: []error{nonRetryableErr},
		})

		attempts := 0
		_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, nonRetryableErr
		})

		if !errors.Is(err, nonRetryableErr) {
			t.Errorf("error = %v, want %v", err, nonRetryableErr)
		}
		if attempts != 1 {
			t.Errorf("attempts = %v, want 1 (should not retry)", attempts)
		}
	})

	t.Run("retries only retryable errors", func(t *testing.T) {
		retryableErr := fortifyerrors.AsRetryable(errors.New("retryable"))

		r := New[int](Config{
			MaxAttempts:  3,
			InitialDelay: 10 * time.Millisecond,
		})

		attempts := 0
		_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, retryableErr
		})

		if attempts != 3 {
			t.Errorf("attempts = %v, want 3", attempts)
		}
		if !errors.Is(err, retryableErr) {
			t.Errorf("error = %v, want %v", err, retryableErr)
		}
	})

	t.Run("uses custom IsRetryable predicate", func(t *testing.T) {
		customErr := errors.New("custom error")

		r := New[int](Config{
			MaxAttempts:  3,
			InitialDelay: 10 * time.Millisecond,
			IsRetryable: func(err error) bool {
				return err.Error() == "should retry"
			},
		})

		attempts := 0
		result, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			if attempts == 1 {
				return 0, customErr // Should not retry
			}
			return 0, errors.New("should retry")
		})

		if attempts != 1 {
			t.Errorf("attempts = %v, want 1 (custom predicate should prevent retry)", attempts)
		}
		if !errors.Is(err, customErr) {
			t.Errorf("error = %v, want %v", err, customErr)
		}
		if result != 0 {
			t.Errorf("result = %v, want 0", result)
		}
	})
}

func TestRetryBackoff(t *testing.T) {
	t.Run("exponential backoff", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:   4,
			InitialDelay:  10 * time.Millisecond,
			Multiplier:    2.0,
			BackoffPolicy: BackoffExponential,
		})

		attempts := 0
		start := time.Now()

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, errors.New("failure")
		})

		duration := time.Since(start)

		// Expected delays: 10ms, 20ms, 40ms = 70ms minimum
		// Allow some variance for execution time
		if duration < 60*time.Millisecond {
			t.Errorf("duration = %v, expected at least 60ms for exponential backoff", duration)
		}
	})

	t.Run("linear backoff", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:   4,
			InitialDelay:  10 * time.Millisecond,
			Multiplier:    1.0,
			BackoffPolicy: BackoffLinear,
		})

		attempts := 0
		start := time.Now()

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, errors.New("failure")
		})

		duration := time.Since(start)

		// Expected delays: 10ms, 20ms, 30ms = 60ms minimum
		if duration < 50*time.Millisecond {
			t.Errorf("duration = %v, expected at least 50ms for linear backoff", duration)
		}
	})

	t.Run("constant backoff", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:   4,
			InitialDelay:  10 * time.Millisecond,
			BackoffPolicy: BackoffConstant,
		})

		attempts := 0
		start := time.Now()

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = r.Do(context.Background(), func(ctx context.Context) (int, error) {
			attempts++
			return 0, errors.New("failure")
		})

		duration := time.Since(start)

		// Expected delays: 10ms, 10ms, 10ms = 30ms minimum
		if duration < 25*time.Millisecond {
			t.Errorf("duration = %v, expected at least 25ms for constant backoff", duration)
		}
	})

	t.Run("respects max delay", func(t *testing.T) {
		r := New[int](Config{
			MaxAttempts:   5,
			InitialDelay:  10 * time.Millisecond,
			MaxDelay:      20 * time.Millisecond,
			Multiplier:    10.0,
			BackoffPolicy: BackoffExponential,
		})

		start := time.Now()

		r.Do(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		duration := time.Since(start)

		// Even with 10x multiplier, max delay should cap at 20ms
		// 4 retries * 20ms max = 80ms maximum
		if duration > 100*time.Millisecond {
			t.Errorf("duration = %v, should respect MaxDelay of 20ms", duration)
		}
	})
}

func TestRetryCallbacks(t *testing.T) {
	t.Run("OnRetry callback", func(t *testing.T) {
		var retries []int

		r := New[int](Config{
			MaxAttempts:  3,
			InitialDelay: 10 * time.Millisecond,
			OnRetry: func(attempt int, err error) {
				retries = append(retries, attempt)
			},
		})

		r.Do(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		// OnRetry should be called for attempts 2 and 3 (not the first attempt)
		if len(retries) != 2 {
			t.Errorf("OnRetry called %d times, want 2", len(retries))
		}
		if len(retries) > 0 && retries[0] != 2 {
			t.Errorf("first retry attempt = %d, want 2", retries[0])
		}
	})
}

func TestRetryWithJitter(t *testing.T) {
	r := New[int](Config{
		MaxAttempts:   10,
		InitialDelay:  10 * time.Millisecond,
		Multiplier:    2.0,
		BackoffPolicy: BackoffExponential,
		Jitter:        true,
	})

	// Run multiple times to verify jitter adds randomness
	durations := make([]time.Duration, 3)
	for i := 0; i < 3; i++ {
		start := time.Now()
		r.Do(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})
		durations[i] = time.Since(start)
	}

	// With jitter, durations should vary (not all identical)
	allSame := durations[0] == durations[1] && durations[1] == durations[2]
	if allSame {
		t.Error("jitter should introduce variance in retry delays")
	}
}
