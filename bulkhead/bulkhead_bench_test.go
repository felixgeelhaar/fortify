package bulkhead

import (
	"context"
	"testing"
)

func BenchmarkBulkheadSuccess(b *testing.B) {
	bh := New[int](Config{
		MaxConcurrent: 100,
		MaxQueue:      0,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})
	}
}

func BenchmarkBulkheadWithQueue(b *testing.B) {
	bh := New[int](Config{
		MaxConcurrent: 10,
		MaxQueue:      100,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return i, nil
		})
	}
}

func BenchmarkBulkheadConcurrent(b *testing.B) {
	bh := New[int](Config{
		MaxConcurrent: 100,
		MaxQueue:      0,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = bh.Execute(ctx, func(ctx context.Context) (int, error) {
				return 42, nil
			})
		}
	})
}