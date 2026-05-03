package adaptive

import (
	"context"
	"errors"
	"testing"

	"pgregory.net/rapid"
)

// TestProperty_LimitWithinBounds asserts that no sequence of Execute calls
// can leave the limit outside [MinLimit, MaxLimit].
func TestProperty_LimitWithinBounds(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		minLim := rapid.IntRange(1, 50).Draw(t, "minLim")
		maxLim := minLim + rapid.IntRange(0, 100).Draw(t, "extra")
		alg := AlgorithmAIMD
		if rapid.Bool().Draw(t, "vegas") {
			alg = AlgorithmVegas
		}

		l := New[int](Config{
			Algorithm:        alg,
			InitialLimit:     rapid.IntRange(minLim, maxLim).Draw(t, "init"),
			MinLimit:         minLim,
			MaxLimit:         maxLim,
			SuccessThreshold: rapid.IntRange(1, 5).Draw(t, "succT"),
			VegasMinSamples:  1,
		}).(*limiter[int])

		ctx := context.Background()
		ops := rapid.SliceOfN(rapid.Bool(), 0, 60).Draw(t, "ops")
		boom := errors.New("boom")
		for _, success := range ops {
			_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
				if success {
					return 0, nil
				}
				return 0, boom
			})

			lim := l.Limit()
			if lim < minLim {
				t.Fatalf("limit %d < MinLimit %d", lim, minLim)
			}
			if lim > maxLim {
				t.Fatalf("limit %d > MaxLimit %d", lim, maxLim)
			}
		}
	})
}

// TestProperty_InFlightNeverExceedsLimitAtObservation asserts that at no
// point does a successful tryAcquire leave inFlight > limit. (Concurrent
// Execute is exercised separately in TestExecute_ConcurrentRespectsLimit.)
func TestProperty_InFlightSettlesToZero(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		l := New[int](Config{
			Algorithm:    AlgorithmAIMD,
			InitialLimit: rapid.IntRange(1, 50).Draw(t, "init"),
			MinLimit:     1,
			MaxLimit:     200,
		}).(*limiter[int])

		ctx := context.Background()
		ops := rapid.IntRange(1, 100).Draw(t, "n")
		for i := 0; i < ops; i++ {
			_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
				return 0, nil
			})
		}
		if got := l.InFlight(); got != 0 {
			t.Fatalf("InFlight after sequential Execute = %d, want 0", got)
		}
	})
}

// TestProperty_PersistentFailureFloorsAtMin asserts that under a long run
// of nothing-but-failures, AIMD floors at MinLimit (and stays there).
func TestProperty_PersistentFailureFloorsAtMin(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		minLim := rapid.IntRange(1, 10).Draw(t, "min")
		l := New[int](Config{
			Algorithm:    AlgorithmAIMD,
			InitialLimit: rapid.IntRange(minLim, minLim+100).Draw(t, "init"),
			MinLimit:     minLim,
			MaxLimit:     500,
		}).(*limiter[int])

		ctx := context.Background()
		boom := errors.New("boom")
		for i := 0; i < 30; i++ {
			_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
				return 0, boom
			})
		}
		if got := l.Limit(); got != minLim {
			t.Fatalf("Limit after persistent failure = %d, want MinLimit=%d", got, minLim)
		}
	})
}
