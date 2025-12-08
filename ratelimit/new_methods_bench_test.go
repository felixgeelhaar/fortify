package ratelimit

import (
	"context"
	"errors"
	"testing"
	"time"
)

// BenchmarkExecute measures the overhead of Execute vs direct Allow + operation.
func BenchmarkExecute(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = limiter.Execute(ctx, "test-key", operation)
	}
}

// BenchmarkExecuteRateLimited measures Execute when rate limited.
func BenchmarkExecuteRateLimited(b *testing.B) {
	limiter := New(&Config{
		Rate:     1,
		Burst:    1,
		Interval: time.Hour,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	// Exhaust bucket
	limiter.Allow(ctx, "test-key")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = limiter.Execute(ctx, "test-key", operation)
	}
}

// BenchmarkExecuteWithError measures Execute when operation returns error.
func BenchmarkExecuteWithError(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return errors.New("test error") }

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = limiter.Execute(ctx, "test-key", operation)
	}
}

// BenchmarkExecuteVsAllowDirect compares Execute overhead vs direct Allow + call.
func BenchmarkExecuteVsAllowDirect(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	b.Run("Execute", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = limiter.Execute(ctx, "test-key", operation)
		}
	})

	b.Run("AllowDirect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if limiter.Allow(ctx, "test-key") {
				_ = operation()
			}
		}
	})
}

// BenchmarkExecuteN measures the overhead of ExecuteN.
func BenchmarkExecuteN(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = limiter.ExecuteN(ctx, "test-key", 10, operation)
	}
}

// BenchmarkExecuteNVsTakeDirect compares ExecuteN overhead vs direct Take + call.
func BenchmarkExecuteNVsTakeDirect(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	b.Run("ExecuteN", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = limiter.ExecuteN(ctx, "test-key", 10, operation)
		}
	})

	b.Run("TakeDirect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if limiter.Take(ctx, "test-key", 10) {
				_ = operation()
			}
		}
	})
}

// BenchmarkBucketCount measures BucketCount performance with varying bucket counts.
func BenchmarkBucketCount(b *testing.B) {
	ctx := context.Background()

	benchmarks := []struct {
		name    string
		buckets int
	}{
		{"Empty", 0},
		{"10Buckets", 10},
		{"100Buckets", 100},
		{"1000Buckets", 1000},
		{"10000Buckets", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			limiter := New(&Config{
				Rate:     1000000,
				Burst:    1000000,
				Interval: time.Second,
			})

			// Create buckets
			for i := 0; i < bm.buckets; i++ {
				limiter.Allow(ctx, keyForInt(i))
			}

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = limiter.BucketCount()
			}
		})
	}
}

// BenchmarkReset measures Reset performance with varying bucket counts.
func BenchmarkReset(b *testing.B) {
	ctx := context.Background()

	benchmarks := []struct {
		name    string
		buckets int
	}{
		{"10Buckets", 10},
		{"100Buckets", 100},
		{"1000Buckets", 1000},
		{"10000Buckets", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			limiter := New(&Config{
				Rate:     1000000,
				Burst:    1000000,
				Interval: time.Second,
			})

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				// Create buckets
				for j := 0; j < bm.buckets; j++ {
					limiter.Allow(ctx, keyForInt(j))
				}
				b.StartTimer()

				// Reset all buckets
				_ = limiter.Reset(ctx)
			}
		})
	}
}

// BenchmarkResetConcurrent measures Reset with concurrent operations.
func BenchmarkResetConcurrent(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()

	// Create initial buckets
	for i := 0; i < 1000; i++ {
		limiter.Allow(ctx, keyForInt(i))
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = limiter.Reset(ctx)
		}
	})
}

// BenchmarkConcurrentExecute measures concurrent Execute performance.
func BenchmarkConcurrentExecute(b *testing.B) {
	limiter := New(&Config{
		Rate:     1000000,
		Burst:    1000000,
		Interval: time.Second,
	})

	ctx := context.Background()
	operation := func() error { return nil }

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = limiter.Execute(ctx, "test-key", operation)
		}
	})
}

// keyForInt is a helper function to generate keys.
func keyForInt(i int) string {
	// Pre-allocate to avoid string concatenation overhead
	const digits = "0123456789"
	if i < 10 {
		return string(digits[i])
	}
	// For larger numbers, use Sprintf which is reasonably fast
	return string(rune('k')) + string(rune('0'+i/100)) + string(rune('0'+(i/10)%10)) + string(rune('0'+i%10))
}
