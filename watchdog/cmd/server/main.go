package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/watchdog/internal/health"
	"github.com/watchdog/internal/httpserver"
	"github.com/watchdog/internal/metrics"
	"github.com/watchdog/internal/status"
	"github.com/watchdog/pkg/respond"
)

func main() {
	addr := env("ADDR", "8083")

	//Compose rouuter
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(httpserver.Metrics)

	stat := status.NewTracker()
	hRunner := health.NewRunnerFromEnv()

	// Liveness: process is healthy if it can answer
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		sum := hRunner.Run(ctx)
		status := http.StatusOK
		if !sum.Overall {
			status = http.StatusServiceUnavailable
		}
		respond.JSON(w, status, sum)
	})

	//Readiness: require extarnal checks to be OK(if any configurted)
	r.Get("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		if !hRunner.HasChecks() {
			respond.JSON(w, http.StatusOK, map[string]any{
				"overall": true,
				"message": "no external checks configured",
				"checks":  []any{},
			})
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		sum := hRunner.Run(ctx)
		if !sum.Overall {
			respond.JSON(w, http.StatusServiceUnavailable, sum)
			return
		}
		respond.JSON(w, http.StatusOK, sum)
	})

	//Status: build/runtime info
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		respond.JSON(w, http.StatusOK, stat.Snapshot())
	})

	//Prometheus
	r.Method(http.MethodGet, "/metrics", metrics.Handler())

	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	//graceful shutdown
	go func() {
		log.Printf("watchdog listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	///SIGINT/SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Println("server stopped")

}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
