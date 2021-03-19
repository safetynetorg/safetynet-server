package database

import "go.mongodb.org/mongo-driver/mongo"

type db struct {
	safetynet *mongo.Database
}
