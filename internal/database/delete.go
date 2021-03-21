package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) Delete(coll string, ctx context.Context, id string) error {
	collection := db.Safetynet.Collection(coll)

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
