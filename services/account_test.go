// services/account_impl_test.go
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"transactions/contracts"
	"transactions/models"
)

// Test CreateAccount method of AccountImpl
func TestCreateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := NewAccountService(mockRepo)

	// Define input and expected output
	req := contracts.CreateAccountRequest{DocumentNumber: "12345678900"}
	expectedAccount := &models.Account{AccountID: 1, DocumentNumber: "12345678900"}

	// Set up expectations
	mockRepo.On("CreateAccount", "12345678900").Return(expectedAccount, nil)

	// Call the method being tested
	account, err := service.CreateAccount(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)
	mockRepo.AssertExpectations(t)
}

// Test GetAccount method of AccountImpl
func TestGetAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := NewAccountService(mockRepo)

	// Define input and expected output
	req := contracts.GetAccountRequest{AccountId: 1}
	expectedAccount := &models.Account{AccountID: 1, DocumentNumber: "12345678900"}

	// Set up expectations
	mockRepo.On("GetAccount", uint(1)).Return(expectedAccount, nil)

	// Call the method being tested
	account, err := service.GetAccount(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)
	mockRepo.AssertExpectations(t)
}
