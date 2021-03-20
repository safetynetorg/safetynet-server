package handlers

import (
	"encoding/json"
	"net/http"
	"safetynet/internal/alerts"

	"github.com/ChristianStefaniw/cgr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAlert(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(cgr.GetParams(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	device, found, err := alerts.CheckAlert(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	go alerts.StopAlertingDevice(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
