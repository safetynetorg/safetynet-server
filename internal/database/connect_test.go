package database

import (
	"safetynet/internal/env"
	"testing"
)

func TestConnect(t *testing.T) {
	env.Load()

	db := Connect()

	if db == nil {
		t.Fatal("Mongodb connection failed!")
	}
}
