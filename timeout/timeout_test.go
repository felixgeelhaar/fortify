package timeout

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestTimeoutExecute(t *testing.T) {
	t.Run("executes function within timeout", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		result, err := tm.Execute(context.Background(), 100*time.Millisecond, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want 42", result)
		}
	})

	t.Run("returns error when function exceeds timeout", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		_, err := tm.Execute(context.Background(), 50*time.Millisecond, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 42, nil
		})

		if err == nil {
			t.Error("expected timeout error")
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("error = %v, want context.DeadlineExceeded", err)
		}
	})

	t.Run("propagates function error", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		expectedErr := errors.New("function error")
		_, err := tm.Execute(context.Background(), time.Second, func(ctx context.Context) (int, error) {
			return 0, expectedErr
		})

		if !errors.Is(err, expectedErr) {
			t.Errorf("error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("respects parent context cancellation", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := tm.Execute(ctx, time.Second, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("error = %v, want context.Canceled", err)
		}
	})

	t.Run("uses default timeout when zero provided", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: 50 * time.Millisecond,
		})

		_, err := tm.Execute(context.Background(), 0, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 42, nil
		})

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("should timeout using default timeout")
		}
	})

	t.Run("function receives context with timeout", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		var receivedDeadline time.Time
		hasDeadline := false

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = tm.Execute(context.Background(), 100*time.Millisecond, func(ctx context.Context) (int, error) {
			receivedDeadline, hasDeadline = ctx.Deadline()
			return 42, nil
		})

		if !hasDeadline {
			t.Error("function context should have deadline")
		}
		if time.Until(receivedDeadline) > 100*time.Millisecond {
			t.Error("deadline should be approximately 100ms from now")
		}
	})
}

func TestTimeoutCallback(t *testing.T) {
	t.Run("calls OnTimeout callback", func(t *testing.T) {
		timeoutCalled := make(chan bool, 1)

		tm := New[int](Config{
			DefaultTimeout: time.Second,
			OnTimeout: func() {
				timeoutCalled <- true
			},
		})

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = tm.Execute(context.Background(), 50*time.Millisecond, func(ctx context.Context) (int, error) {
			time.Sleep(100 * time.Millisecond)
			return 42, nil
		})

		select {
		case <-timeoutCalled:
			// Success
		case <-time.After(200 * time.Millisecond):
			t.Error("OnTimeout callback not called")
		}
	})

	t.Run("does not call OnTimeout on success", func(t *testing.T) {
		timeoutCalled := false

		tm := New[int](Config{
			DefaultTimeout: time.Second,
			OnTimeout: func() {
				timeoutCalled = true
			},
		})

		//nolint:errcheck // intentionally ignoring error in test
		_, _ = tm.Execute(context.Background(), 100*time.Millisecond, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		time.Sleep(50 * time.Millisecond)

		if timeoutCalled {
			t.Error("OnTimeout should not be called on success")
		}
	})
}

func TestTimeoutConcurrent(t *testing.T) {
	t.Run("handles concurrent executions", func(t *testing.T) {
		tm := New[int](Config{
			DefaultTimeout: time.Second,
		})

		ctx := context.Background()
		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func(val int) {
				result, err := tm.Execute(ctx, 100*time.Millisecond, func(ctx context.Context) (int, error) {
					time.Sleep(10 * time.Millisecond)
					return val, nil
				})
				if err != nil || result != val {
					t.Errorf("concurrent execution failed: result=%d, err=%v", result, err)
				}
				done <- true
			}(i)
		}

		// Wait for all goroutines
		for i := 0; i < 10; i++ {
			<-done
		}
	})
}

func TestTimeoutWithSlowFunction(t *testing.T) {
	t.Run("returns timeout error for slow function", func(t *testing.T) {
		tm := New[string](Config{
			DefaultTimeout: time.Second,
		})

		_, err := tm.Execute(context.Background(), 50*time.Millisecond, func(ctx context.Context) (string, error) {
			// Simulate slow operation
			time.Sleep(200 * time.Millisecond)
			return "completed", nil
		})

		// Should get timeout error
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("error = %v, want context.DeadlineExceeded", err)
		}
	})
}

func TestTimeoutDefaults(t *testing.T) {
	t.Run("applies default configuration", func(t *testing.T) {
		tm := New[int](Config{})

		ctx := context.Background()

		// Should have some default timeout
		result, err := tm.Execute(ctx, 0, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error with defaults: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want 42", result)
		}
	})
}
