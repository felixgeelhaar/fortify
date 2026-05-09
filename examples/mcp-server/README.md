# MCP server with Fortify

This example shows how to wrap an MCP (Model Context Protocol) tool
handler with Fortify resilience patterns:

- **Adaptive concurrency** (Vegas) — auto-tunes the bulkhead cap based
  on observed latency.
- **Bulkhead** — caps total in-flight tool invocations server-wide.
- **Per-peer rate limit** — one bucket per remote MCP client identity,
  so a misbehaving client cannot starve the others.
- **Per-call timeout** — bounds the latency of each tool invocation.

## What this is and isn't

The example deliberately does **not** depend on a specific MCP SDK.
It defines a minimal local `ToolHandler` interface so the code builds
standalone in CI and stays readable as a teaching example. To wire it
to a real MCP server, replace the local interface with the equivalent
from your SDK of choice (the official `modelcontextprotocol/go-sdk`,
or community SDKs such as `mark3labs/mcp-go`). The Fortify chain
itself — `WithAdaptive → WithBulkhead → WithRateLimit → WithTimeout`
— does not change.

## Run

```bash
cd examples/mcp-server
go run .
```

You should see two structured log lines, one per synthetic peer call.

## Why this chain order

`WithAdaptive` is outermost so it can shed load before the bulkhead
even acquires a slot. `WithBulkhead` sits between adaptive and rate
limit so per-peer fairness is enforced *after* total concurrency has
been bounded. `WithTimeout` is innermost so it bounds the actual
operation, not the wait time spent in the bulkhead queue or the rate
limiter.

Read `docs/how-to-compose.md` in the repository root for the full
ordering rationale.

## Sensitive payloads

`ToolCall.Args` is opaque to Fortify. The example logger never
serialises it. If you log structured errors via `slog.Any("err", err)`,
the structured error types (`*ferrors.RateLimitError`,
`*ferrors.TimeoutError`, etc.) emit their own `LogValue` groups
without ever inspecting the request payload. See
`docs/PRODUCTION.md` ("Observability and sensitive payloads").
