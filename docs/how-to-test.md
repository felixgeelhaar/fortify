# How to test resilience

Tests for resilience patterns need to verify behavior under fault conditions, not just happy paths. Fortify ships a `testing/` package with chaos utilities.

## Unit and integration tests

Run tests with race detection:

```bash
go test -v -race ./...
```

Coverage:

```bash
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

CI runs a 3-OS matrix (Linux, macOS, Windows) on Go 1.25 with race detection enabled.

## Chaos utilities

```go
import fortifytesting "github.com/felixgeelhaar/fortify/testing"

// Inject errors with configurable probability
errInjector := fortifytesting.NewErrorInjector(0.3, errors.New("service unavailable"))

// Add network latency
latencyInjector := fortifytesting.NewLatencyInjector(10*time.Millisecond, 50*time.Millisecond)

// Simulate timeouts
timeoutSim := fortifytesting.NewTimeoutSimulator(100*time.Millisecond, 0.5)

// Combine all of the above
flakey := fortifytesting.NewFlakeyService(0.3, 10*time.Millisecond, 30*time.Millisecond)
```

Wire one into your test:

```go
func TestCircuitBreakerTripsOnRepeatedFailure(t *testing.T) {
    cb := circuitbreaker.New[string](circuitbreaker.Config{
        ReadyToTrip: func(c circuitbreaker.Counts) bool { return c.ConsecutiveFailures >= 3 },
    })

    flakey := fortifytesting.NewFlakeyService(1.0, 0, 0) // always fails

    for i := 0; i < 5; i++ {
        cb.Execute(ctx, func(ctx context.Context) (string, error) {
            return flakey.Call(ctx)
        })
    }
    if cb.State() != circuitbreaker.StateOpen {
        t.Fatalf("want Open, got %v", cb.State())
    }
}
```

## Property-based tests

Pattern state machines are good candidates for property-based tests with [testing/quick](https://pkg.go.dev/testing/quick) or [pgregory.net/rapid](https://pkg.go.dev/pgregory.net/rapid). Invariants worth checking:

- After any sequence of `Execute` calls, `cb.State()` is one of `{Closed, Open, HalfOpen}`.
- `cb.Counts.Requests >= cb.Counts.TotalSuccesses + cb.Counts.TotalFailures`.
- `rl.BucketCount() <= MaxKeys`.

## Fuzz tests

Fortify includes one fuzz target (`retry/backoff_fuzz_test.go`) for backoff calculation. Run:

```bash
go test -fuzz=FuzzCalculateBackoff ./retry/
```

Add fuzz targets for any code that processes external input (key sanitization, error wrapping, etc.).

## Performance regression testing

Benchmarks live in `*_bench_test.go` files. Run all:

```bash
go test -bench=. -benchmem -run=^$ ./...
```

Run with the helper script:

```bash
./scripts/benchmark.sh run
./scripts/benchmark.sh generate-baseline
./scripts/benchmark.sh check
```

Thresholds: 10% time, 20% allocs, 15% memory. CI runs benchmarks weekly via `.github/workflows/performance.yml`.

See [PERFORMANCE_TESTING.md](PERFORMANCE_TESTING.md) for deeper guidance.

## Race detector

Always run pattern tests with `-race`. The CB lock-free fast path uses `sync/atomic` mirrors that the race detector validates against; running without `-race` will not catch ordering bugs.

```bash
go test -race -count=1 ./...
```

`-count=1` disables the test cache, ensuring tests actually re-run.
