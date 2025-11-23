package httpserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/gofiber/fiber/middleware"
	"github.com/watchdog/internal/metrics"
)

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.InFlight.Inc()
		defer metrics.InFlight.Dec()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		next.ServeHTTP(ww, r)

		dur := time.Since(start)

		// ---- FIXED: correct chi v5 route lookup ----
		rc := chi.RouteContext(r.Context())
		route := rc.RoutePattern()
		if route == "" {
			route = r.URL.Path
		}

		// ---- FIXED: correct metric variable names ----
		metrics.RequestDuration.WithLabelValues(
			r.Method,
			route,
		).Observe(dur.Seconds())

		metrics.RequestsTotal.WithLabelValues(
			r.Method,
			route,
			strconv.Itoa(ww.Status()),
		).Inc()
	})
}
