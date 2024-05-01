package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supitsdu/satur-api/internal/config"
	"github.com/supitsdu/satur-api/internal/handler"
	"github.com/supitsdu/satur-api/internal/repository"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// MongoDB configuration
	connectionString := config.MongoDBURI()
	databaseId := config.MongoDBDatabaseID()
	collectionName := "accounts"

	// MongoDB repository connection
	repo, err := repository.ConnectMongoDBRepo(connectionString, databaseId, collectionName)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	accountHandler := handler.UseAccountHandler(repo)

	// Configure router
	router := mux.NewRouter()
	router.HandleFunc("/account/{username}", accountHandler.GetMethod).Methods("GET")
	router.HandleFunc("/account", accountHandler.PostMethod).Methods("POST")
	router.HandleFunc("/account/{username}", accountHandler.DeleteMethod).Methods("DELETE")

	// Get the address to listen on
	address := config.ServerAddress()

	// Start the HTTP server
	log.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, router))
}
