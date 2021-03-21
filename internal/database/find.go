package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) FindDeviceById(coll string, ctx context.Context, id string) (bson.M, error) {
	var result bson.M
	filter := bson.M{"_id": id}

	collection := db.Safetynet.Collection(coll)

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
