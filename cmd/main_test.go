package main

import (
	"testing"
	"transactions/config"
	"transactions/db"
)

func TestMain(m *testing.M) {
	// Initialize the database
	db.Connect()

	// Initialize the configuration
	config.LoadConfig()

	// Run tests
	m.Run()
}
