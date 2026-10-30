# Architecture

## Current request flow

```text
Client
  |
  v
HTTP server (timeouts and graceful shutdown)
  |
  v
Request logging
  |
  v
Process-local token bucket
  |
  v
API-key authentication
  |
  +---- /healthz -> health handler
  |
  v
Prefix router (longest prefix first)
  |
  v
Reverse proxy
  |
  +---- catalog upstream
  +---- orders upstream
```

## Package boundaries

- `cmd/gateway`: process lifecycle and dependency wiring.
- `internal/config`: route configuration parsing and validation.
- `internal/middleware`: cross-cutting HTTP policies.
- `internal/proxy`: route selection and upstream forwarding.
- `internal/health`: process liveness endpoint.

## Known MVP limitations

This is not a production-ready gateway.

- Configuration is loaded once at startup.
- Authentication uses one shared environment-provided API key.
- Rate limiting is process-local and has no inactive-key eviction yet.
- The health endpoint reports process liveness, not upstream readiness.
- No load balancing, circuit breaker, distributed tracing, or Prometheus metrics yet.
- TLS termination is expected at a trusted ingress; that trust boundary must be configured before deployment.
- Echo services in Compose are only demonstration upstreams.

## Security notes

- Do not commit real credentials or use the example key in production.
- Never trust forwarded headers without a trusted-proxy boundary.
- Keep timeouts bounded.
- Retries need an idempotency policy; blindly replaying non-idempotent requests can duplicate side effects.
