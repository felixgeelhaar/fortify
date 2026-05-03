package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestUnaryClientCircuitBreakerInterceptor_OpenReturnsUnavailable(t *testing.T) {
	cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
		ReadyToTrip: func(c circuitbreaker.Counts) bool { return c.ConsecutiveFailures >= 1 },
		Timeout:     1 * time.Hour,
	})
	defer cb.Close()
	interceptor := UnaryClientCircuitBreakerInterceptor(cb)

	failingInvoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return errors.New("rpc failed")
	}

	// First call: trips breaker.
	_ = interceptor(context.Background(), "/test", nil, nil, nil, failingInvoker)

	// Second call: breaker is open.
	err := interceptor(context.Background(), "/test", nil, nil, nil, failingInvoker)
	if err == nil {
		t.Fatal("want error, got nil")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("not a status error: %v", err)
	}
	if st.Code() != codes.Unavailable {
		t.Fatalf("code = %v, want Unavailable", st.Code())
	}
}

func TestUnaryClientCircuitBreakerInterceptor_PassesThroughOnSuccess(t *testing.T) {
	cb := circuitbreaker.New[interface{}](circuitbreaker.Config{})
	defer cb.Close()
	interceptor := UnaryClientCircuitBreakerInterceptor(cb)

	called := false
	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		called = true
		return nil
	}

	if err := interceptor(context.Background(), "/test", nil, nil, nil, invoker); err != nil {
		t.Fatalf("err = %v", err)
	}
	if !called {
		t.Fatal("invoker not called")
	}
}

func TestUnaryClientRateLimitInterceptor_DeniesOverLimit(t *testing.T) {
	rl := ratelimit.New(ratelimit.Config{Rate: 1, Burst: 1, Interval: time.Hour})
	defer rl.Close()
	interceptor := UnaryClientRateLimitInterceptor(rl, KeyFromMethod())

	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return nil
	}

	// First call admitted.
	if err := interceptor(context.Background(), "/test", nil, nil, nil, invoker); err != nil {
		t.Fatalf("first call err = %v", err)
	}
	// Second call denied.
	err := interceptor(context.Background(), "/test", nil, nil, nil, invoker)
	if err == nil {
		t.Fatal("want denial, got nil")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("not a status error: %v", err)
	}
	if st.Code() != codes.ResourceExhausted {
		t.Fatalf("code = %v, want ResourceExhausted", st.Code())
	}
}

func TestUnaryClientTimeoutInterceptor_FiresOnSlowInvoker(t *testing.T) {
	tm := timeout.New[interface{}](timeout.Config{DefaultTimeout: 10 * time.Millisecond})
	interceptor := UnaryClientTimeoutInterceptor(tm, 10*time.Millisecond)

	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		<-ctx.Done()
		return ctx.Err()
	}

	err := interceptor(context.Background(), "/test", nil, nil, nil, invoker)
	if err == nil {
		t.Fatal("want timeout, got nil")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("not a status error: %v", err)
	}
	if st.Code() != codes.DeadlineExceeded {
		t.Fatalf("code = %v, want DeadlineExceeded", st.Code())
	}
}

func TestStreamClientCircuitBreakerInterceptor_OpenReturnsUnavailable(t *testing.T) {
	cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
		ReadyToTrip: func(c circuitbreaker.Counts) bool { return c.ConsecutiveFailures >= 1 },
		Timeout:     1 * time.Hour,
	})
	defer cb.Close()
	interceptor := StreamClientCircuitBreakerInterceptor(cb)

	failingStreamer := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return nil, errors.New("dial failed")
	}

	// First call: trips breaker.
	_, _ = interceptor(context.Background(), &grpc.StreamDesc{}, nil, "/test", failingStreamer)

	// Second call: breaker is open.
	_, err := interceptor(context.Background(), &grpc.StreamDesc{}, nil, "/test", failingStreamer)
	if err == nil {
		t.Fatal("want error, got nil")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("not a status error: %v", err)
	}
	if st.Code() != codes.Unavailable {
		t.Fatalf("code = %v, want Unavailable", st.Code())
	}
}

func TestKeyFromOutgoingMetadata_ExtractsAndSanitizes(t *testing.T) {
	extractor := KeyFromOutgoingMetadata("x-api-key")

	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-api-key", "key-with-\x00-control")
	got := extractor(ctx, "/test")
	wantPrefix := "key-with--control"
	if got != wantPrefix {
		t.Fatalf("got %q, want %q", got, wantPrefix)
	}
}

func TestKeyFromOutgoingMetadata_MissingHeaderReturnsEmpty(t *testing.T) {
	extractor := KeyFromOutgoingMetadata("missing")
	if got := extractor(context.Background(), "/test"); got != "" {
		t.Fatalf("got %q, want empty", got)
	}
}

func TestKeyFromMethod_ReturnsMethod(t *testing.T) {
	extractor := KeyFromMethod()
	if got := extractor(context.Background(), "/svc/Method"); got != "/svc/Method" {
		t.Fatalf("got %q, want /svc/Method", got)
	}
}
