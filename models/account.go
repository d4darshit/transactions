package models

type Account struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

var Accounts = make(map[int]Account)
