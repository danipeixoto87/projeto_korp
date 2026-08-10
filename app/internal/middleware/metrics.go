package middleware

import (
	"net/http"
	"strconv"
	"time"

	"http-server-projeto-korp/internal/metrics"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *statusResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *statusResponseWriter) Write(body []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	return rw.ResponseWriter.Write(body)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &statusResponseWriter{
			ResponseWriter: w,
			statusCode:     0,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		metrics.RequestsTotal.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(rw.statusCode),
		).Inc()

		metrics.RequestDuration.WithLabelValues(
			r.Method,
			r.URL.Path,
		).Observe(duration.Seconds())
	})
}
