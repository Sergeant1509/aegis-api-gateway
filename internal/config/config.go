// Package config loads and validates gateway route configuration.
package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Route struct {
	PathPrefix string `json:"path_prefix"`
	Upstream string `json:"upstream"`
	StripPrefix bool `json:"strip_prefix"`
}
type File struct { Routes []Route `json:"routes"` }

func Load(path string) (File, error) {
	data, err := os.ReadFile(path)
	if err != nil { return File{}, fmt.Errorf("read route config: %w", err) }
	var cfg File
	if err := json.Unmarshal(data, &cfg); err != nil { return File{}, fmt.Errorf("parse route config: %w", err) }
	if len(cfg.Routes) == 0 { return File{}, fmt.Errorf("at least one route is required") }
	for i, r := range cfg.Routes {
		if !strings.HasPrefix(r.PathPrefix, "/") { return File{}, fmt.Errorf("route %d path_prefix must start with /", i) }
		u, err := url.Parse(r.Upstream)
		if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") { return File{}, fmt.Errorf("route %d has invalid upstream URL %q", i, r.Upstream) }
	}
	return cfg, nil
}
