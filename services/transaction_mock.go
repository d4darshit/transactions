package services

import (
	"github.com/stretchr/testify/mock"
	"transactions/models"
)

// Mock implementation for TransactionRepo
type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) CreateTransaction(accountID uint, operationTypeID int, amount float64) (*models.Transaction, error) {
	args := m.Called(accountID, operationTypeID, amount)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) GetTransactionsByAccount(accountID uint) ([]models.Transaction, error) {
	args := m.Called(accountID)
	return args.Get(0).([]models.Transaction), args.Error(1)
}
