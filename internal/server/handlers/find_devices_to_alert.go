package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/database"
	"safetynet/internal/location"
	"strconv"
)

// called when someone needs to send an alert
func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// add devices within [constants.ALERT_RADIUS] to the alert collection
	// return the number of devices added to the alert collection
	devices_to_alert, err := location.FindDevicesToAlert(context.Background(), &device)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(strconv.Itoa(devices_to_alert)))
}
