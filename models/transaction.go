package models

import (
	"time"
	"transactions/db"
)

type Transaction struct {
	BaseModel                 // default columns
	TransactionID   uint      `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
	AccountID       uint      `gorm:"not null" json:"account_id"`
	OperationTypeID int       `gorm:"not null" json:"operation_type_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	EventDate       time.Time `gorm:"autoCreateTime" json:"event_date"`
}

type TransactionRepo interface {
	CreateTransaction(accountID uint, operationTypeID int, amount float64) (*Transaction, error)
	GetTransactionsByAccount(accountID uint) ([]Transaction, error)
}
type TransactionImpl struct {
}

func NewTransactionRepo() TransactionRepo {
	return TransactionImpl{}
}

func (t TransactionImpl) CreateTransaction(accountID uint, operationTypeID int, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}
	result := db.GetDB().Create(transaction)
	return transaction, result.Error
}

func (t TransactionImpl) GetTransactionsByAccount(accountID uint) ([]Transaction, error) {
	var transactions []Transaction
	result := db.GetDB().Where("account_id = ?", accountID).Find(&transactions)
	return transactions, result.Error
}
