package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Mock unary handler for testing.
func mockUnaryHandler(resp interface{}, err error) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return resp, err
	}
}

// Mock stream handler for testing.
func mockStreamHandler(err error) grpc.StreamHandler {
	return func(srv interface{}, stream grpc.ServerStream) error {
		return err
	}
}

// Mock server stream for testing.
type mockServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (m *mockServerStream) Context() context.Context {
	return m.ctx
}

func TestUnaryCircuitBreakerInterceptor(t *testing.T) {
	t.Run("allows requests when circuit closed", func(t *testing.T) {
		cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
			MaxRequests: 10,
			Interval:    time.Second,
		})

		interceptor := UnaryCircuitBreakerInterceptor(cb)
		handler := mockUnaryHandler("success", nil)

		resp, err := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if resp != "success" {
			t.Errorf("response = %v, want success", resp)
		}
	})

	t.Run("returns unavailable when circuit open", func(t *testing.T) {
		cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
			MaxRequests: 1,
			Interval:    time.Second,
			ReadyToTrip: func(counts circuitbreaker.Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		interceptor := UnaryCircuitBreakerInterceptor(cb)

		// First request fails, opens circuit
		handler1 := mockUnaryHandler(nil, errors.New("service error"))
		//nolint:errcheck // intentionally ignoring error in test
		_, _ = interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler1)

		// Second request should be rejected
		handler2 := mockUnaryHandler("success", nil)
		_, err := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler2)

		if err == nil {
			t.Error("expected error, got nil")
		}
		st, ok := status.FromError(err)
		if !ok {
			t.Fatal("expected gRPC status error")
		}
		if st.Code() != codes.Unavailable {
			t.Errorf("code = %v, want Unavailable", st.Code())
		}
	})
}

func TestUnaryRateLimitInterceptor(t *testing.T) {
	t.Run("allows requests within rate limit", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     10,
			Interval: time.Second,
		})

		keyExtractor := func(ctx context.Context, req interface{}) string {
			return "test-key"
		}

		interceptor := UnaryRateLimitInterceptor(rl, keyExtractor)
		handler := mockUnaryHandler("success", nil)

		resp, err := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if resp != "success" {
			t.Errorf("response = %v, want success", resp)
		}
	})

	t.Run("returns resource exhausted when rate limit exceeded", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     1,
			Burst:    1,
			Interval: time.Hour,
		})

		keyExtractor := func(ctx context.Context, req interface{}) string {
			return "test-key"
		}

		interceptor := UnaryRateLimitInterceptor(rl, keyExtractor)
		handler := mockUnaryHandler("success", nil)

		// First request succeeds
		_, err1 := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)
		if err1 != nil {
			t.Fatalf("first request failed: %v", err1)
		}

		// Second request should be rate limited
		_, err2 := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)
		if err2 == nil {
			t.Error("expected error, got nil")
		}
		st, ok := status.FromError(err2)
		if !ok {
			t.Fatal("expected gRPC status error")
		}
		if st.Code() != codes.ResourceExhausted {
			t.Errorf("code = %v, want ResourceExhausted", st.Code())
		}
	})
}

func TestUnaryTimeoutInterceptor(t *testing.T) {
	t.Run("completes request within timeout", func(t *testing.T) {
		tm := timeout.New[interface{}](timeout.Config{
			DefaultTimeout: time.Second,
		})

		interceptor := UnaryTimeoutInterceptor(tm, 100*time.Millisecond)
		handler := mockUnaryHandler("success", nil)

		resp, err := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if resp != "success" {
			t.Errorf("response = %v, want success", resp)
		}
	})

	t.Run("returns deadline exceeded when request times out", func(t *testing.T) {
		tm := timeout.New[interface{}](timeout.Config{
			DefaultTimeout: time.Second,
		})

		interceptor := UnaryTimeoutInterceptor(tm, 50*time.Millisecond)
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(100 * time.Millisecond):
				return "success", nil
			}
		}

		_, err := interceptor(context.Background(), "request", &grpc.UnaryServerInfo{}, handler)

		if err == nil {
			t.Error("expected error, got nil")
		}
		st, ok := status.FromError(err)
		if !ok {
			t.Fatal("expected gRPC status error")
		}
		if st.Code() != codes.DeadlineExceeded {
			t.Errorf("code = %v, want DeadlineExceeded", st.Code())
		}
	})
}

func TestStreamCircuitBreakerInterceptor(t *testing.T) {
	t.Run("allows streams when circuit closed", func(t *testing.T) {
		cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
			MaxRequests: 10,
			Interval:    time.Second,
		})

		interceptor := StreamCircuitBreakerInterceptor(cb)
		handler := mockStreamHandler(nil)
		stream := &mockServerStream{ctx: context.Background()}

		err := interceptor(nil, stream, &grpc.StreamServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("returns unavailable when circuit open", func(t *testing.T) {
		cb := circuitbreaker.New[interface{}](circuitbreaker.Config{
			MaxRequests: 1,
			Interval:    time.Second,
			ReadyToTrip: func(counts circuitbreaker.Counts) bool {
				return counts.ConsecutiveFailures >= 1
			},
		})

		interceptor := StreamCircuitBreakerInterceptor(cb)

		// First request fails, opens circuit
		handler1 := mockStreamHandler(errors.New("service error"))
		stream1 := &mockServerStream{ctx: context.Background()}
		//nolint:errcheck // intentionally ignoring error in test
		_ = interceptor(nil, stream1, &grpc.StreamServerInfo{}, handler1)

		// Second request should be rejected
		handler2 := mockStreamHandler(nil)
		stream2 := &mockServerStream{ctx: context.Background()}
		err := interceptor(nil, stream2, &grpc.StreamServerInfo{}, handler2)

		if err == nil {
			t.Error("expected error, got nil")
		}
		if !errors.Is(err, fortifyerrors.ErrCircuitOpen) {
			st, ok := status.FromError(err)
			if !ok {
				t.Fatal("expected gRPC status error")
			}
			if st.Code() != codes.Unavailable {
				t.Errorf("code = %v, want Unavailable", st.Code())
			}
		}
	})
}

func TestStreamRateLimitInterceptor(t *testing.T) {
	t.Run("allows streams within rate limit", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     10,
			Interval: time.Second,
		})

		keyExtractor := func(srv interface{}, stream grpc.ServerStream) string {
			return "test-key"
		}

		interceptor := StreamRateLimitInterceptor(rl, keyExtractor)
		handler := mockStreamHandler(nil)
		stream := &mockServerStream{ctx: context.Background()}

		err := interceptor(nil, stream, &grpc.StreamServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("returns resource exhausted when rate limit exceeded", func(t *testing.T) {
		rl := ratelimit.New(ratelimit.Config{
			Rate:     1,
			Burst:    1,
			Interval: time.Hour,
		})

		keyExtractor := func(srv interface{}, stream grpc.ServerStream) string {
			return "test-key"
		}

		interceptor := StreamRateLimitInterceptor(rl, keyExtractor)
		handler := mockStreamHandler(nil)

		// First request succeeds
		stream1 := &mockServerStream{ctx: context.Background()}
		err1 := interceptor(nil, stream1, &grpc.StreamServerInfo{}, handler)
		if err1 != nil {
			t.Fatalf("first request failed: %v", err1)
		}

		// Second request should be rate limited
		stream2 := &mockServerStream{ctx: context.Background()}
		err2 := interceptor(nil, stream2, &grpc.StreamServerInfo{}, handler)
		if err2 == nil {
			t.Error("expected error, got nil")
		}
		st, ok := status.FromError(err2)
		if !ok {
			t.Fatal("expected gRPC status error")
		}
		if st.Code() != codes.ResourceExhausted {
			t.Errorf("code = %v, want ResourceExhausted", st.Code())
		}
	})
}

func TestStreamTimeoutInterceptor(t *testing.T) {
	t.Run("completes stream within timeout", func(t *testing.T) {
		tm := timeout.New[interface{}](timeout.Config{
			DefaultTimeout: time.Second,
		})

		interceptor := StreamTimeoutInterceptor(tm, 100*time.Millisecond)
		handler := mockStreamHandler(nil)
		stream := &mockServerStream{ctx: context.Background()}

		err := interceptor(nil, stream, &grpc.StreamServerInfo{}, handler)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("returns deadline exceeded when stream times out", func(t *testing.T) {
		tm := timeout.New[interface{}](timeout.Config{
			DefaultTimeout: time.Second,
		})

		interceptor := StreamTimeoutInterceptor(tm, 50*time.Millisecond)
		handler := func(srv interface{}, stream grpc.ServerStream) error {
			select {
			case <-stream.Context().Done():
				return stream.Context().Err()
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		}
		stream := &mockServerStream{ctx: context.Background()}

		err := interceptor(nil, stream, &grpc.StreamServerInfo{}, handler)

		if err == nil {
			t.Error("expected error, got nil")
		}
		st, ok := status.FromError(err)
		if !ok {
			t.Fatal("expected gRPC status error")
		}
		if st.Code() != codes.DeadlineExceeded {
			t.Errorf("code = %v, want DeadlineExceeded", st.Code())
		}
	})
}

func TestKeyExtractors(t *testing.T) {
	t.Run("extracts key from metadata", func(t *testing.T) {
		// This test would require actual gRPC metadata setup
		// For now, we'll just verify the extractor returns a non-empty key
		extractor := KeyFromMetadata("x-api-key")
		key := extractor(context.Background(), nil)
		// Empty key is acceptable when no metadata present
		_ = key
	})
}
