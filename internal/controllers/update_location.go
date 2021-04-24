package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"go.mongodb.org/mongo-driver/bson"
)

// update a devices location
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	json.NewDecoder(r.Body).Decode(&device)

	payload := bson.M{"$set": bson.M{"lat": device.Lat, "lon": device.Lon}}
	if err := database.Database.Update(constants.DEVICES_COLL, context.Background(), device.Id, payload); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
