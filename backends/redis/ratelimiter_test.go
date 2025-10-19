package redis

import (
	"context"
	"log/slog"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// setupTestRedis creates an in-memory Redis server and client for testing.
func setupTestRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	t.Helper()

	// Start in-memory Redis server
	mr := miniredis.RunT(t)

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	// Verify connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect to miniredis: %v", err)
	}

	return mr, client
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid configuration",
			config: Config{
				Client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
				Rate:   100,
				Burst:  200,
			},
			wantErr: false,
		},
		{
			name: "missing client",
			config: Config{
				Rate:  100,
				Burst: 200,
			},
			wantErr: true,
		},
		{
			name: "invalid rate",
			config: Config{
				Client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
				Rate:   0,
				Burst:  200,
			},
			wantErr: true,
		},
		{
			name: "invalid burst",
			config: Config{
				Client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
				Rate:   100,
				Burst:  0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     10,
		Burst:    10,
		Interval: time.Second,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Should allow first 10 requests (burst)
	for i := 0; i < 10; i++ {
		if !limiter.Allow(ctx, "test-key") {
			t.Errorf("request %d should be allowed", i)
		}
	}

	// 11th request should be denied (bucket empty)
	if limiter.Allow(ctx, "test-key") {
		t.Error("request 11 should be denied")
	}
}

func TestRateLimiter_Wait(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     10,
		Burst:    2,
		Interval: 100 * time.Millisecond, // Fast refill for testing
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Consume burst
	for i := 0; i < 2; i++ {
		if !limiter.Allow(ctx, "test-key") {
			t.Fatalf("burst request %d should be allowed", i)
		}
	}

	// Wait should block until token available
	start := time.Now()
	if err := limiter.Wait(ctx, "test-key"); err != nil {
		t.Fatalf("Wait() failed: %v", err)
	}
	elapsed := time.Since(start)

	// Should have waited at least some time for refill
	if elapsed < 10*time.Millisecond {
		t.Errorf("Wait() returned too quickly: %v", elapsed)
	}
}

func TestRateLimiter_WaitWithCancellation(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     1,
		Burst:    1,
		Interval: time.Hour, // Very slow refill
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	// Consume the only token
	ctx := context.Background()
	if !limiter.Allow(ctx, "test-key") {
		t.Fatal("first request should be allowed")
	}

	// Create context with short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// Wait should respect context cancellation
	err = limiter.Wait(ctx, "test-key")
	if err == nil {
		t.Error("Wait() should return error on context cancellation")
	}
	if err != context.DeadlineExceeded {
		t.Errorf("Wait() error = %v, want %v", err, context.DeadlineExceeded)
	}
}

func TestRateLimiter_Take(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     10,
		Burst:    10,
		Interval: time.Second,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Take 5 tokens
	if !limiter.Take(ctx, "test-key", 5) {
		t.Error("should be able to take 5 tokens")
	}

	// Take another 5 tokens (total = 10, burst limit)
	if !limiter.Take(ctx, "test-key", 5) {
		t.Error("should be able to take another 5 tokens")
	}

	// Attempt to take 1 more token (should fail)
	if limiter.Take(ctx, "test-key", 1) {
		t.Error("should not be able to take token when bucket is empty")
	}

	// Taking 0 or negative tokens should fail
	if limiter.Take(ctx, "test-key", 0) {
		t.Error("should not allow taking 0 tokens")
	}
	if limiter.Take(ctx, "test-key", -1) {
		t.Error("should not allow taking negative tokens")
	}
}

func TestRateLimiter_MultipleKeys(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     5,
		Burst:    5,
		Interval: time.Second,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Each key should have independent bucket
	for i := 0; i < 5; i++ {
		if !limiter.Allow(ctx, "key1") {
			t.Errorf("key1 request %d should be allowed", i)
		}
		if !limiter.Allow(ctx, "key2") {
			t.Errorf("key2 request %d should be allowed", i)
		}
	}

	// Both buckets should now be empty
	if limiter.Allow(ctx, "key1") {
		t.Error("key1 bucket should be empty")
	}
	if limiter.Allow(ctx, "key2") {
		t.Error("key2 bucket should be empty")
	}
}

func TestRateLimiter_Refill(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     10,
		Burst:    5,
		Interval: 100 * time.Millisecond, // Fast refill for testing
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Consume burst
	for i := 0; i < 5; i++ {
		if !limiter.Allow(ctx, "test-key") {
			t.Fatalf("burst request %d should be allowed", i)
		}
	}

	// Bucket should be empty
	if limiter.Allow(ctx, "test-key") {
		t.Error("bucket should be empty")
	}

	// Wait for refill (10 tokens per 100ms = 1 token per 10ms)
	time.Sleep(150 * time.Millisecond)

	// Should be able to take at least 1 token (actually ~15 tokens refilled)
	if !limiter.Allow(ctx, "test-key") {
		t.Error("should have tokens after refill")
	}
}

func TestRateLimiter_Concurrent(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:   client,
		Rate:     100,
		Burst:    100,
		Interval: time.Second,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()
	const goroutines = 10
	const requestsPerGoroutine = 10

	var wg sync.WaitGroup
	allowedCount := make([]int, goroutines)

	// Launch concurrent goroutines
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				if limiter.Allow(ctx, "concurrent-key") {
					allowedCount[idx]++
				}
			}
		}(i)
	}

	wg.Wait()

	// Count total allowed requests
	total := 0
	for _, count := range allowedCount {
		total += count
	}

	// Should allow exactly burst limit (100)
	if total != 100 {
		t.Errorf("allowed %d requests, want 100 (burst limit)", total)
	}
}

func TestRateLimiter_KeyFunc(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	type contextKey string
	const userIDKey contextKey = "user_id"

	limiter, err := New(Config{
		Client:   client,
		Rate:     10,
		Burst:    10,
		Interval: time.Second,
		KeyFunc: func(ctx context.Context) string {
			if userID, ok := ctx.Value(userIDKey).(string); ok {
				return userID
			}
			return "default"
		},
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	// Create contexts with different user IDs
	ctx1 := context.WithValue(context.Background(), userIDKey, "user1")
	ctx2 := context.WithValue(context.Background(), userIDKey, "user2")

	// Each user should have independent bucket
	for i := 0; i < 10; i++ {
		if !limiter.Allow(ctx1, "") {
			t.Errorf("user1 request %d should be allowed", i)
		}
		if !limiter.Allow(ctx2, "") {
			t.Errorf("user2 request %d should be allowed", i)
		}
	}

	// Both buckets should now be empty
	if limiter.Allow(ctx1, "") {
		t.Error("user1 bucket should be empty")
	}
	if limiter.Allow(ctx2, "") {
		t.Error("user2 bucket should be empty")
	}
}

func TestRateLimiter_Callbacks(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	var limitedKeys []string
	var mu sync.Mutex

	limiter, err := New(Config{
		Client:   client,
		Rate:     1,
		Burst:    1,
		Interval: time.Second,
		OnLimit: func(key string) {
			mu.Lock()
			limitedKeys = append(limitedKeys, key)
			mu.Unlock()
		},
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// First request allowed
	if !limiter.Allow(ctx, "test-key") {
		t.Error("first request should be allowed")
	}

	// Second request denied, should trigger callback
	if limiter.Allow(ctx, "test-key") {
		t.Error("second request should be denied")
	}

	// Give callback time to execute
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	if len(limitedKeys) != 1 || limitedKeys[0] != "test-key" {
		t.Errorf("OnLimit callback not called correctly: got %v", limitedKeys)
	}
	mu.Unlock()
}

func TestRateLimiter_Logger(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	// Create a simple logger (logs to test output)
	logger := slog.New(slog.NewTextHandler(testWriter{t}, nil))

	limiter, err := New(Config{
		Client:   client,
		Rate:     1,
		Burst:    1,
		Interval: time.Second,
		Logger:   logger,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Consume token
	limiter.Allow(ctx, "test-key")

	// This should be logged
	limiter.Allow(ctx, "test-key")
}

func TestRateLimiter_FallbackOnError(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:          client,
		Rate:            10,
		Burst:           10,
		Interval:        time.Second,
		FallbackOnError: true,
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Close Redis to simulate failure
	mr.Close()

	// With FallbackOnError=true, should allow requests on Redis failure
	if !limiter.Allow(ctx, "test-key") {
		t.Error("should allow request when Redis is down and FallbackOnError=true")
	}
}

func TestRateLimiter_DenyOnError(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:          client,
		Rate:            10,
		Burst:           10,
		Interval:        time.Second,
		FallbackOnError: false, // Deny on error
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()

	// Close Redis to simulate failure
	mr.Close()

	// With FallbackOnError=false, should deny requests on Redis failure
	if limiter.Allow(ctx, "test-key") {
		t.Error("should deny request when Redis is down and FallbackOnError=false")
	}
}

func TestRateLimiter_BucketTTL(t *testing.T) {
	mr, client := setupTestRedis(t)
	defer mr.Close()

	limiter, err := New(Config{
		Client:    client,
		Rate:      10,
		Burst:     10,
		Interval:  time.Second,
		BucketTTL: 1 * time.Second, // Short TTL for testing
	})
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	ctx := context.Background()
	key := "ttl-test-key"
	redisKey := "fortify:ratelimit:" + key

	// Make a request to create the bucket
	if !limiter.Allow(ctx, key) {
		t.Error("first request should be allowed")
	}

	// Check that key exists in Redis
	exists := client.Exists(ctx, redisKey).Val()
	if exists == 0 {
		t.Error("bucket key should exist in Redis")
	}

	// Fast-forward time in miniredis
	mr.FastForward(2 * time.Second)

	// Key should have expired
	exists = client.Exists(ctx, redisKey).Val()
	if exists != 0 {
		t.Error("bucket key should have expired after TTL")
	}
}

// testWriter wraps testing.T for slog output.
type testWriter struct {
	t *testing.T
}

func (tw testWriter) Write(p []byte) (n int, err error) {
	tw.t.Log(string(p))
	return len(p), nil
}
