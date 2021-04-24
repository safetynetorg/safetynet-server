package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"safetynet/internal/database"
	"safetynet/internal/firebase"
	"safetynet/internal/location"
)

// called when someone needs to send an alert
func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !firebase.VerifyToken(device.Id, context.Background()) {
		http.Error(w, "Invalid device", 500)
		return
	}

	// alert devices
	devicesAlerted, err := location.FindDevicesToAlert(context.Background(), device)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(strconv.Itoa(devicesAlerted)))
}
