package database

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/keys"
	"testing"
)

func TestFindDeviceId(t *testing.T) {
	keys.Load()
	id := "id"
	db := Connect()
	model := SafetynetDevice{Id: id}
	db.Insert(constants.DEVICES_COLL, context.Background(), model)
	_, err := db.FindDeviceById(constants.DEVICES_COLL, context.Background(), id)
	if err != nil {
		t.Fatal("Could not find document:", err)
	}
}
