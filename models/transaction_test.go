package models

import (
	"testing"
	"time"
	"transactions/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateTransaction(t *testing.T) {
	// Initialize the mock DB
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	// Mock the `SELECT VERSION()` query
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.30"))

	// Set up Gorm DB with the mock DB using MySQL driver
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)
	db.DB = gormDB

	// Set up expectations for the INSERT query with all necessary fields
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `transactions`").
		WithArgs(
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // updated_at
			sqlmock.AnyArg(), // deleted_at
			uint(1),          // account_id
			int(2),           // operation_type_id
			float64(100.0),   // amount
			sqlmock.AnyArg(), // event_date
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function being tested
	transaction, err := TransactionImpl{}.CreateTransaction(1, 2, 100.0)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, uint(1), transaction.AccountID)
	assert.Equal(t, int(2), transaction.OperationTypeID)
	assert.Equal(t, float64(100.0), transaction.Amount)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTransactionsByAccount(t *testing.T) {
	// Initialize the mock DB
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	// Mock the `SELECT VERSION()` query
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.30"))

	// Set up Gorm DB with the mock DB using MySQL driver
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)

	db.DB = gormDB
	// Define fake transactions rows to return
	rows := sqlmock.NewRows([]string{"transaction_id", "account_id", "operation_type_id", "amount", "event_date"}).
		AddRow(1, 1, 2, 100.0, time.Now()).
		AddRow(2, 1, 3, 150.0, time.Now())

	// Set up expectations for the SELECT query
	mock.ExpectQuery("SELECT \\* FROM `transactions` WHERE account_id = \\?").
		WithArgs(uint(1)).
		WillReturnRows(rows)

	// Call the function being tested
	transactions, err := TransactionImpl{}.GetTransactionsByAccount(1)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)
	assert.Equal(t, uint(1), transactions[0].AccountID)
	assert.Equal(t, float64(100.0), transactions[0].Amount)
	assert.Equal(t, float64(150.0), transactions[1].Amount)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
