package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/ready"
)

type ReadyResponse struct {
	Status bool `json:"status"`
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	status := ReadyResponse{Status: ready.IsReady()}

	json.NewEncoder(w).Encode(status)

	ready.ToggleReady()
}
