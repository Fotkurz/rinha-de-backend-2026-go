package handler

import (
	"net/http"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/config"
)

type Ready struct {
}

func NewReadyHandler() Ready {
	return Ready{}
}

func (h Ready) IsReady(w http.ResponseWriter, r *http.Request) {
	if config.Instance().IsReady {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
}
