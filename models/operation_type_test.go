package models

import (
	"testing"
)

func TestOperationType(t *testing.T) {
	if OperationTypes[1].Description != "Normal Purchase" {
		t.Errorf("expected description 'Normal Purchase', got %v", OperationTypes[1].Description)
	}

	if OperationTypes[4].Description != "Credit Voucher" {
		t.Errorf("expected description 'Credit Voucher', got %v", OperationTypes[4].Description)
	}
}
