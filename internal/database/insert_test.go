package database

import (
	"context"
	"safetynet/internal/pkg/helpers"
	"testing"

	"github.com/google/uuid"
)

func TestInsertSafetynet(t *testing.T) {
	helpers.LoadDotEnv()
	db := Connect()
	model := safetynet_device_id{Id: uuid.New()}
	err := db.insert("ids", context.Background(), model)
	if err != nil {
		t.Fatal("Could not insert into db:", err)
	}
}
