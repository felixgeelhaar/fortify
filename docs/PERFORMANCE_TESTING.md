# Performance Regression Testing

This document describes Fortify's performance regression testing framework for tracking and validating benchmark performance over time.

## Overview

Fortify includes a comprehensive performance testing framework that:

- **Tracks Performance**: Records benchmark results with detailed metrics
- **Detects Regressions**: Automatically identifies performance degradation
- **Sets Baselines**: Establishes acceptable performance thresholds
- **Generates Reports**: Creates detailed performance analysis reports
- **CI/CD Integration**: Automated testing in GitHub Actions

## Quick Start

### Running Benchmarks

```bash
# Run all benchmarks
./scripts/benchmark.sh run

# Generate performance baseline from results
./scripts/benchmark.sh generate-baseline

# Check for regressions
./scripts/benchmark.sh check

# Complete workflow (run + check + compare)
./scripts/benchmark.sh all
```

### Manual Benchmark Execution

```bash
# Run benchmarks with standard duration
go test -bench=. -benchmem -benchtime=3s ./...

# Run extended benchmarks
go test -bench=. -benchmem -benchtime=10s -count=3 ./...

# Run specific package benchmarks
go test -bench=. -benchmem ./circuitbreaker
```

## Performance Tracking API

### Setting Up a Tracker

```go
import "github.com/felixgeelhaar/fortify/testing"

// Create performance tracker
tracker := testing.NewPerformanceTracker(".benchmark-results")

// Set custom thresholds
tracker.SetThresholds(testing.RegressionThresholds{
    TimeIncrease:  1.10, // 10% slower is acceptable
    AllocIncrease: 1.20, // 20% more allocations
    BytesIncrease: 1.15, // 15% more memory
})
```

### Adding Baselines

```go
// Manually add baseline
tracker.AddBaseline(testing.PerformanceBaseline{
    Name:        "BenchmarkCircuitBreaker",
    MaxNsPerOp:  1000,
    MaxAllocs:   5,
    MaxBytes:    512,
    Description: "Circuit breaker baseline",
})

// Generate from benchmark results (with 10% safety factor)
results := []testing.BenchmarkResult{...}
tracker.GenerateBaselineFromResults(results, 1.1)

// Save baselines to file
tracker.SaveBaselines("performance-baselines.json")

// Load baselines from file
tracker.LoadBaselines("performance-baselines.json")
```

### Checking for Regressions

```go
// Check current results against baselines
results := []testing.BenchmarkResult{
    {
        Name:        "BenchmarkCircuitBreaker",
        NsPerOp:     950,
        AllocsPerOp: 4,
        BytesPerOp:  480,
        Timestamp:   time.Now(),
    },
}

report := tracker.CheckRegressions(results)

fmt.Printf("Total checks: %d\n", report.TotalChecks)
fmt.Printf("Passed: %d\n", report.Passed)
fmt.Printf("Failed: %d\n", report.Failed)

// Handle regressions
for _, regression := range report.Regressions {
    fmt.Printf("❌ %s: %s increased by %.2f%% (threshold: %.2f%%)\n",
        regression.BenchmarkName,
        regression.Metric,
        regression.Increase,
        regression.Threshold)
}
```

### Saving Reports

```go
// Save benchmark report
report := testing.BenchmarkReport{
    Timestamp: time.Now(),
    Results:   results,
    Metadata: map[string]string{
        "commit": "abc123",
        "branch": "main",
    },
}

tracker.SaveReport(report)

// Load latest report
latest, err := tracker.LoadLatestReport()
```

### Comparing Reports

```go
baseline := testing.BenchmarkReport{...}
current := testing.BenchmarkReport{...}

changes := testing.CompareReports(baseline, current)

for benchmark, metrics := range changes {
    timeChange := metrics["time_change"]
    fmt.Printf("%s: %.2f%% time change\n", benchmark, timeChange)
}
```

## Benchmark Structure

### Benchmark Result Format

```go
type BenchmarkResult struct {
    Name           string    // Benchmark name
    NsPerOp        float64   // Nanoseconds per operation
    AllocsPerOp    uint64    // Allocations per operation
    BytesPerOp     uint64    // Bytes allocated per operation
    Timestamp      time.Time // When benchmark was run
    GitCommit      string    // Git commit hash
    GitBranch      string    // Git branch name
    GoVersion      string    // Go version
    OS             string    // Operating system
    Arch           string    // Architecture
    CPUModel       string    // CPU model
    MemoryTotal    uint64    // Total system memory
    IterationCount int       // Number of iterations
}
```

### Performance Baseline Format

```go
type PerformanceBaseline struct {
    Name        string  // Benchmark name
    MaxNsPerOp  float64 // Maximum acceptable ns/op
    MaxAllocs   uint64  // Maximum acceptable allocations
    MaxBytes    uint64  // Maximum acceptable bytes
    Description string  // Baseline description
}
```

## Regression Thresholds

### Default Thresholds

- **Time**: 10% increase (1.10x)
- **Allocations**: 20% increase (1.20x)
- **Memory**: 15% increase (1.15x)

### Custom Thresholds

```go
tracker.SetThresholds(testing.RegressionThresholds{
    TimeIncrease:       1.05, // 5% time increase
    AllocIncrease:      1.10, // 10% allocation increase
    BytesIncrease:      1.08, // 8% memory increase
    AbsoluteMaxNsPerOp: 10000, // Hard limit at 10µs
})
```

## CI/CD Integration

### GitHub Actions

Fortify includes automated performance testing in CI/CD:

```yaml
# .github/workflows/performance.yml
name: Performance Regression Testing

on:
  pull_request:
    branches: [ main ]
  push:
    branches: [ main ]
```

#### Features

1. **PR Checks**: Runs benchmarks on every pull request
2. **Baseline Comparison**: Compares against main branch
3. **Regression Detection**: Fails CI if regressions detected
4. **Performance Tracking**: Archives results for historical analysis
5. **Automated Comments**: Posts regression warnings on PRs

#### Workflow Steps

1. Run benchmarks on PR code
2. Checkout main branch
3. Run benchmarks on main
4. Compare results using `benchstat`
5. Check against baselines
6. Post results to PR

### Local CI Simulation

```bash
# Simulate CI workflow
./scripts/benchmark.sh all

# View results
cat .benchmark-results/latest-raw.txt
```

## Benchmark Scripts

### benchmark.sh

Main benchmark automation script.

#### Commands

```bash
# Run benchmarks only
./scripts/benchmark.sh run

# Generate baseline from current results
./scripts/benchmark.sh generate-baseline

# Check for regressions
./scripts/benchmark.sh check

# Compare with previous run
./scripts/benchmark.sh compare

# Full workflow
./scripts/benchmark.sh all

# Show help
./scripts/benchmark.sh help
```

#### Environment Variables

```bash
# Customize benchmark duration
export BENCHMARK_TIME=5s
./scripts/benchmark.sh run

# Run multiple times for stability
export BENCHMARK_COUNT=3
./scripts/benchmark.sh run
```

## Performance Analysis

### Viewing Results

```bash
# View raw benchmark output
cat .benchmark-results/latest-raw.txt

# View parsed JSON results
cat .benchmark-results/latest-parsed.json | jq

# View baselines
cat scripts/performance-baselines.json | jq
```

### Using benchstat

Install benchstat for detailed comparison:

```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

Compare two benchmark runs:

```bash
benchstat baseline.txt current.txt
```

Output example:
```
name                    old time/op    new time/op    delta
CircuitBreakerSuccess   850ns ± 2%     920ns ± 3%   +8.24%
RetrySuccess            2.10µs ± 1%    2.25µs ± 2%  +7.14%
```

## Best Practices

### Setting Baselines

1. **Run Multiple Times**: Average 3-5 runs for stability
2. **Use Safety Factor**: Add 10-20% buffer for variance
3. **Update Regularly**: Refresh baselines after optimizations
4. **Document Changes**: Explain baseline adjustments

### Detecting Regressions

1. **Set Appropriate Thresholds**: Balance sensitivity vs noise
2. **Consider System Variance**: Account for CI/CD environment
3. **Review Context**: Not all increases are regressions
4. **Track Trends**: Look for consistent degradation

### Writing Benchmarks

1. **Realistic Workloads**: Mirror production scenarios
2. **Isolate Code**: Minimize external dependencies
3. **Avoid Optimization**: Don't let compiler optimize away code
4. **Use b.ResetTimer**: Exclude setup time
5. **Run Sufficient Iterations**: Ensure statistical significance

Example benchmark:

```go
func BenchmarkCircuitBreaker(b *testing.B) {
    cb := circuitbreaker.New[string](circuitbreaker.Config{
        Timeout: 100 * time.Millisecond,
    })

    b.ResetTimer() // Reset after setup

    for i := 0; i < b.N; i++ {
        _, _ = cb.Execute(context.Background(), func(ctx context.Context) (string, error) {
            return "result", nil
        })
    }
}
```

## Continuous Monitoring

### Historical Tracking

Results are stored in `.benchmark-results/` with timestamps:

```
.benchmark-results/
├── benchmark-20240315-143022.json
├── benchmark-20240315-150432.json
└── latest-raw.txt
```

### Performance Trends

Monitor trends over time:

```bash
# View all historical results
ls -lt .benchmark-results/*.json

# Compare specific dates
benchstat .benchmark-results/benchmark-20240301-*.txt \
          .benchmark-results/benchmark-20240315-*.txt
```

### Alerting

GitHub Actions automatically:
- Fails CI on regressions
- Comments on PRs with warnings
- Archives results for analysis
- Tracks performance over time

## Troubleshooting

### High Variance

**Problem**: Benchmarks show inconsistent results

**Solutions**:
- Increase `benchtime` (e.g., `-benchtime=10s`)
- Run multiple times (`-count=5`)
- Disable CPU frequency scaling
- Close background applications

### False Positives

**Problem**: CI reports regressions incorrectly

**Solutions**:
- Increase threshold tolerance
- Add safety factor to baselines
- Review baseline generation method
- Consider system differences

### Memory Allocations

**Problem**: Unexpected allocation increases

**Solutions**:
- Use `go test -benchmem` for memory profiling
- Run with `-memprofile=mem.prof`
- Analyze with `go tool pprof mem.prof`
- Check for unintended allocations

## Examples

See [testing/example_test.go](../testing/example_test.go) for complete examples:

- `Example_performanceTracking` - Basic tracking usage
- `Example_performanceBaseline` - Generating baselines

## Related Documentation

- [Testing Utilities](TESTING.md) - Chaos engineering tools
- [Metrics](METRICS.md) - Prometheus integration
- [Contributing](../CONTRIBUTING.md) - Development guidelines

## References

- [Go Testing Package](https://pkg.go.dev/testing)
- [Benchmarking Best Practices](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [Benchstat Tool](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)