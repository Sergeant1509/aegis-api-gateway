# Architecture decision records

## ADR-001: Go standard library first
Use `net/http`, `httputil.ReverseProxy`, and standard synchronization primitives. This keeps server, middleware, proxy, and concurrency mechanics visible. Trade-off: behavior must be implemented and tested explicitly.

## ADR-002: Longest-prefix route precedence
Sort routes by descending prefix length. A specific prefix such as `/api/admin/` must match before `/api/`. Exact-match and method-aware routes are not in this version.

## ADR-003: Process-local token bucket first
Use a mutex-protected token bucket before introducing Redis. This keeps refill and burst rules inspectable. Quotas are per process and idle entries are not yet evicted; this is not a multi-replica solution.
