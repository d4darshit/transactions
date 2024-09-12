package models

import (
	"testing"
	"transactions/db"

	"github.com/stretchr/testify/assert"
)

func setup() {
	// Seed initial data if needed
	account := &Account{DocumentNumber: "12345678900"}
	db.DB.Create(account)
}

func TestCreateTransaction(t *testing.T) {
	setup()

	// Test creating a transaction
	transaction, err := CreateTransaction(1, 4, 123.45)

	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, uint(1), transaction.AccountID)
	assert.Equal(t, 4, transaction.OperationTypeID)
	assert.Equal(t, 123.45, transaction.Amount)

	// Verify transaction creation in database
	var createdTransaction Transaction
	result := db.DB.First(&createdTransaction, transaction.TransactionID)

	assert.NoError(t, result.Error)
	assert.Equal(t, uint(1), createdTransaction.AccountID)
	assert.Equal(t, 4, createdTransaction.OperationTypeID)
	assert.Equal(t, 123.45, createdTransaction.Amount)
}

func TestGetTransactionsByAccount(t *testing.T) {
	setup()

	// Create transactions
	_, _ = CreateTransaction(1, 4, 123.45)
	_, _ = CreateTransaction(1, 4, 678.90)

	// Test retrieving transactions
	transactions, err := GetTransactionsByAccount(1)

	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	assert.Equal(t, uint(1), transactions[0].AccountID)
	assert.Equal(t, 4, transactions[0].OperationTypeID)
	assert.Equal(t, 123.45, transactions[0].Amount)

	assert.Equal(t, uint(1), transactions[1].AccountID)
	assert.Equal(t, 4, transactions[1].OperationTypeID)
	assert.Equal(t, 678.90, transactions[1].Amount)
}
