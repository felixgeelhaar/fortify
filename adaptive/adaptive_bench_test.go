package adaptive

import (
	"context"
	"errors"
	"testing"
)

// BenchmarkAIMDSuccess measures the AIMD success hot path: tryAcquire CAS
// + onSuccess counter increment. Steady-state Closed-equivalent.
func BenchmarkAIMDSuccess(b *testing.B) {
	l := New[int](Config{
		Algorithm:        AlgorithmAIMD,
		InitialLimit:     1_000_000,
		MinLimit:         1,
		MaxLimit:         1_000_000,
		SuccessThreshold: 1_000_000, // never trigger increase
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
}

// BenchmarkAIMDFailure measures the failure path: halving + counter reset.
func BenchmarkAIMDFailure(b *testing.B) {
	l := New[int](Config{
		Algorithm:    AlgorithmAIMD,
		InitialLimit: 1_000_000,
		MinLimit:     1_000_000, // floor at InitialLimit so halving is a no-op CAS
		MaxLimit:     1_000_000,
	})
	ctx := context.Background()
	failErr := errors.New("fail")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			return 0, failErr
		})
	}
}

// BenchmarkVegasSuccess measures the Vegas success path including time.Now,
// EMA update, and queue-depth calculation.
func BenchmarkVegasSuccess(b *testing.B) {
	l := New[int](Config{
		Algorithm:       AlgorithmVegas,
		InitialLimit:    1_000_000,
		MinLimit:        1,
		MaxLimit:        1_000_000,
		VegasMinSamples: 1_000_000_000, // never trigger adjustment
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
			return 0, nil
		})
	}
}

// BenchmarkAIMDSuccessParallel measures lock-free admission under contention.
// Should scale with cores since the hot path is pure CAS.
func BenchmarkAIMDSuccessParallel(b *testing.B) {
	l := New[int](Config{
		Algorithm:        AlgorithmAIMD,
		InitialLimit:     1_000_000,
		MinLimit:         1,
		MaxLimit:         1_000_000,
		SuccessThreshold: 1_000_000,
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = l.Execute(ctx, func(ctx context.Context) (int, error) {
				return 0, nil
			})
		}
	})
}
