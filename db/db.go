package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jaavier/dotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	envDatabase   = "DB_NAME"
	envDBHost     = "DB_HOST"
	envDBPort     = "DB_PORT"
	envDBUser     = "DB_USER"
	envDBPassword = "DB_PASSWORD"
)

var (
	client *mongo.Client
	Users  *mongo.Collection
)

func init() {
	loadEnvVariables()

	database := os.Getenv(envDatabase)
	host := os.Getenv(envDBHost)
	port := os.Getenv(envDBPort)
	user := os.Getenv(envDBUser)
	password := os.Getenv(envDBPassword)

	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, database)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("[*] Connected to MongoDB")

	Users = client.Database(database).Collection("users")
}

func loadEnvVariables() {
	if err := dotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	requiredEnvVars := []string{envDatabase, envDBHost, envDBPort, envDBUser, envDBPassword}
	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Environment variable %s is required but not set", envVar)
			os.Exit(1) // Salir con un código de error no cero
		}
	}
}
