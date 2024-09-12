package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transactions/services"
)

var transactionService = services.TransactionService{}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID       uint    `json:"account_id"`
		OperationTypeID int     `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	transaction, err := transactionService.CreateTransaction(req.AccountID, req.OperationTypeID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	id, _ := strconv.Atoi(accountID)
	transactions, err := transactionService.GetTransactionsByAccount(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}
