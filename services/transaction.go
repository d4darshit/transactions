package services

import (
	"errors"
	"time"

	"transactions/models"
)

func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if _, exists := models.Accounts[transaction.AccountID]; !exists {
		return models.Transaction{}, errors.New("account not found")
	}

	transaction.TransactionID = models.GetNextTransactionID()
	transaction.EventDate = time.Now()

	models.Transactions[transaction.TransactionID] = transaction
	return transaction, nil
}
