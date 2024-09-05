package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"transactions/services"
)

func TestCreateTransaction(t *testing.T) {
	// Add a test account
	services.CreateAccount("12345678900")

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

	expected := `{"transaction_id":1,"account_id":1,"operation_type_id":4,"amount":123.45`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want contains %v", rr.Body.String(), expected)
	}
}
