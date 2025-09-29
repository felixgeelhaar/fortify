package bulkhead_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/felixgeelhaar/fortify/bulkhead"
)

// Example demonstrates basic bulkhead usage to limit concurrent operations.
func Example() {
	// Create bulkhead with max 3 concurrent operations
	bh := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 3,
	})

	result, err := bh.Execute(context.Background(), func(ctx context.Context) (string, error) {
		// Your operation here
		return "success", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", result)
	// Output: Result: success
}

// Example_concurrencyLimit demonstrates the concurrency limiting behavior.
func Example_concurrencyLimit() {
	// Allow max 2 concurrent operations
	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 2,
		MaxQueue:      0, // No queueing
	})

	var wg sync.WaitGroup
	results := make(chan string, 5)

	// Launch 5 concurrent operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			_, err := bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
				time.Sleep(time.Millisecond * 100)
				return id, nil
			})

			if err != nil {
				results <- fmt.Sprintf("Task %d: rejected", id)
			} else {
				results <- fmt.Sprintf("Task %d: completed", id)
			}
		}(i)
	}

	wg.Wait()
	close(results)

	completed := 0
	rejected := 0
	for msg := range results {
		if msg[len(msg)-9:] == "completed" {
			completed++
		} else {
			rejected++
		}
	}

	fmt.Printf("Completed: %d, Rejected: %d\n", completed, rejected)
	// Output: Completed: 2, Rejected: 3
}

// Example_withQueue demonstrates bulkhead with request queueing.
func Example_withQueue() {
	// Max 2 concurrent, queue up to 3 more
	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 2,
		MaxQueue:      3,
		QueueTimeout:  time.Second * 5,
	})

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			result, err := bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
				time.Sleep(time.Millisecond * 50)
				return id, nil
			})

			if err != nil {
				fmt.Printf("Task %d: %v\n", id, err)
			} else {
				fmt.Printf("Task %d: completed with result %d\n", id, result)
			}
		}(i)

		time.Sleep(time.Millisecond * 10) // Stagger requests
	}

	wg.Wait()
	fmt.Println("All tasks processed")
	// Output example will vary due to concurrency
}

// Example_queueTimeout demonstrates queue timeout behavior.
func Example_queueTimeout() {
	// Short queue timeout
	bh := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 1,
		MaxQueue:      2,
		QueueTimeout:  time.Millisecond * 100,
	})

	// First request will execute
	go func() {
		bh.Execute(context.Background(), func(ctx context.Context) (string, error) {
			time.Sleep(time.Millisecond * 500)
			return "long operation", nil
		})
	}()

	time.Sleep(time.Millisecond * 50) // Let first request start

	// This request will timeout while queued
	_, err := bh.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "queued request", nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Output: Error: queue timeout exceeded
}

// Example_rejectionCallback demonstrates handling rejected requests.
func Example_rejectionCallback() {
	rejectedCount := 0

	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 1,
		MaxQueue:      0,
		OnRejected: func() {
			rejectedCount++
			fmt.Println("Request rejected - bulkhead full")
		},
	})

	// First request will execute
	go func() {
		bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
			time.Sleep(time.Millisecond * 100)
			return 1, nil
		})
	}()

	time.Sleep(time.Millisecond * 10) // Let first request start

	// Second request will be rejected
	bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 2, nil
	})

	time.Sleep(time.Millisecond * 150) // Wait for completion

	fmt.Printf("Total rejected: %d\n", rejectedCount)
	// Output:
	// Request rejected - bulkhead full
	// Total rejected: 1
}

// Example_monitoring demonstrates monitoring bulkhead state.
func Example_monitoring() {
	active := 0
	completed := 0

	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 5,
		MaxQueue:      10,
	})

	// Execute some operations
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			active++
			_, err := bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
				time.Sleep(time.Millisecond * 50)
				return 0, nil
			})
			active--
			if err == nil {
				completed++
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Completed operations: %d\n", completed)
	// Output:
	// Completed operations: 3
}

// Example_resourceIsolation demonstrates isolating critical operations.
func Example_resourceIsolation() {
	// Create separate bulkheads for different resource pools
	dbBulkhead := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 10, // Max 10 concurrent DB operations
	})

	apiBulkhead := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 20, // Max 20 concurrent API calls
	})

	// Database operation
	dbResult, err := dbBulkhead.Execute(context.Background(), func(ctx context.Context) (string, error) {
		// Simulate DB query
		return "db result", nil
	})

	// API call
	apiResult, err := apiBulkhead.Execute(context.Background(), func(ctx context.Context) (string, error) {
		// Simulate API call
		return "api result", nil
	})

	fmt.Printf("DB: %s, API: %s, Error: %v\n", dbResult, apiResult, err)
	// Output: DB: db result, API: api result, Error: <nil>
}

// Example_contextCancellation demonstrates respecting context cancellation.
func Example_contextCancellation() {
	bh := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	_, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
		// Check context
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-time.After(time.Second):
			return 42, nil
		}
	})

	fmt.Printf("Error: %v\n", err)
	// Output: Error: context deadline exceeded
}

// Example_databaseConnectionPool demonstrates limiting database connections.
func Example_databaseConnectionPool() {
	// Limit to 20 concurrent database operations
	dbBulkhead := bulkhead.New[[]string](bulkhead.Config{
		MaxConcurrent: 20,
		MaxQueue:      50,
		QueueTimeout:  time.Second * 10,
		OnRejected: func() {
			fmt.Println("Database connection pool exhausted")
		},
	})

	// Execute database query
	results, err := dbBulkhead.Execute(context.Background(), func(ctx context.Context) ([]string, error) {
		// Simulate database query
		// rows, err := db.QueryContext(ctx, "SELECT ...")
		return []string{"row1", "row2"}, nil
	})

	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
		return
	}

	fmt.Printf("Query returned %d rows\n", len(results))
	// Output: Query returned 2 rows
}

// Example_httpClientPool demonstrates limiting concurrent HTTP requests.
func Example_httpClientPool() {
	// Limit to 50 concurrent HTTP requests
	httpBulkhead := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 50,
		MaxQueue:      100,
		QueueTimeout:  time.Second * 5,
	})

	// Make HTTP request
	statusCode, err := httpBulkhead.Execute(context.Background(), func(ctx context.Context) (int, error) {
		// Simulate HTTP request
		// req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		// resp, err := httpClient.Do(req)
		// return resp.StatusCode, err
		return 200, nil
	})

	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}

	fmt.Printf("HTTP Status: %d\n", statusCode)
	// Output: HTTP Status: 200
}

// Example_microserviceIsolation demonstrates isolating microservice calls.
func Example_microserviceIsolation() {
	// Separate bulkheads for different services
	userServiceBulkhead := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 15,
	})

	orderServiceBulkhead := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 10,
	})

	// If order service is slow/failing, it won't affect user service calls
	userResult, _ := userServiceBulkhead.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "user data", nil
	})

	orderResult, _ := orderServiceBulkhead.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "order data", nil
	})

	fmt.Printf("User: %s, Order: %s\n", userResult, orderResult)
	// Output: User: user data, Order: order data
}

// Example_threadPoolPattern demonstrates using bulkhead as a thread pool.
func Example_threadPoolPattern() {
	// Create worker pool with 5 workers
	workerPool := bulkhead.New[int](bulkhead.Config{
		MaxConcurrent: 5,
		MaxQueue:      100,
	})

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		taskID := i

		go func() {
			defer wg.Done()

			result, err := workerPool.Execute(context.Background(), func(ctx context.Context) (int, error) {
				// Simulate work
				time.Sleep(time.Millisecond * 10)
				return taskID * 2, nil
			})

			if err == nil {
				fmt.Printf("Task %d completed: %d\n", taskID, result)
			}
		}()
	}

	wg.Wait()
	fmt.Println("All tasks completed")
	// Output: All tasks completed
}