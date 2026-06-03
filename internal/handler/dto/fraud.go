package dto

import "github.com/Fotkurz/rinha-de-backend-2026-go/internal/domain"

type (
	FraudCheckRequest struct {
		ID string `json:"id" validate:"required"`

		Transaction     domain.Transaction      `json:"transaction" validate:"required"`
		Customer        domain.Customer         `json:"customer" validate:"required"`
		Merchant        domain.Merchant         `json:"merchant" validate:"required"`
		Terminal        domain.Terminal         `json:"terminal" validate:"required"`
		LastTransaction *domain.LastTransaction `json:"last_transaction"`
	}
)

type FraudCheckResponse struct {
	Approved   bool    `json:"approved"`
	FraudScore float32 `json:"fraud_score"`
}
