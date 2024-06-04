package models

import (
	"time"

	"gorm.io/gorm"
)

type Deposit struct {
	gorm.Model
	PayerAccount      int       `json:"payer_account"`
	RecipientAccount int        `json:"recipient_account"`
	DepositAmount    float64    `json:"deposit_amount"`
	DepositDate      time.Time  `json:"deposit_date"`
	DepositType      string     `json:"deposit_type"`

}

type AddFunds struct {
	Amount float64 `json:"amount"`
}