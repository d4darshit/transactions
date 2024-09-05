package services

import (
	"errors"
	"transactions/models"
)

func CreateAccount(documentNumber string) models.Account {
	account := models.Account{
		AccountID:      models.GetNextAccountID(),
		DocumentNumber: documentNumber,
	}
	models.Accounts[account.AccountID] = account
	return account
}

func GetAccount(accountId int) (models.Account, error) {
	account, exists := models.Accounts[accountId]
	if !exists {
		return models.Account{}, errors.New("account not found")
	}
	return account, nil
}
