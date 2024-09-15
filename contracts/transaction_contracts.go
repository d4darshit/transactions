package contracts

import "time"

type CreateTransactionRequest struct {
	AccountId       uint    `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}
type TransactionResponse struct {
	AccountId       int       `json:"accountId"`
	OperationTypeId int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	TransactionId   int       `json:"transaction_id"`
	EventDate       time.Time `json:"event_date"`
}
type GetTransactionResponse struct {
	Response []TransactionResponse
}
