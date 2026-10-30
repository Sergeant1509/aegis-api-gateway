# Aegis API Gateway

A learning-focused API gateway written in Go. It starts with HTTP reverse proxying, prefix routing, middleware, API-key authentication, and token-bucket rate limiting, then progresses toward distributed quotas, load balancing, resilience, and observability.

**Status:** MVP foundation. Not production-ready.

## Current features

- HTTP reverse proxy using Go's standard library
- JSON-configured path-prefix routing (longest prefix wins)
- Optional path-prefix stripping
- API-key authentication through `X-API-Key`
- Process-local token-bucket rate limiting
- Request logging and `/healthz` process-liveness endpoint
- Bounded server timeouts and graceful shutdown
- Unit tests, Docker Compose sample, and GitHub Actions CI

## Requirements

Go 1.23 or newer. Docker Engine and Docker Compose are optional for the container demo.

## Run

Set a development key and start the gateway:

```bash
export AEGIS_API_KEY='local-development-only'
go run ./cmd/gateway -config configs/routes.json -listen :8080
```

The checked-in route config targets Docker service names. For a host-run gateway, change upstreams to local addresses and run those services separately.

## Docker Compose demo

```bash
export AEGIS_API_KEY='local-development-only'
docker compose up --build
```

In another terminal:

```bash
curl -i -H 'X-API-Key: local-development-only' http://localhost:8080/catalog/products
curl -i -H 'X-API-Key: local-development-only' http://localhost:8080/orders/42
curl -i -H 'X-API-Key: local-development-only' http://localhost:8080/healthz
```

All routes, including `/healthz`, currently pass through authentication and rate limiting.

## Development checks

```bash
gofmt -w .
go vet ./...
go test -race -cover ./...
go build ./cmd/gateway
```

## Documentation

- [Architecture](docs/ARCHITECTURE.md)
- [Learning path](docs/LEARNING_PATH.md)
- [Architecture decisions](docs/DECISIONS.md)
- [Hands-on exercises](docs/EXERCISES.md)

## Roadmap

- [x] HTTP reverse proxy and prefix routing
- [x] API-key authentication
- [x] Basic in-memory token-bucket limiter
- [ ] Redis-backed distributed quotas
- [ ] Multiple upstreams and health-aware load balancing
- [ ] Circuit breaker and safe retry policy
- [ ] Request IDs, Prometheus metrics, and tracing
- [ ] Config reload, TLS strategy, and security hardening
- [ ] Reproducible benchmarks and performance report

Every feature should have tests, documented trade-offs, and a measurable reason to exist. No performance or production-readiness claims without evidence.
