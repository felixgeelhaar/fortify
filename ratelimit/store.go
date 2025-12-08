package ratelimit

import (
	"context"
	"time"
)

// BucketState represents the state of a token bucket.
// This is the data that gets stored and retrieved by the Store interface.
//
// Implementations of Store should treat BucketState as immutable.
// The updateFn in AtomicUpdate should always return a new BucketState
// instance rather than modifying the input.
type BucketState struct {
	// LastRefill is the timestamp of the last refill calculation.
	// Used to calculate how many tokens to add based on elapsed time.
	LastRefill time.Time

	// Tokens is the current number of available tokens.
	// This can be fractional to allow precise token accumulation.
	Tokens float64
}

// Store defines the storage interface for rate limiter state.
//
// # Stability
//
// This interface is considered stable. Breaking changes (method signature changes,
// new required methods) will only be made in major version releases and will be
// clearly documented in the CHANGELOG. Optional extensions should be added as
// separate interfaces (like HealthChecker) that implementations can optionally
// implement.
//
// # Thread Safety
//
// All Store implementations MUST be safe for concurrent use by multiple goroutines.
// The rate limiter may call Store methods from multiple goroutines simultaneously,
// both for the same key and for different keys. Implementations must handle:
//
//   - Concurrent AtomicUpdate calls on the same key (must serialize correctly)
//   - Concurrent AtomicUpdate calls on different keys (should not block each other)
//   - Concurrent Get/Delete calls mixed with AtomicUpdate calls
//   - Close being called while other operations are in progress
//
// # Implementation Requirements
//
// Implementations must ensure atomic read-modify-write operations
// within AtomicUpdate for correctness in concurrent environments.
// For distributed storage backends (Redis, DynamoDB, etc.), this typically
// requires using transactions, optimistic locking (CAS), or Lua scripts.
//
// # Example Implementations
//
//   - In-memory: Use sync.Map with per-key mutex (see MemoryStore)
//   - Redis: Use WATCH/MULTI/EXEC or Lua scripts
//   - DynamoDB: Use conditional writes with version attributes
//   - PostgreSQL: Use advisory locks or serializable transactions
type Store interface {
	// AtomicUpdate atomically reads, modifies, and writes bucket state.
	//
	// The updateFn receives the current state (or nil if the bucket doesn't exist)
	// and returns the new state to store. The implementation must ensure the entire
	// read-modify-write cycle is atomic to prevent race conditions.
	//
	// The updateFn contains the token bucket algorithm logic (refill and consume).
	// If updateFn returns nil, the operation is considered a no-op and no write occurs.
	//
	// Returns the final state after the update, or an error if the operation failed.
	AtomicUpdate(ctx context.Context, key string, updateFn func(*BucketState) *BucketState) (*BucketState, error)

	// Get retrieves the current bucket state without modification.
	// Returns nil if the bucket doesn't exist. This is useful for monitoring
	// and debugging without triggering state updates.
	Get(ctx context.Context, key string) (*BucketState, error)

	// Delete removes a bucket from the store.
	// Returns nil if the bucket doesn't exist.
	Delete(ctx context.Context, key string) error

	// Close releases any resources held by the store.
	// After Close is called, the store should not be used.
	Close() error
}

// HealthChecker is an optional interface for stores that support health checks.
// Implement this interface to enable health monitoring for distributed stores.
type HealthChecker interface {
	// HealthCheck verifies the store is operational.
	// Returns nil if healthy, or an error describing the issue.
	HealthCheck(ctx context.Context) error
}

// Resetter is an optional interface for stores that support clearing all state.
// Implement this interface to enable Reset() functionality on the rate limiter.
type Resetter interface {
	// Reset clears all buckets from the store.
	// This is useful for testing or administrative purposes.
	//
	// Note: Reset is not atomic with respect to concurrent operations.
	// Entries added during Reset may or may not be removed.
	// For deterministic clearing, ensure no concurrent operations are in progress.
	Reset(ctx context.Context) error
}

// BucketCounter is an optional interface for stores that can report active bucket count.
// Implement this interface to enable BucketCount() functionality on the rate limiter.
type BucketCounter interface {
	// BucketCount returns the current number of active buckets in the store.
	BucketCount() int64
}

// Metrics defines the interface for collecting rate limiter metrics.
// Implement this interface to integrate with your metrics system
// (Prometheus, StatsD, OpenTelemetry, etc.).
//
// All methods receive a context for distributed tracing integration.
// Use the context to extract trace IDs, span context, or other metadata.
type Metrics interface {
	// OnAllow is called when a request is allowed.
	// The context contains any trace context from the original request.
	OnAllow(ctx context.Context, key string)

	// OnDeny is called when a request is denied due to rate limiting.
	// The context contains any trace context from the original request.
	OnDeny(ctx context.Context, key string)

	// OnError is called when a storage error occurs.
	// The context contains any trace context from the original request.
	OnError(ctx context.Context, key string, err error)

	// OnStoreLatency records the latency of store operations.
	// The context contains any trace context from the original request.
	OnStoreLatency(ctx context.Context, operation string, duration time.Duration)
}
