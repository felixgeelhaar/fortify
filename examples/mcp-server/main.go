// Example: protecting a Model Context Protocol (MCP) tool handler with
// Fortify resilience patterns.
//
// This example deliberately uses a minimal local handler interface
// instead of an external SDK so the example builds standalone in CI.
// To wire it to a real MCP server (such as the official
// modelcontextprotocol Go SDK or the community mark3labs/mcp-go),
// replace the ToolHandler interface and the registerTool boilerplate
// with your SDK's equivalents. The Fortify integration — the
// fortifyToolHandler wrapper and how it composes patterns — does not
// change.
//
// What this example demonstrates:
//
//   - Per-peer rate limiting (one bucket per remote MCP client identity).
//   - Bulkhead capping in-flight tool invocations on the server.
//   - A timeout per tool invocation.
//   - Adaptive concurrency tuning the bulkhead under variable load.
//   - Logging of structured errors via slog without leaking tool args.
package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/felixgeelhaar/fortify/adaptive"
	"github.com/felixgeelhaar/fortify/bulkhead"
	"github.com/felixgeelhaar/fortify/ferrors"
	"github.com/felixgeelhaar/fortify/middleware"
	"github.com/felixgeelhaar/fortify/ratelimit"
	"github.com/felixgeelhaar/fortify/timeout"
)

// ToolCall represents a single inbound MCP tool invocation. Real SDKs
// expose their own version of this; the resilience layer only needs
// the peer identity and the tool arguments.
type ToolCall struct {
	// PeerID identifies the remote MCP client. Used as the rate-limit
	// bucket key so a misbehaving client cannot starve the others.
	PeerID string
	// Tool is the name of the requested tool.
	Tool string
	// Args is the tool arguments — opaque to Fortify; never logged.
	Args map[string]any
}

// ToolResult is the protocol-shaped response. Replace with your SDK's
// equivalent.
type ToolResult struct {
	Content string
	IsError bool
}

// ToolHandler is the interface a real MCP SDK would already supply.
// Implementing it as a one-method interface here keeps the example
// readable.
type ToolHandler interface {
	Handle(ctx context.Context, call ToolCall) (ToolResult, error)
}

// HandlerFunc is a convenience adapter, mirroring http.HandlerFunc.
type HandlerFunc func(ctx context.Context, call ToolCall) (ToolResult, error)

func (f HandlerFunc) Handle(ctx context.Context, call ToolCall) (ToolResult, error) {
	return f(ctx, call)
}

// fortifyToolHandler wraps a ToolHandler with the Fortify chain. The
// chain is built once per server and reused across all tool calls.
type fortifyToolHandler struct {
	inner       ToolHandler
	rl          ratelimit.RateLimiter
	bh          bulkhead.Bulkhead[ToolResult]
	tm          timeout.Timeout[ToolResult]
	adaptive    adaptive.Limiter[ToolResult]
	callTimeout time.Duration
	logger      *slog.Logger
}

func (h *fortifyToolHandler) Handle(ctx context.Context, call ToolCall) (ToolResult, error) {
	chain := middleware.New[ToolResult]().
		WithAdaptive(h.adaptive).
		WithBulkhead(h.bh).
		WithRateLimit(h.rl, "peer:"+call.PeerID).
		WithTimeout(h.tm, h.callTimeout)

	return chain.Execute(ctx, func(ctx context.Context) (ToolResult, error) {
		return h.inner.Handle(ctx, call)
	})
}

// businessLogic is a placeholder for real tool work. In a real server
// this might query a database, call an external API, or run a model.
func businessLogic(ctx context.Context, call ToolCall) (ToolResult, error) {
	switch call.Tool {
	case "ping":
		return ToolResult{Content: "pong"}, nil
	case "slow":
		select {
		case <-ctx.Done():
			return ToolResult{}, ctx.Err()
		case <-time.After(5 * time.Second):
			return ToolResult{Content: "done"}, nil
		}
	default:
		return ToolResult{IsError: true, Content: "unknown tool"}, errors.New("unknown tool")
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Per-peer rate limit: 10 calls/sec/peer with a small burst.
	rl := ratelimit.New(ratelimit.Config{
		Rate:     10,
		Burst:    20,
		Interval: time.Second,
	})

	// Bulkhead capping concurrent in-flight calls server-wide.
	bh := bulkhead.New[ToolResult](bulkhead.Config{
		MaxConcurrent: 100,
		MaxQueue:      200,
	})

	// Adaptive Vegas tuning the effective concurrency cap based on
	// observed latency.
	a := adaptive.New[ToolResult](adaptive.Config{
		Algorithm:    adaptive.AlgorithmVegas,
		InitialLimit: 50,
		MinLimit:     10,
		MaxLimit:     500,
	})

	// Per-call timeout.
	tm := timeout.New[ToolResult](timeout.Config{
		DefaultTimeout: 2 * time.Second,
	})

	handler := &fortifyToolHandler{
		inner:       HandlerFunc(businessLogic),
		rl:          rl,
		bh:          bh,
		tm:          tm,
		adaptive:    a,
		callTimeout: 2 * time.Second,
		logger:      logger,
	}

	// In a real server this would be an SDK-supplied loop reading
	// JSON-RPC frames off stdio or a TCP connection. We approximate it
	// with a single synthetic call per peer to demonstrate the wiring
	// without bringing in a transport dependency.
	peers := []string{"peer-a", "peer-b"}
	for _, peer := range peers {
		result, err := handler.Handle(context.Background(), ToolCall{
			PeerID: peer,
			Tool:   "ping",
			Args:   map[string]any{}, // arguments are opaque to Fortify
		})
		logResult(logger, peer, result, err)
	}
}

func logResult(logger *slog.Logger, peer string, result ToolResult, err error) {
	if err == nil {
		logger.Info("tool call ok",
			slog.String("peer", peer),
			slog.String("content", result.Content),
		)
		return
	}

	// Use slog.Any so structured error types emit their LogValue groups
	// automatically. Caller does not have to know which fields belong
	// to which error type.
	switch {
	case errors.Is(err, ferrors.ErrRateLimitExceeded):
		logger.Warn("tool call rate-limited", slog.String("peer", peer), slog.Any("err", err))
	case errors.Is(err, ferrors.ErrBulkheadFull):
		logger.Warn("tool call rejected by bulkhead", slog.String("peer", peer), slog.Any("err", err))
	case errors.Is(err, ferrors.ErrTimeout):
		logger.Warn("tool call timed out", slog.String("peer", peer), slog.Any("err", err))
	default:
		logger.Error("tool call failed", slog.String("peer", peer), slog.Any("err", err))
	}
}
