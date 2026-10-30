# Hands-on exercises

Commit each exercise separately and add tests before moving on.

1. **Add a route:** Add `/users/` and verify forwarded path and response.
2. **Route precedence:** Add overlapping route prefixes and test longest-prefix selection.
3. **Middleware ordering:** Reorder logging, auth, and limiting; record which requests are logged and charged quota.
4. **Bound limiter memory:** Implement idle-bucket eviction and explain its concurrency model.
5. **Introduce an interface:** Decouple middleware from the concrete token bucket.
6. **Shared quotas:** Implement a Redis limiter with atomic token consumption and a two-process test.
7. **Resilience:** Add upstream health checks and circuit breaker state-transition tests.
8. **Benchmark:** Compare direct and proxied requests with identical payloads, concurrency, and environment notes.
