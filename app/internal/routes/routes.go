package routes

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"http-server-projeto-korp/internal/handlers"
	"http-server-projeto-korp/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", handlers.ProjetoKorpHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.Handle("/metrics", promhttp.Handler())

	return middleware.Metrics(mux)
}
