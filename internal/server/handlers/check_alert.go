package handlers

import (
	"encoding/json"
	"net/http"
	"safetynet/internal/alerts"

	"github.com/ChristianStefaniw/cgr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// check if a device is in the alert collection
// the client who requested this endpoint will be sent a push notification if they are in the alert collection
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

	// device is in the alert collection

	// stop alerting the device after [constants.ALERTTIME]
	go alerts.StopAlertingDevice(id)

	// send a json response with the coordinates of the alerter
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
