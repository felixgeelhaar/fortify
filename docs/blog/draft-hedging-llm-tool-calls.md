# Hedging LLM tool calls in Go

> **Status:** draft. This post is gated on the validation gate in
> [POSITIONING.md](../POSITIONING.md). Publish only after at least
> three interviewed adopters confirm the wedge. Replace
> `<measured-numbers>` with real numbers from a representative
> production workload before publishing.

## TL;DR

LLM tool calls in production have a long, ugly tail. Hedge them.

If you're calling OpenAI, Anthropic, or your own model behind an
OpenAI-compatible endpoint from a Go service, the slowest 1% of calls
is dragging your user experience and your error budget. A hedge
fired 800 ms after the primary, with cancel-loser semantics, can
collapse p99 by `<measured-numbers>` while paying for hedges only on
the fraction of calls that need them. With the right cost-budget
ceiling on top, your retry logic stops being able to set fire to
your bill during a provider incident.

## The bill, the tail, and the retry storm

Three failure modes worth naming separately.

**The tail.** LLM responses have a heavy-tailed latency distribution.
A 50th-percentile call takes a couple of seconds; the 99th can take
30. If you're sitting one or two of these calls deep behind a user
interaction, that p99 *is* the user experience.

**The retry storm.** Every Go service hits its first 429 from a
provider, slaps `for i := 0; i < 5; i++` around the call, and ships
it. When the provider has a degraded hour, every replica retries in
parallel; the provider's recovery is delayed, the retries' input
tokens are billed each round, and the bill page from your provider
arrives the following Monday.

**The cost runaway.** A naive retry on 429 doubles input tokens for
the second attempt without producing a different response. A retry
on a long, partial completion bills you for the full prompt twice
and (depending on provider) for whatever output tokens the first
attempt produced before failing. Retry-on-error is a token multiplier
unless you cap it.

These are three different problems. They need three different tools.

## What hedging is, exactly

A hedged request fires a second attempt after a delay if the first
hasn't returned yet. As soon as one attempt succeeds, the others are
cancelled. The delay is the lever: short delays cut tail latency
hard but pay for hedges on most calls; longer delays only pay when
the primary is genuinely slow.

For LLM calls specifically:

- **Same vendor, same model**: hedge for tail-latency hedging on a
  single provider.
- **Same vendor, different model**: primary on `gpt-5`, hedge on
  `gpt-5-mini` if the smaller model can answer the easy cases.
- **Different vendors**: primary on OpenAI, hedge on Anthropic. Pay
  for hedge tokens at both providers when it fires; recover
  availability when one provider has a bad day.

Cross-vendor hedging has a data-residency dimension. The same prompt
goes to two providers, who have two retention and processing
policies. Treat that as a deliberate decision, not a default.

## The wiring

Fortify ships a `LLMCall` preset that wraps the resilience layer
around a single provider, and an `LLMHedge` runner that races
multiple `LLMCall` chains.

A primary chain bound to OpenAI:

```go
import (
    "go.klarlabs.de/fortify/budget"
    "go.klarlabs.de/fortify/middleware"
)

openaiChain, _ := middleware.LLMCall[Response](middleware.LLMCallConfig[Response]{
    Provider:    "openai",
    Model:       "gpt-5",
    CallTimeout: 30 * time.Second,
    MaxRetries:  3,
    Budget: middleware.BudgetConfig[Response]{
        Max: budget.Cost{
            Tokens:    50_000,
            USDMicros: 250_000, // $0.25 ceiling per chain instance
        },
        Charge: chargeFromUsage,
    },
})
```

A second chain bound to Anthropic uses the same shape with a
different `Provider` and `Model`. Then race them:

```go
hedge, _ := middleware.LLMHedge[Response](middleware.LLMHedgeConfig[Response]{
    Attempts: []func(context.Context) (Response, error){
        func(ctx context.Context) (Response, error) {
            return openaiChain.Execute(ctx, callOpenAI)
        },
        func(ctx context.Context) (Response, error) {
            return anthropicChain.Execute(ctx, callAnthropic)
        },
    },
    HedgeAfter:       800 * time.Millisecond,
    AllowCrossVendor: true, // explicit, not default
})

result, err := hedge.Run(ctx)
```

`AllowCrossVendor` is the data-residency gate. The runner refuses to
fire across multiple attempts unless you've either set this flag or
attested via `AllAttemptsSameVendor: true`. The runner cannot detect
provider identity, so this is a deliberate caller assertion.

## Why budget belongs inside retry

The budget sits inside `Retry` in the `LLMCall` chain order, which
means every retry attempt is charged. This is the lever for the
retry-storm problem: when a provider has a degraded hour, the budget
short-circuits remaining attempts on the call that caused the breach
instead of letting the retry loop spend the rest of your hourly
budget on a single bad request.

`budget.ErrBudgetExceeded` is a hard non-retryable error in the
`LLMCall` predicate. Once any dimension of the budget breaches, the
chain returns the structured error verbatim and your handler can
either degrade gracefully (cached response, smaller model fallback,
"sorry, try again later") or surface it to the caller.

## Numbers

> Replace before publishing.
>
> Suggested experiments to source numbers from:
>
> 1. Capture 24 h of production p50/p95/p99/p99.9 latency for the
>    target tool call without hedging.
> 2. Enable `LLMHedge` with `HedgeAfter` swept across 200 ms, 500 ms,
>    800 ms, 1500 ms. Record the same percentiles plus the fraction
>    of calls that fired a hedge and the resulting token cost
>    multiplier.
> 3. Run a controlled provider-degradation simulation (use
>    `fortify/testing/aichaos` once shipped, or a manual 429 injector)
>    with and without budget caps. Record bill delta.

## What hedging is not

It is not a circuit breaker. The chain still has one of those, doing
the right thing per provider. It is not a substitute for caching;
caching avoids the call entirely while hedging accepts the cost in
exchange for speed. It is not safe for non-idempotent side-effects;
the standard rule applies — your operation must be safe to lose its
result if cancelled.

## When not to hedge

If your p99 latency is acceptable, don't hedge. If your token cost
budget is tight, don't hedge across vendors with cancel-loser
semantics — the loser may complete enough tokens to be billed.
If your tool call has side effects (a write, a state change),
absolutely don't hedge.

## Where to start

```bash
go get go.klarlabs.de/fortify
```

Read the [POSITIONING document](../POSITIONING.md) for the wedge and
fit, the [composition guide](../how-to-compose.md) for chain ordering,
and the [Eino integration example](../../examples/eino/) for the full
shape of an LLM-aware chain.

## Discussion

Issues and discussion: <https://github.com/klarlabs-studio/fortify>.
We're collecting adoption stories in
[ADOPTERS.md](../../ADOPTERS.md); a one-line PR helps shape what
ships next.
