# HTTP and gRPC integrations

Fortify ships ready-to-mount middleware for both transports.

## HTTP middleware (`fortify/http`)

Wraps `http.Handler` with each pattern. Compose by stacking decorators.

```go
import (
    "net/http"
    fortifyhttp "github.com/felixgeelhaar/fortify/http"
)

handler := fortifyhttp.RateLimit(rl, fortifyhttp.KeyFromIP)(
    fortifyhttp.Timeout(tm, 5*time.Second)(
        fortifyhttp.CircuitBreaker(cb)(
            http.HandlerFunc(myHandler),
        ),
    ),
)

http.Handle("/api", handler)
```

### Status codes

| Pattern         | Failure status                |
| --------------- | ----------------------------- |
| Circuit breaker | `503 Service Unavailable`     |
| Rate limit      | `429 Too Many Requests`       |
| Timeout         | `504 Gateway Timeout`         |

The CB middleware only writes 503 if the downstream handler hasn't already written a response (avoids double `WriteHeader`).

### Key extractors

For `RateLimit`, supply a `KeyExtractor`:

| Extractor                       | Source                                             |
| ------------------------------- | -------------------------------------------------- |
| `KeyFromIP`                     | `r.RemoteAddr`, parsed and IPv6-zone-stripped      |
| `KeyFromHeader(name)`           | header value, sanitized + truncated to 256 bytes   |
| `KeyFromHeaderWithMaxLen(name, n)` | header value, sanitized + truncated to `n`     |

`KeyFromHeader` panics at construction time on invalid header names — fail-fast for misconfiguration. If you build header names from runtime config, validate them yourself first.

### Caveats

- The `Timeout` middleware does not run the handler in a separate goroutine; non-cooperative handlers (those that don't check `ctx.Done()`) will not be cancellable. The deadline still propagates via the context.
- `responseRecorder` reads `statusCode` under a mutex to defend against handler-spawned goroutines, but cannot truly serialize concurrent body writes from a misbehaving handler.

## gRPC interceptors (`fortify/grpc`)

```go
import (
    fortifygrpc "github.com/felixgeelhaar/fortify/grpc"
    "google.golang.org/grpc"
)

server := grpc.NewServer(
    grpc.ChainUnaryInterceptor(
        fortifygrpc.UnaryCircuitBreakerInterceptor(cb),
        fortifygrpc.UnaryRateLimitInterceptor(rl, fortifygrpc.KeyFromMetadata("x-api-key")),
        fortifygrpc.UnaryTimeoutInterceptor(tm, 5*time.Second),
    ),
    grpc.ChainStreamInterceptor(
        fortifygrpc.StreamCircuitBreakerInterceptor(cb),
        fortifygrpc.StreamRateLimitInterceptor(rl, fortifygrpc.StreamKeyFromMetadata("x-api-key")),
        fortifygrpc.StreamTimeoutInterceptor(tm, 5*time.Second),
    ),
)
```

### Status codes

| Pattern         | gRPC status code                |
| --------------- | ------------------------------- |
| Circuit breaker | `Unavailable`                   |
| Rate limit      | `ResourceExhausted`             |
| Timeout         | `DeadlineExceeded`              |

### Metadata keys

`KeyFromMetadata` and `StreamKeyFromMetadata` sanitize and truncate (256-byte cap) before passing the value to the rate limiter. This protects the underlying Store from `ErrKeyTooLong` errors triggered by adversarial metadata, and bounds metric label cardinality if the key is plumbed downstream.
