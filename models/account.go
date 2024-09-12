package models

import "transactions/db"

type Account struct {
	AccountID      uint   `gorm:"primaryKey;autoIncrement" json:"account_id"`
	DocumentNumber string `gorm:"not null" json:"document_number"`
}

func CreateAccount(documentNumber string) (*Account, error) {
	account := &Account{DocumentNumber: documentNumber}
	result := db.DB.Create(account)
	return account, result.Error
}

func GetAccount(accountID uint) (*Account, error) {
	var account Account
	result := db.DB.First(&account, accountID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
