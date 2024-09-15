package services

import (
	"transactions/contracts"
	"transactions/models"
)

type AccountSvc interface {
	CreateAccount(createAccount contracts.CreateAccountRequest) (*models.Account, error)
	GetAccount(getAccount contracts.GetAccountRequest) (*models.Account, error)
}

type AccountImpl struct {
	Repo models.AccountRepo
}

func NewAccountService(accountRepo models.AccountRepo) AccountSvc {
	return &AccountImpl{Repo: accountRepo}
}

func (as *AccountImpl) CreateAccount(createAccount contracts.CreateAccountRequest) (*models.Account, error) {
	return as.Repo.CreateAccount(createAccount.DocumentNumber)
}

func (as *AccountImpl) GetAccount(getAccount contracts.GetAccountRequest) (*models.Account, error) {
	return as.Repo.GetAccount(uint(getAccount.AccountId))
}
