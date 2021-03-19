package database

import (
	"context"
	"safetynet/internal/constants"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) find_by_device_id(ctx context.Context, id primitive.ObjectID) (bson.M, error) {
	var result bson.M
	filter := bson.M{"deviceid": id}

	collection := db.safetynet.Collection(constants.DEVICES_COLL)

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
