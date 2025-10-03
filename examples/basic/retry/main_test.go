package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/retry"
)

func TestRetryExample(t *testing.T) {
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

	// Verify retry behavior is demonstrated
	output := buf.String()

	// Should show retry attempts
	if output == "" {
		t.Error("expected output from retry example")
	}
}

func TestRetryAttempts(t *testing.T) {
	attempts := 0

	r := retry.New[int](retry.Config{
		MaxAttempts:   3,
		InitialDelay:  time.Millisecond,
		BackoffPolicy: retry.BackoffConstant,
	})

	_, err := r.Do(context.Background(), func(ctx context.Context) (int, error) {
		attempts++
		if attempts < 3 {
			return 0, errors.New("temporary failure")
		}
		return 42, nil
	})

	if err != nil {
		t.Fatalf("expected success after retries, got error: %v", err)
	}

	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}
