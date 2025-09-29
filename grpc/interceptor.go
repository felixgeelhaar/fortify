// Package grpc provides gRPC interceptors for integrating Fortify resilience patterns
// with gRPC services.
//
// This package offers both unary and streaming interceptors that wrap gRPC handlers
// with circuit breakers, rate limiters, timeouts, and other resilience patterns.
//
// Example usage:
//
//	cb := circuitbreaker.New[interface{}](circuitbreaker.Config{...})
//	rl := ratelimit.New(ratelimit.Config{...})
//
//	server := grpc.NewServer(
//	    grpc.UnaryInterceptor(fortifygrpc.UnaryCircuitBreakerInterceptor(cb)),
//	    grpc.StreamInterceptor(fortifygrpc.StreamRateLimitInterceptor(rl, fortifygrpc.KeyFromMetadata("x-api-key"))),
//	)
package grpc

import (
	"context"
	"time"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	fortifyerrors "github.com/felixgeelhaar/fortify/errors"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryKeyExtractor extracts a rate limiting key from a unary RPC context.
type UnaryKeyExtractor func(ctx context.Context, req interface{}) string

// StreamKeyExtractor extracts a rate limiting key from a streaming RPC context.
type StreamKeyExtractor func(srv interface{}, stream grpc.ServerStream) string

// UnaryCircuitBreakerInterceptor wraps a unary RPC handler with circuit breaker protection.
// Returns Unavailable status when the circuit is open.
func UnaryCircuitBreakerInterceptor(cb circuitbreaker.CircuitBreaker[interface{}]) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := cb.Execute(ctx, func(ctx context.Context) (interface{}, error) {
			return handler(ctx, req)
		})

		if err != nil {
			// Convert circuit breaker errors to gRPC status codes
			if err == fortifyerrors.ErrCircuitOpen {
				return nil, status.Error(codes.Unavailable, "circuit breaker is open")
			}
			return nil, status.Error(codes.Unavailable, err.Error())
		}

		return resp, nil
	}
}

// UnaryRateLimitInterceptor wraps a unary RPC handler with rate limiting.
// Returns ResourceExhausted status when the rate limit is exceeded.
func UnaryRateLimitInterceptor(rl ratelimit.RateLimiter, keyExtractor UnaryKeyExtractor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		key := keyExtractor(ctx, req)

		if !rl.Allow(ctx, key) {
			return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}

		return handler(ctx, req)
	}
}

// UnaryTimeoutInterceptor wraps a unary RPC handler with timeout enforcement.
// Returns DeadlineExceeded status when the request times out.
func UnaryTimeoutInterceptor(tm timeout.Timeout[interface{}], duration time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := tm.Execute(ctx, duration, func(ctx context.Context) (interface{}, error) {
			return handler(ctx, req)
		})

		if err != nil {
			// Convert timeout errors to gRPC status codes
			if err == context.DeadlineExceeded {
				return nil, status.Error(codes.DeadlineExceeded, "request timeout")
			}
			return nil, err
		}

		return resp, nil
	}
}

// StreamCircuitBreakerInterceptor wraps a streaming RPC handler with circuit breaker protection.
// Returns Unavailable status when the circuit is open.
func StreamCircuitBreakerInterceptor(cb circuitbreaker.CircuitBreaker[interface{}]) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		_, err := cb.Execute(stream.Context(), func(ctx context.Context) (interface{}, error) {
			return nil, handler(srv, stream)
		})

		if err != nil {
			// Convert circuit breaker errors to gRPC status codes
			if err == fortifyerrors.ErrCircuitOpen {
				return status.Error(codes.Unavailable, "circuit breaker is open")
			}
			return status.Error(codes.Unavailable, err.Error())
		}

		return nil
	}
}

// StreamRateLimitInterceptor wraps a streaming RPC handler with rate limiting.
// Returns ResourceExhausted status when the rate limit is exceeded.
func StreamRateLimitInterceptor(rl ratelimit.RateLimiter, keyExtractor StreamKeyExtractor) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		key := keyExtractor(srv, stream)

		if !rl.Allow(stream.Context(), key) {
			return status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}

		return handler(srv, stream)
	}
}

// StreamTimeoutInterceptor wraps a streaming RPC handler with timeout enforcement.
// Returns DeadlineExceeded status when the stream times out.
func StreamTimeoutInterceptor(tm timeout.Timeout[interface{}], duration time.Duration) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		_, err := tm.Execute(stream.Context(), duration, func(ctx context.Context) (interface{}, error) {
			return nil, handler(srv, &wrappedServerStream{ServerStream: stream, ctx: ctx})
		})

		if err != nil {
			// Convert timeout errors to gRPC status codes
			if err == context.DeadlineExceeded {
				return status.Error(codes.DeadlineExceeded, "stream timeout")
			}
			return err
		}

		return nil
	}
}

// KeyFromMetadata returns a UnaryKeyExtractor that extracts the key from gRPC metadata.
func KeyFromMetadata(header string) UnaryKeyExtractor {
	return func(ctx context.Context, req interface{}) string {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ""
		}
		values := md.Get(header)
		if len(values) == 0 {
			return ""
		}
		return values[0]
	}
}

// StreamKeyFromMetadata returns a StreamKeyExtractor that extracts the key from gRPC metadata.
func StreamKeyFromMetadata(header string) StreamKeyExtractor {
	return func(srv interface{}, stream grpc.ServerStream) string {
		md, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return ""
		}
		values := md.Get(header)
		if len(values) == 0 {
			return ""
		}
		return values[0]
	}
}

// wrappedServerStream wraps a grpc.ServerStream with a custom context.
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}
