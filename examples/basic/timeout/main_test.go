package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/v2/timeout"
)

func TestTimeoutExample(t *testing.T) {
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

	// Verify timeout behavior is demonstrated
	output := buf.String()

	// Should show timeout enforcement
	if output == "" {
		t.Error("expected output from timeout example")
	}
}

func TestTimeoutEnforcement(t *testing.T) {
	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Millisecond * 100,
	})

	_, err := tm.Execute(context.Background(), time.Millisecond*50, func(ctx context.Context) (string, error) {
		select {
		case <-time.After(time.Millisecond * 200):
			return "too slow", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	})

	if err == nil {
		t.Error("expected timeout error")
	}

	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
}
