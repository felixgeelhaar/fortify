window.BENCHMARK_DATA = {
  "lastUpdate": 1765671317188,
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
          "id": "06f161b4e8a3972dd1eaa304baef22c15d1b20d7",
          "message": "docs: refresh Go Report Card badge display\n\nTrigger badge cache refresh to display current A+ grade.",
          "timestamp": "2025-10-03T15:29:42+02:00",
          "tree_id": "b708dd8fa052a4213b423644934beb094a0dc66c",
          "url": "https://github.com/felixgeelhaar/fortify/commit/06f161b4e8a3972dd1eaa304baef22c15d1b20d7"
        },
        "date": 1759498406332,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "36736281 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.54,
            "unit": "ns/op",
            "extra": "36736281 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "36736281 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "36736281 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 115.2,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31396699 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 115.2,
            "unit": "ns/op",
            "extra": "31396699 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31396699 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31396699 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 193.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18141237 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 193.3,
            "unit": "ns/op",
            "extra": "18141237 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18141237 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18141237 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25815471 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.6,
            "unit": "ns/op",
            "extra": "25815471 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25815471 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25815471 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 171.7,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21108733 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 171.7,
            "unit": "ns/op",
            "extra": "21108733 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21108733 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21108733 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52770721 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.54,
            "unit": "ns/op",
            "extra": "52770721 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52770721 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52770721 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 201.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17737590 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 201.6,
            "unit": "ns/op",
            "extra": "17737590 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17737590 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17737590 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.9,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "129202022 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.9,
            "unit": "ns/op",
            "extra": "129202022 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "129202022 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "129202022 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.8,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "90438435 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.8,
            "unit": "ns/op",
            "extra": "90438435 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "90438435 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "90438435 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.39,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "83605647 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.39,
            "unit": "ns/op",
            "extra": "83605647 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "83605647 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "83605647 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "58352067 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.54,
            "unit": "ns/op",
            "extra": "58352067 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "58352067 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "58352067 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 54.14,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66687213 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 54.14,
            "unit": "ns/op",
            "extra": "66687213 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66687213 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66687213 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.84,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61208845 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.84,
            "unit": "ns/op",
            "extra": "61208845 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61208845 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61208845 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.35,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41505399 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.35,
            "unit": "ns/op",
            "extra": "41505399 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41505399 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41505399 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.34,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41367562 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.34,
            "unit": "ns/op",
            "extra": "41367562 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41367562 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41367562 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 87.88,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40896042 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 87.88,
            "unit": "ns/op",
            "extra": "40896042 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40896042 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40896042 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 91.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41182412 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 91.73,
            "unit": "ns/op",
            "extra": "41182412 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41182412 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41182412 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 135,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27702788 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 135,
            "unit": "ns/op",
            "extra": "27702788 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27702788 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27702788 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 65.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54551535 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 65.96,
            "unit": "ns/op",
            "extra": "54551535 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54551535 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54551535 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.485,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481526223 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.485,
            "unit": "ns/op",
            "extra": "481526223 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481526223 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481526223 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3198874,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3198874,
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
            "value": 4253875,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4253875,
            "unit": "ns/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174754340 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.59,
            "unit": "ns/op",
            "extra": "174754340 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174754340 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174754340 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.135,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.135,
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
            "value": 30.81,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.81,
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
            "value": 20.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171492243 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 20.99,
            "unit": "ns/op",
            "extra": "171492243 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171492243 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171492243 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 259414,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 259414,
            "unit": "ns/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 276699,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "12824 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 276699,
            "unit": "ns/op",
            "extra": "12824 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "12824 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12824 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 381.1,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9461944 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 381.1,
            "unit": "ns/op",
            "extra": "9461944 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9461944 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9461944 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 401.2,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8895836 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 401.2,
            "unit": "ns/op",
            "extra": "8895836 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8895836 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8895836 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 187.7,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19744141 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 187.7,
            "unit": "ns/op",
            "extra": "19744141 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19744141 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19744141 times\n4 procs"
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
          "id": "3209172dda66cb14c8c339b2c8a9ef378d08eaa7",
          "message": "feat: add Redis backend for distributed rate limiting\n\nAdd production-ready Redis-backed rate limiter for distributed systems. This maintains Fortify's zero-dependency promise for core by implementing the Redis backend as a separate Go module.\n\nKey Features:\n- Atomic operations via Lua scripts (no race conditions)\n- Same RateLimiter interface as in-memory (drop-in replacement)\n- Support for Redis Cluster, Sentinel, and standalone deployments\n- Configurable fail-open/fail-closed behavior\n- Automatic bucket expiration with TTL\n- Full observability integration (slog, OpenTelemetry)\n- >90% test coverage with miniredis\n\nImplementation:\n- backends/redis/ - Separate Go module with Redis implementation\n- Lua script for atomic token bucket operations\n- Comprehensive test suite (unit, benchmark, examples)\n- Production-grade error handling and connection pooling\n\nDocumentation:\n- Complete README with installation and usage guide\n- Migration guide for moving from in-memory to Redis\n- Working Docker Compose example with 3 instances + Redis",
          "timestamp": "2025-10-19T11:46:23+02:00",
          "tree_id": "f4c65e022509439ba48a39675206c59cadd3dc29",
          "url": "https://github.com/felixgeelhaar/fortify/commit/3209172dda66cb14c8c339b2c8a9ef378d08eaa7"
        },
        "date": 1760867426498,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38149778 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.55,
            "unit": "ns/op",
            "extra": "38149778 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38149778 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38149778 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.6,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31450651 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.6,
            "unit": "ns/op",
            "extra": "31450651 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31450651 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31450651 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 199.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18149322 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 199.1,
            "unit": "ns/op",
            "extra": "18149322 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18149322 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18149322 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25071003 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.7,
            "unit": "ns/op",
            "extra": "25071003 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25071003 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25071003 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 170.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21136417 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 170.2,
            "unit": "ns/op",
            "extra": "21136417 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21136417 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21136417 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 67.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "53094325 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 67.74,
            "unit": "ns/op",
            "extra": "53094325 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "53094325 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "53094325 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 208.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17564700 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 208.1,
            "unit": "ns/op",
            "extra": "17564700 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17564700 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17564700 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.31,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "129286202 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.31,
            "unit": "ns/op",
            "extra": "129286202 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "129286202 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "129286202 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.01,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "87623931 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.01,
            "unit": "ns/op",
            "extra": "87623931 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "87623931 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "87623931 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.4,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "84011307 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.4,
            "unit": "ns/op",
            "extra": "84011307 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "84011307 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "84011307 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 61.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57264010 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 61.67,
            "unit": "ns/op",
            "extra": "57264010 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57264010 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57264010 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 51.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66967976 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 51.73,
            "unit": "ns/op",
            "extra": "66967976 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66967976 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66967976 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.48,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61138442 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.48,
            "unit": "ns/op",
            "extra": "61138442 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61138442 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61138442 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41374254 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.74,
            "unit": "ns/op",
            "extra": "41374254 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41374254 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41374254 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.88,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40984270 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.88,
            "unit": "ns/op",
            "extra": "40984270 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40984270 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40984270 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 89.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40872525 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 89.09,
            "unit": "ns/op",
            "extra": "40872525 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40872525 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40872525 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 91.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41201673 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 91.68,
            "unit": "ns/op",
            "extra": "41201673 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41201673 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41201673 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 136.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27071294 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 136.5,
            "unit": "ns/op",
            "extra": "27071294 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27071294 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27071294 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 65.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54559416 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 65.99,
            "unit": "ns/op",
            "extra": "54559416 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54559416 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54559416 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.489,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "478889420 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.489,
            "unit": "ns/op",
            "extra": "478889420 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "478889420 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "478889420 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3201892,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3201892,
            "unit": "ns/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4236901,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "848 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4236901,
            "unit": "ns/op",
            "extra": "848 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "848 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "848 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174967767 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.59,
            "unit": "ns/op",
            "extra": "174967767 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174967767 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174967767 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.117,
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
            "value": 30.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.91,
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
            "value": 21.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "167218508 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.5,
            "unit": "ns/op",
            "extra": "167218508 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "167218508 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "167218508 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 183593,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "18835 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 183593,
            "unit": "ns/op",
            "extra": "18835 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "18835 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18835 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 187384,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "19459 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 187384,
            "unit": "ns/op",
            "extra": "19459 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "19459 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "19459 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 384.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9301988 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 384.2,
            "unit": "ns/op",
            "extra": "9301988 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9301988 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9301988 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 408.9,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8883945 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 408.9,
            "unit": "ns/op",
            "extra": "8883945 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8883945 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8883945 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 187.3,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19247734 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 187.3,
            "unit": "ns/op",
            "extra": "19247734 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19247734 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19247734 times\n4 procs"
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
          "id": "35630b2a6d3d83b94904b2264de0df026874ce85",
          "message": "chore: bump version to 1.1.0 for Redis backend release",
          "timestamp": "2025-10-19T11:50:01+02:00",
          "tree_id": "09569c4782c7ef853e702560db9d9ad93612d59f",
          "url": "https://github.com/felixgeelhaar/fortify/commit/35630b2a6d3d83b94904b2264de0df026874ce85"
        },
        "date": 1760867614635,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 90.78,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "39629116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 90.78,
            "unit": "ns/op",
            "extra": "39629116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "39629116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "39629116 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.3,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "27393552 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.3,
            "unit": "ns/op",
            "extra": "27393552 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "27393552 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "27393552 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 222.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "16514973 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 222.3,
            "unit": "ns/op",
            "extra": "16514973 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "16514973 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "16514973 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 136.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26373660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 136.6,
            "unit": "ns/op",
            "extra": "26373660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26373660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26373660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 166.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21473834 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 166.2,
            "unit": "ns/op",
            "extra": "21473834 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21473834 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21473834 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52305235 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.85,
            "unit": "ns/op",
            "extra": "52305235 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52305235 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52305235 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 250.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "14337422 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 250.9,
            "unit": "ns/op",
            "extra": "14337422 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "14337422 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "14337422 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 23.51,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "153045373 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 23.51,
            "unit": "ns/op",
            "extra": "153045373 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "153045373 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "153045373 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 37.1,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "97462646 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 37.1,
            "unit": "ns/op",
            "extra": "97462646 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "97462646 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "97462646 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 36.95,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "94448978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 36.95,
            "unit": "ns/op",
            "extra": "94448978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "94448978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "94448978 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 60.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59601000 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 60.02,
            "unit": "ns/op",
            "extra": "59601000 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59601000 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59601000 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 72.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "50235223 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 72.02,
            "unit": "ns/op",
            "extra": "50235223 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "50235223 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "50235223 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 58.33,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61679602 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 58.33,
            "unit": "ns/op",
            "extra": "61679602 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61679602 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61679602 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 77.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48548542 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 77.21,
            "unit": "ns/op",
            "extra": "48548542 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48548542 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48548542 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 75.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47719443 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 75.59,
            "unit": "ns/op",
            "extra": "47719443 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47719443 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47719443 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 76.08,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47703886 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 76.08,
            "unit": "ns/op",
            "extra": "47703886 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47703886 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47703886 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 80.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "44661180 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 80.62,
            "unit": "ns/op",
            "extra": "44661180 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "44661180 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "44661180 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 155.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24680356 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 155.8,
            "unit": "ns/op",
            "extra": "24680356 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24680356 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24680356 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 44.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80393248 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 44.72,
            "unit": "ns/op",
            "extra": "80393248 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80393248 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80393248 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 5.558,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "648265711 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 5.558,
            "unit": "ns/op",
            "extra": "648265711 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "648265711 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "648265711 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3200615,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3200615,
            "unit": "ns/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4247757,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4247757,
            "unit": "ns/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "853 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171747534 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.97,
            "unit": "ns/op",
            "extra": "171747534 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171747534 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171747534 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 2.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 2.62,
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
            "value": 2.064,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.064,
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
            "value": 29.15,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "123602187 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 29.15,
            "unit": "ns/op",
            "extra": "123602187 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123602187 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123602187 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 28.78,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "124621950 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 28.78,
            "unit": "ns/op",
            "extra": "124621950 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "124621950 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "124621950 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 197132,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "18360 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 197132,
            "unit": "ns/op",
            "extra": "18360 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "18360 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18360 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 225276,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "16495 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 225276,
            "unit": "ns/op",
            "extra": "16495 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "16495 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16495 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 472.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "7652209 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 472.2,
            "unit": "ns/op",
            "extra": "7652209 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "7652209 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7652209 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 489.3,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "7461500 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 489.3,
            "unit": "ns/op",
            "extra": "7461500 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "7461500 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "7461500 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 223.5,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "16185900 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 223.5,
            "unit": "ns/op",
            "extra": "16185900 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "16185900 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "16185900 times\n4 procs"
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
          "id": "eded91108b1a5f0ea40a590f1dbdb01d506f8d6c",
          "message": "chore: add go.sum for Redis backend module",
          "timestamp": "2025-10-19T12:16:21+02:00",
          "tree_id": "81434760fe13eab16506c2a51a07ac882480147b",
          "url": "https://github.com/felixgeelhaar/fortify/commit/eded91108b1a5f0ea40a590f1dbdb01d506f8d6c"
        },
        "date": 1760869193208,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.42,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38574289 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.42,
            "unit": "ns/op",
            "extra": "38574289 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38574289 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38574289 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 115,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31273819 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 115,
            "unit": "ns/op",
            "extra": "31273819 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31273819 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31273819 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 199.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18098766 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 199.9,
            "unit": "ns/op",
            "extra": "18098766 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18098766 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18098766 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25664038 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.8,
            "unit": "ns/op",
            "extra": "25664038 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25664038 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25664038 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 170.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20978792 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 170.2,
            "unit": "ns/op",
            "extra": "20978792 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20978792 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20978792 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 69.34,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52837280 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 69.34,
            "unit": "ns/op",
            "extra": "52837280 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52837280 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52837280 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 204.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17632252 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 204.7,
            "unit": "ns/op",
            "extra": "17632252 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17632252 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17632252 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.28,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "129868222 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.28,
            "unit": "ns/op",
            "extra": "129868222 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "129868222 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "129868222 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 37.98,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "90566271 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 37.98,
            "unit": "ns/op",
            "extra": "90566271 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "90566271 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "90566271 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.4,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "81743467 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.4,
            "unit": "ns/op",
            "extra": "81743467 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "81743467 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "81743467 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "56841477 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.74,
            "unit": "ns/op",
            "extra": "56841477 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "56841477 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "56841477 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 53.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66956743 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 53.24,
            "unit": "ns/op",
            "extra": "66956743 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66956743 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66956743 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 60.39,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61086426 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 60.39,
            "unit": "ns/op",
            "extra": "61086426 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61086426 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61086426 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.34,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41525607 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.34,
            "unit": "ns/op",
            "extra": "41525607 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41525607 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41525607 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41300247 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.11,
            "unit": "ns/op",
            "extra": "41300247 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41300247 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41300247 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 88.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40961332 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 88.26,
            "unit": "ns/op",
            "extra": "40961332 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40961332 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40961332 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 91.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41122210 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 91.67,
            "unit": "ns/op",
            "extra": "41122210 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41122210 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41122210 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 137.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27285949 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 137.3,
            "unit": "ns/op",
            "extra": "27285949 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27285949 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27285949 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 66.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54527366 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 66.01,
            "unit": "ns/op",
            "extra": "54527366 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54527366 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54527366 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.844,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "480700477 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.844,
            "unit": "ns/op",
            "extra": "480700477 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "480700477 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "480700477 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3203424,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3203424,
            "unit": "ns/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1123 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4220869,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4220869,
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
            "value": 20.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174703513 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.62,
            "unit": "ns/op",
            "extra": "174703513 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174703513 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174703513 times\n4 procs"
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
            "value": 2.491,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.491,
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
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.87,
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
            "value": 20.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "167117490 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 20.98,
            "unit": "ns/op",
            "extra": "167117490 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "167117490 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "167117490 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 182330,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "18456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 182330,
            "unit": "ns/op",
            "extra": "18456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "18456 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18456 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 191034,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "19809 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 191034,
            "unit": "ns/op",
            "extra": "19809 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "19809 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "19809 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 382,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9439029 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 382,
            "unit": "ns/op",
            "extra": "9439029 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9439029 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9439029 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 406.7,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8918670 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 406.7,
            "unit": "ns/op",
            "extra": "8918670 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8918670 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8918670 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 185,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19165134 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 185,
            "unit": "ns/op",
            "extra": "19165134 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19165134 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19165134 times\n4 procs"
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
          "id": "6bdbd44837103be9dad761b4f44e4133e7d888f2",
          "message": "fix: update go-redis to v9.7.3 to address CVE-2025-29923\n\nUpdates go-redis from v9.7.0 to v9.7.3 to fix a low severity vulnerability (CVE-2025-29923) related to out-of-order responses when CLIENT SETINFO times out during connection establishment.",
          "timestamp": "2025-10-19T12:16:54+02:00",
          "tree_id": "ab0ef782876902faa7fbf35250d09946de5cea10",
          "url": "https://github.com/felixgeelhaar/fortify/commit/6bdbd44837103be9dad761b4f44e4133e7d888f2"
        },
        "date": 1760869232345,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.53,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38416050 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.53,
            "unit": "ns/op",
            "extra": "38416050 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38416050 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38416050 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.6,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31469713 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.6,
            "unit": "ns/op",
            "extra": "31469713 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31469713 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31469713 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 203.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17687550 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "17687550 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17687550 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17687550 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25668691 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140.2,
            "unit": "ns/op",
            "extra": "25668691 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25668691 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25668691 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 173.7,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "19388943 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 173.7,
            "unit": "ns/op",
            "extra": "19388943 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "19388943 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "19388943 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.25,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52766120 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.25,
            "unit": "ns/op",
            "extra": "52766120 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52766120 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52766120 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 202.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17515730 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 202.9,
            "unit": "ns/op",
            "extra": "17515730 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17515730 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17515730 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.71,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "129806948 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.71,
            "unit": "ns/op",
            "extra": "129806948 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "129806948 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "129806948 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.73,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "90956955 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.73,
            "unit": "ns/op",
            "extra": "90956955 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "90956955 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "90956955 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.22,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "83663894 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.22,
            "unit": "ns/op",
            "extra": "83663894 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "83663894 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "83663894 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57438961 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.61,
            "unit": "ns/op",
            "extra": "57438961 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57438961 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57438961 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 52.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "64399624 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 52.74,
            "unit": "ns/op",
            "extra": "64399624 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "64399624 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "64399624 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 60.06,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60497641 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 60.06,
            "unit": "ns/op",
            "extra": "60497641 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60497641 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60497641 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41176323 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.02,
            "unit": "ns/op",
            "extra": "41176323 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41176323 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41176323 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41340516 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.37,
            "unit": "ns/op",
            "extra": "41340516 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41340516 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41340516 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 87.88,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "40955167 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 87.88,
            "unit": "ns/op",
            "extra": "40955167 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "40955167 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "40955167 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 92.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41169344 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 92.21,
            "unit": "ns/op",
            "extra": "41169344 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41169344 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41169344 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 135.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27503940 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 135.6,
            "unit": "ns/op",
            "extra": "27503940 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27503940 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27503940 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 66.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54564812 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 66.03,
            "unit": "ns/op",
            "extra": "54564812 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54564812 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54564812 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.479,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481674496 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.479,
            "unit": "ns/op",
            "extra": "481674496 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481674496 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481674496 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3199914,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3199914,
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
            "value": 4250360,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "867 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4250360,
            "unit": "ns/op",
            "extra": "867 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "867 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "867 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174255897 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.57,
            "unit": "ns/op",
            "extra": "174255897 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174255897 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174255897 times\n4 procs"
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
            "value": 30.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.86,
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
            "value": 21.18,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171552840 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.18,
            "unit": "ns/op",
            "extra": "171552840 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171552840 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171552840 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 194794,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "18441 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 194794,
            "unit": "ns/op",
            "extra": "18441 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "18441 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18441 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 176869,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "22698 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 176869,
            "unit": "ns/op",
            "extra": "22698 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "22698 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "22698 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 386.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9373364 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 386.2,
            "unit": "ns/op",
            "extra": "9373364 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9373364 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9373364 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 414.3,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8659089 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 414.3,
            "unit": "ns/op",
            "extra": "8659089 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8659089 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8659089 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 189.5,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19393195 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 189.5,
            "unit": "ns/op",
            "extra": "19393195 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19393195 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19393195 times\n4 procs"
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
          "id": "a5cc42344e95bd1c380d76152822fec49a7d9621",
          "message": "fix: add go.mod to Redis backend example for CI compatibility\n\nThe example now has its own go.mod with replace directives to reference the local backends/redis and core modules. This fixes CI build failures that occurred when the example tried to import the backends/redis package without proper module resolution.",
          "timestamp": "2025-10-19T12:23:25+02:00",
          "tree_id": "233c71ac02db0fd570a0e293240024fc59a6e367",
          "url": "https://github.com/felixgeelhaar/fortify/commit/a5cc42344e95bd1c380d76152822fec49a7d9621"
        },
        "date": 1760869625016,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 92.69,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38529505 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 92.69,
            "unit": "ns/op",
            "extra": "38529505 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38529505 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38529505 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.6,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31420080 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.6,
            "unit": "ns/op",
            "extra": "31420080 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31420080 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31420080 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 191.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17756678 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 191.9,
            "unit": "ns/op",
            "extra": "17756678 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17756678 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17756678 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25587732 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140.5,
            "unit": "ns/op",
            "extra": "25587732 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25587732 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25587732 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 171.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20926190 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 171.2,
            "unit": "ns/op",
            "extra": "20926190 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20926190 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20926190 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52784527 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.56,
            "unit": "ns/op",
            "extra": "52784527 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52784527 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52784527 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 205.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17473009 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 205.9,
            "unit": "ns/op",
            "extra": "17473009 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17473009 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17473009 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.7,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "127184683 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.7,
            "unit": "ns/op",
            "extra": "127184683 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "127184683 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "127184683 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.14,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91206883 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.14,
            "unit": "ns/op",
            "extra": "91206883 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91206883 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91206883 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.36,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "83597095 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.36,
            "unit": "ns/op",
            "extra": "83597095 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "83597095 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "83597095 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57307088 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.75,
            "unit": "ns/op",
            "extra": "57307088 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57307088 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57307088 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 53.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66897622 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 53.37,
            "unit": "ns/op",
            "extra": "66897622 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66897622 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66897622 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 60.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60574741 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 60.02,
            "unit": "ns/op",
            "extra": "60574741 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60574741 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60574741 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 89.17,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41763043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 89.17,
            "unit": "ns/op",
            "extra": "41763043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41763043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41763043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 87.23,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41254068 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 87.23,
            "unit": "ns/op",
            "extra": "41254068 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41254068 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41254068 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 87.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41067596 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 87.99,
            "unit": "ns/op",
            "extra": "41067596 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41067596 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41067596 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 91.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "41343583 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 91.72,
            "unit": "ns/op",
            "extra": "41343583 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "41343583 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "41343583 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 135.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27159483 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 135.8,
            "unit": "ns/op",
            "extra": "27159483 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27159483 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27159483 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 66.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54540231 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 66.02,
            "unit": "ns/op",
            "extra": "54540231 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54540231 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54540231 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.477,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481355086 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.477,
            "unit": "ns/op",
            "extra": "481355086 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481355086 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481355086 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3201066,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3201066,
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
            "value": 4231819,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4231819,
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
            "value": 20.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174974744 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.59,
            "unit": "ns/op",
            "extra": "174974744 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174974744 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174974744 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.117,
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
            "value": 30.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.89,
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
            "extra": "171818704 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.21,
            "unit": "ns/op",
            "extra": "171818704 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171818704 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171818704 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 209747,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "19640 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 209747,
            "unit": "ns/op",
            "extra": "19640 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "19640 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "19640 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 184886,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "17967 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 184886,
            "unit": "ns/op",
            "extra": "17967 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "17967 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "17967 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 387.3,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9306492 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 387.3,
            "unit": "ns/op",
            "extra": "9306492 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9306492 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9306492 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 412.4,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8711208 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 412.4,
            "unit": "ns/op",
            "extra": "8711208 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8711208 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8711208 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 186.7,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19124035 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 186.7,
            "unit": "ns/op",
            "extra": "19124035 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19124035 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19124035 times\n4 procs"
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
          "id": "8171055bfbc4401f1b8cfd374731c0a2dee90222",
          "message": "fix: update Performance workflow to use Go 1.23\n\nUpdated all Go version references from 1.21 to 1.23 to match project requirements (Go 1.23+). This ensures performance benchmarks run on the correct Go version.",
          "timestamp": "2025-10-19T12:42:09+02:00",
          "tree_id": "715a0ce3745e4df6cdd2a95028116e87fde030f3",
          "url": "https://github.com/felixgeelhaar/fortify/commit/8171055bfbc4401f1b8cfd374731c0a2dee90222"
        },
        "date": 1760870740435,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "39664704 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.07,
            "unit": "ns/op",
            "extra": "39664704 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "39664704 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "39664704 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 111.1,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31936940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 111.1,
            "unit": "ns/op",
            "extra": "31936940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31936940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31936940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 235.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17717580 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 235.7,
            "unit": "ns/op",
            "extra": "17717580 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17717580 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17717580 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 137.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26215440 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 137.3,
            "unit": "ns/op",
            "extra": "26215440 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26215440 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26215440 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 163.3,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "22114621 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 163.3,
            "unit": "ns/op",
            "extra": "22114621 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "22114621 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "22114621 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52181696 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.98,
            "unit": "ns/op",
            "extra": "52181696 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52181696 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52181696 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 246.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "14689660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 246.1,
            "unit": "ns/op",
            "extra": "14689660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "14689660 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "14689660 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 23.49,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "153319057 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 23.49,
            "unit": "ns/op",
            "extra": "153319057 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "153319057 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "153319057 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 36.31,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "99001694 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 36.31,
            "unit": "ns/op",
            "extra": "99001694 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "99001694 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "99001694 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 35.98,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "97221096 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 35.98,
            "unit": "ns/op",
            "extra": "97221096 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "97221096 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "97221096 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 60.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59915902 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 60.59,
            "unit": "ns/op",
            "extra": "59915902 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59915902 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59915902 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 70.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "45529299 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 70.52,
            "unit": "ns/op",
            "extra": "45529299 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "45529299 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "45529299 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 58.41,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61804291 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 58.41,
            "unit": "ns/op",
            "extra": "61804291 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61804291 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61804291 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 79.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48402030 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 79.74,
            "unit": "ns/op",
            "extra": "48402030 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48402030 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48402030 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 75.67,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47836432 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 75.67,
            "unit": "ns/op",
            "extra": "47836432 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47836432 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47836432 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 75.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "47203990 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 75.56,
            "unit": "ns/op",
            "extra": "47203990 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "47203990 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "47203990 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 81.45,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "48025500 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 81.45,
            "unit": "ns/op",
            "extra": "48025500 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "48025500 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "48025500 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 149,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25869237 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 149,
            "unit": "ns/op",
            "extra": "25869237 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25869237 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25869237 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill",
            "value": 44.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "80342546 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - ns/op",
            "value": 44.74,
            "unit": "ns/op",
            "extra": "80342546 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "80342546 times\n4 procs"
          },
          {
            "name": "BenchmarkTokenBucketRefill - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "80342546 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 5.559,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "647208614 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 5.559,
            "unit": "ns/op",
            "extra": "647208614 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "647208614 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "647208614 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3200651,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3200651,
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
            "value": 4165849,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4165849,
            "unit": "ns/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "861 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171724312 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.98,
            "unit": "ns/op",
            "extra": "171724312 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171724312 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171724312 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 2.617,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 2.617,
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
            "value": 2.045,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.045,
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
            "value": 29.14,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "123306375 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 29.14,
            "unit": "ns/op",
            "extra": "123306375 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123306375 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123306375 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 28.94,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "123778346 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 28.94,
            "unit": "ns/op",
            "extra": "123778346 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123778346 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123778346 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 311272,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 311272,
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
            "value": 345505,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 345505,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 475.7,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "7547912 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 475.7,
            "unit": "ns/op",
            "extra": "7547912 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "7547912 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7547912 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 481.7,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "7447518 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 481.7,
            "unit": "ns/op",
            "extra": "7447518 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "7447518 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "7447518 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 229.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "15850938 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 229.2,
            "unit": "ns/op",
            "extra": "15850938 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "15850938 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "15850938 times\n4 procs"
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
          "id": "8a71aa9e4e876ae17731410aea350a1844dd7fc4",
          "message": "feat!: refactor rate limiter to pluggable Store interface\n\nBREAKING CHANGES:\n- Remove backends/redis module - users implement custom Store interface\n- Remove examples/backends/redis - see docs/MIGRATION_REDIS.md for patterns\n\nNew Features:\n- Pluggable Store interface for custom backends (Redis, DynamoDB, etc.)\n- Execute(ctx, key, operation) - combines rate limiting with operation\n- ExecuteN(ctx, key, tokens, operation) - multi-token Execute variant\n- Reset(ctx) - clear all rate limiting state (requires Resetter interface)\n- BucketCount() - monitor active buckets (requires BucketCounter interface)\n- HealthCheck support via optional HealthChecker interface\n- FailOpen configuration for availability over consistency\n\nArchitecture:\n- Store interface with AtomicUpdate for thread-safe operations\n- MemoryStore as default in-memory implementation\n- Optional interfaces: HealthChecker, Resetter, BucketCounter\n- Comprehensive error types with errors.Is() support\n\nQuality:\n- 92.3% test coverage with race detection\n- 16 runnable examples documenting all use cases\n- Accurate performance documentation matching benchmarks\n- Security audit completed with A+ grade\n\nPerformance (Apple M1, Go 1.23):\n- Allow(): ~200ns, 74B, 3 allocs\n- Take(): ~197ns, 65B, 3 allocs\n- BucketCount(): ~3ns, 0 allocs\n- Concurrent: ~395ns with contention\n\nMigration: See docs/MIGRATION_REDIS.md for custom Store implementation patterns.",
          "timestamp": "2025-12-08T14:31:32+01:00",
          "tree_id": "b0a62d91d4074e5f0bd37860056c4b691e3f417d",
          "url": "https://github.com/felixgeelhaar/fortify/commit/8a71aa9e4e876ae17731410aea350a1844dd7fc4"
        },
        "date": 1765201173700,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38721457 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.76,
            "unit": "ns/op",
            "extra": "38721457 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38721457 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38721457 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.8,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31483458 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.8,
            "unit": "ns/op",
            "extra": "31483458 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31483458 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31483458 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 193.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18796851 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 193.1,
            "unit": "ns/op",
            "extra": "18796851 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18796851 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18796851 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25660076 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140,
            "unit": "ns/op",
            "extra": "25660076 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25660076 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25660076 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 169.6,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21003156 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 169.6,
            "unit": "ns/op",
            "extra": "21003156 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21003156 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21003156 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52932564 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.37,
            "unit": "ns/op",
            "extra": "52932564 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52932564 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52932564 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 200.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17890515 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 200.7,
            "unit": "ns/op",
            "extra": "17890515 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17890515 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17890515 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.95,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "127686940 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.95,
            "unit": "ns/op",
            "extra": "127686940 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "127686940 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "127686940 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 40.67,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "87085630 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 40.67,
            "unit": "ns/op",
            "extra": "87085630 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "87085630 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "87085630 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.55,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "84618988 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.55,
            "unit": "ns/op",
            "extra": "84618988 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "84618988 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "84618988 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57761635 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.37,
            "unit": "ns/op",
            "extra": "57761635 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57761635 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57761635 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 53.05,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "67028206 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 53.05,
            "unit": "ns/op",
            "extra": "67028206 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "67028206 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "67028206 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60826868 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.85,
            "unit": "ns/op",
            "extra": "60826868 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60826868 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60826868 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 280.1,
            "unit": "ns/op\t      76 B/op\t       3 allocs/op",
            "extra": "12301422 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 280.1,
            "unit": "ns/op",
            "extra": "12301422 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 76,
            "unit": "B/op",
            "extra": "12301422 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12301422 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 245.5,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "14598766 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 245.5,
            "unit": "ns/op",
            "extra": "14598766 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "14598766 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "14598766 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 271.2,
            "unit": "ns/op\t      66 B/op\t       3 allocs/op",
            "extra": "13142652 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 271.2,
            "unit": "ns/op",
            "extra": "13142652 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 66,
            "unit": "B/op",
            "extra": "13142652 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13142652 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 302.8,
            "unit": "ns/op\t      97 B/op\t       4 allocs/op",
            "extra": "11892001 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 302.8,
            "unit": "ns/op",
            "extra": "11892001 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 97,
            "unit": "B/op",
            "extra": "11892001 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "11892001 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 363.3,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "8490823 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 363.3,
            "unit": "ns/op",
            "extra": "8490823 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "8490823 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "8490823 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate",
            "value": 167.5,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "21363700 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - ns/op",
            "value": 167.5,
            "unit": "ns/op",
            "extra": "21363700 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "21363700 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21363700 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate",
            "value": 209.5,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "17280310 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - ns/op",
            "value": 209.5,
            "unit": "ns/op",
            "extra": "17280310 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17280310 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17280310 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute",
            "value": 285.1,
            "unit": "ns/op\t      76 B/op\t       3 allocs/op",
            "extra": "11943834 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - ns/op",
            "value": 285.1,
            "unit": "ns/op",
            "extra": "11943834 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - B/op",
            "value": 76,
            "unit": "B/op",
            "extra": "11943834 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11943834 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited",
            "value": 247.7,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "14442666 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - ns/op",
            "value": 247.7,
            "unit": "ns/op",
            "extra": "14442666 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "14442666 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "14442666 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError",
            "value": 294.2,
            "unit": "ns/op\t      83 B/op\t       3 allocs/op",
            "extra": "11182533 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - ns/op",
            "value": 294.2,
            "unit": "ns/op",
            "extra": "11182533 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - B/op",
            "value": 83,
            "unit": "B/op",
            "extra": "11182533 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11182533 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute",
            "value": 284.2,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "12028570 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - ns/op",
            "value": 284.2,
            "unit": "ns/op",
            "extra": "12028570 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "12028570 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12028570 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect",
            "value": 282.2,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "12652621 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - ns/op",
            "value": 282.2,
            "unit": "ns/op",
            "extra": "12652621 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "12652621 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12652621 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN",
            "value": 276.8,
            "unit": "ns/op\t      66 B/op\t       3 allocs/op",
            "extra": "12918820 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - ns/op",
            "value": 276.8,
            "unit": "ns/op",
            "extra": "12918820 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - B/op",
            "value": 66,
            "unit": "B/op",
            "extra": "12918820 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12918820 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN",
            "value": 273.9,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "12990646 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - ns/op",
            "value": 273.9,
            "unit": "ns/op",
            "extra": "12990646 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "12990646 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12990646 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect",
            "value": 272.7,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "13179231 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - ns/op",
            "value": 272.7,
            "unit": "ns/op",
            "extra": "13179231 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "13179231 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13179231 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty",
            "value": 4.047,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "886169270 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - ns/op",
            "value": 4.047,
            "unit": "ns/op",
            "extra": "886169270 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "886169270 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "886169270 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets",
            "value": 4.048,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889983279 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - ns/op",
            "value": 4.048,
            "unit": "ns/op",
            "extra": "889983279 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889983279 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889983279 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets",
            "value": 4.053,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888726954 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - ns/op",
            "value": 4.053,
            "unit": "ns/op",
            "extra": "888726954 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888726954 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888726954 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets",
            "value": 4.045,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "878038284 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - ns/op",
            "value": 4.045,
            "unit": "ns/op",
            "extra": "878038284 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "878038284 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "878038284 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets",
            "value": 4.049,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888961844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - ns/op",
            "value": 4.049,
            "unit": "ns/op",
            "extra": "888961844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888961844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888961844 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets",
            "value": 2361,
            "unit": "ns/op\t     496 B/op\t       5 allocs/op",
            "extra": "1414719 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - ns/op",
            "value": 2361,
            "unit": "ns/op",
            "extra": "1414719 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - B/op",
            "value": 496,
            "unit": "B/op",
            "extra": "1414719 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1414719 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets",
            "value": 16105,
            "unit": "ns/op\t    4464 B/op\t       8 allocs/op",
            "extra": "208717 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - ns/op",
            "value": 16105,
            "unit": "ns/op",
            "extra": "208717 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - B/op",
            "value": 4464,
            "unit": "B/op",
            "extra": "208717 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "208717 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets",
            "value": 139800,
            "unit": "ns/op\t   35184 B/op\t      11 allocs/op",
            "extra": "26209 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - ns/op",
            "value": 139800,
            "unit": "ns/op",
            "extra": "26209 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - B/op",
            "value": 35184,
            "unit": "B/op",
            "extra": "26209 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "26209 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets",
            "value": 2083512,
            "unit": "ns/op\t  665969 B/op\t      18 allocs/op",
            "extra": "1681 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - ns/op",
            "value": 2083512,
            "unit": "ns/op",
            "extra": "1681 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - B/op",
            "value": 665969,
            "unit": "B/op",
            "extra": "1681 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "1681 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent",
            "value": 9.184,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "391327116 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - ns/op",
            "value": 9.184,
            "unit": "ns/op",
            "extra": "391327116 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "391327116 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "391327116 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute",
            "value": 372.9,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "9203120 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - ns/op",
            "value": 372.9,
            "unit": "ns/op",
            "extra": "9203120 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "9203120 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9203120 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.468,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "478871594 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.468,
            "unit": "ns/op",
            "extra": "478871594 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "478871594 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "478871594 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3195083,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3195083,
            "unit": "ns/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1126 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4246019,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "854 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4246019,
            "unit": "ns/op",
            "extra": "854 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "854 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "854 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "175138752 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.6,
            "unit": "ns/op",
            "extra": "175138752 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "175138752 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "175138752 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.117,
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
            "value": 30.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.85,
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
            "value": 21.04,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171722905 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.04,
            "unit": "ns/op",
            "extra": "171722905 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171722905 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171722905 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 253029,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "13514 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 253029,
            "unit": "ns/op",
            "extra": "13514 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "13514 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13514 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 275433,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "15189 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 275433,
            "unit": "ns/op",
            "extra": "15189 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "15189 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15189 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 384.3,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9397009 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 384.3,
            "unit": "ns/op",
            "extra": "9397009 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9397009 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9397009 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 408.7,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8840816 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 408.7,
            "unit": "ns/op",
            "extra": "8840816 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8840816 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8840816 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 186.4,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19554535 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 186.4,
            "unit": "ns/op",
            "extra": "19554535 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19554535 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19554535 times\n4 procs"
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
          "id": "462ba329fd3306df6549a1462efb1b68f8f171dc",
          "message": "fix: resolve golangci-lint errors for CI\n\n- Change New(Config) to New(*Config) to avoid hugeParam lint error\n- Reorder struct fields for better memory alignment (fieldalignment)\n- Add nolint:gocyclo comments for complex test functions\n- Add nolint:errcheck comments for test setup code\n- Fix godot lint errors (comments ending with periods)\n- Fix builtinShadow by renaming max parameter to maxKeys\n- Update all callers to use &Config{} instead of Config{}\n\nThis is a breaking API change requiring users to update:\n  ratelimit.New(ratelimit.Config{...})\nto:\n  ratelimit.New(&ratelimit.Config{...})",
          "timestamp": "2025-12-08T18:02:54+01:00",
          "tree_id": "b1e9b26b90d37e6fe209da3f5920961958976c06",
          "url": "https://github.com/felixgeelhaar/fortify/commit/462ba329fd3306df6549a1462efb1b68f8f171dc"
        },
        "date": 1765213860831,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.25,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "37592150 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.25,
            "unit": "ns/op",
            "extra": "37592150 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "37592150 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "37592150 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 118.4,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31362057 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 118.4,
            "unit": "ns/op",
            "extra": "31362057 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31362057 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31362057 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 194.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18303601 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 194.1,
            "unit": "ns/op",
            "extra": "18303601 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18303601 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18303601 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 140.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25226416 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 140.4,
            "unit": "ns/op",
            "extra": "25226416 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25226416 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25226416 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 171.9,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20863207 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 171.9,
            "unit": "ns/op",
            "extra": "20863207 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20863207 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20863207 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52668151 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.64,
            "unit": "ns/op",
            "extra": "52668151 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52668151 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52668151 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 201.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17690364 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 201.9,
            "unit": "ns/op",
            "extra": "17690364 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17690364 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17690364 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.21,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "130259290 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.21,
            "unit": "ns/op",
            "extra": "130259290 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "130259290 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "130259290 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.22,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91386598 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.22,
            "unit": "ns/op",
            "extra": "91386598 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91386598 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91386598 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.35,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "82805178 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.35,
            "unit": "ns/op",
            "extra": "82805178 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "82805178 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "82805178 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57271962 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.57,
            "unit": "ns/op",
            "extra": "57271962 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57271962 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57271962 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 53.78,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66973057 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 53.78,
            "unit": "ns/op",
            "extra": "66973057 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66973057 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66973057 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60759440 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.74,
            "unit": "ns/op",
            "extra": "60759440 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60759440 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60759440 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 280.6,
            "unit": "ns/op\t      76 B/op\t       3 allocs/op",
            "extra": "12200911 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 280.6,
            "unit": "ns/op",
            "extra": "12200911 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 76,
            "unit": "B/op",
            "extra": "12200911 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12200911 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 245,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "14580538 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 245,
            "unit": "ns/op",
            "extra": "14580538 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "14580538 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "14580538 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 271.3,
            "unit": "ns/op\t      66 B/op\t       3 allocs/op",
            "extra": "13211776 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 271.3,
            "unit": "ns/op",
            "extra": "13211776 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 66,
            "unit": "B/op",
            "extra": "13211776 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13211776 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 293.6,
            "unit": "ns/op\t      97 B/op\t       4 allocs/op",
            "extra": "11937105 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 293.6,
            "unit": "ns/op",
            "extra": "11937105 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 97,
            "unit": "B/op",
            "extra": "11937105 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "11937105 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 363.3,
            "unit": "ns/op\t      79 B/op\t       3 allocs/op",
            "extra": "9606938 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 363.3,
            "unit": "ns/op",
            "extra": "9606938 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 79,
            "unit": "B/op",
            "extra": "9606938 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9606938 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate",
            "value": 168,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "21312144 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - ns/op",
            "value": 168,
            "unit": "ns/op",
            "extra": "21312144 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "21312144 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21312144 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate",
            "value": 210,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "17111284 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - ns/op",
            "value": 210,
            "unit": "ns/op",
            "extra": "17111284 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17111284 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17111284 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute",
            "value": 283.4,
            "unit": "ns/op\t      76 B/op\t       3 allocs/op",
            "extra": "12140432 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - ns/op",
            "value": 283.4,
            "unit": "ns/op",
            "extra": "12140432 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - B/op",
            "value": 76,
            "unit": "B/op",
            "extra": "12140432 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12140432 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited",
            "value": 246.7,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "14537216 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - ns/op",
            "value": 246.7,
            "unit": "ns/op",
            "extra": "14537216 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "14537216 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "14537216 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError",
            "value": 293.3,
            "unit": "ns/op\t      83 B/op\t       3 allocs/op",
            "extra": "11214264 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - ns/op",
            "value": 293.3,
            "unit": "ns/op",
            "extra": "11214264 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - B/op",
            "value": 83,
            "unit": "B/op",
            "extra": "11214264 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11214264 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute",
            "value": 282.3,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "12149520 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - ns/op",
            "value": 282.3,
            "unit": "ns/op",
            "extra": "12149520 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "12149520 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12149520 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect",
            "value": 280.4,
            "unit": "ns/op\t      73 B/op\t       3 allocs/op",
            "extra": "12743072 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - ns/op",
            "value": 280.4,
            "unit": "ns/op",
            "extra": "12743072 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - B/op",
            "value": 73,
            "unit": "B/op",
            "extra": "12743072 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12743072 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN",
            "value": 271.8,
            "unit": "ns/op\t      66 B/op\t       3 allocs/op",
            "extra": "13108746 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - ns/op",
            "value": 271.8,
            "unit": "ns/op",
            "extra": "13108746 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - B/op",
            "value": 66,
            "unit": "B/op",
            "extra": "13108746 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13108746 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN",
            "value": 271.8,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "13048825 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - ns/op",
            "value": 271.8,
            "unit": "ns/op",
            "extra": "13048825 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "13048825 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13048825 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect",
            "value": 269.5,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "13255866 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - ns/op",
            "value": 269.5,
            "unit": "ns/op",
            "extra": "13255866 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "13255866 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13255866 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty",
            "value": 4.046,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "884302459 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - ns/op",
            "value": 4.046,
            "unit": "ns/op",
            "extra": "884302459 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "884302459 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "884302459 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets",
            "value": 4.048,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "890301819 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - ns/op",
            "value": 4.048,
            "unit": "ns/op",
            "extra": "890301819 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "890301819 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "890301819 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets",
            "value": 4.045,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888917298 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - ns/op",
            "value": 4.045,
            "unit": "ns/op",
            "extra": "888917298 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888917298 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888917298 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets",
            "value": 4.045,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889658686 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - ns/op",
            "value": 4.045,
            "unit": "ns/op",
            "extra": "889658686 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889658686 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889658686 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets",
            "value": 4.107,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889928406 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - ns/op",
            "value": 4.107,
            "unit": "ns/op",
            "extra": "889928406 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889928406 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889928406 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets",
            "value": 2471,
            "unit": "ns/op\t     496 B/op\t       5 allocs/op",
            "extra": "1513324 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - ns/op",
            "value": 2471,
            "unit": "ns/op",
            "extra": "1513324 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - B/op",
            "value": 496,
            "unit": "B/op",
            "extra": "1513324 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1513324 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets",
            "value": 16708,
            "unit": "ns/op\t    4464 B/op\t       8 allocs/op",
            "extra": "216172 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - ns/op",
            "value": 16708,
            "unit": "ns/op",
            "extra": "216172 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - B/op",
            "value": 4464,
            "unit": "B/op",
            "extra": "216172 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "216172 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets",
            "value": 135720,
            "unit": "ns/op\t   35184 B/op\t      11 allocs/op",
            "extra": "27106 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - ns/op",
            "value": 135720,
            "unit": "ns/op",
            "extra": "27106 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - B/op",
            "value": 35184,
            "unit": "B/op",
            "extra": "27106 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "27106 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets",
            "value": 2042118,
            "unit": "ns/op\t  665968 B/op\t      18 allocs/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - ns/op",
            "value": 2042118,
            "unit": "ns/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - B/op",
            "value": 665968,
            "unit": "B/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "1706 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent",
            "value": 9.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "391106088 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - ns/op",
            "value": 9.61,
            "unit": "ns/op",
            "extra": "391106088 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "391106088 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "391106088 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute",
            "value": 379.2,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "9131570 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - ns/op",
            "value": 379.2,
            "unit": "ns/op",
            "extra": "9131570 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "9131570 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9131570 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.498,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481432323 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.498,
            "unit": "ns/op",
            "extra": "481432323 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481432323 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481432323 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3199056,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3199056,
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
            "value": 4239780,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "819 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4239780,
            "unit": "ns/op",
            "extra": "819 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "819 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "819 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.59,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174617793 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.59,
            "unit": "ns/op",
            "extra": "174617793 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174617793 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174617793 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.138,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.138,
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
            "value": 2.488,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.488,
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
            "value": 30.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.89,
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
            "value": 20.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171794721 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 20.95,
            "unit": "ns/op",
            "extra": "171794721 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171794721 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171794721 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 246132,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "13634 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 246132,
            "unit": "ns/op",
            "extra": "13634 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "13634 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13634 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 247225,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "12940 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 247225,
            "unit": "ns/op",
            "extra": "12940 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "12940 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12940 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 382,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9467017 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 382,
            "unit": "ns/op",
            "extra": "9467017 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9467017 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9467017 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 401.8,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8942014 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 401.8,
            "unit": "ns/op",
            "extra": "8942014 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8942014 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8942014 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 188.3,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19582155 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 188.3,
            "unit": "ns/op",
            "extra": "19582155 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19582155 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19582155 times\n4 procs"
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
          "id": "f21228bf385e98e211ed1a2b401e95d36c2df61f",
          "message": "fix: add comprehensive security hardening for rate limiting\n\nSecurity fixes:\n- HIGH-01: Strip IPv6 zone identifiers in KeyFromIP to prevent rate limit bypass\n  (e.g., fe80::1%eth0 and fe80::1%eth1 now treated as same IP)\n- Add Unicode NFC normalization in SanitizeKey to prevent equivalent-string bypass\n- Fix UTF-8 truncation to use rune count instead of byte count\n- Add RFC 7230 header name validation in KeyFromHeader (fail-fast panics)\n- Add sanitizeLogKey to prevent log injection via control characters\n\nNew tests (30+ test cases):\n- TestKeyFromIP_IPv6ZoneBypass: 8 test cases for zone identifier stripping\n- TestSanitizeKey_UnicodeNormalization: 3 test cases for Unicode handling\n- TestKeyFromHeader_Validation: 5 test cases for header validation panics\n- TestKeyFromHeaderWithMaxLen_Validation: 2 test cases\n- TestSanitizeLogKey: 12 test cases for log injection prevention\n\nDocumentation:\n- Add SECURITY.md with vulnerability reporting process\n- Document security measures, configuration guidance, and fixed vulnerabilities\n\nDependencies:\n- Upgrade golang.org/x/text v0.28.0 → v0.32.0 for Unicode normalization",
          "timestamp": "2025-12-08T22:31:25+01:00",
          "tree_id": "9efed4610bd8ed11ded13febcc2ec7b5dc8e1f00",
          "url": "https://github.com/felixgeelhaar/fortify/commit/f21228bf385e98e211ed1a2b401e95d36c2df61f"
        },
        "date": 1765229972520,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 90.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38885053 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 90.75,
            "unit": "ns/op",
            "extra": "38885053 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38885053 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38885053 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 111,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32533191 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 111,
            "unit": "ns/op",
            "extra": "32533191 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32533191 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32533191 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 227,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "16410280 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 227,
            "unit": "ns/op",
            "extra": "16410280 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "16410280 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "16410280 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26316992 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.8,
            "unit": "ns/op",
            "extra": "26316992 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26316992 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26316992 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 164.1,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "21546892 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 164.1,
            "unit": "ns/op",
            "extra": "21546892 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "21546892 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21546892 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 69.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52204051 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 69.01,
            "unit": "ns/op",
            "extra": "52204051 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52204051 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52204051 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 249.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "14588598 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 249.5,
            "unit": "ns/op",
            "extra": "14588598 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "14588598 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "14588598 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 23.01,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "155872314 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 23.01,
            "unit": "ns/op",
            "extra": "155872314 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "155872314 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "155872314 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 35.93,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 35.93,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 35.77,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "98433679 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 35.77,
            "unit": "ns/op",
            "extra": "98433679 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "98433679 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "98433679 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 59.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59547363 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 59.5,
            "unit": "ns/op",
            "extra": "59547363 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59547363 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59547363 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 70.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51322180 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 70.52,
            "unit": "ns/op",
            "extra": "51322180 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51322180 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51322180 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 57.96,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61718354 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 57.96,
            "unit": "ns/op",
            "extra": "61718354 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61718354 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61718354 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 198.5,
            "unit": "ns/op\t      73 B/op\t       3 allocs/op",
            "extra": "16667043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 198.5,
            "unit": "ns/op",
            "extra": "16667043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 73,
            "unit": "B/op",
            "extra": "16667043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16667043 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 162.7,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "22018762 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 162.7,
            "unit": "ns/op",
            "extra": "22018762 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "22018762 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "22018762 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 193.9,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "18233524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 193.9,
            "unit": "ns/op",
            "extra": "18233524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "18233524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18233524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 223,
            "unit": "ns/op\t      97 B/op\t       4 allocs/op",
            "extra": "15997683 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 223,
            "unit": "ns/op",
            "extra": "15997683 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 97,
            "unit": "B/op",
            "extra": "15997683 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "15997683 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 304.1,
            "unit": "ns/op\t      77 B/op\t       3 allocs/op",
            "extra": "11678893 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 304.1,
            "unit": "ns/op",
            "extra": "11678893 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 77,
            "unit": "B/op",
            "extra": "11678893 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11678893 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate",
            "value": 137.3,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "25791991 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - ns/op",
            "value": 137.3,
            "unit": "ns/op",
            "extra": "25791991 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "25791991 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "25791991 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate",
            "value": 212.2,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "16820084 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - ns/op",
            "value": 212.2,
            "unit": "ns/op",
            "extra": "16820084 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "16820084 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16820084 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute",
            "value": 206.6,
            "unit": "ns/op\t      73 B/op\t       3 allocs/op",
            "extra": "16309826 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - ns/op",
            "value": 206.6,
            "unit": "ns/op",
            "extra": "16309826 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - B/op",
            "value": 73,
            "unit": "B/op",
            "extra": "16309826 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16309826 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited",
            "value": 166.6,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "21609715 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - ns/op",
            "value": 166.6,
            "unit": "ns/op",
            "extra": "21609715 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "21609715 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "21609715 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError",
            "value": 210.3,
            "unit": "ns/op\t      78 B/op\t       3 allocs/op",
            "extra": "15057519 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - ns/op",
            "value": 210.3,
            "unit": "ns/op",
            "extra": "15057519 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - B/op",
            "value": 78,
            "unit": "B/op",
            "extra": "15057519 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15057519 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute",
            "value": 204,
            "unit": "ns/op\t      71 B/op\t       3 allocs/op",
            "extra": "16333375 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - ns/op",
            "value": 204,
            "unit": "ns/op",
            "extra": "16333375 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - B/op",
            "value": 71,
            "unit": "B/op",
            "extra": "16333375 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16333375 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect",
            "value": 202.8,
            "unit": "ns/op\t      71 B/op\t       3 allocs/op",
            "extra": "17852923 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - ns/op",
            "value": 202.8,
            "unit": "ns/op",
            "extra": "17852923 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - B/op",
            "value": 71,
            "unit": "B/op",
            "extra": "17852923 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "17852923 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN",
            "value": 195.6,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "17974509 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - ns/op",
            "value": 195.6,
            "unit": "ns/op",
            "extra": "17974509 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "17974509 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "17974509 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN",
            "value": 194.9,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "18147501 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - ns/op",
            "value": 194.9,
            "unit": "ns/op",
            "extra": "18147501 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "18147501 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18147501 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect",
            "value": 193,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "18605654 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - ns/op",
            "value": 193,
            "unit": "ns/op",
            "extra": "18605654 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "18605654 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "18605654 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty",
            "value": 3.258,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - ns/op",
            "value": 3.258,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets",
            "value": 3.255,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - ns/op",
            "value": 3.255,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets",
            "value": 3.263,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - ns/op",
            "value": 3.263,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets",
            "value": 3.259,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - ns/op",
            "value": 3.259,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets",
            "value": 3.254,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - ns/op",
            "value": 3.254,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets",
            "value": 2640,
            "unit": "ns/op\t     496 B/op\t       5 allocs/op",
            "extra": "1447504 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - ns/op",
            "value": 2640,
            "unit": "ns/op",
            "extra": "1447504 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - B/op",
            "value": 496,
            "unit": "B/op",
            "extra": "1447504 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1447504 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets",
            "value": 18271,
            "unit": "ns/op\t    4464 B/op\t       8 allocs/op",
            "extra": "210264 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - ns/op",
            "value": 18271,
            "unit": "ns/op",
            "extra": "210264 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - B/op",
            "value": 4464,
            "unit": "B/op",
            "extra": "210264 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "210264 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets",
            "value": 160290,
            "unit": "ns/op\t   35184 B/op\t      11 allocs/op",
            "extra": "22621 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - ns/op",
            "value": 160290,
            "unit": "ns/op",
            "extra": "22621 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - B/op",
            "value": 35184,
            "unit": "B/op",
            "extra": "22621 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "22621 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets",
            "value": 2592065,
            "unit": "ns/op\t  665969 B/op\t      18 allocs/op",
            "extra": "1410 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - ns/op",
            "value": 2592065,
            "unit": "ns/op",
            "extra": "1410 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - B/op",
            "value": 665969,
            "unit": "B/op",
            "extra": "1410 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "1410 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent",
            "value": 11.48,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "313437658 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - ns/op",
            "value": 11.48,
            "unit": "ns/op",
            "extra": "313437658 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "313437658 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "313437658 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute",
            "value": 327.4,
            "unit": "ns/op\t      78 B/op\t       3 allocs/op",
            "extra": "10687154 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - ns/op",
            "value": 327.4,
            "unit": "ns/op",
            "extra": "10687154 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - B/op",
            "value": 78,
            "unit": "B/op",
            "extra": "10687154 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10687154 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 5.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "647460204 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 5.56,
            "unit": "ns/op",
            "extra": "647460204 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "647460204 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "647460204 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3209493,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3209493,
            "unit": "ns/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4243528,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "860 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4243528,
            "unit": "ns/op",
            "extra": "860 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "860 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "860 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171843435 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.97,
            "unit": "ns/op",
            "extra": "171843435 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171843435 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171843435 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 2.728,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 2.728,
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
            "extra": "123099852 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 29.23,
            "unit": "ns/op",
            "extra": "123099852 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "123099852 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "123099852 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail",
            "value": 29.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "124304227 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 29.02,
            "unit": "ns/op",
            "extra": "124304227 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "124304227 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "124304227 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 346159,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "9997 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 346159,
            "unit": "ns/op",
            "extra": "9997 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "9997 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9997 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 390059,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "9279 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 390059,
            "unit": "ns/op",
            "extra": "9279 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "9279 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "9279 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 462.7,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "7839580 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 462.7,
            "unit": "ns/op",
            "extra": "7839580 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "7839580 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7839580 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 477.9,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "7582860 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 477.9,
            "unit": "ns/op",
            "extra": "7582860 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "7582860 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "7582860 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 227.2,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "15819220 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 227.2,
            "unit": "ns/op",
            "extra": "15819220 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "15819220 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "15819220 times\n4 procs"
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
          "id": "166a415597f4ea03089e083da32189e4cd4a90b6",
          "message": "perf: add ASCII fast paths for key sanitization\n\nOptimize SanitizeKey and sanitizeLogKey with tiered fast paths:\n- Clean ASCII keys (32-126) return immediately with zero allocation\n- ASCII-only keys use byte-level operations without Unicode overhead\n- Full Unicode NFC normalization reserved for non-ASCII strings\n\nThis addresses performance regression alerts from security hardening\nwhile maintaining all security properties for non-ASCII inputs.",
          "timestamp": "2025-12-08T23:14:27+01:00",
          "tree_id": "0c861abd8e057fac84d8541a204bb18d435dfc33",
          "url": "https://github.com/felixgeelhaar/fortify/commit/166a415597f4ea03089e083da32189e4cd4a90b6"
        },
        "date": 1765232884859,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38127805 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.8,
            "unit": "ns/op",
            "extra": "38127805 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38127805 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38127805 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 114.4,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "30662940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 114.4,
            "unit": "ns/op",
            "extra": "30662940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "30662940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "30662940 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 196.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18177806 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 196.4,
            "unit": "ns/op",
            "extra": "18177806 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18177806 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18177806 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25755850 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.5,
            "unit": "ns/op",
            "extra": "25755850 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25755850 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25755850 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 172,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20967171 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 172,
            "unit": "ns/op",
            "extra": "20967171 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20967171 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20967171 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 68.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52594682 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 68.71,
            "unit": "ns/op",
            "extra": "52594682 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52594682 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52594682 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 202.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17651528 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 202.6,
            "unit": "ns/op",
            "extra": "17651528 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17651528 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17651528 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 28.2,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "130807978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 28.2,
            "unit": "ns/op",
            "extra": "130807978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "130807978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "130807978 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 37.83,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "91254673 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 37.83,
            "unit": "ns/op",
            "extra": "91254673 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "91254673 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "91254673 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 41.46,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "83887411 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 41.46,
            "unit": "ns/op",
            "extra": "83887411 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "83887411 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "83887411 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57218126 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.24,
            "unit": "ns/op",
            "extra": "57218126 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57218126 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57218126 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 52.78,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65921872 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 52.78,
            "unit": "ns/op",
            "extra": "65921872 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65921872 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65921872 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.45,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59715298 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.45,
            "unit": "ns/op",
            "extra": "59715298 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59715298 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59715298 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 227,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "14806676 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 227,
            "unit": "ns/op",
            "extra": "14806676 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "14806676 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14806676 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 191.2,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "18720987 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 191.2,
            "unit": "ns/op",
            "extra": "18720987 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "18720987 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "18720987 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 220.9,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "16029274 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 220.9,
            "unit": "ns/op",
            "extra": "16029274 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "16029274 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16029274 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 242.7,
            "unit": "ns/op\t      97 B/op\t       4 allocs/op",
            "extra": "14361524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 242.7,
            "unit": "ns/op",
            "extra": "14361524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 97,
            "unit": "B/op",
            "extra": "14361524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "14361524 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 309.3,
            "unit": "ns/op\t      77 B/op\t       3 allocs/op",
            "extra": "11244784 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 309.3,
            "unit": "ns/op",
            "extra": "11244784 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 77,
            "unit": "B/op",
            "extra": "11244784 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11244784 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate",
            "value": 167.3,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "21547275 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - ns/op",
            "value": 167.3,
            "unit": "ns/op",
            "extra": "21547275 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "21547275 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21547275 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate",
            "value": 211.7,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "16881238 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - ns/op",
            "value": 211.7,
            "unit": "ns/op",
            "extra": "16881238 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "16881238 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16881238 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute",
            "value": 230.8,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "14624277 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - ns/op",
            "value": 230.8,
            "unit": "ns/op",
            "extra": "14624277 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "14624277 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14624277 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited",
            "value": 195.6,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "18281116 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - ns/op",
            "value": 195.6,
            "unit": "ns/op",
            "extra": "18281116 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "18281116 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "18281116 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError",
            "value": 238.8,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "13344118 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - ns/op",
            "value": 238.8,
            "unit": "ns/op",
            "extra": "13344118 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13344118 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13344118 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute",
            "value": 229.5,
            "unit": "ns/op\t      72 B/op\t       3 allocs/op",
            "extra": "14714604 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - ns/op",
            "value": 229.5,
            "unit": "ns/op",
            "extra": "14714604 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "14714604 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14714604 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect",
            "value": 230.5,
            "unit": "ns/op\t      72 B/op\t       3 allocs/op",
            "extra": "15700814 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - ns/op",
            "value": 230.5,
            "unit": "ns/op",
            "extra": "15700814 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "15700814 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15700814 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN",
            "value": 221.4,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "15996499 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - ns/op",
            "value": 221.4,
            "unit": "ns/op",
            "extra": "15996499 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "15996499 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15996499 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN",
            "value": 221.1,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "15993979 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - ns/op",
            "value": 221.1,
            "unit": "ns/op",
            "extra": "15993979 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "15993979 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15993979 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect",
            "value": 219.5,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "16314249 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - ns/op",
            "value": 219.5,
            "unit": "ns/op",
            "extra": "16314249 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "16314249 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16314249 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty",
            "value": 4.052,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889065844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - ns/op",
            "value": 4.052,
            "unit": "ns/op",
            "extra": "889065844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889065844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889065844 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets",
            "value": 4.044,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889051380 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - ns/op",
            "value": 4.044,
            "unit": "ns/op",
            "extra": "889051380 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889051380 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889051380 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets",
            "value": 4.046,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888225511 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - ns/op",
            "value": 4.046,
            "unit": "ns/op",
            "extra": "888225511 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888225511 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888225511 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets",
            "value": 4.05,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889087784 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - ns/op",
            "value": 4.05,
            "unit": "ns/op",
            "extra": "889087784 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889087784 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889087784 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets",
            "value": 4.047,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888662908 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - ns/op",
            "value": 4.047,
            "unit": "ns/op",
            "extra": "888662908 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888662908 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888662908 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets",
            "value": 2608,
            "unit": "ns/op\t     496 B/op\t       5 allocs/op",
            "extra": "1314411 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - ns/op",
            "value": 2608,
            "unit": "ns/op",
            "extra": "1314411 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - B/op",
            "value": 496,
            "unit": "B/op",
            "extra": "1314411 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1314411 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets",
            "value": 17560,
            "unit": "ns/op\t    4464 B/op\t       8 allocs/op",
            "extra": "187552 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - ns/op",
            "value": 17560,
            "unit": "ns/op",
            "extra": "187552 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - B/op",
            "value": 4464,
            "unit": "B/op",
            "extra": "187552 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "187552 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets",
            "value": 142566,
            "unit": "ns/op\t   35184 B/op\t      11 allocs/op",
            "extra": "25066 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - ns/op",
            "value": 142566,
            "unit": "ns/op",
            "extra": "25066 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - B/op",
            "value": 35184,
            "unit": "B/op",
            "extra": "25066 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "25066 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets",
            "value": 2109228,
            "unit": "ns/op\t  665968 B/op\t      18 allocs/op",
            "extra": "1611 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - ns/op",
            "value": 2109228,
            "unit": "ns/op",
            "extra": "1611 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - B/op",
            "value": 665968,
            "unit": "B/op",
            "extra": "1611 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "1611 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent",
            "value": 9.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "398675841 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - ns/op",
            "value": 9.01,
            "unit": "ns/op",
            "extra": "398675841 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "398675841 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "398675841 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute",
            "value": 325.3,
            "unit": "ns/op\t      78 B/op\t       3 allocs/op",
            "extra": "10585821 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - ns/op",
            "value": 325.3,
            "unit": "ns/op",
            "extra": "10585821 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - B/op",
            "value": 78,
            "unit": "B/op",
            "extra": "10585821 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10585821 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.478,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481943841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.478,
            "unit": "ns/op",
            "extra": "481943841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481943841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481943841 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3200773,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3200773,
            "unit": "ns/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1125 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter",
            "value": 4287621,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4287621,
            "unit": "ns/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174821001 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.57,
            "unit": "ns/op",
            "extra": "174821001 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174821001 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174821001 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.118,
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
            "value": 2.491,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/constant - ns/op",
            "value": 2.491,
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
            "value": 30.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.85,
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
            "value": 21.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171676197 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.12,
            "unit": "ns/op",
            "extra": "171676197 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171676197 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171676197 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 288163,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "12492 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 288163,
            "unit": "ns/op",
            "extra": "12492 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "12492 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "12492 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 297213,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "10915 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 297213,
            "unit": "ns/op",
            "extra": "10915 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "10915 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10915 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 384.6,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9410442 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 384.6,
            "unit": "ns/op",
            "extra": "9410442 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9410442 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9410442 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 407.3,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8860705 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 407.3,
            "unit": "ns/op",
            "extra": "8860705 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8860705 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8860705 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 185,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19636940 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 185,
            "unit": "ns/op",
            "extra": "19636940 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19636940 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19636940 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "committer": {
            "name": "Felix Geelhaar",
            "username": "felixgeelhaar",
            "email": "felix@felixgeelhaar.de"
          },
          "id": "f7b754f623168625dd3dcd0343919f04b5988246",
          "message": "ci: change performance workflow to weekly schedule\n\nRun benchmarks weekly (Sunday midnight UTC) instead of on every push/PR.\nThis reduces CI costs while still tracking performance trends over time.\n\nManual trigger via workflow_dispatch remains available for on-demand runs.",
          "timestamp": "2025-12-09T08:21:38Z",
          "url": "https://github.com/felixgeelhaar/fortify/commit/f7b754f623168625dd3dcd0343919f04b5988246"
        },
        "date": 1765671316966,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBulkheadSuccess",
            "value": 93.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "38537683 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - ns/op",
            "value": 93.2,
            "unit": "ns/op",
            "extra": "38537683 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "38537683 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "38537683 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue",
            "value": 115,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "30830684 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - ns/op",
            "value": 115,
            "unit": "ns/op",
            "extra": "30830684 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "30830684 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadWithQueue - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "30830684 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent",
            "value": 201,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18704236 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - ns/op",
            "value": 201,
            "unit": "ns/op",
            "extra": "18704236 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18704236 times\n4 procs"
          },
          {
            "name": "BenchmarkBulkheadConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18704236 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess",
            "value": 139.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25797565 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - ns/op",
            "value": 139.7,
            "unit": "ns/op",
            "extra": "25797565 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "25797565 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerSuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "25797565 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure",
            "value": 172.2,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "20737000 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - ns/op",
            "value": 172.2,
            "unit": "ns/op",
            "extra": "20737000 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "20737000 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerFailure - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20737000 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen",
            "value": 69.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "52662536 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - ns/op",
            "value": 69.56,
            "unit": "ns/op",
            "extra": "52662536 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "52662536 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerOpen - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "52662536 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent",
            "value": 203.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "17388553 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - ns/op",
            "value": 203.4,
            "unit": "ns/op",
            "extra": "17388553 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "17388553 times\n4 procs"
          },
          {
            "name": "BenchmarkCircuitBreakerConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "17388553 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess",
            "value": 27.95,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "128438697 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - ns/op",
            "value": 27.95,
            "unit": "ns/op",
            "extra": "128438697 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "128438697 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_PrimarySuccess - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "128438697 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered",
            "value": 38.94,
            "unit": "ns/op\t      24 B/op\t       1 allocs/op",
            "extra": "89684521 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - ns/op",
            "value": 38.94,
            "unit": "ns/op",
            "extra": "89684521 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - B/op",
            "value": 24,
            "unit": "B/op",
            "extra": "89684521 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_FallbackTriggered - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "89684521 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks",
            "value": 42.08,
            "unit": "ns/op\t      20 B/op\t       1 allocs/op",
            "extra": "85497928 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - ns/op",
            "value": 42.08,
            "unit": "ns/op",
            "extra": "85497928 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - B/op",
            "value": 20,
            "unit": "B/op",
            "extra": "85497928 times\n4 procs"
          },
          {
            "name": "BenchmarkFallback_WithCallbacks - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "85497928 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest",
            "value": 62.35,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "58233512 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - ns/op",
            "value": 62.35,
            "unit": "ns/op",
            "extra": "58233512 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "58233512 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordCircuitBreakerRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "58233512 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts",
            "value": 53.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66850508 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - ns/op",
            "value": 53.68,
            "unit": "ns/op",
            "extra": "66850508 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66850508 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRetryAttempts - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66850508 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed",
            "value": 59.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61016181 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - ns/op",
            "value": 59.75,
            "unit": "ns/op",
            "extra": "61016181 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61016181 times\n4 procs"
          },
          {
            "name": "BenchmarkRecordRateLimitAllowed - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61016181 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow",
            "value": 230.1,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "14709907 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - ns/op",
            "value": 230.1,
            "unit": "ns/op",
            "extra": "14709907 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "14709907 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllow - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14709907 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited",
            "value": 194.7,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "18581335 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - ns/op",
            "value": 194.7,
            "unit": "ns/op",
            "extra": "18581335 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "18581335 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterAllowRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "18581335 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake",
            "value": 222.8,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "16043660 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - ns/op",
            "value": 222.8,
            "unit": "ns/op",
            "extra": "16043660 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "16043660 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterTake - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16043660 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys",
            "value": 252.8,
            "unit": "ns/op\t      97 B/op\t       4 allocs/op",
            "extra": "14273390 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - ns/op",
            "value": 252.8,
            "unit": "ns/op",
            "extra": "14273390 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - B/op",
            "value": 97,
            "unit": "B/op",
            "extra": "14273390 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterMultipleKeys - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "14273390 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent",
            "value": 303,
            "unit": "ns/op\t      77 B/op\t       3 allocs/op",
            "extra": "11526093 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - ns/op",
            "value": 303,
            "unit": "ns/op",
            "extra": "11526093 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - B/op",
            "value": 77,
            "unit": "B/op",
            "extra": "11526093 times\n4 procs"
          },
          {
            "name": "BenchmarkRateLimiterConcurrent - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11526093 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate",
            "value": 168,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "21249276 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - ns/op",
            "value": 168,
            "unit": "ns/op",
            "extra": "21249276 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "21249276 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreAtomicUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "21249276 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate",
            "value": 210,
            "unit": "ns/op\t      32 B/op\t       1 allocs/op",
            "extra": "17109040 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - ns/op",
            "value": 210,
            "unit": "ns/op",
            "extra": "17109040 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "17109040 times\n4 procs"
          },
          {
            "name": "BenchmarkMemoryStoreConcurrentUpdate - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17109040 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute",
            "value": 232.3,
            "unit": "ns/op\t      74 B/op\t       3 allocs/op",
            "extra": "14462131 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - ns/op",
            "value": 232.3,
            "unit": "ns/op",
            "extra": "14462131 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - B/op",
            "value": 74,
            "unit": "B/op",
            "extra": "14462131 times\n4 procs"
          },
          {
            "name": "BenchmarkExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14462131 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited",
            "value": 196.9,
            "unit": "ns/op\t      33 B/op\t       2 allocs/op",
            "extra": "18147018 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - ns/op",
            "value": 196.9,
            "unit": "ns/op",
            "extra": "18147018 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - B/op",
            "value": 33,
            "unit": "B/op",
            "extra": "18147018 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteRateLimited - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "18147018 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError",
            "value": 240.6,
            "unit": "ns/op\t      80 B/op\t       3 allocs/op",
            "extra": "13155380 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - ns/op",
            "value": 240.6,
            "unit": "ns/op",
            "extra": "13155380 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - B/op",
            "value": 80,
            "unit": "B/op",
            "extra": "13155380 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteWithError - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "13155380 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute",
            "value": 231.4,
            "unit": "ns/op\t      72 B/op\t       3 allocs/op",
            "extra": "14477140 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - ns/op",
            "value": 231.4,
            "unit": "ns/op",
            "extra": "14477140 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "14477140 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/Execute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "14477140 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect",
            "value": 229.7,
            "unit": "ns/op\t      72 B/op\t       3 allocs/op",
            "extra": "15714284 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - ns/op",
            "value": 229.7,
            "unit": "ns/op",
            "extra": "15714284 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "15714284 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteVsAllowDirect/AllowDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15714284 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN",
            "value": 223.8,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "15848298 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - ns/op",
            "value": 223.8,
            "unit": "ns/op",
            "extra": "15848298 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "15848298 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15848298 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN",
            "value": 223.5,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "15863854 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - ns/op",
            "value": 223.5,
            "unit": "ns/op",
            "extra": "15863854 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "15863854 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/ExecuteN - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "15863854 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect",
            "value": 221.6,
            "unit": "ns/op\t      65 B/op\t       3 allocs/op",
            "extra": "16183087 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - ns/op",
            "value": 221.6,
            "unit": "ns/op",
            "extra": "16183087 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - B/op",
            "value": 65,
            "unit": "B/op",
            "extra": "16183087 times\n4 procs"
          },
          {
            "name": "BenchmarkExecuteNVsTakeDirect/TakeDirect - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "16183087 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty",
            "value": 4.047,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888579660 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - ns/op",
            "value": 4.047,
            "unit": "ns/op",
            "extra": "888579660 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888579660 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/Empty - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888579660 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets",
            "value": 4.051,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "889358080 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - ns/op",
            "value": 4.051,
            "unit": "ns/op",
            "extra": "889358080 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "889358080 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "889358080 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets",
            "value": 4.045,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888904566 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - ns/op",
            "value": 4.045,
            "unit": "ns/op",
            "extra": "888904566 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888904566 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/100Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888904566 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets",
            "value": 4.058,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888119352 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - ns/op",
            "value": 4.058,
            "unit": "ns/op",
            "extra": "888119352 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888119352 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/1000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888119352 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets",
            "value": 4.047,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "888262201 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - ns/op",
            "value": 4.047,
            "unit": "ns/op",
            "extra": "888262201 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "888262201 times\n4 procs"
          },
          {
            "name": "BenchmarkBucketCount/10000Buckets - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "888262201 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets",
            "value": 2386,
            "unit": "ns/op\t     496 B/op\t       5 allocs/op",
            "extra": "1481668 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - ns/op",
            "value": 2386,
            "unit": "ns/op",
            "extra": "1481668 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - B/op",
            "value": 496,
            "unit": "B/op",
            "extra": "1481668 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10Buckets - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "1481668 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets",
            "value": 16969,
            "unit": "ns/op\t    4464 B/op\t       8 allocs/op",
            "extra": "217844 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - ns/op",
            "value": 16969,
            "unit": "ns/op",
            "extra": "217844 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - B/op",
            "value": 4464,
            "unit": "B/op",
            "extra": "217844 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/100Buckets - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "217844 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets",
            "value": 147865,
            "unit": "ns/op\t   35184 B/op\t      11 allocs/op",
            "extra": "23792 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - ns/op",
            "value": 147865,
            "unit": "ns/op",
            "extra": "23792 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - B/op",
            "value": 35184,
            "unit": "B/op",
            "extra": "23792 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/1000Buckets - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "23792 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets",
            "value": 2172213,
            "unit": "ns/op\t  665968 B/op\t      18 allocs/op",
            "extra": "1555 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - ns/op",
            "value": 2172213,
            "unit": "ns/op",
            "extra": "1555 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - B/op",
            "value": 665968,
            "unit": "B/op",
            "extra": "1555 times\n4 procs"
          },
          {
            "name": "BenchmarkReset/10000Buckets - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "1555 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent",
            "value": 9.015,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "399641935 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - ns/op",
            "value": 9.015,
            "unit": "ns/op",
            "extra": "399641935 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "399641935 times\n4 procs"
          },
          {
            "name": "BenchmarkResetConcurrent - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "399641935 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute",
            "value": 320.3,
            "unit": "ns/op\t      78 B/op\t       3 allocs/op",
            "extra": "10250770 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - ns/op",
            "value": 320.3,
            "unit": "ns/op",
            "extra": "10250770 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - B/op",
            "value": 78,
            "unit": "B/op",
            "extra": "10250770 times\n4 procs"
          },
          {
            "name": "BenchmarkConcurrentExecute - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "10250770 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess",
            "value": 7.481,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "481512298 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - ns/op",
            "value": 7.481,
            "unit": "ns/op",
            "extra": "481512298 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "481512298 times\n4 procs"
          },
          {
            "name": "BenchmarkRetrySuccess - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "481512298 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure",
            "value": 3199553,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "1124 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryFailure - ns/op",
            "value": 3199553,
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
            "value": 4240811,
            "unit": "ns/op\t     568 B/op\t      10 allocs/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - ns/op",
            "value": 4240811,
            "unit": "ns/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - B/op",
            "value": 568,
            "unit": "B/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkRetryWithJitter - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "838 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential",
            "value": 20.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "174921943 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - ns/op",
            "value": 20.61,
            "unit": "ns/op",
            "extra": "174921943 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "174921943 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/exponential - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "174921943 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear",
            "value": 3.119,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/linear - ns/op",
            "value": 3.119,
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
            "value": 30.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkBackoffCalculation/with_jitter - ns/op",
            "value": 30.82,
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
            "value": 21.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "171489885 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - ns/op",
            "value": 21.03,
            "unit": "ns/op",
            "extra": "171489885 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "171489885 times\n4 procs"
          },
          {
            "name": "BenchmarkErrorInjector_ShouldFail - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "171489885 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay",
            "value": 214381,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "17418 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - ns/op",
            "value": 214381,
            "unit": "ns/op",
            "extra": "17418 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "17418 times\n4 procs"
          },
          {
            "name": "BenchmarkLatencyInjector_Delay - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "17418 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call",
            "value": 224302,
            "unit": "ns/op\t     248 B/op\t       3 allocs/op",
            "extra": "20749 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - ns/op",
            "value": 224302,
            "unit": "ns/op",
            "extra": "20749 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - B/op",
            "value": 248,
            "unit": "B/op",
            "extra": "20749 times\n4 procs"
          },
          {
            "name": "BenchmarkFlakeyService_Call - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "20749 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess",
            "value": 385.6,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "9372111 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - ns/op",
            "value": 385.6,
            "unit": "ns/op",
            "extra": "9372111 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "9372111 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutSuccess - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "9372111 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation",
            "value": 408.7,
            "unit": "ns/op\t     288 B/op\t       5 allocs/op",
            "extra": "8885312 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - ns/op",
            "value": 408.7,
            "unit": "ns/op",
            "extra": "8885312 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - B/op",
            "value": 288,
            "unit": "B/op",
            "extra": "8885312 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutWithShortOperation - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "8885312 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent",
            "value": 183.7,
            "unit": "ns/op\t     272 B/op\t       4 allocs/op",
            "extra": "19539234 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - ns/op",
            "value": 183.7,
            "unit": "ns/op",
            "extra": "19539234 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - B/op",
            "value": 272,
            "unit": "B/op",
            "extra": "19539234 times\n4 procs"
          },
          {
            "name": "BenchmarkTimeoutConcurrent - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "19539234 times\n4 procs"
          }
        ]
      }
    ]
  }
}