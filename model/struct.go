package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

type GeneralAPIResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}

type Wallet struct {
	ID         string     `db:"id" json:"id"`
	OwnedBy    string     `db:"owned_by" json:"owned_by"`
	Balance    float64    `db:"balance" json:"balance"`
	Status     string     `db:"status" json:"status"`
	EnabledAt  *time.Time `db:"enabled_at" json:"enabled_at,omitempty"`
	DisabledAt *time.Time `db:"disabled_at" json:"disabled_at,omitempty"`
}

type Deposit struct {
	ID          string    `db:"id" json:"id"`
	DepositedBy string    `db:"deposited_by" json:"deposited_by"`
	ReferenceID string    `db:"reference_id" json:"reference_id"`
	Amount      float64   `db:"amount" json:"amount"`
	Status      string    `db:"status" json:"status"`
	DepositedAt time.Time `db:"deposited_at" json:"deposited_at"`
}

type InsertDeposit struct {
	DepositedBy string
	ReferenceID string
	Amount      float64
	Status      string
	DepositedAt time.Time
}

type DepositRequest struct {
	Amount      float64 `json:"amount" validate:"min=1"`
	ReferenceID string  `json:"reference_id" validate:"required"`
}

type Withdrawal struct {
	ID          string    `db:"id" json:"id"`
	WithdrawnBy string    `db:"withdrawn_by" json:"withdrawn_by"`
	ReferenceID string    `db:"reference_id" json:"reference_id"`
	Amount      float64   `db:"amount" json:"amount"`
	Status      string    `db:"status" json:"status"`
	WithdrawnAt time.Time `db:"withdrawn_at" json:"withdrawn_at"`
}

type InsertWithdrawal struct {
	WithdrawnBy string
	ReferenceID string
	Amount      float64
	Status      string
	WithdrawnAt time.Time
}

type WithdrawalRequest struct {
	Amount      float64 `json:"amount" validate:"min=1"`
	ReferenceID string  `json:"reference_id" validate:"required"`
}

type InitWalletRequest struct {
	CustomerXID string `json:"customer_xid" validate:"required"`
}

type InitWalletResponse struct {
	Token string `json:"token"`
}

type InsertWallet struct {
	OwnedBy string
	Status  string
}

type UpdateWalletStatus struct {
	CustomerID string
	Status     string
	EnabledAt  *time.Time
	DisabledAt *time.Time
}

type UpdateWalletBalance struct {
	CustomerID string
	Increment  float64
}
