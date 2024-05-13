package ports

import "payment-system-five/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
}
