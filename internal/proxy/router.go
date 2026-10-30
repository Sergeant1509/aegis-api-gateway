// Package proxy implements prefix-based reverse proxy routing.
package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sort"
	"strings"

	"github.com/Sergeant1509/aegis-api-gateway/internal/config"
)

type route struct { prefix string; stripPrefix bool; proxy *httputil.ReverseProxy }
type Router struct { routes []route }

func NewRouter(configs []config.Route) (*Router, error) {
	r := &Router{}
	for _, item := range configs {
		target, err := url.Parse(item.Upstream)
		if err != nil { return nil, fmt.Errorf("parse upstream %q: %w", item.Upstream, err) }
		rp := httputil.NewSingleHostReverseProxy(target)
		rp.ErrorHandler = func(w http.ResponseWriter, _ *http.Request, _ error) { http.Error(w, "upstream unavailable", http.StatusBadGateway) }
		r.routes = append(r.routes, route{prefix:item.PathPrefix, stripPrefix:item.StripPrefix, proxy:rp})
	}
	sort.Slice(r.routes, func(i,j int) bool { return len(r.routes[i].prefix) > len(r.routes[j].prefix) })
	return r, nil
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, item := range r.routes {
		if !strings.HasPrefix(req.URL.Path, item.prefix) { continue }
		if item.stripPrefix {
			clone := req.Clone(req.Context())
			u := *req.URL
			u.Path = strings.TrimPrefix(u.Path, item.prefix)
			if u.Path == "" { u.Path = "/" }
			clone.URL = &u
			item.proxy.ServeHTTP(w, clone)
			return
		}
		item.proxy.ServeHTTP(w, req)
		return
	}
	http.NotFound(w, req)
}
