package models

import (
	"time"
	"transactions/db"
)

type Transaction struct {
	TransactionID   uint      `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
	AccountID       uint      `gorm:"not null" json:"account_id"`
	OperationTypeID int       `gorm:"not null" json:"operation_type_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	EventDate       time.Time `gorm:"autoCreateTime" json:"event_date"`
}

func CreateTransaction(accountID uint, operationTypeID int, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}
	result := db.DB.Create(transaction)
	return transaction, result.Error
}

func GetTransactionsByAccount(accountID uint) ([]Transaction, error) {
	var transactions []Transaction
	result := db.DB.Where("account_id = ?", accountID).Find(&transactions)
	return transactions, result.Error
}
