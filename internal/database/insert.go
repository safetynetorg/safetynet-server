package database

import (
	"context"
)

func (db *db) insert(coll string, ctx context.Context, docs ...interface{}) error {

	collection := db.safetynet.Collection(coll)

	if _, err := collection.InsertMany(ctx, docs); err != nil{
		return err
	}

	return nil
}
