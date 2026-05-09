# Observability demo

Self-contained Prometheus + Grafana stack that brings up Fortify's
metrics end-to-end so you can see the patterns at work without wiring
anything yourself.

## Run

```bash
cd examples/observability-demo
docker compose up --build
```

Then open:

- Grafana — <http://localhost:3000> (anonymous Admin)
- Prometheus — <http://localhost:9090>
- Sample app `/metrics` — <http://localhost:8080/metrics>

The "Fortify overview" dashboard is pre-loaded under General. The
sample app fires one synthetic call every 50 ms through a chain of
`CircuitBreaker → Retry → Timeout` and reports the resulting metrics.
The downstream randomly errors and stalls, so within ~30 seconds you
see panels populate for state, request rate, retry attempts, timeout
durations, and rate-limit allow/deny.

## Capture the README screenshot

The repository's main README references `assets/dashboard.png`. To
refresh it:

1. `docker compose up --build` from this directory.
2. Wait ~60 seconds for graphs to populate.
3. Take a screenshot of the dashboard at <http://localhost:3000>.
4. Save as `assets/dashboard.png` in the repo root.

A maintainer task; CI does not enforce the screenshot.

## What this demonstrates

- The `metrics.Collector` exposes counters, histograms, and gauges
  for every Fortify pattern. The sample app shows the canonical
  wiring: build the collector, register it on a Prometheus registry,
  and call `Record*` from your `OnStateChange` and post-call hooks.
- Sensitive payloads stay out of metrics: labels are pattern names
  and outcomes (`closed` / `open` / `denied` / `success`), never
  request bodies or LLM prompts.
- The chain composes cleanly with the rest of Fortify: replace the
  inline downstream with an HTTP client, an LLM call, or an MCP tool
  invocation; the dashboard does not change.
