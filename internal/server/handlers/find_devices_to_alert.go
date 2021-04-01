package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"safetynet/internal/database"
	"safetynet/internal/location"

	"go.mongodb.org/mongo-driver/bson"
)

// called when someone needs to send an alert
func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var (
		device    database.SafetynetDevice
		err       error
		bsonDoc   bson.M
		bsonBytes []byte
	)
	if err = json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if bsonDoc, err = database.Database.FindDeviceById("devices", context.Background(), device.Id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if bsonBytes, err = bson.Marshal(bsonDoc); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bson.Unmarshal(bsonBytes, &device)

	// alert devices
	devicesAlerted, err := location.FindDevicesToAlert(context.Background(), &device)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(strconv.Itoa(devicesAlerted)))
}
