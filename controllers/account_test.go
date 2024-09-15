// controllers/controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"transactions/contracts"
	"transactions/models"
)

// Mock implementation for AccountSvc
type MockAccountSvc struct {
	mock.Mock
}

func (m *MockAccountSvc) CreateAccount(createAccount contracts.CreateAccountRequest) (*models.Account, error) {
	args := m.Called(createAccount)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountSvc) GetAccount(getAccount contracts.GetAccountRequest) (*models.Account, error) {
	args := m.Called(getAccount)
	return args.Get(0).(*models.Account), args.Error(1)
}

// Test CreateAccount handler
func TestCreateAccount(t *testing.T) {
	mockService := new(MockAccountSvc)
	expectedAccount := &models.Account{AccountID: 1, DocumentNumber: "12345678900"}
	mockService.On("CreateAccount", contracts.CreateAccountRequest{DocumentNumber: "12345678900"}).Return(expectedAccount, nil)

	// Replace global accountService with mock service
	originalService := accountService
	accountService = mockService
	defer func() { accountService = originalService }()

	// Create  request body
	requestBody, _ := json.Marshal(contracts.CreateAccountRequest{DocumentNumber: "12345678900"})
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := contracts.CreateAccountResponse{AccoundId: 1, DocumentNumber: "12345678900"}
	var response contracts.CreateAccountResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
	mockService.AssertExpectations(t)
}

// Test GetAccount handler
func TestGetAccount(t *testing.T) {
	mockService := new(MockAccountSvc)
	expectedAccount := &models.Account{AccountID: 1, DocumentNumber: "12345678900"}
	mockService.On("GetAccount", contracts.GetAccountRequest{AccountId: 1}).Return(expectedAccount, nil)

	// Replace global accountService with mock service
	originalService := accountService
	accountService = mockService
	defer func() { accountService = originalService }()

	// Create request
	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountId:[0-9]+}", GetAccount).Methods("GET")

	// Call the handler
	router.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	expected := contracts.CreateAccountResponse{AccoundId: 1, DocumentNumber: "12345678900"}
	var response contracts.CreateAccountResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
	mockService.AssertExpectations(t)
}
