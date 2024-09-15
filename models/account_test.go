package models

import (
	"testing"
	"transactions/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateAccount(t *testing.T) {
	// Initialize the mock DB
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	// Expect the 'SELECT VERSION()' query from Gorm
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.25"))

	// Set up Gorm DB with the mock DB using MySQL driver
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)
	db.DB = gormDB

	// Set up expectations for the mock DB
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `accounts`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the function being tested
	account, err := AccountImpl{}.CreateAccount("12345678900")

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "12345678900", account.DocumentNumber)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAccount(t *testing.T) {
	// Initialize the mock DB
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	// Expect the 'SELECT VERSION()' query from Gorm
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.25"))

	// Set up Gorm DB with the mock DB using MySQL driver
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)
	db.DB = gormDB

	// Define a fake account row to return
	rows := sqlmock.NewRows([]string{"account_id", "document_number"}).
		AddRow(1, "12345678900")

	// Set up expectations for the SELECT query
	mock.ExpectQuery("SELECT \\* FROM `accounts` WHERE `accounts`.`account_id` = \\? AND `accounts`.`deleted_at` IS NULL ORDER BY `accounts`.`account_id` LIMIT \\?").
		WithArgs(1, 1). // Expect 2 arguments: accountID and LIMIT 1
		WillReturnRows(rows)

	// Call the function being tested
	account, err := AccountImpl{}.GetAccount(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "12345678900", account.DocumentNumber)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
