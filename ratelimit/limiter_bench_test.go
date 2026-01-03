package ratelimit

import (
	"context"
	"testing"
	"time"
)

func BenchmarkRateLimiterAllow(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000, // Very high rate to avoid blocking
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Allow(ctx, "test-key")
	}
}

func BenchmarkRateLimiterAllowRateLimited(b *testing.B) {
	limiter := New(&Config{
		Rate:     1,
		Burst:    1,
		Interval: time.Hour, // No refill during benchmark
	})

	ctx := context.Background()
	limiter.Allow(ctx, "test-key") // Exhaust bucket

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Allow(ctx, "test-key")
	}
}

func BenchmarkRateLimiterTake(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Take(ctx, "test-key", 10)
	}
}

func BenchmarkRateLimiterMultipleKeys(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	keys := []string{"key1", "key2", "key3", "key4", "key5"}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Allow(ctx, keys[i%len(keys)])
	}
}

func BenchmarkRateLimiterConcurrent(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(ctx, "test-key")
		}
	})
}

func BenchmarkMemoryStoreAtomicUpdate(b *testing.B) {
	store := NewMemoryStore()
	ctx := context.Background()

	// Initialize bucket
	//nolint:errcheck // benchmark setup
	_, _ = store.AtomicUpdate(ctx, "test-key", func(s *BucketState) *BucketState {
		return &BucketState{Tokens: 100, LastRefill: time.Now()}
	})

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		//nolint:errcheck // benchmark
		_, _ = store.AtomicUpdate(ctx, "test-key", func(s *BucketState) *BucketState {
			if s == nil {
				return &BucketState{Tokens: 100, LastRefill: time.Now()}
			}
			return &BucketState{Tokens: s.Tokens - 1, LastRefill: time.Now()}
		})
	}
}

func BenchmarkMemoryStoreConcurrentUpdate(b *testing.B) {
	store := NewMemoryStore()
	ctx := context.Background()

	// Initialize bucket
	//nolint:errcheck // benchmark setup
	_, _ = store.AtomicUpdate(ctx, "test-key", func(s *BucketState) *BucketState {
		return &BucketState{Tokens: 1000000, LastRefill: time.Now()}
	})

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//nolint:errcheck // benchmark
			_, _ = store.AtomicUpdate(ctx, "test-key", func(s *BucketState) *BucketState {
				if s == nil {
					return &BucketState{Tokens: 1000000, LastRefill: time.Now()}
				}
				return &BucketState{Tokens: s.Tokens - 1, LastRefill: time.Now()}
			})
		}
	})
}
