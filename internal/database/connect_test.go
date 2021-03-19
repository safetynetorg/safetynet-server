package database

import (
	"safetynet/internal/helpers"
	"testing"
)

func TestConnect(t *testing.T) {
	helpers.LoadDotEnv()

	db := Connect()

	if db == nil {
		t.Fatal("Mongodb connection failed!")
	}
}
