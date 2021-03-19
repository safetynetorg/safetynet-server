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
	db.insert(constants.DEVICES_COLL, context.Background(), model)
	_, err := db.find_by_device_id(context.Background(), id)
	if err != nil {
		t.Fatal("Could not find document:", err)
	}
}
