package models

import (
	"gorm.io/gorm"
	"log"
)

// OperationType represents the operation type entity.
type OperationType struct {
	OperationTypeID int    `gorm:"not null" json:"operation_type_id"`
	Description     string `gorm:"not null" json:"description"`
}

var OperationTypes = map[int]OperationType{
	1: {OperationTypeID: 1, Description: "Normal Purchase"},
	2: {OperationTypeID: 2, Description: "Purchase with Installments"},
	3: {OperationTypeID: 3, Description: "Withdrawal"},
	4: {OperationTypeID: 4, Description: "Credit Voucher"},
}

func SeedOperationTypes(db *gorm.DB) {
	for _, opType := range OperationTypes {
		// Use FirstOrCreate to avoid duplication
		err := db.Where(OperationType{OperationTypeID: opType.OperationTypeID}).
			Attrs(OperationType{Description: opType.Description}).
			FirstOrCreate(&opType).Error

		if err != nil {
			log.Printf("Failed to seed operation type %d: %v\n", opType.OperationTypeID, err)
		}
	}
}
