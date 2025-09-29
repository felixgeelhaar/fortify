package circuitbreaker

import (
	"context"
	"errors"
	"testing"
	"time"

	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

func TestCircuitBreakerStates(t *testing.T) {
	t.Run("starts in closed state", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
			Timeout:     time.Second,
		})

		if cb.State() != StateClosed {
			t.Errorf("initial state = %v, want %v", cb.State(), StateClosed)
		}
	})

	t.Run("transitions to open after consecutive failures", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
			Timeout:     time.Second,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 3
			},
		})

		// Trigger 3 failures
		for i := 0; i < 3; i++ {
			_, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
				return 0, errors.New("failure")
			})
			if err == nil {
				t.Error("expected error, got nil")
			}
		}

		if cb.State() != StateOpen {
			t.Errorf("state after failures = %v, want %v", cb.State(), StateOpen)
		}
	})

	t.Run("rejects requests when open", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
			Timeout:     100 * time.Millisecond,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		// Trigger failure to open circuit
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		// Verify circuit is open
		if cb.State() != StateOpen {
			t.Fatalf("state = %v, want %v", cb.State(), StateOpen)
		}

		// Attempt execution while open
		_, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			t.Error("function should not be called when circuit is open")
			return 0, nil
		})

		if !errors.Is(err, fortifyerrors.ErrCircuitOpen) {
			t.Errorf("error = %v, want %v", err, fortifyerrors.ErrCircuitOpen)
		}
	})

	t.Run("transitions to half-open after timeout", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 2,
			Interval:    time.Second,
			Timeout:     50 * time.Millisecond,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		// Open the circuit
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		if cb.State() != StateOpen {
			t.Fatalf("state = %v, want %v", cb.State(), StateOpen)
		}

		// Wait for timeout
		time.Sleep(60 * time.Millisecond)

		// Next request should attempt execution (half-open)
		executed := false
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			executed = true
			return 42, nil
		})

		if !executed {
			t.Error("function should execute in half-open state")
		}
	})

	t.Run("transitions from half-open to closed on success", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 2,
			Interval:    time.Second,
			Timeout:     50 * time.Millisecond,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		// Open the circuit
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		// Wait for timeout to enter half-open
		time.Sleep(60 * time.Millisecond)

		// Successful requests in half-open
		for i := 0; i < 2; i++ {
			cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
				return 42, nil
			})
		}

		if cb.State() != StateClosed {
			t.Errorf("state after successful half-open requests = %v, want %v", cb.State(), StateClosed)
		}
	})

	t.Run("transitions from half-open to open on failure", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 2,
			Interval:    time.Second,
			Timeout:     50 * time.Millisecond,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		// Open the circuit
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		// Wait for timeout to enter half-open
		time.Sleep(60 * time.Millisecond)

		// Fail in half-open state
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		if cb.State() != StateOpen {
			t.Errorf("state after failed half-open request = %v, want %v", cb.State(), StateOpen)
		}
	})
}

func TestCircuitBreakerExecute(t *testing.T) {
	t.Run("executes function and returns result", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
		})

		result, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want %v", result, 42)
		}
	})

	t.Run("propagates function error", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
		})

		expectedErr := errors.New("test error")
		_, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, expectedErr
		})

		if !errors.Is(err, expectedErr) {
			t.Errorf("error = %v, want %v", err, expectedErr)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
		})

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := cb.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if !errors.Is(err, context.Canceled) {
			t.Errorf("error = %v, want %v", err, context.Canceled)
		}
	})
}

func TestCircuitBreakerReset(t *testing.T) {
	cb := New[int](Config{
		MaxRequests: 1,
		Interval:    time.Second,
		Timeout:     time.Second,
		ReadyToTrip: func(counts Counts) bool {
			return counts.ConsecutiveFailures >= 1
		},
	})

	// Open the circuit
	cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, errors.New("failure")
	})

	if cb.State() != StateOpen {
		t.Fatalf("state = %v, want %v", cb.State(), StateOpen)
	}

	// Reset
	cb.Reset()

	if cb.State() != StateClosed {
		t.Errorf("state after reset = %v, want %v", cb.State(), StateClosed)
	}

	// Should execute successfully after reset
	result, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})

	if err != nil {
		t.Errorf("unexpected error after reset: %v", err)
	}
	if result != 42 {
		t.Errorf("result = %v, want %v", result, 42)
	}
}

func TestCircuitBreakerCallbacks(t *testing.T) {
	t.Run("OnStateChange callback", func(t *testing.T) {
		transitionCh := make(chan struct {
			from State
			to   State
		}, 1)

		cb := New[int](Config{
			MaxRequests: 1,
			Interval:    time.Second,
			Timeout:     50 * time.Millisecond,
			ReadyToTrip: func(counts Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
			OnStateChange: func(from, to State) {
				transitionCh <- struct {
					from State
					to   State
				}{from, to}
			},
		})

		// Trigger state change from Closed to Open
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})

		// Wait for callback with timeout
		select {
		case transition := <-transitionCh:
			if transition.from != StateClosed || transition.to != StateOpen {
				t.Errorf("transition = %v -> %v, want Closed -> Open", transition.from, transition.to)
			}
		case <-time.After(100 * time.Millisecond):
			t.Fatal("timeout waiting for OnStateChange callback")
		}
	})
}

func TestCircuitBreakerCounts(t *testing.T) {
	cb := New[int](Config{
		MaxRequests: 10,
		Interval:    time.Second,
	})

	// Successful requests
	for i := 0; i < 3; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}

	// Failed requests
	for i := 0; i < 2; i++ {
		cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})
	}

	// Verify counts through ReadyToTrip callback
	verified := false
	testCB := New[int](Config{
		MaxRequests: 1,
		Interval:    time.Second,
		ReadyToTrip: func(counts Counts) bool {
			if counts.TotalSuccesses == 3 && counts.TotalFailures == 2 && counts.Requests == 5 {
				verified = true
			}
			return false
		},
	})

	for i := 0; i < 3; i++ {
		testCB.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
	for i := 0; i < 2; i++ {
		testCB.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 0, errors.New("failure")
		})
	}

	if !verified {
		t.Error("counts were not updated correctly")
	}
}
