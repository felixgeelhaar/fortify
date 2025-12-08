# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

#### Pluggable Storage Interface
- **Store interface** for custom storage backends (`ratelimit/store.go`)
  - `AtomicUpdate` method for atomic read-modify-write operations
  - `Get` method for read-only state access without side effects
  - `Delete` and `Close` methods for resource management
  - `BucketState` struct for serializable token bucket state
- **MemoryStore** default in-memory implementation (`ratelimit/memory.go`)
  - Uses `sync.Map` with per-key mutex for thread-safe atomicity
  - Zero-configuration default for single-instance deployments
  - TTL-based automatic cleanup of stale entries (default: 1 hour)
  - Configurable maximum key limit (default: 100,000) for memory protection
  - Configurable maximum key length (default: 1,024) to prevent memory exhaustion
  - Functional options: `WithMaxKeys()`, `WithCleanupInterval()`, `WithEntryTTL()`, `WithMaxKeyLength()`

#### Configuration Options
- **Store** field in Config for custom storage backends
- **FailOpen** field for configurable failure behavior
  - `FailOpen: false` (default) - deny requests when storage fails (consistency)
  - `FailOpen: true` - allow requests when storage fails (availability)
- **MaxTokensPerRequest** field (default: Burst × 10) to prevent DoS via excessive token requests
- **Metrics** field for observability integration

#### Interfaces
- **HealthChecker** interface for distributed store health monitoring
  - `HealthCheck(ctx context.Context) error` method
  - MemoryStore implements HealthChecker by default
- **Metrics** interface for observability hooks (all methods now receive context)
  - `OnAllow(ctx, key)` - called when request is allowed
  - `OnDeny(ctx, key)` - called when request is denied
  - `OnError(ctx, key, err)` - called on storage errors
  - `OnStoreLatency(ctx, operation, duration)` - storage latency tracking
- **RateLimiter.HealthCheck()** method for health monitoring
- **RateLimiter.Close()** method for proper resource cleanup
- **RateLimiter closed flag** - Prevents operations after Close() is called

#### Error Handling
- **Sentinel errors** for better error handling (`ratelimit/errors.go`)
  - `ErrLimitExceeded` - rate limit exceeded
  - `ErrStorageUnavailable` - storage backend unavailable
  - `ErrInvalidTokenCount` - invalid token count in Take()
  - `ErrExcessiveTokens` - token request exceeds MaxTokensPerRequest
  - `ErrStoreClosed` - operation attempted on closed store
  - `ErrKeyLimitExceeded` - maximum key limit reached
  - `ErrKeyTooLong` - key exceeds maximum length
  - `ErrWaitTimeout` - Wait() exceeded maximum iterations or time limit
  - `ErrRateLimiterClosed` - operation attempted on closed rate limiter

#### Documentation
- **Package documentation** moved to `ratelimit/doc.go` per Go conventions
- Comprehensive usage examples in doc.go
- **Store interface stability guarantee** - Documented interface stability and extension patterns
- **KeyFunc key length constraints** - Documented that keys must respect Store's maximum length

### Changed
- Rate limiter now uses pluggable `Store` interface instead of internal `sync.Map`
- Token bucket algorithm logic moved from separate file to `limiter.go`
- In-memory storage now uses `Store` interface (unified architecture)
- `calculateWaitTime()` now uses `Get()` for read-only access (no side effects)
- `refill()` optimized to avoid unnecessary allocations when state unchanged
- **`OnLimit` callback** now receives `context.Context` as first argument for observability
- **Magic numbers extracted** to named constants for maintainability
- **Config validation** now caps Rate, Burst, Interval, and MaxTokensPerRequest to maximum values
- **Metrics interface** now receives context in all methods for distributed tracing
- **KeyFunc** documentation clarified: when set, it takes precedence and the key parameter is ignored
- **Wait()** now has iteration and time limits (10,000 iterations / 5 minutes) to prevent infinite loops

### Fixed
- **Timer leak in Wait()** - Now uses `time.NewTimer` with proper `Stop()` and channel draining
- **Timer reuse optimization** - Wait() now reuses timer with Reset() to reduce allocations
- **Delete() race condition** - Entry lock acquired before deletion to prevent races with AtomicUpdate
- **Delete() key length check** - Now validates key length like other operations
- **Input validation in Take()** - Rejects zero, negative, and excessive token requests
- **calculateWaitTime() side effect** - No longer modifies state via AtomicUpdate
- **TOCTOU race in key creation** - Fixed using pre-increment approach with proper rollback
- **Clear() race condition** - Fixed with two-phase delete (collect then delete)
- **Panic recovery** - Now includes stack trace for debugging
- **Closed rate limiter operations** - Allow/Wait/Take/HealthCheck now check closed flag before proceeding

### Removed

#### Breaking Changes
- **Removed `backends/redis/` module** - Users should implement their own Redis adapter
- **Removed `examples/backends/redis/`** - Docker Compose example removed
- **Removed `ratelimit/tokenbucket.go`** - Logic consolidated into limiter

### Migration
- **In-memory users**: No changes required - works exactly as before
- **Redis users**: Implement custom `Store` adapter (see `docs/MIGRATION_REDIS.md` for examples)

### Documentation
- Updated `docs/MIGRATION_REDIS.md` to be a comprehensive custom storage backend guide
  - Redis implementation example with WATCH/MULTI/EXEC
  - DynamoDB implementation example with conditional writes
  - Testing strategies for custom stores
  - Production deployment best practices

## [1.1.0] - 2025-10-19

### Added

#### Distributed Rate Limiting
- **Redis backend for distributed rate limiting** (`backends/redis`)
  - Separate Go module maintains zero-dependency promise for core
  - Atomic operations via Lua scripts (production-grade, no race conditions)
  - Same `RateLimiter` interface as in-memory (drop-in replacement)
  - Support for Redis Cluster, Redis Sentinel, and standalone Redis
  - Configurable fail-open or fail-closed behavior
  - Automatic bucket expiration with configurable TTL
  - Full observability integration (slog, OpenTelemetry)
  - Comprehensive test suite with miniredis (>90% coverage)
  - Benchmark tests comparing Redis vs in-memory performance

#### Documentation
- **Redis backend README** with installation, configuration, and usage examples
- **Migration guide** (`docs/MIGRATION_REDIS.md`) for moving from in-memory to Redis
- **Complete working example** with Docker Compose demonstrating distributed rate limiting
- Added distributed rate limiting section to main README

#### Examples
- Multi-instance distributed rate limiting example (`examples/backends/redis`)
  - Docker Compose setup with 3 application instances + Redis
  - Makefile for easy testing and demonstration
  - HTTP API with rate limiting across instances
  - Health checks and observability

## [1.0.0] - 2025-10-03

### Added

#### Production Readiness
- **Root package documentation** (doc.go) for improved discoverability
- Comprehensive package-level documentation across all patterns
- Production deployment guide with best practices
- Error handling patterns guide
- Migration guide from other libraries

#### API Improvements
- **Consistent constructor signatures** - All patterns now use Config by value
- **Clean error package naming** - Renamed to `ferrors` to avoid stdlib conflicts
- Zero import aliasing required throughout codebase

#### Critical Bug Fixes
- **Fixed timeout callback goroutine leak** - Callbacks now execute synchronously
- **Fixed bulkhead worker lifecycle** - Clear shutdown semantics with Close()
- **Fixed token bucket edge cases** - Overflow protection and division-by-zero handling
- **Documented concurrency semantics** - Clear thread-safety guarantees

#### Quality Assurance
- Zero goroutine leaks verified with race detector
- Zero race conditions across all packages
- 80%+ test coverage on all packages
- All critical production issues resolved

### Changed
- Constructor signature: `retry.New()` now takes `Config` by value (was pointer)
- Package rename: `errors` → `ferrors` for cleaner imports
- Timeout callbacks execute synchronously to prevent leaks
- Bulkhead `Close()` signals shutdown but doesn't wait (user coordinates)

### Fixed
- Timeout goroutine leak from unbounded callback spawning
- Bulkhead WaitGroup race condition between Execute() and Close()
- Token bucket waitTime() race condition (documented as best-effort)
- Token bucket refill overflow from clock skew or system sleep
- Division-by-zero edge cases in token bucket calculations

### Documentation
- Added root doc.go for pkg.go.dev integration
- V1_RELEASE_SUMMARY.md - Complete v1.0 overview
- V1_FINAL_STATUS.md - Production readiness certification
- Enhanced inline documentation for all edge cases
- Clear resource cleanup semantics

## [0.1.1] - 2025-10-01

### Fixed
- Minor bug fixes and documentation improvements

## [0.1.0] - 2025-10-01

### Added

#### Core Patterns
- **Circuit Breaker** pattern with three states (Closed, Open, Half-Open)
  - Configurable failure thresholds and recovery timeouts
  - State change callbacks for monitoring
  - Request counting with sliding window
  - Type-safe with Go generics

- **Retry** pattern with intelligent backoff strategies
  - Exponential, linear, and constant backoff policies
  - Configurable jitter to prevent thundering herd
  - Custom retry predicates for error handling
  - Maximum attempt and delay limits

- **Rate Limiter** using token bucket algorithm
  - Per-key rate limiting
  - Configurable rate, burst, and interval
  - Both blocking (Wait) and non-blocking (Allow) operations
  - Context-aware cancellation

- **Timeout** pattern with context propagation
  - Configurable default and per-operation timeouts
  - Automatic context deadline enforcement
  - Timeout callbacks for monitoring

- **Bulkhead** pattern for concurrency limiting
  - Semaphore-based concurrency control
  - Queue management with configurable capacity
  - Dual semaphore design preventing race conditions
  - Request statistics tracking

#### Integration
- **Middleware Chain** for composable pattern orchestration
  - Fluent API for building resilience chains
  - Right-to-left composition order
  - Support for all core patterns
  - Type-safe execution

- **HTTP Middleware** for standard http.Handler
  - CircuitBreaker, RateLimit, and Timeout middleware
  - Thread-safe response recording
  - Appropriate HTTP status codes (503, 429, 504)
  - Flexible key extraction (IP, headers)

- **gRPC Interceptors** for unary and streaming RPCs
  - Unary and stream interceptors for all patterns
  - Metadata-based key extraction
  - Standard gRPC status codes
  - Context propagation support

#### Observability
- **Structured Logging** with log/slog integration
  - Pattern-specific log helpers
  - Event and metric logging functions
  - Configurable log levels

- **OpenTelemetry Tracing** support
  - Distributed tracing integration
  - Pattern-specific span creation
  - Attribute and event recording
  - Error tracking

#### Documentation
- Comprehensive README with usage examples
- API documentation with inline comments
- Example programs for all patterns
- Performance benchmarks and metrics

### Performance
- <1µs overhead for all patterns
- Zero allocations in hot paths
- Concurrent execution with race-free guarantees
- Optimized for production workloads

### Testing
- >95% test coverage across all packages
- Race detection on all tests
- Comprehensive benchmarks
- Fuzz testing for backoff algorithms

[Unreleased]: https://github.com/felixgeelhaar/fortify/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/felixgeelhaar/fortify/releases/tag/v1.0.0
[0.1.1]: https://github.com/felixgeelhaar/fortify/releases/tag/v0.1.1
[0.1.0]: https://github.com/felixgeelhaar/fortify/releases/tag/v0.1.0