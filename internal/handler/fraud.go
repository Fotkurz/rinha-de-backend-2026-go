package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/domain"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/handler/dto"
)

type FraudService interface {
	FraudCheck(ctx context.Context, transaction domain.FraudCheck) (approved bool, score float32, err error)
}

type fraudHandler struct {
	fraudService FraudService
}

func NewFraudHandler(fraudService FraudService) *fraudHandler {
	return &fraudHandler{
		fraudService: fraudService,
	}
}

func (h *fraudHandler) FraudCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.FraudCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	check := domain.NewFraudCheck(req.ID, req.Transaction, req.Customer, req.Merchant, req.Terminal, req.LastTransaction)

	approved, score, err := h.fraudService.FraudCheck(ctx, check)
	if err != nil {
		slog.Error("failed to run fraudcheck")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := dto.FraudCheckResponse{
		Approved:   approved,
		FraudScore: score,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
