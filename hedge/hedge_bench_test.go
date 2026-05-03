package hedge

import (
	"context"
	"testing"
	"time"
)

// BenchmarkHedgePrimaryWins measures the hot path where the primary returns
// before HedgeDelay elapses (the dominant case in healthy systems).
// This should be cheap: timer setup + one goroutine + one channel send.
func BenchmarkHedgePrimaryWins(b *testing.B) {
	h := New[int](Config{
		MaxAttempts: 3,
		HedgeDelay:  100 * time.Millisecond,
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = h.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

// BenchmarkHedgeNoHedging measures Execute when MaxAttempts=1 (hedging
// effectively disabled). This is the configuration baseline against which
// to compare the cost of enabling hedging.
func BenchmarkHedgeNoHedging(b *testing.B) {
	h := New[int](Config{
		MaxAttempts: 1,
		HedgeDelay:  time.Hour, // never fires anyway with MaxAttempts=1
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = h.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

// BenchmarkHedgePrimaryWinsParallel exercises the goroutine spawn + channel
// machinery under contention.
func BenchmarkHedgePrimaryWinsParallel(b *testing.B) {
	h := New[int](Config{
		MaxAttempts: 3,
		HedgeDelay:  100 * time.Millisecond,
	})
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = h.Execute(ctx, func(ctx context.Context) (int, error) {
				return 42, nil
			})
		}
	})
}
