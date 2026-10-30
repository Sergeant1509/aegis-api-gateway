# Architecture decision records

## ADR-001: Go standard library first
**Decision:** Use `net/http`, `httputil.ReverseProxy`, and standard synchronization primitives.

**Reason:** This keeps HTTP server, middleware, proxy, and concurrency mechanics visible for learning.

**Trade-off:** More behavior must be implemented and tested explicitly.

## ADR-002: Longest-prefix route precedence
**Decision:** Sort routes by descending prefix length.

**Reason:** A specific prefix such as `/api/admin/` must match before a broad prefix such as `/api/`.

**Trade-off:** Exact-match and method-aware routes are not in the first version.

## ADR-003: Process-local token bucket first
**Decision:** Implement a mutex-protected token bucket before adding Redis.

**Reason:** The refill and burst rules remain inspectable.

**Trade-off:** Quotas are per process and client entries are not yet evicted; this is not a multi-replica solution.
