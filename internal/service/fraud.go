package service

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/config"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/domain"
	"github.com/Fotkurz/rinha-de-backend-2026-go/internal/repository/entity"
	"github.com/Fotkurz/rinha-de-backend-2026-go/pkg/vector"
)

type FraudRepository interface {
	FindSiblingsByEuclidian(ctx context.Context, vector [14]float32, k int) ([]entity.VectorEntity, error)
}

const fraudThreshold = 0.6

type Fraud struct {
	repo FraudRepository
}

func NewFraudService(repo FraudRepository) *Fraud {
	return &Fraud{
		repo: repo,
	}
}

func (s *Fraud) FraudCheck(ctx context.Context, fraudCheck domain.FraudCheck) (approved bool, score float32, err error) {
	v, err := calculateVector(fraudCheck)
	if err != nil {
		slog.Error("failed to calculate vector for request", "err", err)
		return false, 0.0, fmt.Errorf("failed to calculate vector from payload, %w", err)
	}
	fraudCheck.Vector = v

	results, err := s.repo.FindSiblingsByEuclidian(ctx, fraudCheck.Vector, 5)
	if err != nil {
		slog.Error("failed to find siblings in database")
		return false, 0.0, fmt.Errorf("failed to find siblings for vector, err: %w", err)
	}
	if len(results) == 0 {
		slog.Warn("failed to find any neighbor for transaction",
			"id", fraudCheck.ID,
			"amount", fraudCheck.Transaction.Amount,
			"requested_at", fraudCheck.Transaction.RequestedAt,
			"merchant", fraudCheck.Merchant.ID)

		return false, 0, nil
	}

	isFraud, score := s.calculateFraud(results)
	return isFraud, score, nil
}

func (s *Fraud) calculateFraud(vectors []entity.VectorEntity) (bool, float32) {
	fraud := 0
	for _, vector := range vectors {
		if vector.Label == "fraud" {
			fraud++
		}
	}

	score := float32(fraud) / float32(len(vectors))
	return score < fraudThreshold, score
}

// calculateVector
//
//	Generate a 14 positions vector using the transaction request data
//
//	Ref:
//	- 'https://github.com/zanfranceschi/rinha-de-backend-2026/blob/main/docs/br/REGRAS_DE_DETECCAO.md#as-14-dimens%C3%B5es-do-vetor'
func calculateVector(transaction domain.FraudCheck) ([14]float32, error) {
	amount := vector.Normalize(transaction.Transaction.Amount, config.MaxAmount)
	installments := vector.Normalize(transaction.Transaction.Installments, config.MaxInstallments)
	amountVsAvg := vector.Normalize((transaction.Transaction.Amount / transaction.Transaction.Installments), config.AmountVSAvgRatio)

	parsedRequestedAt, err := time.Parse(time.RFC3339, transaction.Transaction.RequestedAt)
	if err != nil {
		return [14]float32{}, fmt.Errorf("not a valid timestamp for transaction.requestedAt: %s", transaction.Transaction.RequestedAt)
	}

	hourOfDay, dayOfWeek := calcRequestedAt(parsedRequestedAt)

	var minutesSinceLastTx float32 = -1
	var kmFromLastTx float32 = -1
	if transaction.LastTransaction != nil {
		minutesSinceLastTx = vector.Normalize(float32(parsedRequestedAt.Minute()), config.MaxMinutes)
		kmFromLastTx = vector.Normalize(float32(transaction.LastTransaction.KmFromCurrent), config.MaxKm)
	}
	kmFromHome := vector.Normalize(transaction.Terminal.KmFromHome, config.MaxKm)
	txCount24h := vector.Normalize(transaction.Customer.TxCount24h, config.MaxTxCount24H)
	var isOnline float32 = 0
	if transaction.Terminal.IsOnline {
		isOnline = 1
	}
	var cardPresent float32 = 0
	if transaction.Terminal.CardPresent {
		cardPresent = 1
	}

	unknownMerchant := float32(isUnknownMerchant(transaction.Customer.KnownMerchants, transaction.Merchant.ID))
	mccRisk := calcMccRisk(transaction.Merchant.MCC)
	merchantAvgAmount := vector.Normalize(transaction.Merchant.AvgAmount, config.MaxMerchantAvgAmount)

	v := [14]float32{
		amount,
		installments,
		amountVsAvg,
		hourOfDay,
		dayOfWeek,
		minutesSinceLastTx,
		kmFromLastTx,
		kmFromHome,
		txCount24h,
		isOnline,
		cardPresent,
		unknownMerchant,
		mccRisk,
		merchantAvgAmount,
	}

	return v, nil
}

func isUnknownMerchant(knownMerchants []string, merchantId string) int {
	if slices.Index(knownMerchants, merchantId) == -1 {
		return 1
	}
	return 0
}

// calcMccRisk
//
//	Based of https://github.com/zanfranceschi/rinha-de-backend-2026/blob/main/resources/mcc_risk.json
func calcMccRisk(mcc string) float32 {
	const defaultValue = 0.5
	reference := config.Instance().Mcc
	v, ok := reference[mcc]
	if !ok {
		return defaultValue
	}

	return v
}

func calcRequestedAt(t time.Time) (hourOfDay float32, dayOfWeek float32) {
	hourOfDay = float32(t.Hour()) / 23

	weekDay := func(wd time.Weekday) float32 {
		// golang format goes from sunday=0 to sat=6, challenge
		// challenge requires going from monday=0 to sunday=6
		if wd == time.Sunday {
			return 6
		}

		return float32(wd - 1)
	}(t.Weekday())
	dayOfWeek = weekDay / 6

	return
}
