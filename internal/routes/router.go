package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supitsdu/satur-api/internal/handler"
	"github.com/supitsdu/satur-api/internal/repository"
)

func UseAccountsRoutes(router *mux.Router, repo *repository.MongoDBRepo) *mux.Router {
	// Create handler instance
	handleAccounts := handler.CreateHandleAccounts(repo)

	// Define routes
	router.HandleFunc("/accounts/{username}", handleAccounts.GetMethod).Methods("GET")
	router.HandleFunc("/accounts", handleAccounts.PostMethod).Methods("POST")
	router.HandleFunc("/accounts/{username}", handleAccounts.DeleteMethod).Methods("DELETE")

	return router
}

func InitializeRouter(repo *repository.MongoDBRepo) http.Handler {
	// Create a new router instance
	router := mux.NewRouter()

	router = UseAccountsRoutes(router, repo)

	return router
}
