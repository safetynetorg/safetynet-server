package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"
)

// Adding contact questions into contact collection
func Contact(w http.ResponseWriter, r *http.Request) {
	var contact database.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	fmt.Println(contact)
	if err := database.Database.Insert(constants.CONTACT_COLL, context.Background(), contact); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}