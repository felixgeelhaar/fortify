package circuitbreaker

import (
	"context"
	"errors"
	"testing"
	"time"
)

func BenchmarkCircuitBreakerSuccess(b *testing.B) {
	cb := New[int](Config{
		MaxRequests: 100,
		Interval:    time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

func BenchmarkCircuitBreakerFailure(b *testing.B) {
	cb := New[int](Config{
		MaxRequests: 100,
		Interval:    time.Second,
		ReadyToTrip: func(counts Counts) bool {
			return false // Never trip
		},
	})

	ctx := context.Background()
	testErr := errors.New("test error")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
			return 0, testErr
		})
	}
}

func BenchmarkCircuitBreakerOpen(b *testing.B) {
	cb := New[int](Config{
		MaxRequests: 1,
		Interval:    time.Second,
		Timeout:     time.Hour, // Keep circuit open
		ReadyToTrip: func(counts Counts) bool {
			return counts.ConsecutiveFailures >= 1
		},
	})

	// Open the circuit
	ctx := context.Background()
	cb.Execute(ctx, func(ctx context.Context) (int, error) {
		return 0, errors.New("failure")
	})

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

func BenchmarkCircuitBreakerConcurrent(b *testing.B) {
	cb := New[int](Config{
		MaxRequests: 100,
		Interval:    time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = cb.Execute(ctx, func(ctx context.Context) (int, error) {
				return 42, nil
			})
		}
	})
}
