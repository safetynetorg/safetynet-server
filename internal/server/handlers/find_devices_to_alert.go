package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/database"
	"safetynet/internal/location"
)

func FindDevicesToAlert(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		w.Write([]byte("errorr"))
		return
	}
	devices_to_alert, err := location.FindDevicesToAlert(context.Background(), &device)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	fmt.Println(devices_to_alert)
	w.Write([]byte("ok"))
}
