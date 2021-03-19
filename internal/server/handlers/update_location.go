package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var device database.SafetynetDevice
	json.NewDecoder(r.Body).Decode(&device)
	payload := bson.M{"$set": bson.M{"lat": device.Lat, "lon": device.Lon}}
	if err := database.Database.Update(constants.DEVICES_COLL, context.TODO(), device.Id, payload); err != nil {
		fmt.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Lat: %f, Lon: %f", device.Lat, device.Lon)))
}
