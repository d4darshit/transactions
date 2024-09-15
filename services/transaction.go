package services

import "transactions/models"

type TransactionSvc interface {
	CreateTransaction(accountID uint, operationTypeID int, amount float64) (*models.Transaction, error)
	GetTransactionsByAccount(accountID uint) ([]models.Transaction, error)
}

type TransactionImpl struct {
	Repo models.TransactionRepo
}

func NewTransactionService(transactionRepo models.TransactionRepo) TransactionSvc {
	return &TransactionImpl{
		Repo: transactionRepo,
	}
}
func (ts *TransactionImpl) CreateTransaction(accountID uint, operationTypeID int, amount float64) (*models.Transaction, error) {
	return ts.Repo.CreateTransaction(accountID, operationTypeID, amount)
}

func (ts *TransactionImpl) GetTransactionsByAccount(accountID uint) ([]models.Transaction, error) {
	return ts.Repo.GetTransactionsByAccount(accountID)
}
