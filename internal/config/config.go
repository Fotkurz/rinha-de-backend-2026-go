package config

import (
	"log"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/Fotkurz/rinha-de-backend-2026-go/pkg/utils"
)

type Config struct {
	Logger  *slog.Logger
	IsReady bool
	Norm    Normalization
	Mcc     map[string]float32
}

type Normalization struct {
	MaxAmount            float32 `json:"max_amount,omitempty"`
	MaxInstallments      float32 `json:"max_installments,omitempty"`
	AmountVSAvgRatio     float32 `json:"amount_vs_avg_ratio,omitempty"`
	MaxTxCount24H        float32 `json:"max_tx_count_24_h,omitempty"`
	MaxMinutes           float32 `json:"max_minutes,omitempty"`
	MaxKm                float32 `json:"max_km,omitempty"`
	MaxMerchantAvgAmount float32 `json:"max_merchant_avg_amount,omitempty"`
}

const (
	MaxAmount            float32 = 10000
	MaxInstallments      float32 = 12
	AmountVSAvgRatio     float32 = 10
	MaxTxCount24H        float32 = 20
	MaxMinutes           float32 = 1440
	MaxKm                float32 = 1000
	MaxMerchantAvgAmount float32 = 10000
)

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		logger := LoadLogger()
		normalization := loadNormalizationDefaults()
		mcc := loadMccRiskReference()

		cfg = &Config{
			IsReady: false,
			Mcc:     mcc,
			Logger:  logger,
			Norm:    normalization,
		}
	})
	return cfg
}

func LoadLogger() *slog.Logger {
	envLevel := utils.ReadEnvOrDefault("LOG_LEVEL", "INFO")

	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(strings.ToUpper(envLevel))); err != nil {
		lvl = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	}))
	slog.SetDefault(logger)
	return logger
}

func loadMccRiskReference() map[string]float32 {
	const defaultMccRiskReferenceFile = "assets/mcc_risk.json"
	mccRiskFilePath := utils.ReadEnvOrDefault("MCC_RISK_FILE", defaultMccRiskReferenceFile)

	var mcc map[string]float32
	if err := utils.LoadStructFromJsonFile(mccRiskFilePath, &mcc); err != nil {
		log.Fatal(err)
	}

	return mcc
}

func loadNormalizationDefaults() Normalization {
	const defaultNormalizationFilePath = "assets/normalization.json"
	normalizationFilePath := utils.ReadEnvOrDefault("NORMALIZATION_FILE", defaultNormalizationFilePath)

	var n Normalization
	if err := utils.LoadStructFromJsonFile(normalizationFilePath, &n); err != nil {
		log.Fatal(err)
	}

	return n
}

func SetReady(v bool) {
	cfg.IsReady = v
}

func Instance() Config {
	if cfg == nil {
		return *LoadConfig()
	}
	return *cfg
}
