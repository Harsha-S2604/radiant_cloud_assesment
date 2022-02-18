package db

import (
	"os"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func ConnectDB() *mongo.Database {
	username, password := os.Getenv("RADIANT_DB_USERNAME"), os.Getenv("RADIANT_DB_PASSWORD")
	if username == "" || password == "" {
		panic("username or password is required")
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, dbErr := mongo.Connect(ctx, clientOptions)
	if dbErr != nil {
		panic("Failed to connect to database " + dbErr.Error())
	}

	dbErr = client.Ping(ctx, nil)
	if dbErr != nil {
		panic("Failed to connect to database " + dbErr.Error())
	}

	databaseName := os.Getenv("RADIANT_DB")
	if databaseName == "" {
		panic("Please provide the database name")
	}

	database := client.Database(databaseName)

	return database
}