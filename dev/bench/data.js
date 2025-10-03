window.BENCHMARK_DATA = {
  "lastUpdate": 1759494437560,
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
      }
    ]
  }
}