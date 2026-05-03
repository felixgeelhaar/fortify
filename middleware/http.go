package middleware

import (
	"context"
	"net/http"
)

// httpRoundTripper wraps an inner http.RoundTripper with a resilience
// middleware chain. The chain runs around each round-trip; the inner
// RoundTripper performs the actual HTTP transport.
type httpRoundTripper struct {
	inner http.RoundTripper
	chain *Chain[*http.Response]
}

// RoundTrip implements http.RoundTripper.
func (h *httpRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return h.chain.Execute(req.Context(), func(ctx context.Context) (*http.Response, error) {
		return h.inner.RoundTrip(req.WithContext(ctx))
	})
}

// HTTPRoundTripper wraps the given RoundTripper with the HTTPClient preset
// chain (CircuitBreaker → Retry → Timeout). If inner is nil, http.DefaultTransport
// is used.
//
// Use it as the Transport on an http.Client:
//
//	rt, err := middleware.HTTPRoundTripper(nil, middleware.HTTPClientConfig{
//	    Timeout:    5 * time.Second,
//	    MaxRetries: 3,
//	})
//	if err != nil {
//	    return err
//	}
//	client := &http.Client{Transport: rt}
//
// Returns an error if Timeout is zero or negative.
//
// IMPORTANT: HTTP retry semantics. The retry policy will retry idempotent
// methods (GET, HEAD, PUT, DELETE) freely. Non-idempotent methods (POST,
// PATCH) carry side-effect risk on retry; the wrapped RoundTripper retries
// any non-CB / non-context error indiscriminately. If your POST handlers
// are not idempotent, override IsRetryable on the underlying retry config
// by building your own Chain with `WithRetry` instead of using this preset.
func HTTPRoundTripper(inner http.RoundTripper, cfg HTTPClientConfig) (http.RoundTripper, error) {
	if inner == nil {
		inner = http.DefaultTransport
	}
	chain, err := HTTPClient(cfg)
	if err != nil {
		return nil, err
	}
	return &httpRoundTripper{inner: inner, chain: chain}, nil
}

// HTTPRoundTripperFromChain wraps inner with an arbitrary user-supplied chain.
// Useful when the HTTPClient preset doesn't fit and you've built a custom
// Chain[*http.Response] yourself.
func HTTPRoundTripperFromChain(inner http.RoundTripper, chain *Chain[*http.Response]) http.RoundTripper {
	if inner == nil {
		inner = http.DefaultTransport
	}
	return &httpRoundTripper{inner: inner, chain: chain}
}
