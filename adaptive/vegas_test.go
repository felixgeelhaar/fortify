package adaptive

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestVegas_ShrinksUnderRisingLatency(t *testing.T) {
	l := New[int](Config{
		Algorithm:       AlgorithmVegas,
		InitialLimit:    20,
		MinLimit:        1,
		MaxLimit:        100,
		VegasAlpha:      3,
		VegasBeta:       6,
		VegasMinSamples: 5,
	}).(*limiter[int])

	ctx := context.Background()

	// First batch: low latency establishes baseline.
	for i := 0; i < 10; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Millisecond)
			return 0, nil
		})
	}
	baseline := l.Limit()

	// Second batch: latency 10x baseline. Vegas should shrink.
	for i := 0; i < 50; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(10 * time.Millisecond)
			return 0, nil
		})
	}
	got := l.Limit()
	if got >= baseline {
		t.Fatalf("Vegas did not shrink: baseline=%d after-saturation=%d", baseline, got)
	}
	if got < 1 {
		t.Fatalf("limit floored below MinLimit: %d", got)
	}
}

func TestVegas_GrowsUnderStableLatency(t *testing.T) {
	l := New[int](Config{
		Algorithm:       AlgorithmVegas,
		InitialLimit:    5,
		MinLimit:        1,
		MaxLimit:        50,
		VegasAlpha:      3,
		VegasBeta:       6,
		VegasMinSamples: 3,
	}).(*limiter[int])

	ctx := context.Background()
	for i := 0; i < 100; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Millisecond)
			return 0, nil
		})
	}
	if got := l.Limit(); got <= 5 {
		t.Fatalf("Vegas did not grow under stable latency: %d", got)
	}
}

func TestVegas_FailureStillHalves(t *testing.T) {
	l := New[int](Config{
		Algorithm:    AlgorithmVegas,
		InitialLimit: 20,
		MinLimit:     1,
		MaxLimit:     100,
	}).(*limiter[int])

	ctx := context.Background()
	wantErr := errors.New("downstream")
	_, err := l.Execute(ctx, func(ctx context.Context) (int, error) {
		return 0, wantErr
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want wantErr", err)
	}
	if got := l.Limit(); got != 10 {
		t.Fatalf("Limit after Vegas failure = %d, want 10 (halved)", got)
	}
}

func TestVegas_RespectsMinSamples(t *testing.T) {
	l := New[int](Config{
		Algorithm:       AlgorithmVegas,
		InitialLimit:    5,
		MinLimit:        1,
		MaxLimit:        50,
		VegasMinSamples: 100, // very high → no adjustments under this test
	}).(*limiter[int])

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Millisecond)
			return 0, nil
		})
	}
	if got := l.Limit(); got != 5 {
		t.Fatalf("Limit changed before MinSamples reached: %d (want 5)", got)
	}
}

func TestVegas_ConfigDefaultsApplied(t *testing.T) {
	l := New[int](Config{Algorithm: AlgorithmVegas}).(*limiter[int])
	if l.config.VegasAlpha != defaultVegasAlpha {
		t.Errorf("VegasAlpha default = %d, want %d", l.config.VegasAlpha, defaultVegasAlpha)
	}
	if l.config.VegasBeta != defaultVegasBeta {
		t.Errorf("VegasBeta default = %d, want %d", l.config.VegasBeta, defaultVegasBeta)
	}
	if l.config.VegasMinSamples != 10 {
		t.Errorf("VegasMinSamples default = %d, want 10", l.config.VegasMinSamples)
	}
}

func TestVegas_BetaBelowAlphaCorrected(t *testing.T) {
	l := New[int](Config{
		Algorithm:  AlgorithmVegas,
		VegasAlpha: 5,
		VegasBeta:  2, // invalid; beta must be > alpha
	}).(*limiter[int])
	if l.config.VegasBeta <= l.config.VegasAlpha {
		t.Fatalf("VegasBeta (%d) must be > VegasAlpha (%d)", l.config.VegasBeta, l.config.VegasAlpha)
	}
}

func TestVegas_LowAllocsOnHotPath(t *testing.T) {
	l := New[int](Config{Algorithm: AlgorithmVegas, InitialLimit: 5, MinLimit: 1, MaxLimit: 5})
	ctx := context.Background()
	var calls atomic.Int64

	allocs := testing.AllocsPerRun(100, func() {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			calls.Add(1)
			return 0, nil
		})
	})
	// Vegas records timestamps; the closure may force one alloc through
	// the generic boundary. We set a low ceiling, not strict zero.
	if allocs > 1 {
		t.Fatalf("Execute hot path allocs/op = %.1f, want <= 1", allocs)
	}
}
