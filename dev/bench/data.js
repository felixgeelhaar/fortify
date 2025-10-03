window.BENCHMARK_DATA = {
  "lastUpdate": 1759498207468,
  "repoUrl": "https://github.com/felixgeelhaar/fortify",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "79f179823b7883500d861f4a7ed87435797bd6ad",
          "message": "fix: add permissions for performance benchmark workflow\n\n- Add contents:write permission to push to gh-pages branch\n- Add deployments:write permission for GitHub Pages deployment\n- Required for benchmark-action/github-action-benchmark to store results",
          "timestamp": "2025-10-03T14:23:36+02:00",
          "tree_id": "37bd3201fe4d3e2188c10c313a135429ab2bdfa4",
          "url": "https://github.com/felixgeelhaar/fortify/commit/79f179823b7883500d861f4a7ed87435797bd6ad"
        },
        "date": 1759494437257,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 96.47,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38515136 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 96.47,
            "unit": "ns/op",
            "extra": "38515136 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38515136 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38515136 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 115.3,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31502426 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 115.3,
            "unit": "ns/op",
            "extra": "31502426 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31502426 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31502426 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 200.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18811897 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 200.8,
            "unit": "ns/op",
            "extra": "18811897 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18811897 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18811897 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25746771 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140.2,
            "unit": "ns/op",
            "extra": "25746771 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25746771 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25746771 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 171.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21139125 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 171.2,
            "unit": "ns/op",
            "extra": "21139125 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21139125 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21139125 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52636140 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.4,
            "unit": "ns/op",
            "extra": "52636140 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52636140 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52636140 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 200.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17750090 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 200.4,
            "unit": "ns/op",
            "extra": "17750090 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17750090 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17750090 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.4,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "128734766 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.4,
            "unit": "ns/op",
            "extra": "128734766 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "128734766 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "128734766 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.23,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91375500 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.23,
            "unit": "ns/op",
            "extra": "91375500 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91375500 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91375500 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.46,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "85291510 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.46,
            "unit": "ns/op",
            "extra": "85291510 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "85291510 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "85291510 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.04,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57181725 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.04,
            "unit": "ns/op",
            "extra": "57181725 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57181725 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57181725 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 52.69,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66362080 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 52.69,
            "unit": "ns/op",
            "extra": "66362080 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66362080 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66362080 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60708709 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.79,
            "unit": "ns/op",
            "extra": "60708709 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60708709 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60708709 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.39,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41149388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.39,
            "unit": "ns/op",
            "extra": "41149388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41149388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41149388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.38,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41173626 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.38,
            "unit": "ns/op",
            "extra": "41173626 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41173626 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41173626 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 88.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40851511 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 88.21,
            "unit": "ns/op",
            "extra": "40851511 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40851511 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40851511 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 92.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41333318 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 92.12,
            "unit": "ns/op",
            "extra": "41333318 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41333318 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41333318 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 134.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28042550 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 134.1,
            "unit": "ns/op",
            "extra": "28042550 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28042550 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28042550 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 65.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54595735 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 65.96,
            "unit": "ns/op",
            "extra": "54595735 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54595735 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54595735 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.486,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481625542 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.486,
            "unit": "ns/op",
            "extra": "481625542 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481625542 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481625542 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3201558,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3201558,
            "unit": "ns/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4260968,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "826 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4260968,
            "unit": "ns/op",
            "extra": "826 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "826 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "826 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174967906 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.57,
            "unit": "ns/op",
            "extra": "174967906 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174967906 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174967906 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.116,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant",
            "value": 2.489,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.489,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter",
            "value": 31.08,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 31.08,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 21.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "169441830 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.57,
            "unit": "ns/op",
            "extra": "169441830 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "169441830 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "169441830 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 238096,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "13779 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 238096,
            "unit": "ns/op",
            "extra": "13779 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "13779 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13779 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 262355,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "14377 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 262355,
            "unit": "ns/op",
            "extra": "14377 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "14377 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14377 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 384.8,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9368564 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 384.8,
            "unit": "ns/op",
            "extra": "9368564 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9368564 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9368564 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 410.4,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8863375 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 410.4,
            "unit": "ns/op",
            "extra": "8863375 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8863375 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8863375 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 182.4,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19591866 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 182.4,
            "unit": "ns/op",
            "extra": "19591866 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19591866 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19591866 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "ca8c54099e60f97395980c8b41177b9af0ef011d",
          "message": "fix: skip HTTP example main() test to prevent port conflicts\n\nThe TestHTTPExample test calls main() which starts a real HTTP server\non port 8080. When running with -count > 1 (used in extended benchmarks),\nthe second run fails with 'address already in use'.\n\nSolution: Skip the test since it only verifies the example compiles.\nThe middleware functionality is thoroughly tested in the other unit tests\n(TestHTTPMiddleware, TestHTTPRateLimit, TestHTTPTimeout).\n\nFixes performance regression testing workflow failures.",
          "timestamp": "2025-10-03T15:00:12+02:00",
          "tree_id": "ba7b88c8ee864ba0e33c8992fbaeaeef9731bd06",
          "url": "https://github.com/felixgeelhaar/fortify/commit/ca8c54099e60f97395980c8b41177b9af0ef011d"
        },
        "date": 1759496683504,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 90.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "39531360 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 90.76,
            "unit": "ns/op",
            "extra": "39531360 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "39531360 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "39531360 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 110,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32393672 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 110,
            "unit": "ns/op",
            "extra": "32393672 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32393672 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32393672 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 236.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "15949273 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 236.9,
            "unit": "ns/op",
            "extra": "15949273 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "15949273 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "15949273 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 137.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26259591 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 137.9,
            "unit": "ns/op",
            "extra": "26259591 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26259591 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26259591 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 166.3,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21876030 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 166.3,
            "unit": "ns/op",
            "extra": "21876030 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21876030 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21876030 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 69.29,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52205341 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 69.29,
            "unit": "ns/op",
            "extra": "52205341 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52205341 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52205341 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 245.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "14509760 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 245.5,
            "unit": "ns/op",
            "extra": "14509760 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "14509760 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "14509760 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 23.51,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "156009927 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 23.51,
            "unit": "ns/op",
            "extra": "156009927 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "156009927 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "156009927 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 35.45,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "98918474 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 35.45,
            "unit": "ns/op",
            "extra": "98918474 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "98918474 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "98918474 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 35.69,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "98925852 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 35.69,
            "unit": "ns/op",
            "extra": "98925852 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "98925852 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "98925852 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 59.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60377662 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 59.82,
            "unit": "ns/op",
            "extra": "60377662 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60377662 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60377662 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 70.28,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51292615 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 70.28,
            "unit": "ns/op",
            "extra": "51292615 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51292615 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51292615 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 58.23,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61841732 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 58.23,
            "unit": "ns/op",
            "extra": "61841732 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61841732 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61841732 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 77.34,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48435834 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 77.34,
            "unit": "ns/op",
            "extra": "48435834 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48435834 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48435834 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 75.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47666605 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 75.58,
            "unit": "ns/op",
            "extra": "47666605 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47666605 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47666605 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 75.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47717630 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 75.58,
            "unit": "ns/op",
            "extra": "47717630 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47717630 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47717630 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 79.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45721712 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 79.97,
            "unit": "ns/op",
            "extra": "45721712 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45721712 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45721712 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 154.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24662854 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 154.9,
            "unit": "ns/op",
            "extra": "24662854 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24662854 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24662854 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 44.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "79913922 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 44.74,
            "unit": "ns/op",
            "extra": "79913922 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "79913922 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "79913922 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 5.559,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "636794424 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 5.559,
            "unit": "ns/op",
            "extra": "636794424 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "636794424 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "636794424 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3200066,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1122 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3200066,
            "unit": "ns/op",
            "extra": "1122 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1122 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1122 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4234072,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4234072,
            "unit": "ns/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 21.06,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171986984 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 21.06,
            "unit": "ns/op",
            "extra": "171986984 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171986984 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171986984 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 2.715,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 2.715,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant",
            "value": 2.023,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.023,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter",
            "value": 29.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "123638605 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 29.21,
            "unit": "ns/op",
            "extra": "123638605 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123638605 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123638605 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 28.94,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "124851234 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 28.94,
            "unit": "ns/op",
            "extra": "124851234 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "124851234 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "124851234 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 322823,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 322823,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 342554,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "12748 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 342554,
            "unit": "ns/op",
            "extra": "12748 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "12748 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12748 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 464.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "7734219 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 464.2,
            "unit": "ns/op",
            "extra": "7734219 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "7734219 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7734219 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 482.9,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "7575085 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 482.9,
            "unit": "ns/op",
            "extra": "7575085 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "7575085 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "7575085 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 226.6,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "15927523 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 226.6,
            "unit": "ns/op",
            "extra": "15927523 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "15927523 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "15927523 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "d6d5ff4b2b2a0917531b8aff797008f1ad22dcb3",
          "message": "docs: remove temporary v1.0 status tracking files\n\nRemove temporary documentation created during v1.0 development:\n- PROJECT_SUMMARY.md: Initial project planning (archived)\n- V1_FINAL_STATUS.md: v1.0 status tracking (info in CHANGELOG)\n- V1_RELEASE_SUMMARY.md: v1.0 release notes (in GitHub release)\n\nPermanent documentation remains in docs/ directory.",
          "timestamp": "2025-10-03T15:05:28+02:00",
          "tree_id": "f1844c8aaad39f95673eba40557a239352fbe90c",
          "url": "https://github.com/felixgeelhaar/fortify/commit/d6d5ff4b2b2a0917531b8aff797008f1ad22dcb3"
        },
        "date": 1759496962306,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "32823795 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.64,
            "unit": "ns/op",
            "extra": "32823795 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "32823795 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "32823795 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.4,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31451810 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.4,
            "unit": "ns/op",
            "extra": "31451810 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31451810 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31451810 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 204,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17405410 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 204,
            "unit": "ns/op",
            "extra": "17405410 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17405410 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17405410 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25781320 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.5,
            "unit": "ns/op",
            "extra": "25781320 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25781320 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25781320 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 172,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20943753 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 172,
            "unit": "ns/op",
            "extra": "20943753 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20943753 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20943753 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52611008 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.54,
            "unit": "ns/op",
            "extra": "52611008 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52611008 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52611008 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 203.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17652108 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 203.8,
            "unit": "ns/op",
            "extra": "17652108 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17652108 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17652108 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.8,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "128598860 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.8,
            "unit": "ns/op",
            "extra": "128598860 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "128598860 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "128598860 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 37.86,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91056558 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 37.86,
            "unit": "ns/op",
            "extra": "91056558 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91056558 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91056558 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 42.09,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "85109649 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 42.09,
            "unit": "ns/op",
            "extra": "85109649 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "85109649 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "85109649 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 61.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57890244 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 61.83,
            "unit": "ns/op",
            "extra": "57890244 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57890244 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57890244 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 52.18,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66830498 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 52.18,
            "unit": "ns/op",
            "extra": "66830498 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66830498 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66830498 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57625773 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.56,
            "unit": "ns/op",
            "extra": "57625773 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57625773 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57625773 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41553249 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.01,
            "unit": "ns/op",
            "extra": "41553249 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41553249 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41553249 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "37825971 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.24,
            "unit": "ns/op",
            "extra": "37825971 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "37825971 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "37825971 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 95.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40969060 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 95.54,
            "unit": "ns/op",
            "extra": "40969060 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40969060 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40969060 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 95.41,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40772431 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 95.41,
            "unit": "ns/op",
            "extra": "40772431 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40772431 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40772431 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 134.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27246418 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 134.7,
            "unit": "ns/op",
            "extra": "27246418 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27246418 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27246418 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 66,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54574024 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 66,
            "unit": "ns/op",
            "extra": "54574024 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54574024 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54574024 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.476,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "480517969 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.476,
            "unit": "ns/op",
            "extra": "480517969 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "480517969 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "480517969 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3198600,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3198600,
            "unit": "ns/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4268669,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4268669,
            "unit": "ns/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174860516 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.57,
            "unit": "ns/op",
            "extra": "174860516 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174860516 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174860516 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant",
            "value": 2.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.49,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter",
            "value": 30.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "115753386 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.87,
            "unit": "ns/op",
            "extra": "115753386 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "115753386 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "115753386 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171889148 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21,
            "unit": "ns/op",
            "extra": "171889148 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171889148 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171889148 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 283693,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "14456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 283693,
            "unit": "ns/op",
            "extra": "14456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "14456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14456 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 275725,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "13176 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 275725,
            "unit": "ns/op",
            "extra": "13176 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "13176 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13176 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 384.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9245974 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 384.2,
            "unit": "ns/op",
            "extra": "9245974 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9245974 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9245974 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 405.7,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8848309 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 405.7,
            "unit": "ns/op",
            "extra": "8848309 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8848309 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8848309 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 188.9,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19653205 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 188.9,
            "unit": "ns/op",
            "extra": "19653205 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19653205 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19653205 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "c765d5256bda98af4c36f8194b9aa4151087a4a0",
          "message": "chore: move CLAUDE.md to local-only configuration\n\nCLAUDE.md contains project-specific instructions for Claude Code\nand should not be version controlled. It's now in .gitignore.\n\nUsers can create their own local CLAUDE.md if needed for\nproject-specific Claude Code configuration.",
          "timestamp": "2025-10-03T15:09:08+02:00",
          "tree_id": "42f21ee2eaffbe5c20cd71ca33e5e416a010d20a",
          "url": "https://github.com/felixgeelhaar/fortify/commit/c765d5256bda98af4c36f8194b9aa4151087a4a0"
        },
        "date": 1759497175701,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 90.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "39440083 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 90.67,
            "unit": "ns/op",
            "extra": "39440083 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "39440083 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "39440083 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 113.1,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32800304 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 113.1,
            "unit": "ns/op",
            "extra": "32800304 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32800304 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32800304 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 236.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "15323163 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 236.5,
            "unit": "ns/op",
            "extra": "15323163 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "15323163 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "15323163 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 136.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26288210 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 136.3,
            "unit": "ns/op",
            "extra": "26288210 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26288210 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26288210 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 166.7,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21449448 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 166.7,
            "unit": "ns/op",
            "extra": "21449448 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21449448 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21449448 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52081554 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.89,
            "unit": "ns/op",
            "extra": "52081554 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52081554 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52081554 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 242,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "14921672 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 242,
            "unit": "ns/op",
            "extra": "14921672 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "14921672 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "14921672 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 23.79,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "154376260 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 23.79,
            "unit": "ns/op",
            "extra": "154376260 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "154376260 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "154376260 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 35.65,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "99754885 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 35.65,
            "unit": "ns/op",
            "extra": "99754885 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "99754885 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "99754885 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 35.91,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "96766423 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 35.91,
            "unit": "ns/op",
            "extra": "96766423 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "96766423 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "96766423 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 59.53,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60402645 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 59.53,
            "unit": "ns/op",
            "extra": "60402645 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60402645 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60402645 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 70.46,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51166066 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 70.46,
            "unit": "ns/op",
            "extra": "51166066 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51166066 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51166066 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 58.14,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61869979 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 58.14,
            "unit": "ns/op",
            "extra": "61869979 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61869979 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61869979 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 77.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48474722 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 77.52,
            "unit": "ns/op",
            "extra": "48474722 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48474722 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48474722 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 85.35,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47721492 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 85.35,
            "unit": "ns/op",
            "extra": "47721492 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47721492 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47721492 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 76.39,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47800388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 76.39,
            "unit": "ns/op",
            "extra": "47800388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47800388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47800388 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 81.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45495181 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 81.54,
            "unit": "ns/op",
            "extra": "45495181 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45495181 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45495181 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 159.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24782480 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 159.3,
            "unit": "ns/op",
            "extra": "24782480 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24782480 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24782480 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 44.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80419200 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 44.67,
            "unit": "ns/op",
            "extra": "80419200 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80419200 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80419200 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 5.555,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "651589138 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 5.555,
            "unit": "ns/op",
            "extra": "651589138 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "651589138 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "651589138 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3201310,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3201310,
            "unit": "ns/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1120 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4261298,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4261298,
            "unit": "ns/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.93,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "172090123 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.93,
            "unit": "ns/op",
            "extra": "172090123 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "172090123 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "172090123 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 2.679,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 2.679,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant",
            "value": 2.025,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.025,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter",
            "value": 29.23,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "123048398 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 29.23,
            "unit": "ns/op",
            "extra": "123048398 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123048398 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123048398 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 28.84,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "124397785 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 28.84,
            "unit": "ns/op",
            "extra": "124397785 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "124397785 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "124397785 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 208471,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "17206 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 208471,
            "unit": "ns/op",
            "extra": "17206 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "17206 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "17206 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 212985,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 212985,
            "unit": "ns/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15520 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 471,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "7646766 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 471,
            "unit": "ns/op",
            "extra": "7646766 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "7646766 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7646766 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 482.2,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "7481803 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 482.2,
            "unit": "ns/op",
            "extra": "7481803 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "7481803 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "7481803 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 224.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "16141136 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 224.2,
            "unit": "ns/op",
            "extra": "16141136 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "16141136 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "16141136 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "committer": {
            "email": "felix@felixgeelhaar.de",
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar"
          },
          "distinct": true,
          "id": "997bf5d2d638dff1f701a38591578d2bda31b6ef",
          "message": "chore: remove project configuration file from version control\n\nMove project-specific configuration to local-only files.\nConfiguration file is now in .gitignore for local development only.",
          "timestamp": "2025-10-03T15:26:26+02:00",
          "tree_id": "42f21ee2eaffbe5c20cd71ca33e5e416a010d20a",
          "url": "https://github.com/felixgeelhaar/fortify/commit/997bf5d2d638dff1f701a38591578d2bda31b6ef"
        },
        "date": 1759498207060,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 97.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "35471116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 97.7,
            "unit": "ns/op",
            "extra": "35471116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "35471116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "35471116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 115,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31119444 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 115,
            "unit": "ns/op",
            "extra": "31119444 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31119444 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31119444 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 197.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18239256 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 197.3,
            "unit": "ns/op",
            "extra": "18239256 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18239256 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18239256 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25492653 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140.4,
            "unit": "ns/op",
            "extra": "25492653 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25492653 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25492653 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 172.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20564808 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 172.2,
            "unit": "ns/op",
            "extra": "20564808 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20564808 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20564808 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.93,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52321239 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.93,
            "unit": "ns/op",
            "extra": "52321239 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52321239 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52321239 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 200.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17873624 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 200.2,
            "unit": "ns/op",
            "extra": "17873624 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17873624 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17873624 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.07,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "128279635 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.07,
            "unit": "ns/op",
            "extra": "128279635 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "128279635 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "128279635 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.08,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91571536 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.08,
            "unit": "ns/op",
            "extra": "91571536 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91571536 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91571536 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 42.31,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "83754717 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 42.31,
            "unit": "ns/op",
            "extra": "83754717 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "83754717 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "83754717 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "56620322 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.2,
            "unit": "ns/op",
            "extra": "56620322 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "56620322 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "56620322 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 51.45,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "69807064 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 51.45,
            "unit": "ns/op",
            "extra": "69807064 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "69807064 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "69807064 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 60.16,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61475035 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 60.16,
            "unit": "ns/op",
            "extra": "61475035 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61475035 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61475035 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41645557 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89,
            "unit": "ns/op",
            "extra": "41645557 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41645557 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41645557 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41369876 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.21,
            "unit": "ns/op",
            "extra": "41369876 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41369876 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41369876 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 88.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41009599 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 88.02,
            "unit": "ns/op",
            "extra": "41009599 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41009599 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41009599 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 91.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41211858 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 91.09,
            "unit": "ns/op",
            "extra": "41211858 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41211858 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41211858 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 134.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27561748 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 134.1,
            "unit": "ns/op",
            "extra": "27561748 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27561748 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27561748 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 65.93,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54584899 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 65.93,
            "unit": "ns/op",
            "extra": "54584899 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54584899 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54584899 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.479,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "480878361 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.479,
            "unit": "ns/op",
            "extra": "480878361 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "480878361 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "480878361 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3200072,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3200072,
            "unit": "ns/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4243883,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "849 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4243883,
            "unit": "ns/op",
            "extra": "849 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "849 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "849 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174854048 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.59,
            "unit": "ns/op",
            "extra": "174854048 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174854048 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174854048 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant",
            "value": 2.493,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.493,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter",
            "value": 30.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.95,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 21.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171542260 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.21,
            "unit": "ns/op",
            "extra": "171542260 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171542260 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171542260 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 229504,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "14660 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 229504,
            "unit": "ns/op",
            "extra": "14660 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "14660 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14660 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 220419,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "14760 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 220419,
            "unit": "ns/op",
            "extra": "14760 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "14760 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14760 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 386.4,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9296656 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 386.4,
            "unit": "ns/op",
            "extra": "9296656 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9296656 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9296656 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 411.8,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8815329 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 411.8,
            "unit": "ns/op",
            "extra": "8815329 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8815329 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8815329 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 184,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19299976 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 184,
            "unit": "ns/op",
            "extra": "19299976 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19299976 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19299976 times\n4 procs"
          }
        ]
      }
    ]
  }
}