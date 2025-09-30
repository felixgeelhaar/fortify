package ratelimit

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRateLimiterAllow(t *testing.T) {
	t.Run("allows requests within rate limit", func(t *testing.T) {
		limiter := New(Config{
			Rate:     5,
			Burst:    5,
			Interval: time.Second,
		})

		ctx := context.Background()
		for i := 0; i < 5; i++ {
			if !limiter.Allow(ctx, "test-key") {
				t.Errorf("request %d should be allowed", i+1)
			}
		}
	})

	t.Run("blocks requests exceeding burst", func(t *testing.T) {
		limiter := New(Config{
			Rate:     5,
			Burst:    3,
			Interval: time.Second,
		})

		ctx := context.Background()
		// First 3 should succeed (burst limit)
		for i := 0; i < 3; i++ {
			if !limiter.Allow(ctx, "test-key") {
				t.Errorf("request %d should be allowed (within burst)", i+1)
			}
		}

		// 4th request should be blocked
		if limiter.Allow(ctx, "test-key") {
			t.Error("request should be blocked (exceeds burst)")
		}
	})

	t.Run("separate keys have independent limits", func(t *testing.T) {
		limiter := New(Config{
			Rate:     2,
			Burst:    2,
			Interval: time.Second,
		})

		ctx := context.Background()

		// Exhaust key1
		limiter.Allow(ctx, "key1")
		limiter.Allow(ctx, "key1")

		// key2 should still have full quota
		if !limiter.Allow(ctx, "key2") {
			t.Error("key2 should be allowed (independent quota)")
		}
		if !limiter.Allow(ctx, "key2") {
			t.Error("key2 should be allowed (independent quota)")
		}
	})

	t.Run("refills tokens over time", func(t *testing.T) {
		limiter := New(Config{
			Rate:     10,
			Burst:    2,
			Interval: 100 * time.Millisecond,
		})

		ctx := context.Background()

		// Exhaust burst
		limiter.Allow(ctx, "test-key")
		limiter.Allow(ctx, "test-key")

		// Should be blocked
		if limiter.Allow(ctx, "test-key") {
			t.Error("should be blocked initially")
		}

		// Wait for refill (10 tokens per 100ms = 1 token per 10ms)
		time.Sleep(15 * time.Millisecond)

		// Should have ~1.5 tokens refilled
		if !limiter.Allow(ctx, "test-key") {
			t.Error("should be allowed after refill")
		}
	})
}

func TestRateLimiterWait(t *testing.T) {
	t.Run("waits for token availability", func(t *testing.T) {
		limiter := New(Config{
			Rate:     10,
			Burst:    1,
			Interval: 100 * time.Millisecond,
		})

		ctx := context.Background()

		// Exhaust burst
		limiter.Allow(ctx, "test-key")

		// Wait should block until token available
		start := time.Now()
		err := limiter.Wait(ctx, "test-key")
		duration := time.Since(start)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Should have waited ~10ms for 1 token refill
		if duration < 5*time.Millisecond {
			t.Errorf("should have waited for refill, duration: %v", duration)
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		limiter := New(Config{
			Rate:     1,
			Burst:    1,
			Interval: time.Second,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		// Exhaust burst
		limiter.Allow(ctx, "test-key")

		// Wait should fail due to context timeout
		err := limiter.Wait(ctx, "test-key")
		if err == nil {
			t.Error("expected context deadline exceeded error")
		}
		if ctx.Err() == nil {
			t.Error("context should be cancelled")
		}
	})
}

func TestRateLimiterTake(t *testing.T) {
	t.Run("takes multiple tokens at once", func(t *testing.T) {
		limiter := New(Config{
			Rate:     10,
			Burst:    10,
			Interval: time.Second,
		})

		ctx := context.Background()

		// Take 5 tokens
		if !limiter.Take(ctx, "test-key", 5) {
			t.Error("should allow taking 5 tokens")
		}

		// Take 5 more tokens
		if !limiter.Take(ctx, "test-key", 5) {
			t.Error("should allow taking 5 more tokens")
		}

		// Should be exhausted
		if limiter.Take(ctx, "test-key", 1) {
			t.Error("should be exhausted")
		}
	})

	t.Run("rejects when insufficient tokens", func(t *testing.T) {
		limiter := New(Config{
			Rate:     10,
			Burst:    5,
			Interval: time.Second,
		})

		ctx := context.Background()

		// Try to take more than burst
		if limiter.Take(ctx, "test-key", 10) {
			t.Error("should reject request exceeding burst")
		}

		// Burst should still be available
		if !limiter.Take(ctx, "test-key", 5) {
			t.Error("burst should still be available after rejected request")
		}
	})
}

func TestRateLimiterKeyFunc(t *testing.T) {
	type contextKey string
	const userIDKey contextKey = "user_id"

	t.Run("uses custom key function", func(t *testing.T) {
		limiter := New(Config{
			Rate:     2,
			Burst:    2,
			Interval: time.Second,
			KeyFunc: func(ctx context.Context) string {
				userID := ctx.Value(userIDKey)
				if userID == nil {
					return "anonymous"
				}
				//nolint:errcheck // type assertion safe here
				return userID.(string)
			},
		})

		// Use context with user_id
		ctx1 := context.WithValue(context.Background(), userIDKey, "user1")
		ctx2 := context.WithValue(context.Background(), userIDKey, "user2")

		// Exhaust user1's quota
		limiter.Allow(ctx1, "")
		limiter.Allow(ctx1, "")

		// user2 should have full quota
		if !limiter.Allow(ctx2, "") {
			t.Error("user2 should have independent quota")
		}
	})
}

func TestRateLimiterConcurrent(t *testing.T) {
	t.Run("handles concurrent requests safely", func(t *testing.T) {
		limiter := New(Config{
			Rate:     100,
			Burst:    100,
			Interval: time.Second,
		})

		ctx := context.Background()
		allowed := atomic.Int32{}
		denied := atomic.Int32{}

		var wg sync.WaitGroup
		for i := 0; i < 200; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if limiter.Allow(ctx, "test-key") {
					allowed.Add(1)
				} else {
					denied.Add(1)
				}
			}()
		}

		wg.Wait()

		// Should allow exactly burst amount
		if int(allowed.Load()) != 100 {
			t.Errorf("allowed = %d, want 100", allowed.Load())
		}
		if int(denied.Load()) != 100 {
			t.Errorf("denied = %d, want 100", denied.Load())
		}
	})
}

func TestRateLimiterCallbacks(t *testing.T) {
	t.Run("calls OnLimit callback", func(t *testing.T) {
		limitedKeys := make(chan string, 10)

		limiter := New(Config{
			Rate:     1,
			Burst:    1,
			Interval: time.Second,
			OnLimit: func(key string) {
				limitedKeys <- key
			},
		})

		ctx := context.Background()

		// Exhaust quota
		limiter.Allow(ctx, "test-key")

		// This should trigger OnLimit
		limiter.Allow(ctx, "test-key")

		select {
		case key := <-limitedKeys:
			if key != "test-key" {
				t.Errorf("OnLimit called with key = %v, want test-key", key)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("OnLimit callback not called")
		}
	})
}

func TestRateLimiterDefaults(t *testing.T) {
	t.Run("applies default configuration", func(t *testing.T) {
		limiter := New(Config{})

		ctx := context.Background()

		// Should have some default rate
		allowed := 0
		for i := 0; i < 1000; i++ {
			if limiter.Allow(ctx, "test-key") {
				allowed++
			}
		}

		if allowed == 0 {
			t.Error("default config should allow some requests")
		}
		if allowed == 1000 {
			t.Error("default config should have some limit")
		}
	})
}
