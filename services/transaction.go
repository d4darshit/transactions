package services

import "transactions/models"

type TransactionService struct{}

func (ts *TransactionService) CreateTransaction(accountID uint, operationTypeID int, amount float64) (*models.Transaction, error) {
	return models.CreateTransaction(accountID, operationTypeID, amount)
}

func (ts *TransactionService) GetTransactionsByAccount(accountID uint) ([]models.Transaction, error) {
	return models.GetTransactionsByAccount(accountID)
}
