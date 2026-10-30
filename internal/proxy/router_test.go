package proxy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sergeant1509/aegis-api-gateway/internal/config"
)

func TestRouterLongestPrefixWins(t *testing.T) {
	broad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "broad")
	}))
	defer broad.Close()
	specific := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "specific")
	}))
	defer specific.Close()

	router, err := NewRouter([]config.Route{
		{PathPrefix: "/api/", Upstream: broad.URL},
		{PathPrefix: "/api/admin/", Upstream: specific.URL},
	})
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/api/admin/users", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if got := res.Body.String(); got != "specific" {
		t.Fatalf("got %q, want specific", got)
	}
}

func TestRouterNotFound(t *testing.T) {
	router, err := NewRouter([]config.Route{{PathPrefix: "/known/", Upstream: "http://localhost:8081"}})
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusNotFound {
		t.Fatalf("status=%d, want 404", res.Code)
	}
}
