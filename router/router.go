package router

import (
	"transactions/controllers"

	"github.com/gorilla/mux"
)

// SetupRouter configures the routes and returns a new router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Account routes
	r.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{accountId}", controllers.GetAccount).Methods("GET")

	// Transaction routes
	r.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	return r
}
