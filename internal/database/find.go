package database

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) find_by_device_id(ctx context.Context, id uuid.UUID) (bson.M, error) {
	var result bson.M
	filter := bson.M{"deviceid": id}

	collection := db.safetynet.Collection("ids")

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
