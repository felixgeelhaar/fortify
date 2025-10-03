package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/middleware"
	"github.com/felixgeelhaar/fortify/retry"
	"github.com/felixgeelhaar/fortify/timeout"
)

func TestCompositionExample(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the example
	main()

	// Restore stdout and read output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	//nolint:errcheck // ignoring error in test
	_, _ = io.Copy(&buf, r)

	// Verify composition behavior is demonstrated
	output := buf.String()

	// Should show middleware composition in action
	if output == "" {
		t.Error("expected output from composition example")
	}
}

func TestMiddlewareComposition(t *testing.T) {
	r := retry.New[string](retry.Config{
		MaxAttempts:   2,
		InitialDelay:  time.Millisecond,
		BackoffPolicy: retry.BackoffConstant,
	})

	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Millisecond * 100,
	})

	operation := func(ctx context.Context) (string, error) {
		return "success", nil
	}

	// Build middleware chain using the Chain builder pattern
	chain := middleware.New[string]().
		WithRetry(r).
		WithTimeout(tm, time.Millisecond*100)

	result, err := chain.Execute(context.Background(), operation)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if result != "success" {
		t.Errorf("expected 'success', got %q", result)
	}
}
