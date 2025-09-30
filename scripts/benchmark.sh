#!/bin/bash
# Performance benchmark and regression testing script

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
RESULTS_DIR="$PROJECT_ROOT/.benchmark-results"
BASELINES_FILE="$PROJECT_ROOT/scripts/performance-baselines.json"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Create results directory
mkdir -p "$RESULTS_DIR"

# Function to print colored output
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to run benchmarks
run_benchmarks() {
    print_info "Running benchmarks..."

    cd "$PROJECT_ROOT"

    # Run benchmarks with memory profiling
    go test -bench=. -benchmem -benchtime=3s ./... > "$RESULTS_DIR/latest-raw.txt" 2>&1

    print_info "Benchmark results saved to $RESULTS_DIR/latest-raw.txt"
}

# Function to parse benchmark results
parse_benchmarks() {
    print_info "Parsing benchmark results..."

    # Extract benchmark data using awk
    awk '/^Benchmark/ {
        name=$1
        ops=$2
        ns_per_op=$3
        if ($5 == "B/op") {
            bytes_per_op=$4
            allocs_per_op=$6
        } else if ($7 == "B/op") {
            bytes_per_op=$6
            allocs_per_op=$8
        }
        printf "{\"name\":\"%s\",\"ns_per_op\":%s,\"allocs_per_op\":%s,\"bytes_per_op\":%s}\n",
            name, ns_per_op, allocs_per_op, bytes_per_op
    }' "$RESULTS_DIR/latest-raw.txt" > "$RESULTS_DIR/latest-parsed.json"
}

# Function to generate baseline from current results
generate_baseline() {
    print_info "Generating performance baseline..."

    if [ ! -f "$RESULTS_DIR/latest-parsed.json" ]; then
        print_error "No benchmark results found. Run benchmarks first."
        exit 1
    fi

    # Add safety factor of 1.1 (10% buffer)
    python3 -c "
import json
import sys

try:
    with open('$RESULTS_DIR/latest-parsed.json', 'r') as f:
        results = [json.loads(line) for line in f if line.strip()]

    baselines = []
    for result in results:
        baseline = {
            'name': result['name'],
            'max_ns_per_op': result['ns_per_op'] * 1.1,
            'max_allocs': int(result['allocs_per_op'] * 1.1),
            'max_bytes': int(result['bytes_per_op'] * 1.1),
            'description': 'Generated baseline with 10% safety factor'
        }
        baselines.append(baseline)

    with open('$BASELINES_FILE', 'w') as f:
        json.dump(baselines, f, indent=2)

    print(f'Generated {len(baselines)} baselines')
except Exception as e:
    print(f'Error: {e}', file=sys.stderr)
    sys.exit(1)
"

    print_info "Baseline saved to $BASELINES_FILE"
}

# Function to check for regressions
check_regressions() {
    print_info "Checking for performance regressions..."

    if [ ! -f "$BASELINES_FILE" ]; then
        print_warn "No baseline file found at $BASELINES_FILE"
        print_warn "Run '$0 generate-baseline' to create one"
        return 0
    fi

    if [ ! -f "$RESULTS_DIR/latest-parsed.json" ]; then
        print_error "No benchmark results found. Run benchmarks first."
        exit 1
    fi

    # Check for regressions using Python
    python3 -c "
import json
import sys

try:
    # Load baselines
    with open('$BASELINES_FILE', 'r') as f:
        baselines = {b['name']: b for b in json.load(f)}

    # Load current results
    with open('$RESULTS_DIR/latest-parsed.json', 'r') as f:
        results = [json.loads(line) for line in f if line.strip()]

    # Thresholds
    TIME_THRESHOLD = 1.10  # 10% slower
    ALLOC_THRESHOLD = 1.20  # 20% more allocations
    BYTES_THRESHOLD = 1.15  # 15% more memory

    regressions = []
    passed = 0

    for result in results:
        name = result['name']
        if name not in baselines:
            continue

        baseline = baselines[name]

        # Check time regression
        if result['ns_per_op'] > baseline['max_ns_per_op']:
            increase = (result['ns_per_op'] / baseline['max_ns_per_op'] - 1) * 100
            regressions.append({
                'benchmark': name,
                'metric': 'time',
                'baseline': baseline['max_ns_per_op'],
                'current': result['ns_per_op'],
                'increase': increase
            })
        # Check allocation regression
        elif result['allocs_per_op'] > baseline['max_allocs'] * ALLOC_THRESHOLD:
            increase = (result['allocs_per_op'] / baseline['max_allocs'] - 1) * 100
            regressions.append({
                'benchmark': name,
                'metric': 'allocs',
                'baseline': baseline['max_allocs'],
                'current': result['allocs_per_op'],
                'increase': increase
            })
        # Check memory regression
        elif result['bytes_per_op'] > baseline['max_bytes'] * BYTES_THRESHOLD:
            increase = (result['bytes_per_op'] / baseline['max_bytes'] - 1) * 100
            regressions.append({
                'benchmark': name,
                'metric': 'bytes',
                'baseline': baseline['max_bytes'],
                'current': result['bytes_per_op'],
                'increase': increase
            })
        else:
            passed += 1

    # Print results
    print(f'\nPerformance Regression Check Results:')
    print(f'  Total benchmarks: {len(results)}')
    print(f'  Passed: {passed}')
    print(f'  Failed: {len(regressions)}')

    if regressions:
        print(f'\nRegressions detected:')
        for reg in regressions:
            print(f'  ❌ {reg[\"benchmark\"]}')
            print(f'     Metric: {reg[\"metric\"]}')
            print(f'     Baseline: {reg[\"baseline\"]:.2f}')
            print(f'     Current: {reg[\"current\"]:.2f}')
            print(f'     Increase: {reg[\"increase\"]:.2f}%')
        sys.exit(1)
    else:
        print(f'\n✅ All benchmarks passed!')
        sys.exit(0)

except Exception as e:
    print(f'Error: {e}', file=sys.stderr)
    sys.exit(1)
"
}

# Function to compare with previous run
compare_with_previous() {
    print_info "Comparing with previous benchmark run..."

    # Find previous result file
    PREVIOUS=$(ls -t "$RESULTS_DIR"/benchmark-*.txt 2>/dev/null | head -2 | tail -1)

    if [ -z "$PREVIOUS" ]; then
        print_warn "No previous benchmark results found for comparison"
        return 0
    fi

    print_info "Comparing with: $PREVIOUS"

    # Use benchcmp if available
    if command -v benchcmp &> /dev/null; then
        benchcmp "$PREVIOUS" "$RESULTS_DIR/latest-raw.txt"
    else
        print_warn "benchcmp not installed. Install with: go install golang.org/x/tools/cmd/benchcmp@latest"
    fi
}

# Function to show usage
show_usage() {
    cat << EOF
Usage: $0 [command]

Commands:
    run                 Run all benchmarks
    generate-baseline   Generate performance baseline from current results
    check              Check for performance regressions against baseline
    compare            Compare with previous benchmark run
    all                Run benchmarks, check regressions, and compare
    help               Show this help message

Examples:
    $0 run                  # Run benchmarks
    $0 generate-baseline    # Create baseline from results
    $0 check                # Check for regressions
    $0 all                  # Complete benchmark workflow

Environment Variables:
    BENCHMARK_TIME      Benchmark duration (default: 3s)
    BENCHMARK_COUNT     Number of benchmark runs (default: 1)

EOF
}

# Main execution
case "${1:-all}" in
    run)
        run_benchmarks
        parse_benchmarks
        ;;
    generate-baseline)
        if [ ! -f "$RESULTS_DIR/latest-raw.txt" ]; then
            print_warn "No benchmark results found. Running benchmarks first..."
            run_benchmarks
            parse_benchmarks
        fi
        generate_baseline
        ;;
    check)
        if [ ! -f "$RESULTS_DIR/latest-raw.txt" ]; then
            print_warn "No benchmark results found. Running benchmarks first..."
            run_benchmarks
            parse_benchmarks
        fi
        check_regressions
        ;;
    compare)
        compare_with_previous
        ;;
    all)
        run_benchmarks
        parse_benchmarks
        check_regressions
        compare_with_previous
        ;;
    help|--help|-h)
        show_usage
        ;;
    *)
        print_error "Unknown command: $1"
        show_usage
        exit 1
        ;;
esac