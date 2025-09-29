# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2025-01-XX

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

[Unreleased]: https://github.com/felixgeelhaar/fortify/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/felixgeelhaar/fortify/releases/tag/v0.1.0