# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.x     | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

If you discover a security vulnerability in Fortify, please report it by emailing:

**security@felixgeelhaar.de**

Please include the following information in your report:

- Type of vulnerability (e.g., rate limit bypass, DoS, injection)
- Full path to the source file(s) related to the vulnerability
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact assessment of the vulnerability

You should receive an initial response within 48 hours acknowledging receipt. We will work with you to understand and validate the issue, and aim to release a fix within 90 days depending on severity.

## Security Measures

Fortify implements several security measures to protect against common attack vectors:

### Rate Limit Bypass Prevention

1. **IPv6 Zone Identifier Stripping**: Prevents bypass by varying zone identifiers (e.g., `fe80::1%eth0` vs `fe80::1%eth1`)
2. **Unicode Normalization (NFC)**: Prevents bypass using Unicode equivalent strings
3. **Key Sanitization**: Validates and sanitizes all rate limiting keys

### Injection Prevention

1. **Log Injection Protection**: Control characters are sanitized from log output
2. **HTTP Header Validation**: Header names validated per RFC 7230 at construction time

### Denial of Service Prevention

1. **Key Length Limits**: Prevents memory exhaustion from long keys (`DefaultMaxKeyLength: 256`)
2. **Maximum Key Count**: Limits total buckets in memory store (`DefaultMaxKeys: 100,000`)
3. **Token Limits**: `MaxTokensPerRequest` prevents integer overflow attacks
4. **Rate/Burst Caps**: Configuration values capped to prevent overflow

### Concurrency Safety

1. **Thread-Safe Operations**: All public methods are safe for concurrent use
2. **Atomic Operations**: Uses `sync.Map`, `atomic.Int64`, and per-key mutexes
3. **Race Detector**: All tests pass with Go's race detector enabled

## Security Configuration

### Fail-Open vs Fail-Closed

```go
// Fail-closed (default, recommended for security-critical applications)
limiter := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    200,
    FailOpen: false, // Deny requests on storage errors
})

// Fail-open (for high-availability scenarios)
limiter := ratelimit.New(&ratelimit.Config{
    Rate:     100,
    Burst:    200,
    FailOpen: true, // Allow requests on storage errors
})
```

### Reset() Authorization

The `Reset()` method clears all rate limiting state. Applications **MUST** implement authorization checks before calling `Reset()`:

```go
// WRONG: Exposed without authorization
http.HandleFunc("/admin/reset", func(w http.ResponseWriter, r *http.Request) {
    limiter.Reset(r.Context()) // SECURITY RISK: No authorization!
})

// CORRECT: With proper authorization
http.HandleFunc("/admin/reset", func(w http.ResponseWriter, r *http.Request) {
    if !isAdmin(r) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    limiter.Reset(r.Context())
})
```

## Known Security Considerations

### Clock Manipulation

The token bucket algorithm relies on system time. System clock adjustments (NTP, manual changes) may affect rate limiting accuracy:

- **Forward jump**: May allow temporary rate limit bypass
- **Backward jump**: May temporarily block legitimate requests

Mitigation: The implementation caps elapsed time to prevent extreme effects from clock skew.

### Memory Store Limitations

The in-memory store is suitable for single-instance deployments. For distributed systems:

- Implement a custom `Store` using Redis, DynamoDB, or similar
- Consider fail-open behavior for store availability issues
- Monitor key count to prevent exhaustion attacks

## Fixed Vulnerabilities

### v1.1.0 (December 2024)

- **HIGH**: IPv6 zone identifier bypass in `KeyFromIP` - attackers could bypass rate limits by varying IPv6 zone identifiers
- **MEDIUM**: Unicode normalization missing in `SanitizeKey` - equivalent Unicode strings could bypass rate limits
- **MEDIUM**: UTF-8 truncation issue - byte-based truncation could split multi-byte characters
- **LOW**: Log injection via control characters in keys

### Dependencies

- Updated `go-redis` to v9.7.3 to address CVE-2025-29923 (Redis backend module)

## Security Audits

The codebase has been reviewed for:

- OWASP Top 10 vulnerabilities
- CWE Top 25 issues
- Go-specific security patterns
- Concurrency safety (race detector)

See `SECURITY_AUDIT_REPORT.md` for the latest audit report.

## Responsible Disclosure

We believe in responsible disclosure and will:

1. Acknowledge receipt of vulnerability reports within 48 hours
2. Provide regular updates on remediation progress
3. Credit reporters in release notes (unless anonymity is requested)
4. Not pursue legal action against good-faith security researchers

Thank you for helping keep Fortify secure!
