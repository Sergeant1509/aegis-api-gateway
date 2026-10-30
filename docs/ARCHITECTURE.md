# Architecture

## Request flow

```text
Client -> HTTP server -> request logging -> API-key auth -> token bucket
       -> /healthz handler OR longest-prefix router -> reverse proxy -> upstream
```

## Package boundaries
- `cmd/gateway`: process lifecycle and dependency wiring.
- `internal/config`: route configuration parsing and validation.
- `internal/middleware`: cross-cutting HTTP policies.
- `internal/proxy`: route selection and upstream forwarding.
- `internal/health`: process liveness endpoint.

## MVP limitations

This is not production-ready. Configuration loads once at startup. Authentication uses one shared environment key. Rate limiting is process-local and inactive client entries are not yet evicted. The health endpoint reports process liveness, not upstream readiness. There is no load balancing, circuit breaker, distributed tracing, or Prometheus metrics. TLS termination is expected at a trusted ingress and that trust boundary must be configured before deployment. Compose echo services are demonstration upstreams only.

## Security notes

Do not commit real credentials. Never trust forwarded headers without a trusted-proxy boundary. Keep timeouts bounded. Retries need an idempotency policy because replaying non-idempotent requests can duplicate side effects.
