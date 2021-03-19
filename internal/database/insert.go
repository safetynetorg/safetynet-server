package database

import (
	"context"
)

func (db *db) insert(coll string, ctx context.Context, docs ...interface{}) error {

	var err error

	collection := db.safetynet.Collection(coll)

	for _, doc := range docs {
		_, err = collection.InsertOne(ctx, doc)
		if err != nil {
			return err
		}
	}
	return nil
}
