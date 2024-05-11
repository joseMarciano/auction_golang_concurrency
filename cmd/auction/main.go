package main

import (
	"auction_golang_concurrency/configuration/database/mongodb"
	"auction_golang_concurrency/configuration/logger"
	"context"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	_, err := mongodb.NewMongoDbConnection(context.TODO())
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	logger.Info("DEU BOA")
}
