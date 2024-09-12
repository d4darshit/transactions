package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"transactions/services"

	"github.com/gorilla/mux"
)

var accountService = services.AccountService{}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DocumentNumber string `json:"document_number"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	account, err := accountService.CreateAccount(req.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Use mux to get URL parameters
	idStr := vars["accountId"]
	fmt.Println(r.URL.Query())
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accountService.GetAccount(uint(accountID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}
