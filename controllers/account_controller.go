package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"transactions/models"
	"transactions/services"

	"github.com/gorilla/mux"
)

// CreateAccount handles POST /accounts
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdAccount := services.CreateAccount(account.DocumentNumber)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAccount)
}

// GetAccount handles GET /accounts/{accountId}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, err := strconv.Atoi(vars["accountId"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := services.GetAccount(accountId)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
