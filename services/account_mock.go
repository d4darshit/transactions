package services

import (
	"github.com/stretchr/testify/mock"
	"transactions/models"
)

// Mock implementation for AccountRepo
type MockAccountRepo struct {
	mock.Mock
}

func NewAccountRepoMock() models.AccountRepo {
	return &MockAccountRepo{}
}

func (m *MockAccountRepo) CreateAccount(documentNumber string) (*models.Account, error) {
	args := m.Called(documentNumber)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountRepo) GetAccount(accountID uint) (*models.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).(*models.Account), args.Error(1)
}
