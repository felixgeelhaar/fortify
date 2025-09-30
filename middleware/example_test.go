package middleware_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/bulkhead"
	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/middleware"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/retry"
	"github.com/felixgeelhaar/fortify/timeout"
)

// Example demonstrates basic middleware chain composition.
func Example() {
	// Create individual patterns
	cb := circuitbreaker.New[string](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second * 30,
	})

	r := retry.New[string](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 100,
	})

	// Compose them into a chain
	chain := middleware.New[string]().
		WithCircuitBreaker(cb).
		WithRetry(r)

	// Execute through the chain
	result, err := chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: success
}

// Example_fullStack demonstrates a complete resilience stack.
func Example_fullStack() {
	// Create all patterns
	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 10,
		MaxQueue:      20,
	})

	rl := ratelimit.New(ratelimit.Config{
		Rate:     100,
		Burst:    200,
		Interval: time.Second,
	})

	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 5,
	})

	cb := circuitbreaker.New[int](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second * 30,
	})

	r := retry.New[int](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 100,
	})

	// Compose in optimal order:
	// 1. Bulkhead - Limit concurrency first
	// 2. Rate Limit - Check quotas
	// 3. Timeout - Enforce time limits
	// 4. Circuit Breaker - Check service health
	// 5. Retry - Retry on failures
	chain := middleware.New[int]().
		WithBulkhead(bh).
		WithRateLimit(rl, "api-key").
		WithTimeout(tm, time.Second*2).
		WithCircuitBreaker(cb).
		WithRetry(r)

	result, err := chain.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %d\n", result)
	// Output: Result: 42
}

// Example_retryWithCircuitBreaker demonstrates retry with circuit breaker.
func Example_retryWithCircuitBreaker() {
	cb := circuitbreaker.New[string](circuitbreaker.Config{
		MaxRequests: 5,
		Interval:    time.Second * 10,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	})

	r := retry.New[string](retry.Config{
		MaxAttempts:  5,
		InitialDelay: time.Millisecond * 100,
	})

	// Circuit breaker before retry - stops retrying if circuit opens
	chain := middleware.New[string]().
		WithCircuitBreaker(cb).
		WithRetry(r)

	attempt := 0
	chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		attempt++
		return "", errors.New("service failure")
	})

	fmt.Printf("Attempts before circuit opened: %d\n", attempt)
	fmt.Printf("Error: circuit breaker opened\n")
	// Circuit breaker will open and stop retries
}

// Example_timeoutWithRetry demonstrates timeout on each retry attempt.
func Example_timeoutWithRetry() {
	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	r := retry.New[int](retry.Config{
		MaxAttempts:  3,
		InitialDelay: time.Millisecond * 100,
	})

	// Timeout before retry - each attempt gets a fresh timeout
	chain := middleware.New[int]().
		WithTimeout(tm, time.Millisecond*500).
		WithRetry(r)

	attempt := 0
	_, err := chain.Execute(context.Background(), func(ctx context.Context) (int, error) {
		attempt++
		if attempt < 3 {
			time.Sleep(time.Millisecond * 600) // Exceeds timeout
			return 0, errors.New("timeout")
		}
		return 42, nil
	})

	if err != nil {
		fmt.Printf("Failed after %d attempts\n", attempt)
	}
}

// Example_rateLimitWithBulkhead demonstrates rate limiting and concurrency control.
func Example_rateLimitWithBulkhead() {
	rl := ratelimit.New(ratelimit.Config{
		Rate:     10,
		Burst:    20,
		Interval: time.Second,
	})

	bh := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 5,
	})

	// Bulkhead before rate limit - limit concurrency, then check quota
	chain := middleware.New[string]().
		WithBulkhead(bh).
		WithRateLimit(rl, "user-123")

	result, err := chain.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "processed", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: processed
}

// Example_apiClient demonstrates a complete API client protection stack.
func Example_apiClient() {
	// Configure patterns for external API calls
	bh := bulkhead.New[map[string]interface{}](bulkhead.Config{
		MaxConcurrent: 20,
		MaxQueue:      50,
	})

	rl := ratelimit.New(ratelimit.Config{
		Rate:     1000,
		Burst:    1500,
		Interval: time.Hour, // 1000 requests per hour
	})

	tm := timeout.New[map[string]interface{}](timeout.Config{
		DefaultTimeout: time.Second * 30,
	})

	cb := circuitbreaker.New[map[string]interface{}](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second * 60,
		Timeout:     time.Second * 30,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			failureRate := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 10 && failureRate >= 0.5
		},
	})

	r := retry.New[map[string]interface{}](retry.Config{
		MaxAttempts:   3,
		InitialDelay:  time.Millisecond * 500,
		BackoffPolicy: retry.BackoffExponential,
		Jitter:        true,
	})

	// Build API client middleware stack
	apiClient := middleware.New[map[string]interface{}]().
		WithBulkhead(bh).
		WithRateLimit(rl, "api-client").
		WithTimeout(tm, time.Second*10).
		WithCircuitBreaker(cb).
		WithRetry(r)

	// Make API call
	response, err := apiClient.Execute(context.Background(), func(ctx context.Context) (map[string]interface{}, error) {
		// Simulate API call
		// resp, err := http.Get("https://api.example.com/data")
		return map[string]interface{}{"status": "ok"}, nil
	})

	if err != nil {
		fmt.Printf("API call failed: %v\n", err)
		return
	}

	fmt.Printf("API response: %v\n", response["status"])
	// Output: API response: ok
}

// Example_databaseClient demonstrates database operation protection.
func Example_databaseClient() {
	// Limit concurrent database connections
	bh := bulkhead.New[[]string](bulkhead.Config{
		MaxConcurrent: 50,
		MaxQueue:      100,
	})

	// Database operation timeout
	tm := timeout.New[[]string](timeout.Config{
		DefaultTimeout: time.Second * 30,
	})

	// Circuit breaker for database
	cb := circuitbreaker.New[[]string](circuitbreaker.Config{
		MaxRequests: 20,
		Interval:    time.Second * 60,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 5
		},
	})

	// Retry transient database errors
	r := retry.New[[]string](retry.Config{
		MaxAttempts:  2,
		InitialDelay: time.Millisecond * 500,
	})

	dbClient := middleware.New[[]string]().
		WithBulkhead(bh).
		WithTimeout(tm, time.Second*10).
		WithCircuitBreaker(cb).
		WithRetry(r)

	results, err := dbClient.Execute(context.Background(), func(ctx context.Context) ([]string, error) {
		// Simulate database query
		// rows, err := db.QueryContext(ctx, query)
		return []string{"row1", "row2", "row3"}, nil
	})

	if err != nil {
		fmt.Printf("Database query failed: %v\n", err)
		return
	}

	fmt.Printf("Query returned %d rows\n", len(results))
	// Output: Query returned 3 rows
}

// Example_microserviceCall demonstrates protecting service-to-service calls.
func Example_microserviceCall() {
	// Per-service bulkhead for isolation
	bh := bulkhead.New[interface{}](bulkhead.Config{
		MaxConcurrent: 15,
	})

	// Aggressive timeout for microservices
	tm := timeout.New[interface{}](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	// Fast-fail circuit breaker
	cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
		MaxRequests: 5,
		Interval:    time.Second * 30,
		Timeout:     time.Second * 15,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	})

	// Quick retries for transient failures
	r := retry.New[interface{}](retry.Config{
		MaxAttempts:   2,
		InitialDelay:  time.Millisecond * 200,
		BackoffPolicy: retry.BackoffConstant,
	})

	serviceClient := middleware.New[interface{}]().
		WithBulkhead(bh).
		WithTimeout(tm, time.Second*3).
		WithCircuitBreaker(cb).
		WithRetry(r)

	response, err := serviceClient.Execute(context.Background(), func(ctx context.Context) (interface{}, error) {
		// Call downstream service
		return map[string]string{"result": "success"}, nil
	})

	if err != nil {
		fmt.Printf("Service call failed: %v\n", err)
		return
	}

	fmt.Printf("Service response received: %v\n", response != nil)
	// Output: Service response received: true
}

// Example_backgroundJob demonstrates resilience for background jobs.
func Example_backgroundJob() {
	// Higher concurrency for background jobs
	bh := bulkhead.New[bool](bulkhead.Config{
		MaxConcurrent: 50,
	})

	// Longer timeout for batch operations
	tm := timeout.New[bool](timeout.Config{
		DefaultTimeout: time.Minute * 5,
	})

	// More aggressive retries for batch jobs
	r := retry.New[bool](retry.Config{
		MaxAttempts:   5,
		InitialDelay:  time.Second,
		BackoffPolicy: retry.BackoffExponential,
		Multiplier:    2.0,
	})

	// No circuit breaker - jobs should always attempt
	jobRunner := middleware.New[bool]().
		WithBulkhead(bh).
		WithTimeout(tm, time.Minute*2).
		WithRetry(r)

	success, err := jobRunner.Execute(context.Background(), func(ctx context.Context) (bool, error) {
		// Process batch job
		return true, nil
	})

	if err != nil {
		fmt.Printf("Job failed: %v\n", err)
		return
	}

	fmt.Printf("Job completed: %v\n", success)
	// Output: Job completed: true
}
