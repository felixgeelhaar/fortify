# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Fortify is a production-ready Go resilience package providing circuit breakers, retries, rate limiting, timeouts, and bulkheads. The package is designed for Go 1.23+ with zero external dependencies (core), full observability support, and modern Go best practices.

## Architecture

### Package Organization

The codebase follows a modular package structure with separate concerns:

- **`circuitbreaker/`**: State machine implementation (Closed/Open/Half-Open) with configurable failure thresholds
- **`retry/`**: Intelligent retry logic with multiple backoff strategies (exponential, linear, constant) and error classification
- **`ratelimit/`**: Token bucket algorithm for per-key rate limiting with wait/reject modes
- **`timeout/`**: Context-based timeout enforcement with deadline propagation
- **`bulkhead/`**: Semaphore-based concurrency limiting with queue support
- **`middleware/`**: Composable patterns for HTTP (`net/http`) and gRPC with chaining support
- **`otel/`**: OpenTelemetry W3C TraceContext integration with automatic span creation
- **`slog/`**: Structured logging handlers with trace context enrichment
- **`errors/`**: Standard error types and retryable error interface

### Design Patterns

**Generics-First**: All patterns use Go 1.23+ generics for type safety:
```go
CircuitBreaker[T any].Execute(ctx context.Context, fn func(context.Context) (T, error)) (T, error)
```

**Context Propagation**: All operations accept `context.Context` as first parameter for cancellation, deadlines, and trace propagation.

**Middleware Composition**: Patterns can be composed using the middleware chain pattern:
```go
handler := middleware.Chain(
    circuitbreaker.AsMiddleware(cb),
    retry.AsMiddleware(retry),
    timeout.AsMiddleware(timeout),
)(operation)
```

**Functional Options**: Configuration uses both builder and functional options patterns for flexibility.

## Development Commands

### Testing

```bash
# Run all tests with race detection
go test -race ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. -benchmem ./...

# Run fuzz tests (Go 1.18+)
go test -fuzz=Fuzz -fuzztime=30s ./...

# Run specific package tests
go test ./circuitbreaker/
```

### Code Quality

```bash
# Format code
gofmt -w .

# Run linter (requires golangci-lint)
golangci-lint run

# Check for race conditions
go test -race ./...

# Generate test coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Building

```bash
# Build all packages
go build ./...

# Verify all dependencies
go mod verify

# Tidy dependencies
go mod tidy
```

## Key Implementation Details

### State Management (Circuit Breaker)

The circuit breaker maintains thread-safe state transitions using atomic operations:
- **Closed → Open**: When `ReadyToTrip` function returns true (typically based on consecutive failures)
- **Open → Half-Open**: After timeout period expires
- **Half-Open → Closed**: When `MaxRequests` succeed
- **Half-Open → Open**: On any failure during half-open state

### Backoff Strategies (Retry)

Implement backoff calculation in `retry/backoff.go`:
- **Exponential**: `delay = initialDelay * multiplier^attempt`
- **Linear**: `delay = initialDelay * attempt`
- **Constant**: `delay = initialDelay`
- **Jitter**: Add random jitter to prevent thundering herd: `delay += random(0, delay * 0.1)`

### Token Bucket Algorithm (Rate Limiter)

Token bucket implementation in `ratelimit/tokenbucket.go`:
- Refill tokens at `Rate` per `Interval`
- Allow burst up to `Burst` tokens
- Use mutex for thread-safe token management
- Support per-key buckets using sync.Map

### Observability Integration

**OpenTelemetry Spans**: Create spans for all operations with attributes:
```go
span.SetAttributes(
    attribute.String("pattern", "circuit_breaker"),
    attribute.String("state", state.String()),
)
```

**Structured Logging**: Use `log/slog` with context enrichment:
```go
logger.InfoContext(ctx, "operation",
    slog.String("pattern", "retry"),
    slog.Int("attempt", attempt),
)
```

## Testing Requirements

### Unit Tests

- Test all state transitions (circuit breaker)
- Verify backoff calculations (retry)
- Validate token bucket refill logic (rate limiter)
- Test context cancellation handling
- Verify thread safety with race detector

### Benchmark Tests

Target performance metrics:
- Overhead < 1µs for fast-path operations
- Zero allocations in hot paths
- Throughput > 1M ops/sec

### Fuzz Tests

Required for:
- Backoff calculation edge cases
- Configuration validation
- Concurrent state transitions

## Error Handling

### Standard Errors

Use package-level sentinel errors:
```go
var (
    ErrCircuitOpen        = errors.New("circuit breaker is open")
    ErrRateLimitExceeded  = errors.New("rate limit exceeded")
    ErrBulkheadFull       = errors.New("bulkhead at capacity")
)
```

### Retryable Errors

Implement `RetryableError` interface:
```go
type RetryableError interface {
    error
    Retryable() bool
}
```

Classify errors using `errors.Is()` for comparison against configured retryable/non-retryable error lists.

## Configuration Patterns

Support both builder and functional options:

**Builder Pattern**:
```go
cb := circuitbreaker.NewBuilder[*Response]().
    WithMaxRequests(5).
    WithTimeout(60 * time.Second).
    Build()
```

**Functional Options**:
```go
cb := circuitbreaker.New[*Response](
    circuitbreaker.WithMaxRequests(5),
    circuitbreaker.WithTimeout(60 * time.Second),
)
```

## Critical Implementation Rules

1. **Zero External Dependencies**: Core packages must use only Go standard library
2. **Context-First**: All operations must accept and respect `context.Context`
3. **Thread-Safe**: Use atomic operations or mutexes for state management
4. **Generic Type Parameters**: Use generics for type safety, avoid `interface{}`
5. **Observability Optional**: Logging and tracing must be optional (nil-safe)
6. **Fast-Path Optimization**: Minimize allocations in success paths
7. **Error Classification**: Support both error type comparison and custom predicates

## Performance Considerations

- Use `atomic` package for counters and flags
- Minimize allocations with sync.Pool for frequently allocated objects
- Avoid defer in hot paths (< 1µs overhead requirement)
- Use pointer receivers for large structs
- Consider CPU cache locality for frequently accessed fields

## Documentation Standards

All exported types must have:
- Comprehensive godoc comments
- Usage examples in package documentation
- Performance characteristics documented
- Thread-safety guarantees specified

## Migration Path

When implementing patterns, ensure compatibility with popular alternatives:
- Circuit Breaker: Compatible with `sony/gobreaker` migration
- Retry: Similar API to `avast/retry-go`
- Rate Limiter: Compatible with `golang.org/x/time/rate`

Provide migration examples in package documentation.