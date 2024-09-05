package models

import (
	"testing"
)

func TestGetNextAccountID(t *testing.T) {
	initialID := accountCounter
	nextID := GetNextAccountID()
	if nextID != initialID {
		t.Errorf("GetNextAccountID() = %v; want %v", nextID, initialID)
	}
	if accountCounter != initialID+1 {
		t.Errorf("accountCounter = %v; want %v", accountCounter, initialID+1)
	}
}

func TestGetNextTransactionID(t *testing.T) {
	initialID := transactionCounter
	nextID := GetNextTransactionID()
	if nextID != initialID {
		t.Errorf("GetNextTransactionID() = %v; want %v", nextID, initialID)
	}
	if transactionCounter != initialID+1 {
		t.Errorf("transactionCounter = %v; want %v", transactionCounter, initialID+1)
	}
}
