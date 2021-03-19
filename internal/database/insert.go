package database

import (
	"context"
)

func (db *db) Insert(coll string, ctx context.Context, docs ...interface{}) error {

	collection := db.Safetynet.Collection(coll)

	if _, err := collection.InsertMany(ctx, docs); err != nil {
		return err
	}

	return nil
}
