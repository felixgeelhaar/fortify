package ratelimit

import (
	"context"
	"testing"
	"time"
)

func BenchmarkRateLimiterAllow(b *testing.B) {
	limiter := New(Config{
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
	limiter := New(Config{
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
	limiter := New(Config{
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
	limiter := New(Config{
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
	limiter := New(Config{
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

func BenchmarkTokenBucketRefill(b *testing.B) {
	bucket := newTokenBucket(100, 100, time.Second)

	// Exhaust bucket
	for i := 0; i < 100; i++ {
		bucket.allow()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bucket.refill()
	}
}
