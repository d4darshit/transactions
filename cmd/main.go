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
	db.GetDB().AutoMigrate(&models.Account{}, &models.Transaction{}, &models.OperationType{})

	//seed operations table
	models.SeedOperationTypes(db.GetDB())
	// Set up the router
	r := router.SetupRouter()

	// Start the server

	appPort := config.AppConfig.Server.Port

	fmt.Println("listening on ", appPort)

	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
