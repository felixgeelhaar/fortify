package timeout

import (
	"context"
	"testing"
	"time"
)

func BenchmarkTimeoutSuccess(b *testing.B) {
	tm := New[int](Config{
		DefaultTimeout: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = tm.Execute(ctx, 100*time.Millisecond, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

func BenchmarkTimeoutWithShortOperation(b *testing.B) {
	tm := New[int](Config{
		DefaultTimeout: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = tm.Execute(ctx, time.Second, func(ctx context.Context) (int, error) {
			// Very short operation
			return i, nil
		})
	}
}

func BenchmarkTimeoutConcurrent(b *testing.B) {
	tm := New[int](Config{
		DefaultTimeout: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = tm.Execute(ctx, 100*time.Millisecond, func(ctx context.Context) (int, error) {
				return 42, nil
			})
		}
	})
}