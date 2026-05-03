package middleware_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/middleware"
)

func TestHTTPClient_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.HTTPClient(middleware.HTTPClientConfig{Timeout: 0})
	if err == nil {
		t.Fatal("want error for zero Timeout, got nil")
	}
}

func TestHTTPClient_ExecutesOperation(t *testing.T) {
	chain, err := middleware.HTTPClient(middleware.HTTPClientConfig{Timeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	want := &http.Response{StatusCode: 200}
	got, err := chain.Execute(context.Background(), func(ctx context.Context) (*http.Response, error) {
		return want, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestHTTPClient_RetriesOnTransientError(t *testing.T) {
	chain, err := middleware.HTTPClient(middleware.HTTPClientConfig{
		Timeout:           500 * time.Millisecond,
		MaxRetries:        3,
		RetryInitialDelay: 1 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	attempts := 0
	transient := errors.New("transient")
	got, err := chain.Execute(context.Background(), func(ctx context.Context) (*http.Response, error) {
		attempts++
		if attempts < 3 {
			return nil, transient
		}
		return &http.Response{StatusCode: 200}, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", got.StatusCode)
	}
	if attempts != 3 {
		t.Fatalf("attempts = %d, want 3", attempts)
	}
}

func TestDatabaseQuery_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.DatabaseQuery(middleware.DatabaseQueryConfig{QueryTimeout: 0})
	if err == nil {
		t.Fatal("want error for zero QueryTimeout, got nil")
	}
}

func TestDatabaseQuery_ExecutesOperation(t *testing.T) {
	chain, err := middleware.DatabaseQuery(middleware.DatabaseQueryConfig{QueryTimeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	got, err := chain.Execute(context.Background(), func(ctx context.Context) (any, error) {
		return "row-data", nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got.(string) != "row-data" {
		t.Fatalf("got %v, want row-data", got)
	}
}

func TestRPCDownstream_RejectsZeroTimeout(t *testing.T) {
	_, err := middleware.RPCDownstream[string](middleware.RPCDownstreamConfig{CallTimeout: 0})
	if err == nil {
		t.Fatal("want error for zero CallTimeout, got nil")
	}
}

func TestRPCDownstream_GenericTypePropagated(t *testing.T) {
	chain, err := middleware.RPCDownstream[int](middleware.RPCDownstreamConfig{CallTimeout: 100 * time.Millisecond})
	if err != nil {
		t.Fatalf("constructor err = %v", err)
	}

	got, err := chain.Execute(context.Background(), func(ctx context.Context) (int, error) {
		return 42, nil
	})
	if err != nil {
		t.Fatalf("Execute err = %v", err)
	}
	if got != 42 {
		t.Fatalf("got %d, want 42", got)
	}
}
