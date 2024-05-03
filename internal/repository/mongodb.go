package repository

import (
	"context"

	"github.com/supitsdu/satur-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ErrNotFound represents a not found error

type Actions interface {
	GetAccount(id string) (*model.AccountPersonalData, error)
	CreateAccount(account *model.AccountPersonalData) error
	DeleteAccount(id string) error
}

type MongoDBRepo struct {
	collection *mongo.Collection
}

// SetupCollection creates a new MongoDB collection
func SetupCollection(client *mongo.Client, dbName, collectionName string) (*mongo.Collection, error) {
	database := client.Database(dbName)
	collection := database.Collection(collectionName)
	return collection, nil
}

// Attempts to estabilish a new instance of MongoDBRepo with the provide arguments
func ConnectMongoDBRepo(connectionString, databaseId, collectionName string) (*MongoDBRepo, error) {
	opts := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	collection, err := SetupCollection(client, databaseId, collectionName)
	if err != nil {
		return nil, err
	}

	return &MongoDBRepo{
		collection: collection,
	}, nil
}

// GetAccount retrieves an account by ID
func (r *MongoDBRepo) GetAccount(username string) (*model.AccountPersonalData, error) {
	var account model.AccountPersonalData
	filter := bson.M{"username": username}
	err := r.collection.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}
	return &account, nil
}

// CreateAccount creates a new account
func (r *MongoDBRepo) CreateAccount(account *model.AccountPersonalData) error {
	_, err := r.collection.InsertOne(context.Background(), account)
	return err
}

// DeleteAccount deletes an account by ID
func (r *MongoDBRepo) DeleteAccount(username string) error {
	filter := bson.M{"username": username}
	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
