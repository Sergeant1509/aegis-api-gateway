package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sergeant1509/aegis-api-gateway/internal/config"
	"github.com/Sergeant1509/aegis-api-gateway/internal/health"
	"github.com/Sergeant1509/aegis-api-gateway/internal/middleware"
	"github.com/Sergeant1509/aegis-api-gateway/internal/proxy"
)

func main() {
	configPath := flag.String("config", "configs/routes.json", "route configuration file")
	addr := flag.String("listen", ":8080", "HTTP listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "aegis ", log.LstdFlags|log.LUTC)
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Fatalf("configuration: %v", err)
	}
	router, err := proxy.NewRouter(cfg.Routes)
	if err != nil {
		logger.Fatalf("router: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/healthz", health.Handler())
	mux.Handle("/", router)

	apiKey := os.Getenv("AEGIS_API_KEY")
	if apiKey == "" {
		logger.Fatal("AEGIS_API_KEY must be set")
	}
	limiter := middleware.NewTokenBucket(5, 20)
	handler := middleware.Chain(
		mux,
		middleware.RequestLog(logger),
		middleware.RateLimit(limiter),
		middleware.APIKey(apiKey),
	)

	server := &http.Server{
		Addr: *addr, Handler: handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		logger.Printf("listening on %s", *addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("server: %v", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Printf("graceful shutdown: %v", err)
	}
}
