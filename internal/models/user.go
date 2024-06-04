package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	Password         string  `json:"password"`
	DateOfBirth      string  `json:"date_of_birth"`
	Email            string  `json:"email"`
	AccountNo        int     `json:"account_no"`
	AvailableBalance float64 `json:"available_balance"`
	Phone            string  `json:"phone"`
	Address          string  `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Transaction struct {
	gorm.Model
	PayerAccount      int       `json:"payer_account"`
	RecipientAccount  int       `json:"recipient_account"`
	TransactionAmount float64   `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
	TransactionType   string    `json:"transaction_type"`
}

type TransferFunds struct {
	AccountNo int     `json:"account_no"`
	Amount    float64 `json:"amount"`
}

type AddFunds struct {
	Amount float64 `json:"amount"`
}