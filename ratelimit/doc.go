// Package ratelimit provides token bucket based rate limiting for controlling
// request rates and preventing resource exhaustion.
//
// The rate limiter uses a token bucket algorithm where tokens are added at a
// constant rate up to a maximum burst capacity. Each request consumes one or more
// tokens. When the bucket is empty, requests are either rejected (Allow) or wait
// for tokens to become available (Wait).
//
// # Basic Usage
//
// Create a rate limiter with default in-memory storage:
//
//	limiter := ratelimit.New(&ratelimit.Config{
//	    Rate:     100,           // 100 tokens per second
//	    Burst:    150,           // Allow bursts up to 150
//	    Interval: time.Second,
//	})
//
//	if limiter.Allow(ctx, "user-123") {
//	    // Process request
//	} else {
//	    // Return 429 Too Many Requests
//	}
//
// # Storage Backends
//
// The rate limiter uses a pluggable Store interface for state management.
// By default, an in-memory store is used. For distributed rate limiting
// across multiple instances, provide a custom Store implementation backed
// by Redis, DynamoDB, or another distributed backend.
//
//	limiter := ratelimit.New(&ratelimit.Config{
//	    Rate:     100,
//	    Burst:    150,
//	    Interval: time.Second,
//	    Store:    myRedisStore,  // Custom Store implementation
//	    FailOpen: true,          // Allow requests on storage failure
//	})
//
// # Key Extraction
//
// Use KeyFunc to extract rate limiting keys from context:
//
//	limiter := ratelimit.New(&ratelimit.Config{
//	    Rate: 100,
//	    KeyFunc: func(ctx context.Context) string {
//	        return ctx.Value("user_id").(string)
//	    },
//	})
//
// # Observability
//
// The rate limiter supports structured logging via slog and metrics collection
// via the optional Metrics interface:
//
//	limiter := ratelimit.New(&ratelimit.Config{
//	    Rate:    100,
//	    Logger:  slog.Default(),
//	    Metrics: myMetricsRecorder,
//	})
//
// # Resource Management
//
// Always close the rate limiter when done to release resources:
//
//	limiter := ratelimit.New(config)
//	defer limiter.Close()
//
// This is especially important when using distributed stores that maintain
// connections (Redis, database connections, etc.).
package ratelimit
