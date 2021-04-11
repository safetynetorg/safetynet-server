package handlers

import (
	"context"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"github.com/ChristianStefaniw/cgr"
)

// delete a registered device
func DeleteDevice(w http.ResponseWriter, r *http.Request) {

	id := cgr.GetParams(r)["id"]

	if err := database.Database.Delete(constants.DEVICES_COLL, context.Background(), id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
