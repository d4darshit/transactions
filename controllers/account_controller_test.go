package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"transactions/db"
)

func init() {
	db.Connect()
}
func TestCreateAccount(t *testing.T) {
	req, err := http.NewRequest("POST", "/accounts", strings.NewReader(`{"document_number": "12345678900"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
