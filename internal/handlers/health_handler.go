package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheckHandler retorna o status da API.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Cria uma resposta JSON simples
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
}
