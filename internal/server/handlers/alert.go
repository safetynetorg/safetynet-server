package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Alert(w http.ResponseWriter, r *http.Request) {
	var id *database.AlertThisId

	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		w.Write([]byte("error"))
		return
	}

	found, err := checkAlert(id.Id)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte(strconv.FormatBool(found)))

}

func checkAlert(id primitive.ObjectID) (bool, error) {
	count, err := database.Database.Count(constants.ALERT_IDS_COLL, context.Background(), id)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
