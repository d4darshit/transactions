package main

import (
	"fmt"
	"log"
	"net/http"
	"transactions/config"
	"transactions/db"
	"transactions/models"
	"transactions/router"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Connect to the database
	db.Connect()

	// Auto-migrate the models (will create tables if they don't exist)
	db.DB.AutoMigrate(&models.Account{}, &models.Transaction{})

	// Set up the router
	r := router.SetupRouter()

	// Start the server
	fmt.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8081", r))
}
