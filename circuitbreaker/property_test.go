package circuitbreaker

import (
	"context"
	"errors"
	"testing"
	"time"

	"pgregory.net/rapid"
)

// TestProperty_StateAlwaysValid asserts that no sequence of Execute calls
// can leave the breaker in an unrecognized state.
func TestProperty_StateAlwaysValid(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		cb := New[int](Config{
			MaxRequests: uint32(rapid.IntRange(1, 10).Draw(t, "maxRequests")),
			Interval:    time.Duration(rapid.IntRange(0, 1000).Draw(t, "intervalMs")) * time.Millisecond,
			Timeout:     time.Duration(rapid.IntRange(1, 1000).Draw(t, "timeoutMs")) * time.Millisecond,
			ReadyToTrip: func(c Counts) bool {
				return c.ConsecutiveFailures >= uint32(rapid.IntRange(1, 5).Draw(t, "trip"))
			},
		})
		defer cb.Close()

		ctx := context.Background()
		ops := rapid.SliceOfN(rapid.Bool(), 0, 50).Draw(t, "ops")

		boomErr := errors.New("boom")
		for _, success := range ops {
			_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
				if success {
					return 0, nil
				}
				return 0, boomErr
			})
		}

		state := cb.State()
		if state != StateClosed && state != StateOpen && state != StateHalfOpen {
			t.Fatalf("unrecognized state: %v", state)
		}
	})
}

// TestProperty_CountsCoherent asserts the basic accounting invariant on
// counts in Closed state. After any sequence of operations, while the
// breaker has been in Closed long enough not to reset, totals must match
// individual counters.
func TestProperty_CountsCoherent(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Use Interval=0 so counts never reset under us.
		cb := New[int](Config{
			Interval: 0,
			ReadyToTrip: func(c Counts) bool {
				// Effectively never trip; we want to observe Closed counts.
				return c.ConsecutiveFailures >= 100000
			},
		}).(*circuitBreaker[int])
		defer cb.Close()

		ctx := context.Background()
		ops := rapid.SliceOfN(rapid.Bool(), 0, 30).Draw(t, "ops")

		boomErr := errors.New("boom")
		var successes, failures uint32
		for _, success := range ops {
			_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
				if success {
					return 0, nil
				}
				return 0, boomErr
			})
			if success {
				successes++
			} else {
				failures++
			}
		}

		cb.mu.RLock()
		got := cb.counts
		cb.mu.RUnlock()

		if got.TotalSuccesses != successes {
			t.Fatalf("TotalSuccesses = %d, want %d", got.TotalSuccesses, successes)
		}
		if got.TotalFailures != failures {
			t.Fatalf("TotalFailures = %d, want %d", got.TotalFailures, failures)
		}
		if got.Requests != successes+failures {
			t.Fatalf("Requests = %d, want %d", got.Requests, successes+failures)
		}
	})
}

// TestProperty_GenerationMonotonic asserts the generation counter never
// goes backwards across any sequence of Execute or Reset calls.
func TestProperty_GenerationMonotonic(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		cb := New[int](Config{
			MaxRequests: 2,
			Interval:    50 * time.Millisecond,
			Timeout:     50 * time.Millisecond,
			ReadyToTrip: func(c Counts) bool { return c.ConsecutiveFailures >= 2 },
		}).(*circuitBreaker[int])
		defer cb.Close()

		ctx := context.Background()
		actions := rapid.SliceOfN(rapid.IntRange(0, 3), 0, 40).Draw(t, "actions")

		boomErr := errors.New("boom")
		var prevGen uint64
		for _, action := range actions {
			switch action {
			case 0:
				_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) { return 0, nil })
			case 1:
				_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) { return 0, boomErr })
			case 2:
				cb.Reset()
			case 3:
				time.Sleep(5 * time.Millisecond) // let timers progress
			}

			cb.mu.RLock()
			gen := cb.generation
			cb.mu.RUnlock()
			if gen < prevGen {
				t.Fatalf("generation went backwards: %d -> %d", prevGen, gen)
			}
			prevGen = gen
		}
	})
}

// TestProperty_OpenStateRejectsUntilTimeout asserts that once Open, all
// requests reject (with ErrCircuitOpen wrap) until at least Timeout has
// elapsed. After Timeout, the breaker may admit trial requests in HalfOpen.
func TestProperty_OpenStateRejectsUntilTimeout(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		timeoutMs := rapid.IntRange(20, 100).Draw(t, "timeoutMs")
		cb := New[int](Config{
			MaxRequests: 1,
			Timeout:     time.Duration(timeoutMs) * time.Millisecond,
			ReadyToTrip: func(c Counts) bool { return c.ConsecutiveFailures >= 1 },
		})
		defer cb.Close()

		ctx := context.Background()
		boomErr := errors.New("boom")

		// Trip the breaker.
		_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) { return 0, boomErr })
		if cb.State() != StateOpen {
			t.Fatalf("breaker not Open after trip; state=%v", cb.State())
		}

		// Within the open window (well below Timeout), all requests must reject.
		halfTimeout := time.Duration(timeoutMs/2) * time.Millisecond
		deadline := time.Now().Add(halfTimeout)
		for time.Now().Before(deadline) {
			_, err := cb.Execute(ctx, func(ctx context.Context) (int, error) {
				t.Fatal("operation ran while breaker open")
				return 0, nil
			})
			if err == nil {
				t.Fatal("expected error while open")
			}
		}
	})
}
