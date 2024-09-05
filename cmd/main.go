package main

import (
	"fmt"
	"log"
	"net/http"
	"transactions/router"
)

func main() {
	fmt.Println("Application started")

	r := router.SetupRouter()

	// Start server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
