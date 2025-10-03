# Error Handling Guide

This guide explains error handling patterns and best practices when using Fortify resilience patterns.

## Table of Contents

- [Overview](#overview)
- [Standard Errors](#standard-errors)
- [Error Classification](#error-classification)
- [Pattern-Specific Errors](#pattern-specific-errors)
- [Error Wrapping](#error-wrapping)
- [Best Practices](#best-practices)

## Overview

Fortify provides consistent error handling across all resilience patterns:

1. **Sentinel Errors**: Standard errors for common failure modes
2. **Error Classification**: Determine which errors should trigger resilience behavior
3. **Context Preservation**: Maintain error context through the stack
4. **Type Safety**: Use Go generics for type-safe error handling

## Standard Errors

All Fortify standard errors are defined in the `errors` package:

```go
import fortifyerrors "github.com/felixgeelhaar/fortify/errors"
```

### Circuit Breaker Errors

```go
// ErrCircuitOpen indicates the circuit breaker is in the open state
fortifyerrors.ErrCircuitOpen

// ErrTooManyRequests indicates too many requests in half-open state
fortifyerrors.ErrTooManyRequests
```

### Rate Limiter Errors

```go
// ErrRateLimitExceeded indicates rate limit has been exceeded
fortifyerrors.ErrRateLimitExceeded
```

### Bulkhead Errors

```go
// ErrBulkheadFull indicates the bulkhead is at capacity
fortifyerrors.ErrBulkheadFull
```

### Fallback Errors

```go
// ErrFallbackFailed indicates both primary and fallback operations failed
fortifyerrors.ErrFallbackFailed
```

### Usage Example

```go
result, err := cb.Execute(ctx, operation)
if err != nil {
    switch {
    case errors.Is(err, fortifyerrors.ErrCircuitOpen):
        // Circuit is open - fail fast
        return nil, fmt.Errorf("service unavailable: %w", err)

    case errors.Is(err, context.DeadlineExceeded):
        // Timeout occurred
        return nil, fmt.Errorf("operation timed out: %w", err)

    default:
        // Other error
        return nil, fmt.Errorf("operation failed: %w", err)
    }
}
```

## Error Classification

### Retry Error Classification

Determine which errors should trigger retries:

```go
// Method 1: Error Lists
r := retry.New[T](&retry.Config{
    // Retry these specific errors
    RetryableErrors: []error{
        ErrTemporary,
        ErrRateLimited,
        ErrNetworkTimeout,
    },

    // Never retry these errors
    NonRetryableErrors: []error{
        ErrInvalidInput,
        ErrUnauthorized,
        ErrNotFound,
    },
})

// Method 2: Custom Function
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        // Retry on temporary or timeout errors
        if errors.Is(err, ErrTemporary) {
            return true
        }
        if errors.Is(err, context.DeadlineExceeded) {
            return true
        }

        // Check for retryable HTTP status codes
        var httpErr *HTTPError
        if errors.As(err, &httpErr) {
            return httpErr.StatusCode >= 500
        }

        return false
    },
})
```

### Circuit Breaker Success Classification

Determine what constitutes a successful request:

```go
cb := circuitbreaker.New[T](circuitbreaker.Config{
    IsSuccessful: func(err error) bool {
        // Success if no error
        if err == nil {
            return true
        }

        // Don't count client errors as failures
        var httpErr *HTTPError
        if errors.As(err, &httpErr) {
            return httpErr.StatusCode < 500
        }

        // Ignore context cancellation
        if errors.Is(err, context.Canceled) {
            return true
        }

        return false
    },
})
```

## Pattern-Specific Errors

### Circuit Breaker

**Error Flow:**

1. **Closed State**: Returns operation error or `nil`
2. **Open State**: Returns `ErrCircuitOpen` immediately
3. **Half-Open State**: Returns operation error, `ErrCircuitOpen`, or `ErrTooManyRequests`

**Example:**

```go
result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    resp, err := apiClient.Call(ctx)
    if err != nil {
        return nil, fmt.Errorf("api call failed: %w", err)
    }
    return resp, nil
})

if err != nil {
    if errors.Is(err, fortifyerrors.ErrCircuitOpen) {
        // Circuit is open - use fallback or cached response
        return cachedResponse, nil
    }
    return nil, err
}
```

### Retry

**Error Flow:**

1. **All Retries Exhausted**: Returns last error from operation
2. **Context Cancelled**: Returns `context.Canceled` or `context.DeadlineExceeded`
3. **Non-Retryable Error**: Returns error immediately without retrying

**Example:**

```go
result, err := r.Do(ctx, func(ctx context.Context) (*Response, error) {
    resp, err := apiClient.Call(ctx)
    if err != nil {
        // Wrap error with context
        return nil, fmt.Errorf("attempt failed: %w", err)
    }
    return resp, nil
})

if err != nil {
    // Check if all retries were exhausted
    var maxAttemptsErr *retry.MaxAttemptsError
    if errors.As(err, &maxAttemptsErr) {
        logger.Error("all retry attempts failed",
            slog.Int("attempts", maxAttemptsErr.Attempts),
            slog.String("error", err.Error()),
        )
        return nil, fmt.Errorf("operation failed after %d attempts: %w",
            maxAttemptsErr.Attempts, err)
    }
    return nil, err
}
```

### Rate Limiter

**Error Flow:**

1. **Allow**: Returns `true` or `false` (no error)
2. **Wait**: Blocks until token available or context cancelled
3. **Take**: Returns `true` or `false` for n tokens

**Example:**

```go
// Non-blocking check
if !rl.Allow(ctx, userID) {
    return nil, fortifyerrors.ErrRateLimitExceeded
}

// Blocking wait
if err := rl.Wait(ctx, userID); err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        return nil, fortifyerrors.ErrRateLimitExceeded
    }
    return nil, fmt.Errorf("rate limit wait failed: %w", err)
}

// Multi-token operation
if !rl.Take(ctx, userID, 5) {
    return nil, fmt.Errorf("insufficient rate limit quota")
}
```

### Timeout

**Error Flow:**

1. **Within Timeout**: Returns operation result/error
2. **Timeout Exceeded**: Returns `context.DeadlineExceeded`
3. **Parent Context Cancelled**: Returns `context.Canceled`

**Example:**

```go
result, err := tm.Execute(ctx, 5*time.Second, func(ctx context.Context) (*Response, error) {
    // Ensure operation respects context
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    case result := <-makeRequest(ctx):
        return result, nil
    }
})

if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        logger.Warn("operation timed out after 5s")
        return nil, fmt.Errorf("timeout: %w", err)
    }
    return nil, err
}
```

### Bulkhead

**Error Flow:**

1. **Capacity Available**: Executes and returns operation result/error
2. **At Capacity, Queue Available**: Queues and returns operation result/error
3. **Queue Full/Timeout**: Returns `ErrBulkheadFull`
4. **After Close()**: Returns `ErrBulkheadFull`

**Example:**

```go
result, err := bh.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return processRequest(ctx)
})

if err != nil {
    if errors.Is(err, fortifyerrors.ErrBulkheadFull) {
        logger.Warn("bulkhead at capacity",
            slog.String("operation", "process_request"),
        )
        return nil, fmt.Errorf("service overloaded: %w", err)
    }
    return nil, err
}
```

## Error Wrapping

### Preserving Error Context

Always wrap errors to preserve context:

```go
result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    resp, err := apiClient.Call(ctx)
    if err != nil {
        // Wrap with context about what failed
        return nil, fmt.Errorf("API call to %s failed: %w", apiClient.BaseURL, err)
    }

    if resp.StatusCode != 200 {
        // Create contextual error
        return nil, fmt.Errorf("API returned status %d: %s",
            resp.StatusCode, resp.Status)
    }

    return resp, nil
})
```

### Custom Error Types

Define domain-specific errors:

```go
// Custom error type
type ServiceError struct {
    Service   string
    Operation string
    Err       error
    Retryable bool
}

func (e *ServiceError) Error() string {
    return fmt.Sprintf("%s.%s failed: %v", e.Service, e.Operation, e.Err)
}

func (e *ServiceError) Unwrap() error {
    return e.Err
}

// Usage with retry
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        var serviceErr *ServiceError
        if errors.As(err, &serviceErr) {
            return serviceErr.Retryable
        }
        return false
    },
})
```

### Error Chains

Use error chains for detailed context:

```go
result, err := middleware.New[*Response]().
    WithCircuitBreaker(cb).
    WithRetry(r).
    WithTimeout(tm, 10*time.Second).
    Execute(ctx, func(ctx context.Context) (*Response, error) {
        resp, err := apiClient.Call(ctx)
        if err != nil {
            return nil, &ServiceError{
                Service:   "api-client",
                Operation: "Call",
                Err:       err,
                Retryable: isRetryable(err),
            }
        }
        return resp, nil
    })

// Error chain might be:
// - timeout.Execute error
// - retry.Do error (after 3 attempts)
// - circuitbreaker.Execute error
// - ServiceError
// - original HTTP error
```

## Best Practices

### 1. Use Sentinel Errors

```go
// Good: Use standard errors for comparison
if errors.Is(err, fortifyerrors.ErrCircuitOpen) {
    // Handle circuit open
}

// Bad: String comparison
if err.Error() == "circuit breaker is open" {
    // Fragile and error-prone
}
```

### 2. Wrap Errors with Context

```go
// Good: Preserve error chain with context
return nil, fmt.Errorf("failed to process order %s: %w", orderID, err)

// Bad: Lose error chain
return nil, fmt.Errorf("failed to process order %s: %v", orderID, err)
```

### 3. Classify Errors Correctly

```go
// Good: Specific error classification
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        // Only retry specific transient errors
        return errors.Is(err, ErrNetworkTimeout) ||
               errors.Is(err, ErrServiceUnavailable)
    },
})

// Bad: Retry everything
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        return true  // Will retry permanent failures
    },
})
```

### 4. Log Errors Appropriately

```go
// Good: Structured logging with context
result, err := cb.Execute(ctx, operation)
if err != nil {
    logger.Error("circuit breaker execution failed",
        slog.String("pattern", "circuit_breaker"),
        slog.String("operation", "api_call"),
        slog.Any("error", err),
        slog.String("user_id", userID),
    )
    return nil, err
}

// Bad: Print to stdout
if err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

### 5. Handle Context Cancellation

```go
// Good: Check context before operation
result, err := tm.Execute(ctx, timeout, func(ctx context.Context) (*Response, error) {
    // Respect context cancellation
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        return operation(ctx)
    }
})

// Bad: Ignore context
result, err := tm.Execute(ctx, timeout, func(ctx context.Context) (*Response, error) {
    // This will continue even if context is cancelled
    return longRunningOperation()
})
```

### 6. Fail Fast for Permanent Errors

```go
// Good: Don't retry permanent errors
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        // Don't retry 4xx client errors
        var httpErr *HTTPError
        if errors.As(err, &httpErr) {
            return httpErr.StatusCode >= 500
        }
        return true
    },
})

// Bad: Retry everything including permanent errors
r := retry.New[T](&retry.Config{
    IsRetryable: func(err error) bool {
        return err != nil  // Will retry 404, 401, etc.
    },
})
```

### 7. Use Fallbacks for Non-Critical Operations

```go
// Good: Graceful degradation with fallback
fb := fallback.New[*Response](fallback.Config{
    Fallback: func(ctx context.Context, err error) (*Response, error) {
        // Return cached data or degraded response
        if cached := cache.Get(ctx, key); cached != nil {
            return cached, nil
        }
        return defaultResponse, nil
    },
})

result, err := fb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return apiClient.Call(ctx)
})
// Always returns a response, even on error
```

### 8. Monitor Error Rates

```go
// Track error types for monitoring
cb := circuitbreaker.New[T](circuitbreaker.Config{
    OnStateChange: func(from, to circuitbreaker.State) {
        errorRateMetric.WithLabelValues(to.String()).Inc()
    },
})

r := retry.New[T](&retry.Config{
    OnRetry: func(attempt int, err error) {
        retryMetric.WithLabelValues(getErrorType(err)).Inc()
    },
})
```

### 9. Document Error Behavior

```go
// Document what errors callers should expect
// processOrder processes an order with resilience patterns.
// Returns:
//   - fortifyerrors.ErrCircuitOpen if the payment service circuit is open
//   - fortifyerrors.ErrRateLimitExceeded if the rate limit is exceeded
//   - fortifyerrors.ErrBulkheadFull if the processing queue is full
//   - context.DeadlineExceeded if the operation times out
//   - Other errors from the underlying operation
func processOrder(ctx context.Context, order *Order) (*Receipt, error) {
    // Implementation
}
```

### 10. Test Error Paths

```go
func TestCircuitBreaker_ErrorHandling(t *testing.T) {
    t.Run("returns circuit open error", func(t *testing.T) {
        cb := circuitbreaker.New[string](circuitbreaker.Config{
            MaxRequests: 1,
            ReadyToTrip: func(counts circuitbreaker.Counts) bool {
                return counts.ConsecutiveFailures >= 1
            },
        })

        // Trigger circuit to open
        _, _ = cb.Execute(ctx, func(ctx context.Context) (string, error) {
            return "", errors.New("failure")
        })

        // Verify circuit open error
        _, err := cb.Execute(ctx, func(ctx context.Context) (string, error) {
            return "success", nil
        })

        if !errors.Is(err, fortifyerrors.ErrCircuitOpen) {
            t.Errorf("expected ErrCircuitOpen, got %v", err)
        }
    })
}
```

## Error Handling Patterns

### Pattern: Retry with Exponential Backoff

```go
r := retry.New[*Response](&retry.Config{
    MaxAttempts:   5,
    InitialDelay:  100 * time.Millisecond,
    MaxDelay:      10 * time.Second,
    Multiplier:    2.0,
    BackoffPolicy: retry.BackoffExponential,
    Jitter:        true,
    IsRetryable: func(err error) bool {
        return errors.Is(err, ErrTemporary) ||
               errors.Is(err, context.DeadlineExceeded)
    },
    OnRetry: func(attempt int, err error) {
        logger.Warn("retrying operation",
            slog.Int("attempt", attempt),
            slog.String("error", err.Error()),
        )
    },
})
```

### Pattern: Circuit Breaker with Fallback

```go
cb := circuitbreaker.New[*Response](circuitbreaker.Config{
    MaxRequests: 3,
    Timeout:     30 * time.Second,
    ReadyToTrip: func(counts circuitbreaker.Counts) bool {
        return counts.ConsecutiveFailures >= 5
    },
})

result, err := cb.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return primaryService.Call(ctx)
})

if err != nil {
    if errors.Is(err, fortifyerrors.ErrCircuitOpen) {
        // Use fallback when circuit is open
        return fallbackService.Call(ctx)
    }
    return nil, err
}
```

### Pattern: Timeout with Cleanup

```go
tm := timeout.New[*Response](timeout.Config{
    DefaultTimeout: 5 * time.Second,
})

result, err := tm.Execute(ctx, 10*time.Second, func(ctx context.Context) (*Response, error) {
    // Start operation
    op := startOperation(ctx)

    // Wait for completion or timeout
    select {
    case result := <-op.Done():
        return result, nil
    case <-ctx.Done():
        // Clean up on timeout
        op.Cancel()
        return nil, ctx.Err()
    }
})
```

### Pattern: Bulkhead with Graceful Overload

```go
bh := bulkhead.New[*Response](bulkhead.Config{
    MaxConcurrent: 100,
    MaxQueue:      500,
    QueueTimeout:  30 * time.Second,
})

result, err := bh.Execute(ctx, func(ctx context.Context) (*Response, error) {
    return processRequest(ctx)
})

if err != nil {
    if errors.Is(err, fortifyerrors.ErrBulkheadFull) {
        // Return degraded response instead of error
        return &Response{
            Status:  http.StatusServiceUnavailable,
            Message: "Service temporarily overloaded, please retry",
        }, nil
    }
    return nil, err
}
```

## Debugging Error Flows

### Enable Verbose Logging

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
}))

cb := circuitbreaker.New[T](circuitbreaker.Config{
    Logger: logger,
    OnStateChange: func(from, to circuitbreaker.State) {
        logger.Debug("state change",
            slog.String("from", from.String()),
            slog.String("to", to.String()),
        )
    },
})
```

### Error Stack Traces

```go
import "github.com/pkg/errors"

result, err := operation(ctx)
if err != nil {
    // Add stack trace
    err = errors.WithStack(err)
    logger.Error("operation failed",
        slog.Any("error", err),
        slog.String("stack", fmt.Sprintf("%+v", err)),
    )
    return nil, err
}
```

## Additional Resources

- [Production Deployment Guide](./PRODUCTION.md)
- [API Documentation](https://pkg.go.dev/github.com/felixgeelhaar/fortify)
- [Examples](../examples/)
- [Standard Errors Package](../errors/)
