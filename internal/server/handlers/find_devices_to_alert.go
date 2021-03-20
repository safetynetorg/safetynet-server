package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/database"
	"safetynet/internal/location"
	"strconv"
)

func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	devices_to_alert, err := location.FindDevicesToAlert(context.Background(), &device)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(strconv.Itoa(len(devices_to_alert))))
}
