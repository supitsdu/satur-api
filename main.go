package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Account represents an account entity
type Account struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	BirthDate    string `json:"birthDate"`
	EmailAddress string `json:"emailAddress"`
}

// Repository is an interface for data storage operations
type Repository interface {
	GetAccount(id string) (*Account, error)
	CreateAccount(account *Account) error
	DeleteAccount(id string) error
}

// MongoDBRepo is a MongoDB implementation of Repository
type MongoDBRepo struct {
	collection *mongo.Collection
}

// NewMongoDBRepo creates a new instance of MongoDBRepo
func NewMongoDBRepo(connectionString, databaseId, collectionName string) (*MongoDBRepo, error) {
	opts := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(databaseId)
	collection := db.Collection(collectionName)

	return &MongoDBRepo{
		collection: collection,
	}, nil
}

// GetAccount retrieves an account by ID
func (r *MongoDBRepo) GetAccount(id string) (*Account, error) {
	var account Account
	filter := bson.M{"id": id}
	err := r.collection.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

// CreateAccount creates a new account
func (r *MongoDBRepo) CreateAccount(account *Account) error {
	_, err := r.collection.InsertOne(context.Background(), account)
	return err
}

// DeleteAccount deletes an account by ID
func (r *MongoDBRepo) DeleteAccount(id string) error {
	filter := bson.M{"id": id}
	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

// APIError represents an error response
type APIError struct {
	Error string `json:"error"`
}

// Handler is the HTTP handler for the API
type Handler struct {
	repo Repository
}

// NewHandler creates a new instance of Handler
func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// GetAccountHandler handles GET requests for account information
func (h *Handler) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	account, err := h.repo.GetAccount(id)
	if err != nil {
		if err == ErrNotFound {
			WriteError(w, http.StatusNotFound, "Account not found")
			return
		}
		WriteError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	WriteJSON(w, http.StatusOK, account)
}

// CreateAccountHandler handles POST requests to create a new account
func (h *Handler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.repo.CreateAccount(&account); err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to create account")
		return
	}

	WriteJSON(w, http.StatusCreated, "Account successfully created")
}

// DeleteAccountHandler handles DELETE requests to delete an account
func (h *Handler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.DeleteAccount(id); err != nil {
		if err == ErrNotFound {
			WriteError(w, http.StatusNotFound, "Account not found")
			return
		}
		WriteError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}

// ErrNotFound represents a not found error
var ErrNotFound = mongo.ErrNoDocuments

// WriteError writes an error response
func WriteError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIError{Error: message})
}

// WriteJSON writes a JSON response
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// MongoDB configuration
	connectionString := os.Getenv("MONGODB_URI")
	databaseId := os.Getenv("MONGODB_ID")
	collectionName := "accounts"

	// Create MongoDB repository
	repo, err := NewMongoDBRepo(connectionString, databaseId, collectionName)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Create HTTP handler
	handler := NewHandler(repo)

	// Configure router
	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", handler.GetAccountHandler).Methods("GET")
	router.HandleFunc("/account", handler.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/account/{id}", handler.DeleteAccountHandler).Methods("DELETE")

	// Specify the address to listen on
	address := "localhost:8080"

	// Start the HTTP server
	log.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, router))
}
