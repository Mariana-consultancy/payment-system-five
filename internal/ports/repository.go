package ports

import "payment-system-five/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	FindAdminByEmail(email string) (*models.Admin, error)
	CreateAdmin(admin *models.Admin) error
	UpdateAdmin(admin *models.Admin) error
	FindUserByAccountNumber(account_no int) (*models.User, error)
	TransferFunds(user *models.User, recipient *models.User, amount float64) error
	//ViewUserBalance(user *models.User) error
	//ViewUserTransactionsHistory(transaction *models.Transaction) error
	GetTransactionByAccountNumber(account_no int) (*[]models.Transaction, error)
	//AdminViewAllTransactionsHistory(transaction *models.Transaction) error
	//AdminViewUserBalance(user *models.User) error
}
