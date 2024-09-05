package models

import (
	"testing"
	"time"
)

func TestTransactionCreation(t *testing.T) {
	transaction := Transaction{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -50.0,
		EventDate:       time.Now(),
	}

	Transactions[transaction.TransactionID] = transaction

	if Transactions[1].Amount != -50.0 {
		t.Errorf("expected amount -50.0, got %v", Transactions[1].Amount)
	}
}
