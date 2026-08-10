package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	},
	[]string{"method", "path", "status"},
)

var RequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request duration in seconds.",
	},
	[]string{"method", "path"},
)

func Register() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)
}
