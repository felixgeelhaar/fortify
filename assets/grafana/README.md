# Fortify Grafana Dashboard

`fortify-dashboard.json` is a ready-to-import Grafana dashboard for Fortify's Prometheus metrics.

## Import

1. Grafana UI → Dashboards → New → Import
2. Upload `fortify-dashboard.json` or paste its contents
3. Select your Prometheus datasource when prompted (variable `DS_PROMETHEUS`)
4. Save

## Panels

| Row             | Panels                                                              |
| --------------- | ------------------------------------------------------------------- |
| Circuit Breaker | state-by-instance, failure ratio (1m), state changes (5m)           |
| Retry           | attempts per call (p50/p99), success/failure rate                   |
| Rate Limit      | allowed/denied rate, p99 wait time                                  |
| Timeout         | execution rate vs exceeded rate, p99 duration                       |
| Bulkhead        | active/queued gauges, rejection rate                                |

## Required metric exposure

Wire `metrics.MustRegister(prometheus.DefaultRegisterer)` in your service and configure each pattern with the corresponding `OnStateChange`/`OnRetry`/etc. callback to the package-level `metrics.DefaultCollector()`. See `docs/how-to-observe.md`.

## Tweaking

- All panels label by `name` — set the same `name` on the metrics callbacks as on your pattern instances for clean grouping.
- Failure ratio panel has no thresholds applied; add red bands at 0.5/0.9 to alert visually.
- Timeranges default to `now-30m`. Adjust per Grafana setup.

## Compatible Grafana versions

Tested against Grafana 10.x. Schema version 39. Older Grafana may need manual datasource binding.
