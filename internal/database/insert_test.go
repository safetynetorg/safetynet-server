package database

import (
	"context"
	"fmt"
	"safetynet/internal/constants"
	"safetynet/internal/env"
	"testing"
)

func TestInsert(t *testing.T) {
	env.Load()
	db := Connect()
	id := "id"
	fmt.Println(id)
	//lat: 43.649632  lon: -79.483017
	model := SafetynetDevice{Id: id, Lat: 43.650761, Lon: -79.483131}
	err := db.Insert(constants.DEVICES_COLL, context.Background(), model)
	if err != nil {
		t.Fatal("Could not insert into db:", err)
	}
}
