package services

import (
	"testing"
)

func TestCreateAccountService(t *testing.T) {
	documentNumber := "12345678900"
	accountSvc := AccountService{}
	account, _ := accountSvc.CreateAccount(documentNumber)

	if account.DocumentNumber != documentNumber {
		t.Errorf("expected document number %v, got %v", documentNumber, account.DocumentNumber)
	}

	if account.AccountID == 0 {
		t.Errorf("expected a valid account ID, got %v", account.AccountID)
	}
}

func TestGetAccountService(t *testing.T) {
	documentNumber := "12345678900"
	accountSvc := AccountService{}

	createdAccount, err := accountSvc.CreateAccount(documentNumber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	account, err := accountSvc.GetAccount(createdAccount.AccountID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if account.AccountID != createdAccount.AccountID {
		t.Errorf("expected account ID %v, got %v", createdAccount.AccountID, account.AccountID)
	}

	if account.DocumentNumber != createdAccount.DocumentNumber {
		t.Errorf("expected document number %v, got %v", createdAccount.DocumentNumber, account.DocumentNumber)
	}
}

func TestGetAccountService_NotFound(t *testing.T) {
	accountSvc := AccountService{}
	_, err := accountSvc.GetAccount(999)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != "account not found" {
		t.Errorf("expected 'account not found', got %v", err.Error())
	}
}
