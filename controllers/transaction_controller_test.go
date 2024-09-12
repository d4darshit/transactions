package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"transactions/db"
)

func init() {
	db.Connect()
}
func TestCreateTransaction(t *testing.T) {

	req, err := http.NewRequest("POST", "/transactions", strings.NewReader(`{"account_id": 1, "operation_type_id": 4, "amount": 123.45}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to parse response: %v", err)
	}

	if transactionID, ok := response["transaction_id"]; !ok || transactionID != float64(1) {
		t.Errorf("handler returned unexpected body: got %v want contains %v", response, "transaction_id:1")
	}
}
