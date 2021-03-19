package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *db) Count(coll string, ctx context.Context, id primitive.ObjectID) (int64, error) {

	collection := db.Safetynet.Collection(coll)

	count, err := collection.CountDocuments(ctx, bson.M{"_id": id})

	if err != nil {
		return 0, err
	}
	return count, nil
}
