package services

import (
	"testing"
	"transactions/models"
)

func TestCreateTransactionService(t *testing.T) {
	accountSvc := AccountService{}
	account, err := accountSvc.CreateAccount("12345678900")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	transaction := models.Transaction{
		AccountID:       account.AccountID,
		OperationTypeID: 4,
		Amount:          123.45,
	}

	transactionSvc := TransactionService{}
	createdTransaction, err := transactionSvc.CreateTransaction(transaction.AccountID, transaction.OperationTypeID, transaction.Amount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if createdTransaction.TransactionID == 0 {
		t.Errorf("expected a valid transaction ID, got %v", createdTransaction.TransactionID)
	}

	if createdTransaction.AccountID != transaction.AccountID {
		t.Errorf("expected account ID %v, got %v", transaction.AccountID, createdTransaction.AccountID)
	}

	if createdTransaction.Amount != transaction.Amount {
		t.Errorf("expected amount %v, got %v", transaction.Amount, createdTransaction.Amount)
	}
}

func TestCreateTransactionService_AccountNotFound(t *testing.T) {
	transaction := models.Transaction{
		AccountID:       999, // non-existent account ID
		OperationTypeID: 1,
		Amount:          -50.0,
	}

	transactionSvc := TransactionService{}
	_, err := transactionSvc.CreateTransaction(transaction.AccountID, transaction.OperationTypeID, transaction.Amount)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != "account not found" {
		t.Errorf("expected 'account not found', got %v", err.Error())
	}
}
