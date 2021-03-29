package database

import (
	"context"
	"safetynet/internal/constants"

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

func (db *db) SignupEmailExists(ctx context.Context, email string) (bool, error) {
	var result bson.M
	filter := bson.M{"email": email}
	collection := db.Safetynet.Collection("sign-up")

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err.Error() == constants.NO_DOC_FOUND {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
