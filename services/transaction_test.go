// services/transaction_impl_test.go
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"transactions/models"
)

// Test CreateTransaction method of TransactionImpl
func TestCreateTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionImpl{Repo: mockRepo}

	// Define input and expected output
	accountID := uint(1)
	operationTypeID := 2
	amount := 100.0
	expectedTransaction := &models.Transaction{
		TransactionID:   1,
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}

	// Set up expectations
	mockRepo.On("CreateTransaction", accountID, operationTypeID, amount).Return(expectedTransaction, nil)

	// Call the method being tested
	transaction, err := service.CreateTransaction(accountID, operationTypeID, amount)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, transaction)
	mockRepo.AssertExpectations(t)
}

// Test GetTransactionsByAccount method of TransactionImpl
func TestGetTransactionsByAccount(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	service := &TransactionImpl{Repo: mockRepo}

	// Define input and expected output
	accountID := uint(1)
	expectedTransactions := []models.Transaction{
		{TransactionID: 1, AccountID: accountID, OperationTypeID: 2, Amount: 100.0},
		{TransactionID: 2, AccountID: accountID, OperationTypeID: 3, Amount: 200.0},
	}

	// Set up expectations
	mockRepo.On("GetTransactionsByAccount", accountID).Return(expectedTransactions, nil)

	// Call the method being tested
	transactions, err := service.GetTransactionsByAccount(accountID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
	mockRepo.AssertExpectations(t)
}
