package routes

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"http-server-projeto-korp/internal/handlers"
)

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", handlers.ProjetoKorpHandler)

	mux.HandleFunc("/health", handlers.HealthHandler)

	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
