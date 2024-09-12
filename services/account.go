package services

import "transactions/models"

type AccountService struct{}

func (as *AccountService) CreateAccount(documentNumber string) (*models.Account, error) {
	return models.CreateAccount(documentNumber)
}

func (as *AccountService) GetAccount(accountID uint) (*models.Account, error) {
	return models.GetAccount(accountID)
}
