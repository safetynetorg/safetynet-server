package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) Update(coll string, ctx context.Context, id string, payload bson.M) error {
	collection := db.Safetynet.Collection(coll)

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, payload)
	if err != nil {
		return err
	}
	return nil
}
