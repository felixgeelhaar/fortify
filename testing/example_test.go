package testing_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/retry"
	fortifyTesting "github.com/felixgeelhaar/fortify/testing"
)

// Example demonstrates basic error injection for testing.
func Example() {
	// Create an error injector with 50% failure rate
	injector := fortifyTesting.NewErrorInjector(0.5, errors.New("simulated failure"))

	successCount := 0
	failureCount := 0

	// Simulate 10 calls
	for i := 0; i < 10; i++ {
		err := injector.Execute(func() error {
			return nil // Your actual operation
		})

		if err != nil {
			failureCount++
		} else {
			successCount++
		}
	}

	calls, failures, rate := injector.Stats()
	fmt.Printf("Calls: %d, Failures: %d, Rate: %.2f\n", calls, failures, rate)
	fmt.Printf("Success: %d, Failure: %d\n", successCount, failureCount)
	// Output will vary due to randomness
}

// Example_latencyInjection demonstrates simulating network latency.
func Example_latencyInjection() {
	// Create latency injector with 10-50ms delays
	injector := fortifyTesting.NewLatencyInjector(10*time.Millisecond, 50*time.Millisecond)

	start := time.Now()

	// Add latency to operation
	err := injector.Delay(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Operation took %v\n", elapsed > 10*time.Millisecond)
	// Output: Operation took true
}

// Example_timeoutSimulation demonstrates timeout scenario testing.
func Example_timeoutSimulation() {
	// Create timeout simulator with 100ms timeout, 100% probability
	sim := fortifyTesting.NewTimeoutSimulator(100*time.Millisecond, 1.0)

	ctx, cancel := sim.Context(context.Background())
	defer cancel()

	// Simulate slow operation
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}
	// Output: Operation timed out
}

// Example_flakeyService demonstrates testing with unreliable services.
func Example_flakeyService() {
	// Create flakey service with 30% error rate and 5-15ms latency
	service := fortifyTesting.NewFlakeyService(0.3, 5*time.Millisecond, 15*time.Millisecond)

	var successCount, errorCount int

	// Make 10 calls to the service
	for i := 0; i < 10; i++ {
		err := service.Call(context.Background(), func() error {
			// Your actual operation
			return nil
		})

		if err != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	stats := service.Stats()
	//nolint:errcheck // type assertion in example, error handled by test
	errorRate, _ := stats["error_rate"].(float64)
	//nolint:errcheck // type assertion in example, error handled by test
	avgLatency, _ := stats["avg_latency"].(time.Duration)
	hasErrors := errorRate > 0.0
	hasLatency := avgLatency > 0
	fmt.Printf("Error rate present: %v\n", hasErrors)
	fmt.Printf("Average latency present: %v\n", hasLatency)
	// Output will vary due to randomness
}

// Example_testingCircuitBreaker demonstrates testing circuit breaker with chaos.
func Example_testingCircuitBreaker() {
	// Create error injector that fails 60% of the time
	injector := fortifyTesting.NewErrorInjector(0.6, errors.New("service unavailable"))

	// Create circuit breaker
	cb := circuitbreaker.New[string](circuitbreaker.Config{
		Timeout: 100 * time.Millisecond,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	})

	var closedCount, openCount int

	// Make calls until circuit breaker opens
	for i := 0; i < 20; i++ {
		_, err := cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
			// Simulate service call with injected errors
			if injector.ShouldFail() {
				return "", injector.Error()
			}
			return "success", nil
		})

		state := cb.State()
		switch state {
		case circuitbreaker.StateClosed:
			closedCount++
		case circuitbreaker.StateOpen:
			openCount++
		}

		if err != nil && state == circuitbreaker.StateOpen {
			fmt.Printf("Circuit breaker opened after failures\n")
			break
		}

		time.Sleep(10 * time.Millisecond)
	}

	calls, failures, _ := injector.Stats()
	fmt.Printf("Service calls: %d, Failures: %d\n", calls, failures)
	// Output will vary but should show circuit breaker opening
}

// Example_testingRetryPolicy demonstrates testing retry behavior with chaos.
func Example_testingRetryPolicy() {
	// Create error injector that fails first 2 attempts, then succeeds
	attemptCount := 0
	failureThreshold := 2

	r := retry.New[string](retry.Config{
		MaxAttempts:  5,
		InitialDelay: 10 * time.Millisecond,
		MaxDelay:     100 * time.Millisecond,
		Multiplier:   2.0,
	})

	result, err := r.Do(context.Background(), func(ctx context.Context) (string, error) {
		attemptCount++
		if attemptCount <= failureThreshold {
			return "", errors.New("temporary failure")
		}
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s after %d attempts\n", result, attemptCount)
	// Output: Result: success after 3 attempts
}

// Example_gradualFailure demonstrates increasing failure rate over time.
func Example_gradualFailure() {
	injector := fortifyTesting.NewErrorInjector(0.0, errors.New("degraded service"))

	// Gradually increase failure rate
	for phase := 0; phase < 3; phase++ {
		failureProb := float64(phase) * 0.3 // 0%, 30%, 60%
		injector.SetProbability(failureProb)

		// Make 10 calls at this failure rate
		var failures int
		for i := 0; i < 10; i++ {
			if injector.ShouldFail() {
				failures++
			}
		}

		fmt.Printf("Phase %d (%.0f%% failure): %d failures\n", phase+1, failureProb*100, failures)
	}
	// Output will vary but should show increasing failures
}

// Example_contextCancellation demonstrates testing context cancellation with latency.
func Example_contextCancellation() {
	injector := fortifyTesting.NewLatencyInjector(100*time.Millisecond, 200*time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := injector.Delay(ctx)
	if err == context.DeadlineExceeded {
		fmt.Println("Operation canceled due to timeout")
	}
	// Output: Operation canceled due to timeout
}

// Example_testingWithMetrics demonstrates combining chaos testing with monitoring.
func Example_testingWithMetrics() {
	service := fortifyTesting.NewFlakeyService(0.4, 10*time.Millisecond, 30*time.Millisecond)

	// Simulate load
	for i := 0; i < 50; i++ {
		//nolint:errcheck // intentionally ignoring error in example
		_ = service.Call(context.Background(), func() error {
			return nil
		})
	}

	// Get statistics
	stats := service.Stats()
	//nolint:errcheck // type assertion in example, error handled by test
	errorCalls, _ := stats["error_calls"].(int64)
	//nolint:errcheck // type assertion in example, error handled by test
	errorRate, _ := stats["error_rate"].(float64)
	//nolint:errcheck // type assertion in example, error handled by test
	avgLatency := stats["avg_latency"].(time.Duration)

	fmt.Printf("Total calls: %d\n", errorCalls)
	fmt.Printf("Error rate: %.2f\n", errorRate)
	fmt.Printf("Average latency: %v\n", avgLatency > 0)
	// Output will vary due to randomness
}

// Example_intermittentFailures demonstrates testing with occasional failures.
func Example_intermittentFailures() {
	// Create injector with low failure rate (10%)
	injector := fortifyTesting.NewErrorInjector(0.1, errors.New("intermittent error"))

	cb := circuitbreaker.New[int](circuitbreaker.Config{
		Timeout: 100 * time.Millisecond,
		ReadyToTrip: func(counts circuitbreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 5
		},
	})

	var total, successes, failures int

	// Make many calls - circuit should stay closed with low error rate
	for i := 0; i < 100; i++ {
		total++
		_, err := cb.Execute(context.Background(), func(ctx context.Context) (int, error) {
			if injector.ShouldFail() {
				return 0, injector.Error()
			}
			return i, nil
		})

		if err != nil {
			failures++
		} else {
			successes++
		}
	}

	fmt.Printf("Completed %d calls: %d successes, %d failures\n", total, successes, failures)
	fmt.Printf("Circuit state: %s\n", cb.State())
	// Output will vary but circuit should stay closed
}

// Example_performanceTracking demonstrates tracking benchmark performance.
func Example_performanceTracking() {
	tracker := fortifyTesting.NewPerformanceTracker(".benchmark-results")

	// Add performance baselines
	tracker.AddBaseline(fortifyTesting.PerformanceBaseline{
		Name:        "BenchmarkCircuitBreaker",
		MaxNsPerOp:  1000,
		MaxAllocs:   5,
		MaxBytes:    512,
		Description: "Circuit breaker baseline",
	})

	// Simulate benchmark results
	results := []fortifyTesting.BenchmarkResult{
		{
			Name:        "BenchmarkCircuitBreaker",
			NsPerOp:     950,
			AllocsPerOp: 4,
			BytesPerOp:  480,
			Timestamp:   time.Now(),
		},
	}

	// Check for regressions
	report := tracker.CheckRegressions(results)

	fmt.Printf("Total checks: %d\n", report.TotalChecks)
	fmt.Printf("Passed: %d\n", report.Passed)
	fmt.Printf("Failed: %d\n", report.Failed)
	// Output:
	// Total checks: 1
	// Passed: 1
	// Failed: 0
}

// Example_performanceBaseline demonstrates generating performance baselines.
func Example_performanceBaseline() {
	tracker := fortifyTesting.NewPerformanceTracker(".benchmark-results")

	// Simulate benchmark results
	results := []fortifyTesting.BenchmarkResult{
		{
			Name:        "BenchmarkRetry",
			NsPerOp:     2000,
			AllocsPerOp: 10,
			BytesPerOp:  1024,
			Timestamp:   time.Now(),
		},
	}

	// Generate baseline with 10% safety factor
	tracker.GenerateBaselineFromResults(results, 1.1)

	// Check how many baselines were created
	report := tracker.CheckRegressions(results)
	fmt.Printf("Baselines generated, checks performed: %d\n", report.TotalChecks)
	// Output: Baselines generated, checks performed: 1
}
