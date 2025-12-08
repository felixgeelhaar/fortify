package ratelimit

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// DefaultMaxKeys is the default maximum number of keys allowed in MemoryStore.
	DefaultMaxKeys = 100000

	// DefaultCleanupInterval is the default interval for cleaning up stale entries.
	DefaultCleanupInterval = time.Minute

	// DefaultEntryTTL is the default TTL for bucket entries.
	DefaultEntryTTL = time.Hour

	// DefaultMaxKeyLength is the maximum length of a rate limiting key.
	// Keys longer than this are rejected to prevent memory exhaustion.
	DefaultMaxKeyLength = 1024
)

// MemoryStoreOption configures a MemoryStore.
type MemoryStoreOption func(*MemoryStore)

// WithMaxKeys sets the maximum number of keys allowed in the store.
// When the limit is reached, new keys will be rejected with ErrKeyLimitExceeded.
func WithMaxKeys(maxKeys int) MemoryStoreOption {
	return func(m *MemoryStore) {
		if maxKeys > 0 {
			m.maxKeys = int64(maxKeys)
		}
	}
}

// WithCleanupInterval sets the interval for cleaning up stale entries.
// Set to 0 to disable automatic cleanup.
func WithCleanupInterval(interval time.Duration) MemoryStoreOption {
	return func(m *MemoryStore) {
		m.cleanupInterval = interval
	}
}

// WithEntryTTL sets the TTL for bucket entries.
// Entries not accessed within this duration will be removed during cleanup.
func WithEntryTTL(ttl time.Duration) MemoryStoreOption {
	return func(m *MemoryStore) {
		if ttl > 0 {
			m.entryTTL = ttl
		}
	}
}

// WithMaxKeyLength sets the maximum allowed length for rate limiting keys.
// Keys longer than this limit are rejected with ErrKeyTooLong.
func WithMaxKeyLength(maxLen int) MemoryStoreOption {
	return func(m *MemoryStore) {
		if maxLen > 0 {
			m.maxKeyLength = maxLen
		}
	}
}

// MemoryStore is an in-memory implementation of the Store interface.
// It uses sync.Map for concurrent bucket access and per-key mutexes
// for atomic updates.
//
// This is the default store used when no custom Store is provided
// in the rate limiter configuration.
//
// MemoryStore is safe for concurrent use by multiple goroutines.
// However, state is not persisted and is lost when the process exits.
// For distributed rate limiting across multiple instances, implement
// a custom Store using Redis, DynamoDB, or another distributed backend.
//
// Features:
//   - TTL-based automatic cleanup of stale entries
//   - Configurable maximum key limit to prevent memory exhaustion
//   - Configurable maximum key length to prevent memory exhaustion from long keys
//   - Health check support via HealthChecker interface
//
//nolint:govet // fieldalignment: optimized for cache efficiency, not minimal size
type MemoryStore struct {
	// Hot fields (cache line 1) - accessed on every operation
	buckets  sync.Map     // map[string]*bucketEntry - 48 bytes
	keyCount atomic.Int64 // 8 bytes
	closed   atomic.Bool  // 1 byte + 7 padding

	// Warm fields (cache line 2) - accessed occasionally
	maxKeys      int64          // 8 bytes
	maxKeyLength int            // 8 bytes
	done         chan struct{}  // 8 bytes
	wg           sync.WaitGroup // 12 bytes

	// Cold fields - rarely accessed (only during cleanup)
	cleanupInterval time.Duration // 8 bytes
	entryTTL        time.Duration // 8 bytes
}

// bucketEntry holds the state and mutex for a single bucket.
type bucketEntry struct {
	state      *BucketState
	lastAccess atomic.Int64 // Unix nano timestamp
	mu         sync.Mutex
}

// NewMemoryStore creates a new in-memory store with default settings.
func NewMemoryStore() *MemoryStore {
	return NewMemoryStoreWithOptions()
}

// NewMemoryStoreWithOptions creates a new in-memory store with the given options.
func NewMemoryStoreWithOptions(opts ...MemoryStoreOption) *MemoryStore {
	m := &MemoryStore{
		maxKeys:         DefaultMaxKeys,
		maxKeyLength:    DefaultMaxKeyLength,
		cleanupInterval: DefaultCleanupInterval,
		entryTTL:        DefaultEntryTTL,
		done:            make(chan struct{}),
	}

	for _, opt := range opts {
		opt(m)
	}

	// Start cleanup goroutine if interval is set
	if m.cleanupInterval > 0 {
		m.wg.Add(1)
		go m.cleanupLoop()
	}

	return m
}

// cleanupLoop periodically removes stale entries.
func (m *MemoryStore) cleanupLoop() {
	defer m.wg.Done()

	ticker := time.NewTicker(m.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.cleanup()
		case <-m.done:
			return
		}
	}
}

// cleanup removes entries that haven't been accessed within the TTL.
func (m *MemoryStore) cleanup() {
	now := time.Now().UnixNano()
	ttlNanos := m.entryTTL.Nanoseconds()

	m.buckets.Range(func(key, value interface{}) bool {
		entry, ok := value.(*bucketEntry)
		if !ok {
			return true // Skip invalid entry (should never happen)
		}
		lastAccess := entry.lastAccess.Load()

		// If entry hasn't been accessed within TTL, remove it
		if now-lastAccess > ttlNanos {
			// Lock entry before deletion to prevent race with AtomicUpdate
			entry.mu.Lock()
			// Double-check after acquiring lock
			if now-entry.lastAccess.Load() > ttlNanos {
				// Use LoadAndDelete to atomically check if we're the one deleting.
				// This prevents double-decrement if Delete() was called concurrently.
				if _, deleted := m.buckets.LoadAndDelete(key); deleted {
					m.keyCount.Add(-1)
				}
			}
			entry.mu.Unlock()
		}
		return true
	})
}

// AtomicUpdate atomically reads, modifies, and writes bucket state.
// It uses a per-key mutex to ensure thread safety.
func (m *MemoryStore) AtomicUpdate(ctx context.Context, key string, updateFn func(*BucketState) *BucketState) (*BucketState, error) {
	// Check if store is closed
	if m.closed.Load() {
		return nil, ErrStoreClosed
	}

	// Check context before proceeding
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Check key length limit
	if len(key) > m.maxKeyLength {
		return nil, ErrKeyTooLong
	}

	// Try to load existing entry first
	entryI, exists := m.buckets.Load(key)

	if !exists {
		// Pre-increment key count atomically to fix TOCTOU race.
		// If we exceed the limit, decrement and return error.
		newCount := m.keyCount.Add(1)
		if newCount > m.maxKeys {
			m.keyCount.Add(-1)
			return nil, ErrKeyLimitExceeded
		}

		// Create new entry
		newEntry := &bucketEntry{}
		newEntry.lastAccess.Store(time.Now().UnixNano())

		// Try to store, handle race condition where another goroutine
		// stored the same key between our Load and LoadOrStore
		entryI, exists = m.buckets.LoadOrStore(key, newEntry)
		if exists {
			// Another goroutine already created this entry, decrement our count
			m.keyCount.Add(-1)
		}
	}

	entry := entryI.(*bucketEntry) //nolint:errcheck // type assertion always succeeds

	// Lock this specific bucket for atomic update
	entry.mu.Lock()
	defer entry.mu.Unlock()

	// Check context again after acquiring lock
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Update last access time
	entry.lastAccess.Store(time.Now().UnixNano())

	// Apply the update function
	newState := updateFn(entry.state)

	// Only update if updateFn returned a non-nil state
	if newState != nil {
		entry.state = newState
	}

	return entry.state, nil
}

// Get retrieves the current bucket state without modification.
// Returns nil if the bucket doesn't exist.
func (m *MemoryStore) Get(ctx context.Context, key string) (*BucketState, error) {
	if m.closed.Load() {
		return nil, ErrStoreClosed
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Check key length limit
	if len(key) > m.maxKeyLength {
		return nil, ErrKeyTooLong
	}

	entryI, exists := m.buckets.Load(key)
	if !exists {
		return nil, nil
	}

	entry := entryI.(*bucketEntry) //nolint:errcheck // type assertion always succeeds

	entry.mu.Lock()
	defer entry.mu.Unlock()

	// Don't update lastAccess for read-only Get
	// This allows the entry to expire if only being monitored

	if entry.state == nil {
		return nil, nil
	}

	// Return a copy to prevent external mutation
	return &BucketState{
		Tokens:     entry.state.Tokens,
		LastRefill: entry.state.LastRefill,
	}, nil
}

// Delete removes a bucket from the store.
func (m *MemoryStore) Delete(ctx context.Context, key string) error {
	if m.closed.Load() {
		return ErrStoreClosed
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	// Check key length limit for consistency with other operations
	if len(key) > m.maxKeyLength {
		return ErrKeyTooLong
	}

	// Load entry first to acquire lock before deleting
	entryI, exists := m.buckets.Load(key)
	if !exists {
		return nil // Already deleted or never existed
	}

	entry := entryI.(*bucketEntry) //nolint:errcheck // type assertion always succeeds

	// Lock entry before deletion to prevent race with AtomicUpdate
	entry.mu.Lock()
	defer entry.mu.Unlock()

	// Use LoadAndDelete to atomically check if we're the one deleting.
	// This prevents double-decrement if cleanup() runs concurrently.
	if _, deleted := m.buckets.LoadAndDelete(key); deleted {
		m.keyCount.Add(-1)
	}

	return nil
}

// Close stops the cleanup goroutine and releases resources.
// After Close is called, the store should not be used.
func (m *MemoryStore) Close() error {
	if m.closed.Swap(true) {
		return nil // Already closed
	}

	close(m.done)
	m.wg.Wait()
	return nil
}

// Reset implements the Resetter interface.
// It removes all buckets from the store.
// This is useful for testing or resetting the rate limiter state.
//
// Note: Reset is not atomic with respect to concurrent operations.
// Entries added during Reset may or may not be removed.
// For deterministic clearing, ensure no concurrent operations are in progress.
func (m *MemoryStore) Reset(ctx context.Context) error {
	if m.closed.Load() {
		return ErrStoreClosed
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	// Collect keys first to avoid modifying during iteration
	var keysToDelete []interface{}
	m.buckets.Range(func(key, value interface{}) bool {
		keysToDelete = append(keysToDelete, key)
		return true
	})

	// Delete collected keys
	for _, key := range keysToDelete {
		// Check context periodically for large stores
		if err := ctx.Err(); err != nil {
			return err
		}

		entryI, exists := m.buckets.Load(key)
		if !exists {
			continue
		}
		entry := entryI.(*bucketEntry) //nolint:errcheck // type assertion always succeeds
		entry.mu.Lock()
		// Use LoadAndDelete to atomically check if we're the one deleting.
		// This prevents double-decrement if cleanup() or Delete() runs concurrently.
		if _, deleted := m.buckets.LoadAndDelete(key); deleted {
			m.keyCount.Add(-1)
		}
		entry.mu.Unlock()
	}

	return nil
}

// Clear removes all buckets from the store.
// Deprecated: Use Reset(ctx) instead for context support.
// This method is kept for backwards compatibility.
func (m *MemoryStore) Clear() {
	_ = m.Reset(context.Background()) //nolint:errcheck // backwards compatibility method
}

// BucketCount implements the BucketCounter interface.
// It returns the current number of active buckets in the store.
func (m *MemoryStore) BucketCount() int64 {
	return m.keyCount.Load()
}

// KeyCount returns the current number of keys in the store.
// Deprecated: Use BucketCount() instead.
func (m *MemoryStore) KeyCount() int64 {
	return m.BucketCount()
}

// HealthCheck verifies the store is operational.
// For MemoryStore, this always returns nil as in-memory operations don't fail.
func (m *MemoryStore) HealthCheck(ctx context.Context) error {
	if m.closed.Load() {
		return ErrStoreClosed
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	return nil
}

// Compile-time interface checks.
var (
	_ Store         = (*MemoryStore)(nil)
	_ HealthChecker = (*MemoryStore)(nil)
	_ Resetter      = (*MemoryStore)(nil)
	_ BucketCounter = (*MemoryStore)(nil)
)
