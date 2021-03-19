package database

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindDeviceId(t *testing.T) {
	helpers.LoadDotEnv()
	id := primitive.NewObjectID()
	db := Connect()
	model := SafetynetDevice{Id: id}
	db.Insert(constants.DEVICES_COLL, context.Background(), model)
	_, err := db.FindDeviceById(constants.DEVICES_COLL, context.Background(), id)
	if err != nil {
		t.Fatal("Could not find document:", err)
	}
}
