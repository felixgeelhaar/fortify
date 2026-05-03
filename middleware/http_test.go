package middleware_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/middleware"
)

// fakeTransport allows controlling RoundTrip outcomes per-call for tests.
type fakeTransport struct {
	calls   atomic.Int32
	handler func(*http.Request) (*http.Response, error)
}

func (f *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	f.calls.Add(1)
	return f.handler(req)
}

func TestHTTPRoundTripper_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.HTTPRoundTripper(nil, middleware.HTTPClientConfig{Timeout: 0})
	if err == nil {
		t.Fatal("want error for zero Timeout, got nil")
	}
}

func TestHTTPRoundTripper_DefaultTransportFallback(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	rt, err := middleware.HTTPRoundTripper(nil, middleware.HTTPClientConfig{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	client := &http.Client{Transport: rt}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("Get err = %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
}

func TestHTTPRoundTripper_RetriesOnTransientError(t *testing.T) {
	wantOK := &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}

	ft := &fakeTransport{}
	ft.handler = func(req *http.Request) (*http.Response, error) {
		if ft.calls.Load() < 3 {
			return nil, errors.New("transient")
		}
		return wantOK, nil
	}

	rt, err := middleware.HTTPRoundTripper(ft, middleware.HTTPClientConfig{
		Timeout:           500 * time.Millisecond,
		MaxRetries:        5,
		RetryInitialDelay: 1 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.invalid/", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip err = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if c := ft.calls.Load(); c != 3 {
		t.Fatalf("calls = %d, want 3 (2 transient + 1 success)", c)
	}
}

func TestHTTPRoundTripperFromChain_UsesProvidedChain(t *testing.T) {
	chain := middleware.New[*http.Response]()

	wantResp := &http.Response{StatusCode: http.StatusTeapot, Body: http.NoBody}
	ft := &fakeTransport{
		handler: func(*http.Request) (*http.Response, error) { return wantResp, nil },
	}

	rt := middleware.HTTPRoundTripperFromChain(ft, chain)
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.invalid/", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip err = %v", err)
	}
	if resp.StatusCode != http.StatusTeapot {
		t.Fatalf("status = %d, want 418", resp.StatusCode)
	}
}
