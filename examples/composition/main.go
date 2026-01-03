// Package main demonstrates middleware chain composition.
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/felixgeelhaar/fortify/v2/bulkhead"
	"github.com/felixgeelhaar/fortify/v2/circuitbreaker"
	"github.com/felixgeelhaar/fortify/v2/middleware"
	"github.com/felixgeelhaar/fortify/v2/ratelimit"
	"github.com/felixgeelhaar/fortify/v2/retry"
	"github.com/felixgeelhaar/fortify/v2/timeout"
)

// APIResponse represents a typical API response
type APIResponse struct {
	Data      string
	Timestamp time.Time
	RequestID string
}

// ExternalAPIClient simulates calls to an external API
type ExternalAPIClient struct {
	failureRate float64
	latency     time.Duration
}

func (c *ExternalAPIClient) Call(ctx context.Context, requestID string) (*APIResponse, error) {
	// Simulate network latency
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(c.latency):
	}

	// Simulate random failures
	//nolint:gosec // G404: using math/rand intentionally for non-security simulation
	if rand.Float64() < c.failureRate {
		return nil, errors.New("external API error")
	}

	return &APIResponse{
		Data:      "Success response from external API",
		Timestamp: time.Now(),
		RequestID: requestID,
	}, nil
}

func main() {
	// Create an unreliable external API client
	// 30% failure rate, 100-500ms latency
	apiClient := &ExternalAPIClient{
		failureRate: 0.3,
		//nolint:gosec // G404: using math/rand intentionally for simulation latency
		latency: time.Millisecond * time.Duration(100+rand.Intn(400)),
	}

	// Build a comprehensive resilience chain
	chain := buildResilientChain()

	// Simulate multiple API calls
	fmt.Println("Making 20 API calls through resilience chain...")
	successCount := 0
	failureCount := 0

	for i := 0; i < 20; i++ {
		requestID := fmt.Sprintf("req-%d", i+1)
		fmt.Printf("\n[%s] Starting...\n", requestID)

		// Execute through the chain
		response, err := chain.Execute(context.Background(), func(ctx context.Context) (*APIResponse, error) {
			return apiClient.Call(ctx, requestID)
		})

		if err != nil {
			failureCount++
			fmt.Printf("[%s] ✗ Failed: %v\n", requestID, err)
		} else {
			successCount++
			fmt.Printf("[%s] ✓ Success: %s (at %s)\n",
				requestID,
				response.Data,
				response.Timestamp.Format("15:04:05.000"))
		}

		// Small delay between requests
		time.Sleep(time.Millisecond * 200)
	}

	// Print summary
	fmt.Println("\n=== Summary ===")
	fmt.Printf("Total requests: 20\n")
	fmt.Printf("Successful: %d (%.1f%%)\n", successCount, float64(successCount)/20*100)
	fmt.Printf("Failed: %d (%.1f%%)\n", failureCount, float64(failureCount)/20*100)
}

func buildResilientChain() *middleware.Chain[*APIResponse] {
	// 1. Circuit Breaker - Prevent cascading failures
	cb := circuitbreaker.New[*APIResponse](circuitbreaker.Config{
		MaxRequests: 5,
		Interval:    time.Second * 10,
		Timeout:     time.Second * 5, // Time in open state before half-open
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			// Open after 3 consecutive failures or 50% failure rate with at least 5 requests
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.ConsecutiveFailures >= 3 ||
				(counts.Requests >= 5 && failureRatio >= 0.5)
		},
		OnStateChange: func(from, to circuitbreaker.State) {
			fmt.Printf("\n🔄 Circuit Breaker: %s -> %s\n", from, to)
		},
	})

	// 2. Retry - Handle transient failures
	r := retry.New[*APIResponse](retry.Config{
		MaxAttempts:   3,
		InitialDelay:  time.Millisecond * 100,
		MaxDelay:      time.Second * 2,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
		IsRetryable: func(err error) bool {
			// Only retry on specific errors
			return err != nil && err.Error() == "external API error"
		},
		OnRetry: func(attempt int, err error) {
			fmt.Printf("   ⟳ Retry attempt %d: %v\n", attempt, err)
		},
	})

	// 3. Rate Limiter - Prevent overwhelming the API
	rl := ratelimit.New(&ratelimit.Config{
		Rate:     5,
		Burst:    10,
		Interval: time.Second,
	})

	// 4. Timeout - Enforce response time SLA
	tm := timeout.New[*APIResponse](timeout.Config{
		DefaultTimeout: time.Second * 5,
		OnTimeout: func() {
			fmt.Println("   ⏱ Timeout occurred")
		},
	})

	// 5. Bulkhead - Limit concurrent requests
	bh := bulkhead.New[*APIResponse](bulkhead.Config{
		MaxConcurrent: 3,
		MaxQueue:      5,
		QueueTimeout:  time.Second * 2,
		OnRejected: func() {
			fmt.Println("   ⛔ Rejected: bulkhead full")
		},
	})

	// Compose the chain
	// Order matters: Bulkhead -> Rate Limit -> Timeout -> Circuit Breaker -> Retry
	return middleware.New[*APIResponse]().
		WithBulkhead(bh).               // First: Limit concurrency
		WithRateLimit(rl, "api-key").   // Second: Rate limiting
		WithTimeout(tm, time.Second*2). // Third: Enforce timeout
		WithCircuitBreaker(cb).         // Fourth: Circuit breaker
		WithRetry(r)                    // Last: Retry on failures
}

// Alternative composition for different scenarios
// nolint:unused // example showing alternative pattern
func buildFastFailChain() *middleware.Chain[*APIResponse] {
	// For scenarios where fast failure is preferred over retries
	cb := circuitbreaker.New[*APIResponse](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second * 5,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 2
		},
	})

	tm := timeout.New[*APIResponse](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	bh := bulkhead.New[*APIResponse](bulkhead.Config{
		MaxConcurrent: 5,
		MaxQueue:      0, // No queuing - immediate rejection
	})

	return middleware.New[*APIResponse]().
		WithBulkhead(bh).
		WithTimeout(tm, time.Millisecond*500). // Aggressive timeout
		WithCircuitBreaker(cb)
	// No retry - fail fast
}

// Background job composition
// nolint:unused // example showing alternative pattern
func buildBackgroundJobChain() *middleware.Chain[*APIResponse] {
	// For background jobs that can tolerate longer delays and more retries
	r := retry.New[*APIResponse](retry.Config{
		MaxAttempts:   10, // More attempts for background jobs
		InitialDelay:  time.Second,
		MaxDelay:      time.Minute,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
	})

	tm := timeout.New[*APIResponse](timeout.Config{
		DefaultTimeout: time.Minute, // Longer timeout
	})

	return middleware.New[*APIResponse]().
		WithTimeout(tm, time.Second*30).
		WithRetry(r)
	// No circuit breaker or rate limit for background jobs
}
