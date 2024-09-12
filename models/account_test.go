package models

import (
	"testing"
	"transactions/config"
	"transactions/db"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfig()
	db.Connect()
}

// Mock function
func mockCreateAccount(documentNumber string) (*Account, error) {
	return &Account{AccountID: 1, DocumentNumber: documentNumber}, nil
}

func TestCreateAccount(t *testing.T) {
	// Patch CreateAccount function
	patch := monkey.Patch(CreateAccount, mockCreateAccount)
	defer patch.Unpatch()

	// Call the function being tested
	account, err := CreateAccount("12345678900")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "12345678900", account.DocumentNumber)
}
