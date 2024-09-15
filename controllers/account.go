package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transactions/contracts"
	"transactions/models"
	"transactions/services"

	"github.com/gorilla/mux"
)

var accountService = services.NewAccountService(models.NewAccountRepo())

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	req := contracts.CreateAccountRequest{}

	//parse request
	json.NewDecoder(r.Body).Decode(&req)

	account, err := accountService.CreateAccount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// convert model to response object
	res := contracts.CreateAccountResponse{
		AccoundId:      int(account.AccountID),
		DocumentNumber: account.DocumentNumber,
	}

	json.NewEncoder(w).Encode(res)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Use mux to get URL parameters
	idStr := vars["accountId"]

	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := contracts.GetAccountRequest{AccountId: accountID}
	res := contracts.GetAccountResponse{}

	account, err := accountService.GetAccount(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// convert model to response object
	res.AccountId = int(account.AccountID)
	res.DocumentNumber = account.DocumentNumber

	json.NewEncoder(w).Encode(res)
}
