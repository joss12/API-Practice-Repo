package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	registerOnce sync.Once

	ActiveWSConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "chat_ws_active_connections",
			Help: "Number of active WebSocket connections",
		},
	)

	MessagesPerRoom = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "chat_messages_total",
			Help: "Total chat messages per room",
		},
		[]string{"room"},
	)

	RedisLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "chat_redis_latency_seconds",
			Help:    "Redis operation latency",
			Buckets: prometheus.DefBuckets,
		},
	)

	DBLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "chat_db_latency_seconds",
			Help:    "Database query latency",
			Buckets: prometheus.DefBuckets,
		},
	)

	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "chat_http_requests_total",
			Help: "Total HTTP requests received",
		},
		[]string{"method", "path", "status"},
	)
)

func RegisterMetrics() {
	registerOnce.Do(func() {
		prometheus.MustRegister(ActiveWSConnections)
		prometheus.MustRegister(MessagesPerRoom)
		prometheus.MustRegister(RedisLatency)
		prometheus.MustRegister(DBLatency)
		prometheus.MustRegister(RequestCounter)
	})
}
