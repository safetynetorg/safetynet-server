package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewDevice(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	device.Id = primitive.NewObjectID()
	json.NewDecoder(r.Body).Decode(&device)

	if err := database.Database.Insert(constants.DEVICES_COLL, context.Background(), device); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write([]byte(device.Id.Hex()))
}
