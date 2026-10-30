// Package middleware contains composable HTTP middleware.
package middleware

import (
	"crypto/subtle"
	"log"
	"net/http"
	"sync"
	"time"
)

type Handler func(http.Handler) http.Handler

func Chain(h http.Handler, handlers ...Handler) http.Handler {
	for i := len(handlers) - 1; i >= 0; i-- {
		h = handlers[i](h)
	}
	return h
}

func RequestLog(logger *log.Logger) Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Printf("method=%s path=%q duration=%s", r.Method, r.URL.Path, time.Since(start))
		})
	}
}

func APIKey(expected string) Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			provided := r.Header.Get("X-API-Key")
			if expected == "" || len(provided) != len(expected) ||
				subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) != 1 {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

type bucket struct {
	tokens float64
	last   time.Time
}

// TokenBucket is process-local. Multiple replicas need a shared store.
type TokenBucket struct {
	mu       sync.Mutex
	buckets  map[string]*bucket
	rate     float64
	capacity float64
}

func NewTokenBucket(ratePerSecond float64, burst int) *TokenBucket {
	if ratePerSecond < 0 {
		ratePerSecond = 0
	}
	if burst < 0 {
		burst = 0
	}
	return &TokenBucket{buckets: make(map[string]*bucket), rate: ratePerSecond, capacity: float64(burst)}
}

func (tb *TokenBucket) Allow(key string) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	b, ok := tb.buckets[key]
	if !ok {
		b = &bucket{tokens: tb.capacity, last: now}
		tb.buckets[key] = b
	}
	b.tokens += now.Sub(b.last).Seconds() * tb.rate
	if b.tokens > tb.capacity {
		b.tokens = tb.capacity
	}
	b.last = now
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

func RateLimit(tb *TokenBucket) Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")
			if key == "" {
				key = r.RemoteAddr
			}
			if !tb.Allow(key) {
				w.Header().Set("Retry-After", "1")
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
