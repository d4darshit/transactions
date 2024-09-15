package models

import (
	"transactions/db"
)

type Account struct {
	BaseModel             // default columns
	AccountID      uint   `gorm:"primaryKey;autoIncrement" json:"account_id"`
	DocumentNumber string `gorm:"not null" json:"document_number"`
}

// Define an interface for the account model operations
type AccountRepo interface {
	CreateAccount(documentNumber string) (*Account, error)
	GetAccount(accountID uint) (*Account, error)
}
type AccountImpl struct {
}

func NewAccountRepo() AccountRepo {
	return AccountImpl{}
}

func (a AccountImpl) CreateAccount(documentNumber string) (*Account, error) {
	account := &Account{DocumentNumber: documentNumber}
	result := db.GetDB().Create(account)
	return account, result.Error
}

func (a AccountImpl) GetAccount(accountID uint) (*Account, error) {
	var account Account
	result := db.GetDB().First(&account, accountID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
