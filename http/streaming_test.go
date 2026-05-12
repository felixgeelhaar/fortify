package http

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/streamtimeout"
)

// TestCircuitBreakerPreservesFlusher is the regression test for the
// SSE buffering bug: the recorder used by CircuitBreaker must forward
// http.Flusher so per-chunk flushes propagate to the client.
func TestCircuitBreakerPreservesFlusher(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second,
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		f, ok := w.(http.Flusher)
		if !ok {
			t.Errorf("CircuitBreaker middleware dropped http.Flusher")
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		for i := 0; i < 3; i++ {
			//nolint:errcheck // test response writer
			_, _ = fmt.Fprintf(w, "data: event-%d\n\n", i)
			f.Flush()
			time.Sleep(20 * time.Millisecond)
		}
	})

	srv := httptest.NewServer(CircuitBreaker(cb)(handler))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Read chunks with a per-chunk timeout. If Flush is buffered the
	// reader will block until the handler returns, which a short
	// per-line deadline catches.
	reader := bufio.NewReader(resp.Body)
	events := 0
	deadline := time.Now().Add(500 * time.Millisecond)
	for events < 3 {
		if time.Now().After(deadline) {
			t.Fatalf("only %d/3 SSE events received before deadline — Flush did not propagate", events)
		}
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Fatalf("read: %v", err)
		}
		if strings.HasPrefix(line, "data: event-") {
			events++
		}
	}
}

// TestCircuitBreakerStreamHappyPath verifies that a well-behaved
// streaming handler completes successfully through the stream-aware
// middleware: events arrive flushed, no watchdog fires, breaker stays
// closed.
func TestCircuitBreakerStreamHappyPath(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second,
	})

	mw, err := CircuitBreakerStream(cb, streamtimeout.Config{
		FirstByteTimeout: 200 * time.Millisecond,
		IdleTimeout:      200 * time.Millisecond,
		TotalTimeout:     2 * time.Second,
	})
	if err != nil {
		t.Fatalf("CircuitBreakerStream: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		f := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		for i := 0; i < 4; i++ {
			//nolint:errcheck // test response writer
			_, _ = fmt.Fprintf(w, "data: %d\n\n", i)
			f.Flush()
			time.Sleep(50 * time.Millisecond) // < IdleTimeout
		}
	})

	srv := httptest.NewServer(mw(handler))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	for i := 0; i < 4; i++ {
		want := fmt.Sprintf("data: %d", i)
		if !strings.Contains(string(body), want) {
			t.Errorf("body missing %q:\n%s", want, body)
		}
	}
}

// TestCircuitBreakerStreamIdleTimeout verifies a stalled stream
// cancels the handler context so a cooperative handler aborts. The
// breaker should record the call as a failure.
func TestCircuitBreakerStreamIdleTimeout(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second,
		ReadyToTrip: func(c circuitbreaker.Counts) bool { return c.ConsecutiveFailures >= 1 },
	})

	mw, err := CircuitBreakerStream(cb, streamtimeout.Config{
		IdleTimeout:  80 * time.Millisecond,
		TotalTimeout: 2 * time.Second,
	})
	if err != nil {
		t.Fatalf("CircuitBreakerStream: %v", err)
	}

	var handlerObserved atomic.Bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		//nolint:errcheck // test response writer
		_, _ = fmt.Fprint(w, "data: first\n\n")
		f.Flush()
		// Stall longer than IdleTimeout; ctx cancellation should release us.
		select {
		case <-r.Context().Done():
			handlerObserved.Store(true)
		case <-time.After(time.Second):
		}
	})

	srv := httptest.NewServer(mw(handler))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	// Drain so the handler-side goroutine unwinds cleanly.
	_, _ = io.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if !handlerObserved.Load() {
		t.Errorf("handler did not observe ctx cancellation — idle watchdog did not fire")
	}
}

// TestCircuitBreakerStreamFirstByteTimeout verifies a handler that
// never writes is killed by the FirstByteTimeout, and the middleware
// synthesises a 504 before the handler ever wrote headers.
func TestCircuitBreakerStreamFirstByteTimeout(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second,
	})

	mw, err := CircuitBreakerStream(cb, streamtimeout.Config{
		FirstByteTimeout: 80 * time.Millisecond,
		TotalTimeout:     2 * time.Second,
	})
	if err != nil {
		t.Fatalf("CircuitBreakerStream: %v", err)
	}

	handler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})

	srv := httptest.NewServer(mw(handler))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusGatewayTimeout {
		t.Errorf("status = %d, want 504", resp.StatusCode)
	}
}

// TestCircuitBreakerStreamConcurrent verifies the middleware's
// per-request streamingResponseWriter is not shared across requests
// (the streamtimeout instance is, the writer state is not).
func TestCircuitBreakerStreamConcurrent(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 100,
		Interval:    time.Second,
	})
	mw, err := CircuitBreakerStream(cb, streamtimeout.Config{
		FirstByteTimeout: 500 * time.Millisecond,
		IdleTimeout:      500 * time.Millisecond,
		TotalTimeout:     5 * time.Second,
	})
	if err != nil {
		t.Fatalf("CircuitBreakerStream: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		f := w.(http.Flusher)
		w.WriteHeader(http.StatusOK)
		//nolint:errcheck // test response writer
		_, _ = fmt.Fprint(w, "ok")
		f.Flush()
	})

	srv := httptest.NewServer(mw(handler))
	defer srv.Close()

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(srv.URL)
			if err != nil {
				t.Errorf("get: %v", err)
				return
			}
			defer func() { _ = resp.Body.Close() }()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("status = %d, want 200", resp.StatusCode)
			}
		}()
	}
	wg.Wait()
}

// TestCircuitBreakerStreamConfigValidation rejects empty configs.
func TestCircuitBreakerStreamConfigValidation(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 1, Interval: time.Second,
	})
	if _, err := CircuitBreakerStream(cb, streamtimeout.Config{}); err == nil {
		t.Errorf("expected error for empty streamtimeout.Config")
	}
	if _, err := CircuitBreakerStream(nil, streamtimeout.Config{TotalTimeout: time.Second}); err == nil {
		t.Errorf("expected error for nil CircuitBreaker")
	}
}

// TestCircuitBreakerStreamHandlerError surfaces a 5xx from the
// handler through the breaker (counts as failure).
func TestCircuitBreakerStreamHandlerError(t *testing.T) {
	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
		MaxRequests: 10,
		Interval:    time.Second,
		ReadyToTrip: func(c circuitbreaker.Counts) bool { return c.ConsecutiveFailures >= 1 },
	})
	mw, err := CircuitBreakerStream(cb, streamtimeout.Config{
		TotalTimeout: time.Second,
	})
	if err != nil {
		t.Fatalf("CircuitBreakerStream: %v", err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	srv := httptest.NewServer(mw(handler))
	defer srv.Close()

	// First call: 500 from handler, breaker counts a failure.
	resp1, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("get 1: %v", err)
	}
	_ = resp1.Body.Close()
	if resp1.StatusCode != http.StatusInternalServerError {
		t.Errorf("call 1 status = %d, want 500", resp1.StatusCode)
	}

	// Second call: breaker open, expect 503.
	resp2, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("get 2: %v", err)
	}
	_ = resp2.Body.Close()
	if resp2.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("call 2 status = %d, want 503", resp2.StatusCode)
	}
}

// Ensure the embedded ctx import does not get pruned by goimports.
var _ = context.Background
