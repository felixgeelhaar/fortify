package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func BenchmarkRetrySuccess(b *testing.B) {
	r := New[int](Config{
		MaxAttempts:  3,
		InitialDelay: 10 * time.Millisecond,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = r.Do(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

func BenchmarkRetryFailure(b *testing.B) {
	r := New[int](Config{
		MaxAttempts:  3,
		InitialDelay: 1 * time.Millisecond, // Small delay for benchmarking
	})

	ctx := context.Background()
	testErr := errors.New("test error")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = r.Do(ctx, func(ctx context.Context) (int, error) {
			return 0, testErr
		})
	}
}

func BenchmarkRetryWithJitter(b *testing.B) {
	r := New[int](Config{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Millisecond,
		Jitter:        true,
		BackoffPolicy: BackoffExponential,
	})

	ctx := context.Background()
	testErr := errors.New("test error")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = r.Do(ctx, func(ctx context.Context) (int, error) {
			return 0, testErr
		})
	}
}

func BenchmarkBackoffCalculation(b *testing.B) {
	b.Run("exponential", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			calculateBackoff(BackoffExponential, 5, 100*time.Millisecond, 5*time.Second, 2.0, false)
		}
	})

	b.Run("linear", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			calculateBackoff(BackoffLinear, 5, 100*time.Millisecond, 5*time.Second, 2.0, false)
		}
	})

	b.Run("constant", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			calculateBackoff(BackoffConstant, 5, 100*time.Millisecond, 5*time.Second, 2.0, false)
		}
	})

	b.Run("with_jitter", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			calculateBackoff(BackoffExponential, 5, 100*time.Millisecond, 5*time.Second, 2.0, true)
		}
	})
}