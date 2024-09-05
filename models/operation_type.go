package models

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"`
	Description     string `json:"description"`
}

var OperationTypes = map[int]OperationType{
	1: {OperationTypeID: 1, Description: "Normal Purchase"},
	2: {OperationTypeID: 2, Description: "Purchase with Installments"},
	3: {OperationTypeID: 3, Description: "Withdrawal"},
	4: {OperationTypeID: 4, Description: "Credit Voucher"},
}
