package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"github.com/ChristianStefaniw/cgr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice

	id, err := primitive.ObjectIDFromHex(cgr.GetParams(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	device.Id = id

	json.NewDecoder(r.Body).Decode(&device)

	if err := database.Database.Delete(constants.DEVICES_COLL, context.Background(), device.Id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
