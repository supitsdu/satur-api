package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func MongoDBURI() string {
	return os.Getenv("MONGODB_URI")
}

func MongoDBDatabaseID() string {
	return os.Getenv("MONGODB_ID")
}

func ServerAddress() string {
	return os.Getenv("SERVER_ADDRESS")
}
