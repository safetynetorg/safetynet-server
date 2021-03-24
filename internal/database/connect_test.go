package database

import (
	"safetynet/internal/keys"
	"testing"
)

func TestConnect(t *testing.T) {
	keys.Load()

	db := Connect()

	if db == nil {
		t.Fatal("Mongodb connection failed!")
	}
}
