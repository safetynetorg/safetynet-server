package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := Connect()

	if db == nil {
		t.Fatal("Mongodb connection failed!")
	}
}