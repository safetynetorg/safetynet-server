package database

import (
	"context"
	"safetynet/internal/pkg/helpers"
	"testing"

	"github.com/google/uuid"
)

func TestFindDeviceId(t *testing.T) {
	helpers.LoadDotEnv()
	id := uuid.New()
	db := Connect()
	model := safetynet_device{Id: id}
	db.insert("ids", context.Background(), model)
	_, err := db.find_by_device_id(context.Background(), id)
	if err != nil {
		t.Fatal("Could not find document:", err)
	}
}
