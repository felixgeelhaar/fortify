package fallback

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestFallback_PrimarySuccess(t *testing.T) {
	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			t.Fatal("Fallback should not be called when primary succeeds")
			return 0, nil
		},
	})

	result, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != 42 {
		t.Errorf("Expected result 42, got %d", result)
	}
}

func TestFallback_FallbackSuccess(t *testing.T) {
	primaryErr := errors.New("primary failed")

	fb := New[string](Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			if err != primaryErr {
				t.Errorf("Expected error %v, got %v", primaryErr, err)
			}
			return "fallback value", nil
		},
	})

	result, err := fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", primaryErr
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != "fallback value" {
		t.Errorf("Expected 'fallback value', got %q", result)
	}
}

func TestFallback_FallbackFailure(t *testing.T) {
	primaryErr := errors.New("primary failed")
	fallbackErr := errors.New("fallback failed")

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 0, fallbackErr
		},
	})

	result, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, primaryErr
	})

	// Should return original error, not fallback error
	if err != primaryErr {
		t.Errorf("Expected original error %v, got %v", primaryErr, err)
	}
	if result != 0 {
		t.Errorf("Expected zero value, got %d", result)
	}
}

func TestFallback_ShouldFallback(t *testing.T) {
	retryableErr := errors.New("retryable error")
	nonRetryableErr := errors.New("non-retryable error")

	var fallbackCalled bool

	fb := New[string](Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			fallbackCalled = true
			return "fallback", nil
		},
		ShouldFallback: func(err error) bool {
			return err == retryableErr
		},
	})

	// Test retryable error - should trigger fallback
	fallbackCalled = false
	result, err := fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", retryableErr
	})
	if err != nil {
		t.Errorf("Expected no error for retryable error, got %v", err)
	}
	if !fallbackCalled {
		t.Error("Expected fallback to be called for retryable error")
	}
	if result != "fallback" {
		t.Errorf("Expected 'fallback', got %q", result)
	}

	// Test non-retryable error - should not trigger fallback
	fallbackCalled = false
	result, err = fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", nonRetryableErr
	})
	if err != nonRetryableErr {
		t.Errorf("Expected non-retryable error, got %v", err)
	}
	if fallbackCalled {
		t.Error("Fallback should not be called for non-retryable error")
	}
}

func TestFallback_OnFallbackCallback(t *testing.T) {
	primaryErr := errors.New("primary failed")
	var callbackErr error

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 100, nil
		},
		OnFallback: func(err error) {
			callbackErr = err
		},
	})

	_, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, primaryErr
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if callbackErr != primaryErr {
		t.Errorf("Expected callback error %v, got %v", primaryErr, callbackErr)
	}
}

func TestFallback_OnSuccessCallback(t *testing.T) {
	var successCalled bool

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 0, nil
		},
		OnSuccess: func() {
			successCalled = true
		},
	})

	_, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !successCalled {
		t.Error("Expected OnSuccess callback to be called")
	}
}

func TestFallback_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 0, nil
		},
	})

	_, err := fb.Execute(ctx, func(ctx context.Context) (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 42, nil
	})

	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
}

func TestFallback_NilFallback(t *testing.T) {
	fb := New[int](Config[int]{
		Fallback: nil,
	})

	primaryErr := errors.New("primary failed")
	_, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 0, primaryErr
	})

	if err != primaryErr {
		t.Errorf("Expected primary error %v, got %v", primaryErr, err)
	}
}

func TestFallback_ConcurrentExecution(t *testing.T) {
	var primaryCount, fallbackCount atomic.Int32

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			fallbackCount.Add(1)
			return 100, nil
		},
	})

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			_, err := fb.Execute(context.Background(), func(ctx context.Context) (int, error) {
				primaryCount.Add(1)
				if id%2 == 0 {
					return id, nil
				}
				return 0, errors.New("failed")
			})

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}(i)
	}

	wg.Wait()

	if primaryCount.Load() != numGoroutines {
		t.Errorf("Expected %d primary calls, got %d", numGoroutines, primaryCount.Load())
	}

	// Half should fail and trigger fallback
	expectedFallback := int32(numGoroutines / 2)
	if fallbackCount.Load() != expectedFallback {
		t.Errorf("Expected %d fallback calls, got %d", expectedFallback, fallbackCount.Load())
	}
}

func TestFallback_Logging(t *testing.T) {
	// Create a buffer to capture log output
	var logBuffer []string
	logger := slog.New(slog.NewTextHandler(&testLogWriter{lines: &logBuffer}, nil))

	fb := New[string](Config[string]{
		Fallback: func(ctx context.Context, err error) (string, error) {
			return "fallback", nil
		},
		Logger: logger,
	})

	// Test primary success
	logBuffer = nil
	_, _ = fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "success", nil
	})

	if len(logBuffer) == 0 {
		t.Error("Expected log entry for primary success")
	}

	// Test fallback triggered
	logBuffer = nil
	_, _ = fb.Execute(context.Background(), func(ctx context.Context) (string, error) {
		return "", errors.New("failed")
	})

	if len(logBuffer) < 2 {
		t.Errorf("Expected at least 2 log entries (fallback triggered + success), got %d", len(logBuffer))
	}
}

// testLogWriter is a simple writer for capturing log output in tests
type testLogWriter struct {
	lines *[]string
}

func (w *testLogWriter) Write(p []byte) (n int, err error) {
	*w.lines = append(*w.lines, string(p))
	return len(p), nil
}

func BenchmarkFallback_PrimarySuccess(b *testing.B) {
	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 0, nil
		},
	})

	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = fb.Execute(ctx, func(ctx context.Context) (int, error) {
			return i, nil
		})
	}
}

func BenchmarkFallback_FallbackTriggered(b *testing.B) {
	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 100, nil
		},
	})

	ctx := context.Background()
	err := errors.New("error")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = fb.Execute(ctx, func(ctx context.Context) (int, error) {
			return 0, err
		})
	}
}

func BenchmarkFallback_WithCallbacks(b *testing.B) {
	var count int

	fb := New[int](Config[int]{
		Fallback: func(ctx context.Context, err error) (int, error) {
			return 100, nil
		},
		OnFallback: func(err error) {
			count++
		},
		OnSuccess: func() {
			count++
		},
	})

	ctx := context.Background()
	err := errors.New("error")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			_, _ = fb.Execute(ctx, func(ctx context.Context) (int, error) {
				return i, nil
			})
		} else {
			_, _ = fb.Execute(ctx, func(ctx context.Context) (int, error) {
				return 0, err
			})
		}
	}
}