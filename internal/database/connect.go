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

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(ctx, clientOptions)
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
