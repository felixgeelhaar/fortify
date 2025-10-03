package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/ratelimit"
)

func TestRateLimitExample(t *testing.T) {
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

	// Verify rate limiting behavior is demonstrated
	output := buf.String()

	// Should show rate limiting in action
	if output == "" {
		t.Error("expected output from rate limit example")
	}
}

func TestRateLimitEnforcement(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{
		Rate:     2,
		Burst:    2,
		Interval: time.Second,
	})

	allowed := 0
	for i := 0; i < 5; i++ {
		if rl.Allow(context.Background(), "test") {
			allowed++
		}
	}

	// Should allow only burst amount immediately
	if allowed > 2 {
		t.Errorf("expected max 2 allowed immediately, got %d", allowed)
	}
}
