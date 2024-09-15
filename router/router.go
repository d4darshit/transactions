package router

import (
	"github.com/gorilla/mux"
	"transactions/controllers"
)

// SetupRouter configures the routes and returns a new router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// API route group
	apiRouter := r.PathPrefix("/").Subrouter()

	// Account routes
	accountRouter := apiRouter.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("", controllers.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/{accountId}", controllers.GetAccount).Methods("GET")

	// Transaction routes
	transactionRouter := apiRouter.PathPrefix("/transactions").Subrouter()
	transactionRouter.HandleFunc("", controllers.CreateTransaction).Methods("POST")
	transactionRouter.HandleFunc("/{accountId}", controllers.GetTransactions).Methods("GET")

	return r
}
