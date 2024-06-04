package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	PayerAccount      int        `json:"payer_account"`
	RecipientAccount  int        `json:"recipient_account"`
	TransactionAmount float64    `json:"transaction_amount"`
	TransactionDate   time.Time  `json:"transaction_date"`
	TransactionType   string     `json:"transaction_type"`
}

type TransferFunds struct {
	RecipienctACC int     `json:"recipient_acc"`
	Amount       float64  `json:"amount"`
}
