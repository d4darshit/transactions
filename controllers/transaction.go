package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"transactions/contracts"
	"transactions/models"
	"transactions/services"
)

var transactionService = services.NewTransactionService(models.NewTransactionRepo())

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	req := contracts.CreateTransactionRequest{}
	res := contracts.TransactionResponse{}
	json.NewDecoder(r.Body).Decode(&req)

	transaction, err := transactionService.CreateTransaction(req.AccountId, req.OperationTypeId, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.TransactionId = int(transaction.TransactionID)
	res.Amount = transaction.Amount
	res.OperationTypeId = transaction.OperationTypeID
	res.AccountId = int(transaction.AccountID)
	res.EventDate = transaction.EventDate

	json.NewEncoder(w).Encode(res)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Use mux to get URL parameters
	accountID := vars["accountId"]
	id, _ := strconv.Atoi(accountID)

	transactions, err := transactionService.GetTransactionsByAccount(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := convertGetTransactionsResponse(transactions)

	json.NewEncoder(w).Encode(res.Response)
}
func convertGetTransactionsResponse(txnModels []models.Transaction) contracts.GetTransactionResponse {
	res := contracts.GetTransactionResponse{}
	txnArray := make([]contracts.TransactionResponse, 0)
	for _, transaction := range txnModels {
		newTransaction := contracts.TransactionResponse{
			AccountId:       int(transaction.AccountID),
			OperationTypeId: transaction.OperationTypeID,
			Amount:          transaction.Amount,
			TransactionId:   int(transaction.TransactionID),
			EventDate:       transaction.EventDate,
		}
		txnArray = append(txnArray, newTransaction)
	}
	res.Response = txnArray
	return res
}
