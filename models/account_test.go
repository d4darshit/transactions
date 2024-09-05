package models

import (
	"testing"
)

func TestAccountCreation(t *testing.T) {
	account := Account{
		AccountID:      1,
		DocumentNumber: "12345678900",
	}

	Accounts[account.AccountID] = account

	if Accounts[1].DocumentNumber != "12345678900" {
		t.Errorf("expected document number 12345678900, got %v", Accounts[1].DocumentNumber)
	}
}
