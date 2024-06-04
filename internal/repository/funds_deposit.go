package repository

import "payment-system-one/internal/models"

func (p *Postgres) FindUserDepositByAccountNumber(account_no int) (*models.Deposit, error) {
	deposit := &models.Deposit{}

	if err := p.DB.Where("account_no = ?", account_no).First(&deposit).Error; err != nil {
		return nil, err
	}
	return deposit, nil
}

// create a funds deposit in the database
func (p *Postgres) DepositFunds(deposit *models.Deposit) error {
	if err := p.DB.Create(deposit).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateDepositFunds(deposit *models.Deposit) error {
	if err := p.DB.Save(deposit).Error; err != nil {
		return err
	}
	return nil
}