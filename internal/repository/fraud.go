package repository

import (
	"context"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/repository/entity"
)

type Fraud struct{}

func NewFraudRepository() *Fraud {
	return &Fraud{}
}

func (r Fraud) FindSiblingsByEuclidian(ctx context.Context, vector [14]float32, k int) ([]entity.VectorEntity, error) {

	return []entity.VectorEntity{}, nil
}
