
## Fix README/doc.go inconsistencies

Critical landing-page trust gap. README says Go 1.25, doc.go says 1.23+. doc.go lists 6 patterns (missing hedge + adaptive). Perf numbers diverge (~70ns vs ~100ns CB closed). Single source of truth = go.mod. README and doc.go reference one canonical perf table. Add hedge + adaptive sections to doc.go Core Patterns + Package Organization. Drop emoji from doc.go for stdlib aesthetic.

---

## Rewrite README — code above the fold

Currently violates Hick's Law (4 paths to first program) and Jakob's Law (Go libs lead with code). Reorder: tagline → install → 60-second quick-start using a real Response struct (not [string]). Single quick-start path. "Why Fortify" + comparison table below the fold. Add a 5-line "Which pattern do I need?" decision snippet at top of docs/concepts.md and link from README.

---

## slog.LogValuer on structured errors

Implement LogValue() slog.Value on CircuitOpenError, RateLimitError, TimeoutError so logs emit structured fields automatically. Delivers the observability promise without user boilerplate. Cheap; big DX win. Add tests verifying slog output shape.

---

## Write docs/POSITIONING.md

April Dunford 5-component positioning brief. Best-fit customer = Go teams shipping LLM/agent backends or MCP servers. Category = Go agent resilience layer. Headline = production-grade resilience for Go services calling LLMs and tools. Head-to-head against sony/gobreaker, failsafe-go, x/time/rate, ad-hoc retry. Document wedge selection rationale, validation gate, and fallback wedge (distributed RL).

---

## Run 10 discovery interviews

Story-based JTBD interviews with Go AI/agent shops. Bar to commit AI wedge: >=6 of 10 ship Go service calling LLM/MCP in prod; >=4 of 10 had a cost or tail-latency incident in last 90 days; >=3 of 10 say they would try Fortify in next sprint after seeing LLMCall preset. Channels: AI Engineer Slack #go, MCP Discord, Gophers Slack #ai, /r/golang. Below bar => pivot to distributed-RL wedge.

---

## Cost budget primitive

New interceptor (budget/ subpackage). Halts execution when token / dollar / call budget exceeded. Returns *BudgetExceededError with actuals. Composable into any chain. Pluggable Counter interface so callers wire their own tokenizer. Default-safe (low caps) when wired into LLMCall preset.

---

## Streaming-aware timeout

Standard timeout breaks streaming LLM responses. New primitive in timeout/ exposing 3 dimensions: FirstByteTimeout, IdleTimeout (max gap between chunks), TotalTimeout. Compose with retry / CB / hedge. Reference adapter for io.Reader and SSE chunk callback shape.

---

## middleware.LLMCall preset

Provider-aware preset wrapping CB + RL + 3-dim timeout + retry-with-Retry-After + per-provider bulkhead + cost budget. Built-ins for OpenAI, Anthropic, generic OpenAI-compatible (configurable baseURL). Honors Retry-After. Default-safe budgets and idempotency-off semantics. PII redaction in observability hooks by default; opt-in unsafe-log flag.

---

## middleware.LLMHedge preset

Race two providers/models. Delayed-hedge default (HedgeAfter ~800ms) to save tokens when primary is fast. Cancel-loser semantics. DataResidencyPolicy field requiring explicit opt-in for cross-vendor hedging with documented data-flow warning.

---

## MCP server reference example

End-to-end working MCP server example using middleware.MCPServer (per-peer rate limit, bulkhead, adaptive concurrency). Uses official modelcontextprotocol Go SDK. Lives under examples/mcp/. PR link to upstream mcp-go-examples list once shipped.

---

## Grafana dashboard screenshot in README

Visceral evidence of observability wedge. Image at assets/dashboard.png with caption "5 minutes from go get to this". Generated from a docker-compose demo (Prometheus + Grafana + sample service with Fortify chain). Compose lives under examples/observability-demo/.

---

## Blog post — Hedging LLM tool calls in Go

Long-form Lightning Strike anchoring AI wedge. Concrete numbers: latency distribution before/after hedge, cost impact, code walkthrough. Submit to AI Engineer / Latent Space, lobste.rs, /r/golang, GoWeekly. Conditional on AI wedge passing validation gate.

---

## PII / prompt redaction by default

slog and OTel hooks must NOT log raw prompt content or completion tokens by default. Opt-in via WithUnsafeLogPrompts(true). Document in docs/PRODUCTION.md and otel/ readme. Necessary safety default for AI-wedge presets.

---

## Eino integration reference

Reference example wrapping Eino LLM model with a Fortify chain. Demonstrates Fortify is complementary, not competitive, to Go agent frameworks. Lives under examples/eino/. Optional blog post co-marketing once shipped.

---
