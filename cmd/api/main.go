package main

import (
	"log"
	"net/http"

	"github.com/supitsdu/satur-api/internal/config"
	"github.com/supitsdu/satur-api/internal/repository"
	"github.com/supitsdu/satur-api/internal/routes"
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

	// Get the address to listen on
	address := config.ServerAddress()

	// Start the HTTP server
	log.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, routes.InitializeRouter(repo)))
}
