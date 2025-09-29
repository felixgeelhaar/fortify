package timeout_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/felixgeelhaar/fortify/timeout"
)

// Example demonstrates basic timeout usage.
func Example() {
	// Create timeout manager with 5 second default
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 5,
	})

	// Execute with specific timeout
	result, err := tm.Execute(context.Background(), time.Second, func(ctx context.Context) (string, error) {
		// Fast operation completes before timeout
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: success
}

// Example_timeoutExceeded demonstrates handling timeout errors.
func Example_timeoutExceeded() {
	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 5,
	})

	_, err := tm.Execute(context.Background(), time.Millisecond*100, func(ctx context.Context) (int, error) {
		// Simulate slow operation
		time.Sleep(time.Millisecond * 500)
		return 42, nil
	})

	if err != nil {
		fmt.Printf("Operation timed out: %v\n", err)
	}
	// Output: Operation timed out: context deadline exceeded
}

// Example_defaultTimeout demonstrates using the default timeout.
func Example_defaultTimeout() {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 3,
	})

	// Use default timeout by passing 0
	result, err := tm.Execute(context.Background(), 0, func(ctx context.Context) (string, error) {
		time.Sleep(time.Millisecond * 100)
		return "completed within default timeout", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: completed within default timeout
}

// Example_contextAwareness demonstrates respecting context cancellation.
func Example_contextAwareness() {
	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 5,
	})

	_, err := tm.Execute(context.Background(), time.Second, func(ctx context.Context) (int, error) {
		// Check context periodically
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Operation cancelled")
				return 0, ctx.Err()
			default:
				time.Sleep(time.Millisecond * 150)
			}
		}
		return 42, nil
	})

	fmt.Printf("Error: %v\n", err)
	// Output:
	// Operation cancelled
	// Error: context deadline exceeded
}

// Example_callback demonstrates timeout event callbacks.
func Example_callback() {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 5,
		OnTimeout: func() {
			fmt.Println("Operation timed out")
		},
	})

	_, err := tm.Execute(context.Background(), time.Millisecond*100, func(ctx context.Context) (string, error) {
		time.Sleep(time.Millisecond * 200)
		return "too slow", nil
	})

	if err != nil {
		fmt.Println("Timeout occurred")
	}
	// Output:
	// Operation timed out
	// Timeout occurred
}

// Example_databaseQuery demonstrates timeout with database operations.
func Example_databaseQuery() {
	tm := timeout.New[[]string](timeout.Config{
		DefaultTimeout: time.Second * 30,
	})

	// Execute database query with 5 second timeout
	results, err := tm.Execute(context.Background(), time.Second*5, func(ctx context.Context) ([]string, error) {
		// Simulate database query
		// db.QueryContext(ctx, "SELECT ...")
		time.Sleep(time.Millisecond * 100)
		return []string{"row1", "row2", "row3"}, nil
	})

	if err != nil {
		fmt.Printf("Database query failed: %v\n", err)
		return
	}

	fmt.Printf("Retrieved %d results\n", len(results))
	// Output: Retrieved 3 results
}

// Example_httpClient demonstrates timeout with HTTP requests.
func Example_httpClient() {
	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 30,
		OnTimeout: func() {
			fmt.Println("HTTP request timed out")
		},
	})

	// Execute HTTP request with 2 second timeout
	statusCode, err := tm.Execute(context.Background(), time.Second*2, func(ctx context.Context) (int, error) {
		// Simulate HTTP request
		// req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
		// resp, err := httpClient.Do(req)
		// if err != nil {
		//     return 0, err
		// }
		// return resp.StatusCode, nil
		time.Sleep(time.Millisecond * 100)
		return 200, nil
	})

	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}

	fmt.Printf("HTTP Status: %d\n", statusCode)
	// Output: HTTP Status: 200
}

// Example_microserviceCall demonstrates timeout for service-to-service calls.
func Example_microserviceCall() {
	tm := timeout.New[map[string]interface{}](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	// Call downstream service with 3 second timeout
	response, err := tm.Execute(context.Background(), time.Second*3, func(ctx context.Context) (map[string]interface{}, error) {
		// Simulate service call
		time.Sleep(time.Millisecond * 500)
		return map[string]interface{}{
			"status": "ok",
			"data":   "service response",
		}, nil
	})

	if err != nil {
		fmt.Printf("Service call failed: %v\n", err)
		return
	}

	fmt.Printf("Service status: %v\n", response["status"])
	// Output: Service status: ok
}

// Example_parentContext demonstrates using a parent context with timeout.
func Example_parentContext() {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	// Parent context with 500ms timeout
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	// Execute with 2 second timeout (but parent context will cancel first)
	_, err := tm.Execute(parentCtx, time.Second*2, func(ctx context.Context) (string, error) {
		time.Sleep(time.Second * 1)
		return "result", nil
	})

	// Parent context timeout takes precedence
	fmt.Printf("Error: %v\n", err)
	// Output: Error: context deadline exceeded
}

// Example_gracefulShutdown demonstrates timeout for graceful shutdown operations.
func Example_gracefulShutdown() {
	tm := timeout.New[bool](timeout.Config{
		DefaultTimeout: time.Second * 30,
	})

	// Give shutdown operations 15 seconds to complete
	success, err := tm.Execute(context.Background(), time.Second*15, func(ctx context.Context) (bool, error) {
		// Simulate graceful shutdown tasks
		fmt.Println("Closing connections...")
		time.Sleep(time.Millisecond * 100)

		fmt.Println("Flushing buffers...")
		time.Sleep(time.Millisecond * 100)

		fmt.Println("Saving state...")
		time.Sleep(time.Millisecond * 100)

		return true, nil
	})

	if err != nil {
		fmt.Printf("Shutdown timed out: %v\n", err)
		return
	}

	if success {
		fmt.Println("Graceful shutdown completed")
	}
	// Output:
	// Closing connections...
	// Flushing buffers...
	// Saving state...
	// Graceful shutdown completed
}

// Example_batchProcessing demonstrates timeout for batch operations.
func Example_batchProcessing() {
	tm := timeout.New[int](timeout.Config{
		DefaultTimeout: time.Second * 60,
	})

	// Process batch with 30 second timeout
	processed, err := tm.Execute(context.Background(), time.Second*30, func(ctx context.Context) (int, error) {
		count := 0
		for i := 0; i < 100; i++ {
			// Check if context was cancelled
			select {
			case <-ctx.Done():
				return count, ctx.Err()
			default:
				// Process item
				time.Sleep(time.Millisecond)
				count++
			}
		}
		return count, nil
	})

	if err != nil {
		fmt.Printf("Batch processing interrupted: processed %d items\n", processed)
		return
	}

	fmt.Printf("Batch processing completed: %d items\n", processed)
	// Output: Batch processing completed: 100 items
}

// Example_retryWithTimeout demonstrates combining timeout with retry logic.
func Example_retryWithTimeout() {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 10,
	})

	attempt := 0
	// Each attempt has a 500ms timeout
	result, err := tm.Execute(context.Background(), time.Millisecond*500, func(ctx context.Context) (string, error) {
		attempt++

		// First two attempts fail
		if attempt < 3 {
			return "", errors.New("temporary failure")
		}

		// Third attempt succeeds quickly
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return
	}

	fmt.Printf("Result: %s after %d attempts\n", result, attempt)
	// Output: Result: success after 3 attempts
}

// Example_asyncOperation demonstrates timeout for asynchronous operations.
func Example_asyncOperation() {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 5,
	})

	// Start async operation with 2 second timeout
	result, err := tm.Execute(context.Background(), time.Second*2, func(ctx context.Context) (string, error) {
		// Simulate async operation
		done := make(chan string, 1)

		go func() {
			time.Sleep(time.Millisecond * 300)
			done <- "async result"
		}()

		// Wait for result or timeout
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case res := <-done:
			return res, nil
		}
	})

	if err != nil {
		fmt.Printf("Async operation failed: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: async result
}