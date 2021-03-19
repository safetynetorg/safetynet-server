package database

import (
	"context"
	"safetynet/internal/constants"
)

func (db *db) insert(ctx context.Context, docs ...interface{}) error {

	collection := db.safetynet.Collection(constants.DEVICES_COLL)

	if _, err := collection.InsertMany(ctx, docs); err != nil {
		return err
	}

	return nil
}
