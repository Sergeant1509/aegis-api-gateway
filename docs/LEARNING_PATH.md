# Aegis learning path

This project is a guided backend-systems exercise. Complete each milestone in order and explain the trade-offs before adding complexity.

## 1. HTTP and reverse proxy (current)
**Concepts:** HTTP lifecycle, handlers, reverse proxies, prefix routing, timeouts, graceful shutdown.

Read `cmd/gateway/main.go`, `internal/proxy/router.go`, and `internal/config/config.go`.

Exercises:
- Add a third upstream and verify forwarding with curl.
- Test overlapping routes and prove longest-prefix matching.
- Stop an upstream and confirm the gateway returns 502.
- Explain how a reverse proxy changes the request's network destination.

## 2. Middleware and authentication
**Concepts:** middleware composition, constant-time comparison, authentication versus authorization, key rotation.

Exercises:
- Trace the middleware order in `Chain`.
- Add key identifiers and a key registry.
- Design overlapping key rotation without downtime.
- Keep credentials out of Git.

## 3. Rate limiting
**Concepts:** token bucket, burst capacity, refill rate, races, per-consumer quotas.

Exercises:
- Derive the token refill equation.
- Inject a clock for deterministic tests.
- Evict inactive client buckets; the current map can grow without bound.
- Explain why process-local quotas diverge across replicas.

## 4. Redis-backed distributed limits
**Concepts:** atomic Lua scripts, TTL, data modelling, failure policy.

Exercises:
- Introduce a limiter interface.
- Implement shared token consumption atomically with Redis.
- Decide fail-open versus fail-closed when Redis is unavailable.
- Run two gateway instances against one shared quota.

## 5. Load balancing and resilience
**Concepts:** round robin, health state, circuit breakers, retry budgets, idempotency.

Exercises:
- Support multiple upstream instances per route.
- Implement closed/open/half-open circuit states.
- Retry only safe operations by default; replaying POST may duplicate side effects.

## 6. Observability
**Concepts:** structured logs, RED metrics, cardinality, tracing.

Exercises:
- Add request IDs and JSON logs.
- Export Prometheus counters and latency histograms.
- Avoid API keys and unbounded URL values as metric labels.

## 7. Production hardening
**Concepts:** TLS, config reload, body limits, header policy, threat modelling, container security.

Exercises:
- Define the TLS termination boundary.
- Add body-size limits and method policy.
- Add dependency and container scanning.
- Write operational runbooks for key rotation and upstream/Redis outages.

## 8. Performance engineering
**Concepts:** benchmarks, p50/p95/p99, throughput, profiling.

Exercises:
- Benchmark router and limiter independently.
- Compare direct and proxied traffic with the same payload and concurrency.
- Record CPU, memory, Go version, warm-up, and duration.
- Never publish a performance claim without reproducible data.
