package grpc

import (
	"context"
	"errors"
	"time"

	"go.klarlabs.de/fortify/circuitbreaker"
	"go.klarlabs.de/fortify/ferrors"
	"go.klarlabs.de/fortify/ratelimit"
	"go.klarlabs.de/fortify/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryClientKeyExtractor extracts a rate-limit key from a unary client call.
type UnaryClientKeyExtractor func(ctx context.Context, method string) string

// StreamClientKeyExtractor extracts a rate-limit key from a streaming client call.
type StreamClientKeyExtractor func(ctx context.Context, method string) string

// UnaryClientCircuitBreakerInterceptor wraps unary client calls with circuit
// breaker protection. When the breaker is open, returns Unavailable without
// invoking the RPC.
func UnaryClientCircuitBreakerInterceptor(cb circuitbreaker.CircuitBreaker[interface{}]) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, err := cb.Execute(ctx, func(ctx context.Context) (interface{}, error) {
			return nil, invoker(ctx, method, req, reply, cc, opts...)
		})
		if err == nil {
			return nil
		}
		if errors.Is(err, ferrors.ErrCircuitOpen) {
			return status.Error(codes.Unavailable, err.Error())
		}
		return err
	}
}

// UnaryClientRateLimitInterceptor wraps unary client calls with rate limiting.
// When the limit is exceeded, returns ResourceExhausted without invoking the RPC.
func UnaryClientRateLimitInterceptor(rl ratelimit.RateLimiter, keyExtractor UnaryClientKeyExtractor) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		key := keyExtractor(ctx, method)
		if !rl.Allow(ctx, key) {
			return status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// UnaryClientTimeoutInterceptor wraps unary client calls with a timeout.
// When the deadline is exceeded by Fortify's timeout (not the parent ctx),
// returns DeadlineExceeded.
func UnaryClientTimeoutInterceptor(tm timeout.Timeout[interface{}], duration time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, err := tm.Execute(ctx, duration, func(ctx context.Context) (interface{}, error) {
			return nil, invoker(ctx, method, req, reply, cc, opts...)
		})
		if err == nil {
			return nil
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return status.Error(codes.DeadlineExceeded, err.Error())
		}
		return err
	}
}

// StreamClientCircuitBreakerInterceptor wraps streaming client calls with
// circuit breaker protection. The breaker only gates the stream creation;
// once the stream is established, individual messages are not gated.
func StreamClientCircuitBreakerInterceptor(cb circuitbreaker.CircuitBreaker[interface{}]) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		var stream grpc.ClientStream
		_, err := cb.Execute(ctx, func(ctx context.Context) (interface{}, error) {
			s, e := streamer(ctx, desc, cc, method, opts...)
			if e != nil {
				return nil, e
			}
			stream = s
			return nil, nil
		})
		if err == nil {
			return stream, nil
		}
		if errors.Is(err, ferrors.ErrCircuitOpen) {
			return nil, status.Error(codes.Unavailable, err.Error())
		}
		return nil, err
	}
}

// StreamClientRateLimitInterceptor wraps streaming client calls with rate
// limiting. Only the stream creation consumes a token; individual messages
// are not gated.
func StreamClientRateLimitInterceptor(rl ratelimit.RateLimiter, keyExtractor StreamClientKeyExtractor) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		key := keyExtractor(ctx, method)
		if !rl.Allow(ctx, key) {
			return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// KeyFromOutgoingMetadata extracts a rate-limit key from the outgoing
// metadata context, sanitizing and truncating to a safe length.
func KeyFromOutgoingMetadata(header string) UnaryClientKeyExtractor {
	return func(ctx context.Context, _ string) string {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			return ""
		}
		values := md.Get(header)
		if len(values) == 0 {
			return ""
		}
		return sanitizeMetadataKey(values[0])
	}
}

// StreamKeyFromOutgoingMetadata is the streaming-client equivalent of
// KeyFromOutgoingMetadata.
func StreamKeyFromOutgoingMetadata(header string) StreamClientKeyExtractor {
	return func(ctx context.Context, _ string) string {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			return ""
		}
		values := md.Get(header)
		if len(values) == 0 {
			return ""
		}
		return sanitizeMetadataKey(values[0])
	}
}

// KeyFromMethod extracts the gRPC method name as the rate-limit key.
// Useful for per-method rate limiting on the client side.
func KeyFromMethod() UnaryClientKeyExtractor {
	return func(_ context.Context, method string) string {
		return method
	}
}

// StreamKeyFromMethod is the streaming-client equivalent of KeyFromMethod.
func StreamKeyFromMethod() StreamClientKeyExtractor {
	return func(_ context.Context, method string) string {
		return method
	}
}
