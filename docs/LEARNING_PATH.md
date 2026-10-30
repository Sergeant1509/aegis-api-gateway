# Aegis learning path

Complete each milestone in order. Explain trade-offs and add tests before advancing.

## 1. HTTP and reverse proxy (current)
Concepts: HTTP lifecycle, handlers, reverse proxying, prefix routing, timeouts, graceful shutdown. Read `cmd/gateway/main.go`, `internal/proxy/router.go`, and `internal/config/config.go`.
Exercises: add a third upstream; prove longest-prefix routing with a test; stop an upstream and verify 502; trace forwarded request headers.

## 2. Middleware and authentication
Concepts: middleware composition, constant-time comparison, authentication vs authorization, key rotation. Exercises: trace middleware order; add key identifiers; support overlapping keys during rotation; keep secrets out of Git.

## 3. Rate limiting
Concepts: token bucket, burst capacity, refill rate, races, per-consumer quotas. Exercises: derive the refill equation; inject a clock for deterministic tests; evict inactive buckets; explain replica divergence.

## 4. Redis-backed distributed limits
Concepts: atomic Lua scripts, TTL, data modelling, failure policy. Exercises: define a limiter interface; implement atomic shared quotas; decide fail-open vs fail-closed for Redis outage; test two gateway instances.

## 5. Load balancing and resilience
Concepts: round robin, health state, circuit breakers, retry budgets, idempotency. Exercises: multiple upstreams; closed/open/half-open breaker; retry only safe operations by default.

## 6. Observability
Concepts: structured logs, RED metrics, cardinality, tracing. Exercises: request IDs; Prometheus counters and histograms; avoid secrets and unbounded URLs as labels.

## 7. Production hardening
Concepts: TLS, config reload, body limits, header policy, threat modelling, container security. Exercises: define TLS boundary; configure body and method policy; scan dependencies/images; write outage and key-rotation runbooks.

## 8. Performance engineering
Concepts: benchmarks, p50/p95/p99, throughput, profiling. Exercises: benchmark router and limiter; compare direct vs proxied requests; record CPU, memory, Go version, concurrency and duration; publish no unsupported performance claims.
