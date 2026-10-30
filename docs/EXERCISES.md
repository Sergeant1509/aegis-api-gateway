# Hands-on exercises

Commit each exercise separately and add tests.

1. Add `/users/` and verify the forwarded path and response.
2. Add overlapping route prefixes and test longest-prefix selection.
3. Reorder logging, auth, and limiting; record logging and quota behavior.
4. Implement idle-bucket eviction and explain its concurrency model.
5. Decouple middleware from the concrete token bucket with an interface.
6. Implement Redis quotas atomically and test two gateway processes.
7. Add upstream health checks and circuit-breaker state-transition tests.
8. Compare direct and proxied requests under identical load and publish the environment and raw data.
