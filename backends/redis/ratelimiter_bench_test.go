package redis

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"

	"github.com/felixgeelhaar/fortify/ratelimit"
)

// BenchmarkRedisAllow measures performance of Redis-backed Allow().
func BenchmarkRedisAllow(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	limiter, err := New(Config{
		Client:   client,
		Rate:     1000000, // High limit to avoid blocking
		Burst:    1000000,
		Interval: time.Second,
	})
	if err != nil {
		b.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Allow(ctx, "benchmark-key")
	}
}

// BenchmarkInMemoryAllow measures performance of in-memory Allow() for comparison.
func BenchmarkInMemoryAllow(b *testing.B) {
	limiter := ratelimit.New(ratelimit.Config{
		Rate:     1000000, // High limit to avoid blocking
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Allow(ctx, "benchmark-key")
	}
}

// BenchmarkRedisTake measures performance of Redis-backed Take().
func BenchmarkRedisTake(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	limiter, err := New(Config{
		Client:   client,
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})
	if err != nil {
		b.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Take(ctx, "benchmark-key", 1)
	}
}

// BenchmarkInMemoryTake measures performance of in-memory Take() for comparison.
func BenchmarkInMemoryTake(b *testing.B) {
	limiter := ratelimit.New(ratelimit.Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		limiter.Take(ctx, "benchmark-key", 1)
	}
}

// BenchmarkRedisAllowConcurrent measures concurrent performance.
func BenchmarkRedisAllowConcurrent(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	limiter, err := New(Config{
		Client:   client,
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})
	if err != nil {
		b.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(ctx, "benchmark-key")
		}
	})
}

// BenchmarkInMemoryAllowConcurrent measures concurrent performance of in-memory.
func BenchmarkInMemoryAllowConcurrent(b *testing.B) {
	limiter := ratelimit.New(ratelimit.Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(ctx, "benchmark-key")
		}
	})
}

// BenchmarkRedisMultipleKeys measures performance with different keys.
func BenchmarkRedisMultipleKeys(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	limiter, err := New(Config{
		Client:   client,
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})
	if err != nil {
		b.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	// Simulate different users
	keys := []string{"user1", "user2", "user3", "user4", "user5"}

	for i := 0; i < b.N; i++ {
		key := keys[i%len(keys)]
		limiter.Allow(ctx, key)
	}
}

// BenchmarkInMemoryMultipleKeys measures performance with different keys.
func BenchmarkInMemoryMultipleKeys(b *testing.B) {
	limiter := ratelimit.New(ratelimit.Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	keys := []string{"user1", "user2", "user3", "user4", "user5"}

	for i := 0; i < b.N; i++ {
		key := keys[i%len(keys)]
		limiter.Allow(ctx, key)
	}
}

// BenchmarkRedisScriptExecution measures raw Lua script execution.
func BenchmarkRedisScriptExecution(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	limiter, err := New(Config{
		Client:   client,
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})
	if err != nil {
		b.Fatalf("failed to create limiter: %v", err)
	}

	rl := limiter.(*redisRateLimiter)
	ctx := context.Background()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		rl.executeScript(ctx, "fortify:ratelimit:bench", 1)
	}
}

// BenchmarkRedisConnectionPool measures connection pool overhead.
func BenchmarkRedisConnectionPool(b *testing.B) {
	mr := miniredis.RunT(b)
	defer mr.Close()

	// Test different pool sizes
	poolSizes := []int{5, 10, 20, 50}

	for _, poolSize := range poolSizes {
		b.Run("poolsize_"+string(rune(poolSize+'0')), func(b *testing.B) {
			client := redis.NewClient(&redis.Options{
				Addr:     mr.Addr(),
				PoolSize: poolSize,
			})

			limiter, err := New(Config{
				Client:   client,
				Rate:     1000000,
				Burst:    1000000,
				Interval: time.Second,
			})
			if err != nil {
				b.Fatalf("failed to create limiter: %v", err)
			}

			ctx := context.Background()

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				limiter.Allow(ctx, "benchmark-key")
			}
		})
	}
}
