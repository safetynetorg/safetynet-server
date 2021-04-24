package database

import (
	"context"
	"log"
	"safetynet/internal/constants"
	"safetynet/internal/keys"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb
func Connect() *db {
	var (
		ctx      context.Context
		mongoURI string
	)

	ctx = context.Background()
	mongoURI = keys.MONGO_URI

	client_options := options.Client().ApplyURI(mongoURI)

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
