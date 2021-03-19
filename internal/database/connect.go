package database

import (
	"context"
	"log"
	"os"
	"safetynet/internal/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb
func Connect() *db {
	var (
		ctx      context.Context
		mongoUrl string
	)

	ctx = context.Background()
	mongoUrl = os.Getenv("MONGO_URL")

	client_options := options.Client().ApplyURI(mongoUrl)

	client, err := mongo.Connect(ctx, client_options)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	db := &db{Safetynet: client.Database(constants.DATABASE)}
	Database = db

	return db
}
