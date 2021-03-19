package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb
func Connect() *db {
	var (
		ctx       context.Context
		mongo_url string
	)

	ctx = context.Background()
	mongo_url = os.Getenv("MONGO_URL")

	client_options := options.Client().ApplyURI(mongo_url)

	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &db{safetynet: client.Database("safetynet")}
}
