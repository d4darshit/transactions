package models

var (
	accountCounter     = 1
	transactionCounter = 1
)

func GetNextAccountID() int {
	id := accountCounter
	accountCounter++
	return id
}

func GetNextTransactionID() int {
	id := transactionCounter
	transactionCounter++
	return id
}
