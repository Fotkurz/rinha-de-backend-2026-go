package domain

type (
	Transaction struct {
		Amount       float32 `json:"amount" validate:"required"`
		Installments float32 `json:"installments" validate:"required"`
		RequestedAt  string  `json:"requested_at" validate:"required"`
	}

	Customer struct {
		AvgAmount      float32  `json:"avg_amount" validate:"required"`
		TxCount24h     float32  `json:"tx_count_24h" validate:"required"`
		KnownMerchants []string `json:"known_merchants" validate:"required"`
	}

	Merchant struct {
		ID        string  `json:"id" validate:"required"`
		MCC       string  `json:"mcc" validate:"required"`
		AvgAmount float32 `json:"avg_amount" validate:"required"`
	}

	Terminal struct {
		IsOnline    bool    `json:"is_online" validate:"required"`
		CardPresent bool    `json:"card_present" validate:"required"`
		KmFromHome  float32 `json:"km_from_home" validate:"required"`
	}

	LastTransaction struct {
		Timestamp     string  `json:"timestamp" validate:"required"`
		KmFromCurrent float32 `json:"km_from_current" validate:"required"`
	}
)

type FraudCheck struct {
	ID string `json:"id"`

	Transaction     Transaction      `json:"transaction"`
	Customer        Customer         `json:"customer"`
	Merchant        Merchant         `json:"merchant"`
	Terminal        Terminal         `json:"terminal"`
	LastTransaction *LastTransaction `json:"last_transaction"`

	Vector [14]float32 `json:"vector"`
}

func NewFraudCheck(id string, transaction Transaction, customer Customer, merchant Merchant, terminal Terminal, lastTransaction *LastTransaction) FraudCheck {
	f := FraudCheck{
		ID:              id,
		Transaction:     transaction,
		Customer:        customer,
		Merchant:        merchant,
		Terminal:        terminal,
		LastTransaction: lastTransaction,
		Vector:          [14]float32{},
	}

	return f
}
