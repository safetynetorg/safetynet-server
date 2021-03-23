package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/database"
	"safetynet/internal/location"
)

// called when someone needs to send an alert
func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// alert devices
	go location.FindDevicesToAlert(context.Background(), &device)

	w.WriteHeader(http.StatusAccepted)
}
