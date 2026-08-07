package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"http-server-projeto-korp/internal/metrics"
	"http-server-projeto-korp/internal/models"
)

func ProjetoKorpHandler(w http.ResponseWriter, r *http.Request) {

	metrics.RequestsTotal.Inc()

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	response := models.Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
