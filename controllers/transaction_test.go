// controllers/transaction_controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"transactions/models"
)

// Mock implementation for TransactionSvc
type MockTransactionSvc struct {
	mock.Mock
}

func (m *MockTransactionSvc) CreateTransaction(accountID uint, operationTypeID int, amount float64) (*models.Transaction, error) {
	args := m.Called(accountID, operationTypeID, amount)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionSvc) GetTransactionsByAccount(accountID uint) ([]models.Transaction, error) {
	args := m.Called(accountID)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

// Test CreateTransaction handler
func TestCreateTransaction(t *testing.T) {
	mockService := new(MockTransactionSvc)
	expectedTransaction := &models.Transaction{
		TransactionID:   1,
		AccountID:       0,
		OperationTypeID: 1,
		Amount:          100.0,
	}
	mockService.On("CreateTransaction", uint(1), 1, 100.0).Return(expectedTransaction, nil)

	// Replace global transactionService with mock service
	originalService := transactionService
	transactionService = mockService
	defer func() { transactionService = originalService }()

	// Create request body
	requestBody, _ := json.Marshal(struct {
		AccountID       uint    `json:"account_id"`
		OperationTypeID int     `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100.0,
	})
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.Transaction
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, *expectedTransaction, response)
	mockService.AssertExpectations(t)
}
