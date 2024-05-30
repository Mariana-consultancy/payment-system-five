package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Password          string  `json:"password"`
	DateOfBirth       string  `json:"date_of_birth"`
	Email             string  `json:"email"`
	AccountNo         int     `json:"account_no"`
	AvailableBalance  float64 `json:"available_balance"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`	
	
}

//type UserProfile struct {
//	gorm.Model
//	ValidIdentity string `json:"valid_identity"`
//	PassPort string `json:"passport"`
//
//}

type Transaction struct {
	gorm.Model
	UserID          uint    `json:"user_id"`
	Amount          float64 `json:"amount"`
	Reference       string  `json:"reference"`
	TransactionType string  `json:"transaction_type"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
