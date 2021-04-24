package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"
)

// register a new device
func NewDevice(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	json.NewDecoder(r.Body).Decode(&device)

	if err := database.Database.Insert(constants.DEVICES_COLL, context.Background(), device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(device.Id))
}
