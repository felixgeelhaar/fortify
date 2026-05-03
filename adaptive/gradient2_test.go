package adaptive

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestGradient2_ShrinksUnderRTTInflation(t *testing.T) {
	l := New[int](Config{
		Algorithm:          AlgorithmGradient2,
		InitialLimit:       50,
		MinLimit:           1,
		MaxLimit:           500,
		GradientMinSamples: 5,
	}).(*limiter[int])

	ctx := context.Background()

	// Establish baseline at low RTT.
	for i := 0; i < 20; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Millisecond)
			return 0, nil
		})
	}
	baseline := l.Limit()

	// Inflate RTT 10x; limit should shrink (gradient < 1).
	for i := 0; i < 30; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(10 * time.Millisecond)
			return 0, nil
		})
	}
	got := l.Limit()
	if got >= baseline {
		t.Fatalf("Gradient2 did not shrink under RTT inflation: baseline=%d after=%d", baseline, got)
	}
}

func TestGradient2_GrowsAtBaselineRTT(t *testing.T) {
	l := New[int](Config{
		Algorithm:          AlgorithmGradient2,
		InitialLimit:       5,
		MinLimit:           1,
		MaxLimit:           100,
		GradientMinSamples: 3,
	}).(*limiter[int])

	ctx := context.Background()
	for i := 0; i < 30; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(1 * time.Millisecond)
			return 0, nil
		})
	}
	if got := l.Limit(); got <= 5 {
		t.Fatalf("Gradient2 did not grow under stable RTT: %d", got)
	}
}

func TestGradient2_FailureStillHalves(t *testing.T) {
	l := New[int](Config{
		Algorithm:    AlgorithmGradient2,
		InitialLimit: 20,
		MinLimit:     1,
		MaxLimit:     100,
	}).(*limiter[int])

	wantErr := errors.New("downstream")
	_, _ = l.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, wantErr
	})
	if got := l.Limit(); got != 10 {
		t.Fatalf("Limit after Gradient2 failure = %d, want 10 (halved)", got)
	}
}

func TestGradient2_DefaultsApplied(t *testing.T) {
	l := New[int](Config{Algorithm: AlgorithmGradient2}).(*limiter[int])
	if l.config.GradientMinSamples != 10 {
		t.Errorf("GradientMinSamples default = %d, want 10", l.config.GradientMinSamples)
	}
	if l.config.GradientSmoothing != 0.2 {
		t.Errorf("GradientSmoothing default = %v, want 0.2", l.config.GradientSmoothing)
	}
}

func TestGradient2_FloorAtMinLimit(t *testing.T) {
	l := New[int](Config{
		Algorithm:          AlgorithmGradient2,
		InitialLimit:       4,
		MinLimit:           3,
		MaxLimit:           100,
		GradientMinSamples: 1,
	}).(*limiter[int])

	ctx := context.Background()
	// Establish min RTT.
	_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
		time.Sleep(time.Millisecond)
		return 0, nil
	})

	// Inflate RTT massively; gradient → 0.5 (clamp).
	for i := 0; i < 50; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			time.Sleep(50 * time.Millisecond)
			return 0, nil
		})
	}
	if got := l.Limit(); got < 3 {
		t.Fatalf("Limit dropped below MinLimit: %d", got)
	}
}
