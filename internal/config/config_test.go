package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadValidConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "routes.json")
	content := `{"routes":[{"path_prefix":"/api/","upstream":"http://localhost:8081"}]}`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Routes) != 1 || cfg.Routes[0].PathPrefix != "/api/" {
		t.Fatalf("unexpected config: %#v", cfg)
	}
}

func TestLoadRejectsInvalidUpstream(t *testing.T) {
	path := filepath.Join(t.TempDir(), "routes.json")
	content := `{"routes":[{"path_prefix":"/api/","upstream":"not-a-url"}]}`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Fatal("expected invalid upstream error")
	}
}
