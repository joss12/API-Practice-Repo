package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "watchdog",
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests",
		},
		[]string{"method", "route", "status"},
	)

	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "watchdog",
			Name:      "http_request_duration_seconds",
			Help:      "Request duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	InFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "watchdog",
			Name:      "in_flight_requests",
			Help:      "Number of in-flight HTTP requests",
		},
	)
)

func Handler() http.Handler {
	return promhttp.Handler()
}
