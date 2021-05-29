package controllers

import (
	"context"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"github.com/ChristianStefaniw/cgr-v2"
)

// delete a registered device
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id := cgr.GetParam(r, "id")

	if err := database.Database.Delete(constants.DEVICES_COLL, context.Background(), id); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
