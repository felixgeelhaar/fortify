package http

import (
	"bufio"
	"context"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/felixgeelhaar/fortify/circuitbreaker"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/felixgeelhaar/fortify/streamtimeout"
)

// CircuitBreakerStream pairs a CircuitBreaker with a StreamTimeout so
// long-lived responses (Server-Sent Events, chunked transfer, gRPC-Web
// streaming) get per-chunk health signals instead of a single
// after-the-fact outcome.
//
// The middleware:
//
//   - forwards http.Flusher / http.Hijacker / http.Pusher from the
//     underlying ResponseWriter so SSE clients receive each frame as
//     soon as the handler flushes it,
//   - calls streamtimeout.Mark on every Write so the FirstByte and
//     Idle watchdogs observe the live data path, and
//   - treats a fired watchdog as a breaker failure: the configured
//     CircuitBreaker sees ErrTimeout and counts the call accordingly.
//
// The handler MUST honour ctx cancellation. Once a watchdog fires the
// streamtimeout cancels the context passed to ServeHTTP; a handler
// that ignores ctx will keep streaming until upstream closes, leaving
// the breaker unable to abort the in-flight call (the breaker decision
// still records correctly, but the client sees the truncated stream).
//
// Returns an error when cfg has no positive timeout — a zero-config
// StreamTimeout would be a no-op and defeat the purpose of this
// middleware. The same Config rules as streamtimeout.New apply.
//
// Example: SSE pass-through with 5s TTFB, 2s idle, 5m total cap.
//
//	cb := circuitbreaker.New[*http.Response](circuitbreaker.Config{
//	    MaxRequests: 100,
//	    Interval:    time.Minute,
//	})
//	mw, err := fortifyhttp.CircuitBreakerStream(cb, streamtimeout.Config{
//	    FirstByteTimeout: 5 * time.Second,
//	    IdleTimeout:      2 * time.Second,
//	    TotalTimeout:     5 * time.Minute,
//	})
//	if err != nil { log.Fatal(err) }
//	mux.Handle("/v1/messages", mw(upstreamProxy))
func CircuitBreakerStream(
	cb circuitbreaker.CircuitBreaker[*http.Response],
	cfg streamtimeout.Config,
) (func(http.Handler) http.Handler, error) {
	st, err := streamtimeout.New[*http.Response](cfg)
	if err != nil {
		return nil, err
	}
	if cb == nil {
		return nil, errors.New("fortify/http: CircuitBreakerStream requires a non-nil CircuitBreaker")
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sw := &streamingResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			_, cbErr := cb.Execute(r.Context(), func(ctx context.Context) (*http.Response, error) {
				_, tErr := st.Execute(ctx, func(ictx context.Context, mark streamtimeout.Mark) (*http.Response, error) {
					sw.setMark(mark)
					next.ServeHTTP(sw, r.WithContext(ictx))
					return &http.Response{StatusCode: sw.status()}, nil
				})
				if tErr != nil {
					// Watchdog fired or context cancelled — surface as a
					// generic timeout so the breaker counts the failure.
					return &http.Response{StatusCode: http.StatusGatewayTimeout}, tErr
				}
				if sc := sw.status(); sc >= 500 {
					return &http.Response{StatusCode: sc}, ferrors.ErrCircuitOpen
				}
				return &http.Response{StatusCode: sw.status()}, nil
			})

			if cbErr == nil {
				return
			}
			// Only synthesise an error response when the handler hasn't
			// already written headers. Once headers are out the only
			// honest signal we can give the client is to close the
			// connection (which the deferred cancel + upstream closing
			// will do naturally).
			if sw.wroteHeader() {
				return
			}
			// Streamtimeout firings must be classified first because
			// StreamTimeoutError unwraps to context.DeadlineExceeded —
			// the generic context check below would swallow them
			// otherwise.
			var stErr *streamtimeout.StreamTimeoutError
			if errors.As(cbErr, &stErr) {
				http.Error(w, "Gateway Timeout", http.StatusGatewayTimeout)
				return
			}
			if errors.Is(cbErr, ferrors.ErrCircuitOpen) {
				http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
				return
			}
			// Client disconnect / parent-ctx cancellation isn't a
			// server-side failure — there's no one left to receive a
			// synthetic error response, and writing one would skew
			// access-log status codes. Drop silently.
			if errors.Is(cbErr, context.Canceled) || errors.Is(cbErr, context.DeadlineExceeded) {
				return
			}
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		})
	}, nil
}

// streamingResponseWriter wraps http.ResponseWriter, forwards Flusher
// / Hijacker / Pusher to the underlying writer, and pings a
// streamtimeout.Mark on every Write so per-chunk activity satisfies
// the idle watchdog.
type streamingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
	mu         sync.Mutex
	mark       streamtimeout.Mark
}

func (s *streamingResponseWriter) setMark(m streamtimeout.Mark) {
	s.mu.Lock()
	s.mark = m
	s.mu.Unlock()
}

func (s *streamingResponseWriter) status() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.statusCode
}

func (s *streamingResponseWriter) wroteHeader() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.written
}

func (s *streamingResponseWriter) WriteHeader(statusCode int) {
	s.mu.Lock()
	if !s.written {
		s.statusCode = statusCode
		s.written = true
		s.ResponseWriter.WriteHeader(statusCode)
	}
	s.mu.Unlock()
}

func (s *streamingResponseWriter) Write(b []byte) (int, error) {
	s.mu.Lock()
	if !s.written {
		s.statusCode = http.StatusOK
		s.written = true
		s.ResponseWriter.WriteHeader(http.StatusOK)
	}
	mark := s.mark
	s.mu.Unlock()

	n, err := s.ResponseWriter.Write(b)
	if n > 0 && mark != nil {
		mark()
	}
	return n, err
}

// Flush forwards to the underlying ResponseWriter when it implements
// http.Flusher. The first Flush also marks the stream, so handlers
// that flush before writing any body (e.g. to push status + headers)
// still satisfy FirstByteTimeout.
func (s *streamingResponseWriter) Flush() {
	s.mu.Lock()
	mark := s.mark
	s.mu.Unlock()
	if f, ok := s.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
	if mark != nil {
		mark()
	}
}

// Hijack forwards to the underlying ResponseWriter when it implements
// http.Hijacker. Hijacking transfers connection ownership to the
// caller, which means the streamtimeout watchdogs no longer observe
// data flow — callers are responsible for their own deadlines after
// hijack.
//
// Returns http.ErrNotSupported when the underlying writer is not
// hijackable, matching the net/http convention.
func (s *streamingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := s.ResponseWriter.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, http.ErrNotSupported
}

// Push forwards to the underlying ResponseWriter when it implements
// http.Pusher (HTTP/2 server push). Returns http.ErrNotSupported
// otherwise, matching net/http convention.
func (s *streamingResponseWriter) Push(target string, opts *http.PushOptions) error {
	if p, ok := s.ResponseWriter.(http.Pusher); ok {
		return p.Push(target, opts)
	}
	return http.ErrNotSupported
}
