package bulkhead

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
)

func TestBulkheadExecute(t *testing.T) {
	t.Run("executes function within concurrency limit", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 2,
		})

		result, err := bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want 42", result)
		}
	})

	t.Run("enforces max concurrency limit", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 2,
			MaxQueue:      0, // No queue
		})

		ctx := context.Background()
		executing := atomic.Int32{}
		maxConcurrent := atomic.Int32{}

		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bh.Execute(ctx, func(ctx context.Context) (int, error) {
					current := executing.Add(1)
					defer executing.Add(-1)

					// Track max concurrent executions
					for {
						max := maxConcurrent.Load()
						if current <= max || maxConcurrent.CompareAndSwap(max, current) {
							break
						}
					}

					time.Sleep(50 * time.Millisecond)
					return 42, nil
				})
			}()
		}

		wg.Wait()

		if maxConcurrent.Load() > 2 {
			t.Errorf("max concurrent = %d, want <= 2", maxConcurrent.Load())
		}
	})

	t.Run("rejects when bulkhead full and no queue", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      0,
		})

		ctx := context.Background()
		started := make(chan bool)
		done := make(chan bool)

		// Start first request (fills bulkhead)
		go func() {
			bh.Execute(ctx, func(ctx context.Context) (int, error) {
				started <- true
				<-done
				return 42, nil
			})
		}()

		<-started // Wait for first request to start

		// Second request should be rejected
		_, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
			t.Error("should not execute when bulkhead full")
			return 0, nil
		})

		close(done)

		if !errors.Is(err, fortifyerrors.ErrBulkheadFull) {
			t.Errorf("error = %v, want ErrBulkheadFull", err)
		}
	})

	t.Run("queues requests when bulkhead full", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      2,
		})

		ctx := context.Background()
		results := make(chan int, 3)
		started := make(chan bool, 1)

		// First request (fills bulkhead)
		go func() {
			result, _ := bh.Execute(ctx, func(ctx context.Context) (int, error) {
				started <- true
				time.Sleep(50 * time.Millisecond)
				return 1, nil
			})
			results <- result
		}()

		<-started // Wait for first to start

		// Second and third requests (should queue)
		for i := 2; i <= 3; i++ {
			val := i
			go func() {
				result, _ := bh.Execute(ctx, func(ctx context.Context) (int, error) {
					return val, nil
				})
				results <- result
			}()
		}

		// Collect all results
		collected := make(map[int]bool)
		for i := 0; i < 3; i++ {
			collected[<-results] = true
		}

		// Should have received all 3 unique results
		if len(collected) != 3 {
			t.Errorf("got %d unique results, want 3", len(collected))
		}
	})

	t.Run("rejects when queue full", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      1,
		})

		ctx := context.Background()
		started := make(chan bool)
		release := make(chan bool)

		// Fill bulkhead
		go func() {
			bh.Execute(ctx, func(ctx context.Context) (int, error) {
				started <- true
				<-release
				return 1, nil
			})
		}()

		<-started

		// Fill queue
		go func() {
			bh.Execute(ctx, func(ctx context.Context) (int, error) {
				return 2, nil
			})
		}()

		time.Sleep(20 * time.Millisecond)

		// This should be rejected (queue full)
		_, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 3, nil
		})

		close(release)

		if !errors.Is(err, fortifyerrors.ErrBulkheadFull) {
			t.Errorf("error = %v, want ErrBulkheadFull", err)
		}
	})

	t.Run("respects context cancellation while queued", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      10,
			QueueTimeout:  time.Second,
		})

		started := make(chan bool)
		release := make(chan bool)

		// Fill bulkhead
		go func() {
			bh.Execute(context.Background(), func(ctx context.Context) (int, error) {
				started <- true
				<-release
				return 1, nil
			})
		}()

		<-started

		// Try with cancelled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 2, nil
		})

		close(release)

		if !errors.Is(err, context.Canceled) {
			t.Errorf("error = %v, want context.Canceled", err)
		}
	})

	t.Run("respects queue timeout", func(t *testing.T) {
		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      5,
			QueueTimeout:  50 * time.Millisecond,
		})

		ctx := context.Background()
		started := make(chan bool)

		// Fill bulkhead with long-running operation
		go func() {
			bh.Execute(ctx, func(ctx context.Context) (int, error) {
				started <- true
				time.Sleep(200 * time.Millisecond)
				return 1, nil
			})
		}()

		<-started

		// This should timeout while queued
		_, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 2, nil
		})

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("error = %v, want context.DeadlineExceeded", err)
		}
	})
}

func TestBulkheadCallback(t *testing.T) {
	t.Run("calls OnRejected callback", func(t *testing.T) {
		rejectedCh := make(chan bool, 1)

		bh := New[int](Config{
			MaxConcurrent: 1,
			MaxQueue:      0,
			OnRejected: func() {
				rejectedCh <- true
			},
		})

		ctx := context.Background()
		started := make(chan bool)
		done := make(chan bool)

		// Fill bulkhead
		go func() {
			bh.Execute(ctx, func(ctx context.Context) (int, error) {
				started <- true
				<-done
				return 1, nil
			})
		}()

		<-started

		// This should be rejected
		bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 2, nil
		})

		close(done)

		select {
		case <-rejectedCh:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Error("OnRejected callback not called")
		}
	})
}

func TestBulkheadDefaults(t *testing.T) {
	t.Run("applies default configuration", func(t *testing.T) {
		bh := New[int](Config{})

		ctx := context.Background()

		result, err := bh.Execute(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("result = %v, want 42", result)
		}
	})
}