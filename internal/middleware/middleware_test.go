package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIKey(t *testing.T) {
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusNoContent)
	})
	h := APIKey("secret")(next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	h.ServeHTTP(res, req)
	if res.Code != http.StatusUnauthorized || called {
		t.Fatalf("missing key: status=%d called=%v", res.Code, called)
	}

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-API-Key", "secret")
	res = httptest.NewRecorder()
	h.ServeHTTP(res, req)
	if res.Code != http.StatusNoContent || !called {
		t.Fatalf("valid key: status=%d called=%v", res.Code, called)
	}
}

func TestTokenBucketBurst(t *testing.T) {
	tb := NewTokenBucket(0, 2)
	if !tb.Allow("client") || !tb.Allow("client") {
		t.Fatal("initial burst should be allowed")
	}
	if tb.Allow("client") {
		t.Fatal("third immediate request should be rejected")
	}
}
