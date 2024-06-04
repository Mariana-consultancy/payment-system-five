package repository

import "payment-system-one/internal/models"

func (p *Postgres) FindUserByAccountNumber(account_no int) (*models.Transaction, error) {
	transaction := &models.Transaction{}

	if err := p.DB.Where("account_no = ?", account_no).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// create a funds transfer in the database
func (p *Postgres) TransferFunds(transaction *models.Transaction) error {
	if err := p.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateFunds(transaction *models.Transaction) error {
	if err := p.DB.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}