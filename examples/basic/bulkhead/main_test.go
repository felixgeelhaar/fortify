package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/bulkhead"
)

func TestBulkheadExample(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the example
	main()

	// Restore stdout and read output
	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	//nolint:errcheck // ignoring error in test
	_, _ = io.Copy(&buf, r)

	// Verify bulkhead behavior is demonstrated
	output := buf.String()

	// Should show concurrent execution limits
	if output == "" {
		t.Error("expected output from bulkhead example")
	}
}

func TestBulkheadConcurrencyLimit(t *testing.T) {
	bh := bulkhead.New[string](bulkhead.Config{
		MaxConcurrent: 2,
	})
	defer func() { _ = bh.Close() }()

	var concurrent int32
	var mu sync.Mutex
	var maxConcurrent int32

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//nolint:errcheck // ignoring error in test
			_, _ = bh.Execute(context.Background(), func(ctx context.Context) (string, error) {
				mu.Lock()
				concurrent++
				if concurrent > maxConcurrent {
					maxConcurrent = concurrent
				}
				mu.Unlock()

				time.Sleep(time.Millisecond * 10)

				mu.Lock()
				concurrent--
				mu.Unlock()

				return "done", nil
			})
		}()
	}

	wg.Wait()

	if maxConcurrent > 2 {
		t.Errorf("expected max concurrent <= 2, got %d", maxConcurrent)
	}
}
