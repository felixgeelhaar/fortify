# Positioning

This document records Fortify's positioning, the wedge it is anchored on, and
the validation gate that decides whether to commit further investment in that
wedge or pivot. It exists so future maintainer decisions stay coherent with
strategy instead of drifting toward feature accretion.

It is a working document, not a marketing artifact. Update it when evidence
moves.

## Headline

**Production-grade resilience for Go services calling LLMs and tools. One
library. One observability story. Patterns composable in three lines.**

## April Dunford's 5 components

### 1. Competitive alternatives

What a Go team building an LLM or agent backend would otherwise reach for:

- Ad-hoc retry loops + `context.WithTimeout` (the dominant status quo).
- `sony/gobreaker` for circuit breaking, bolted onto a manual retry loop.
- `hashicorp/go-retryablehttp` for HTTP-only retry semantics.
- `golang.org/x/time/rate` for in-process rate limiting.
- `failsafe-go` — the closest scope match; v0.x; no AI-specific presets.
- Provider SDK built-in retries (OpenAI, Anthropic Go clients) — no CB, no RL,
  no hedge, no cost budget.
- Reaching across the stack to a Python or TypeScript framework for the AI
  service to side-step Go entirely.

### 2. Differentiated capabilities

- Eight composable patterns under one API: circuit breaker, retry, rate limit,
  timeout, bulkhead, fallback, hedge, adaptive concurrency.
- Built-in observability: OpenTelemetry spans, Prometheus histograms, slog
  structured logging, with structured error types that implement
  `slog.LogValuer` so logs surface fields automatically.
- Generics on result types — no `any` casts in user code.
- Lock-free fast path on the closed-state circuit breaker; alloc-free hot
  paths on retry, timeout, and bulkhead.
- Pluggable rate-limit `Store` for distributed (multi-replica, multi-region)
  rate limiting. Atomic update protocol; fail-open / fail-closed configurable.
- Zero external dependencies in core.
- Planned (Q1, gated by validation): cost-budget primitive, streaming-aware
  3-dimensional timeout (FirstByte / Idle / Total), `LLMCall` and `LLMHedge`
  presets with default-safe budgets and PII redaction.

### 3. Differentiated value

- **Cap LLM bill blowouts during provider incidents.** Cost-budget +
  retry-storm protection (CB) limits the dollar cost of a degraded provider,
  not just the latency cost.
- **Cut p99 of LLM tool-call latency** with delayed hedged requests across
  providers or models, with cancel-loser semantics so the saved tokens are
  real.
- **Ship one import of resilience** instead of stitching four libraries with
  four observability stories.
- **Get production-grade Grafana out of the box** via the bundled metrics
  package, instead of writing exporters by hand.

### 4. Best-fit target customer

The wedge ECP (Maya Voje's *Early Customer Profile*):

- Go shops shipping at least one production service that calls an LLM
  provider (OpenAI, Anthropic, Gemini, Mistral, self-hosted) **or** acts as a
  Model Context Protocol server / client.
- Have an SRE-grade observability stack already (OTel collector, Prometheus,
  Grafana, or equivalent).
- Have experienced at least one cost or tail-latency incident attributable to
  an LLM provider in the last 90 days.
- Solo or small platform / infra team owning the service; willing to adopt a
  v1.x library from a small maintainer base.

This is intentionally narrower than "any Go team that wants resilience."
Narrow is the point.

### 5. Market category

**Go agent resilience layer** — a sub-category of AI infrastructure
middleware, distinct from agent frameworks (Eino, langchaingo, genkit-go) and
from generic resilience libraries (gobreaker, failsafe-go).

Fortify sits below an agent framework and above a provider SDK. It is
complementary to agent frameworks; it does not own prompts, memory, planning,
tool selection, or evaluation.

## Wedge selection rationale

Three candidate wedges considered, ranked by struggling-moment urgency and
proximity:

| Wedge | Pain urgency | Reach | Fortify fit | Verdict |
|---|---|---|---|---|
| Go LLM/agent backends + MCP servers | High (cost + tail latency, 2026) | Medium (AI Eng Slack, MCP Discord, Gophers Slack #ai) | High — hedge + cost budget + per-provider CB are uniquely valuable | **Selected** |
| Distributed rate limit for multi-region Go services | Medium (real but quiet) | Low (no anchored community) | High — pluggable Store is unique | Fallback wedge if AI gate fails |
| HTTP outbound for fintech / payments | Medium (compliance-driven) | Low (gated communities, slow procurement) | Medium | Deferred; bad ECP, possible later ICP |

## Validation gate

Before committing Q2 scope to the AI wedge, run **10 discovery interviews**
with engineers who match the ECP. Story-based JTBD format ("tell me about the
last time an LLM provider caused you pain"). No demo, no pitch.

Bar to commit:

- ≥ 6 of 10 ship at least one Go service in production that calls an LLM
  provider or speaks MCP.
- ≥ 4 of 10 describe a cost or tail-latency incident in the last 90 days.
- ≥ 3 of 10 say "I'd try Fortify in the next sprint" after seeing the
  `LLMCall` preset spec on paper (no working implementation needed).

If under bar:

- Pivot to the **distributed-rate-limit wedge**. Smaller TAM, proven existing
  pain, lower differentiation effort. Reposition `docs/POSITIONING.md`,
  redirect Q2 roadmap accordingly.

## Head-to-head matrix

For the AI / agent backend ECP. Each row is a struggling moment.

| Struggling moment | Fortify | gobreaker | failsafe-go | x/time/rate | retryablehttp | Provider SDK retries |
|---|---|---|---|---|---|---|
| Per-provider circuit breaker | Yes | Yes (CB only; no per-key) | Yes | — | — | No |
| Retry that honors `Retry-After` | Yes (planned in `LLMCall`) | — | Partial | — | Yes | Yes (sometimes) |
| Cost / token budget halt | Yes (planned) | — | — | — | — | No |
| Streaming-aware 3-dim timeout | Yes (planned) | — | — | — | — | No |
| Hedge across providers | Yes (`hedge/`) | — | Yes | — | — | No |
| Distributed rate limit | Yes (pluggable Store) | — | No | No (in-process) | — | No |
| OTel + Prometheus + slog out of box | Yes | — | Hooks only | — | — | — |
| Structured error types with `LogValuer` | Yes | No | Partial | — | — | — |
| MCP server middleware preset | Yes (planned) | — | — | — | — | — |
| Generics on result type | Yes | No | Yes | — | — | — |

## Anti-positioning

To stay coherent with the wedge, Fortify will **not** be positioned as:

- "The best Go circuit breaker." That's `sony/gobreaker`. Don't fight there.
- "The all-in-one resilience kit for any Go service." Too broad, no enemy.
- "An AI agent framework." That's Eino, langchaingo, genkit-go. Stay below.
- "A managed SaaS." Not until ≥ 20 production adopters validate demand.

## Channels (community-led, primary)

In priority order, for the AI wedge:

1. AI Engineer Slack / Latent Space — long-form blog post on hedging LLM
   tool calls in Go.
2. Conference CFPs — GopherCon, FOSDEM AI track, AI Engineer Summit.
3. MCP community Discord and the modelcontextprotocol Go SDK examples list —
   reference MCP server using `middleware.MCPServer`.
4. lobste.rs and `/r/golang` — comparison post against ad-hoc patterns; do
   *not* lead with a head-to-head against gobreaker.
5. Gophers Slack `#show-and-tell` and `#ai` — light cadence, one drop per
   month.

Skipped: paid acquisition, Twitter/X, ProductHunt, HN front-page chase.

## North-star metric and OKR sketch

NSM: **distinct Go modules importing `go.klarlabs.de/fortify`** in
production (proxy via pkg.go.dev import graph and Sourcegraph search).

Q1 (90 days): validate.

- 10 discovery interviews complete, transcripts archived.
- 1 wedge chosen and recorded in this document.
- 5 named adopters in `ADOPTERS.md`.
- NSM defined and instrumented; baseline + target recorded.

Q2 (90 days): anchor.

- 3 published case studies / blog posts from adopters.
- ≥ 50% of GitHub issues filed by non-maintainer accounts.
- 1 conference talk submitted.
- One wedge-specific preset (`middleware.LLMCall` if AI wedge wins) shipped
  with reference docs.
- NSM grows ≥ 3× from Q1 baseline.

## Owner and review cadence

Solo maintainer. Re-read this document at the end of every quarter. Update
when evidence (interviews, adopter count, NSM movement) contradicts what's
written here. If the wedge changes, supersede the document — do not edit in
place.
