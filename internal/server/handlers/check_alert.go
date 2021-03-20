package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"github.com/ChristianStefaniw/cgr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAlert(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(cgr.GetParams(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	device, found, err := checkAlert(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if !found {
		w.Write([]byte("none"))
		return
	}

	b, err := json.Marshal(device)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(b)

}

func checkAlert(id primitive.ObjectID) (*database.SafetynetDevice, bool, error) {
	var device *database.SafetynetDevice
	doc, err := database.Database.FindDeviceById(constants.ALERT_COLL, context.Background(), id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, false, nil
		}
		return nil, false, err
	}

	bsonBytes, _ := bson.Marshal(doc)
	bson.Unmarshal(bsonBytes, &device)

	return device, true, err
}
