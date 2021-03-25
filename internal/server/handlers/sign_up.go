package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"
)

// Adding email into sign up collection
func SignUp(w http.ResponseWriter, r *http.Request) {
	var email database.Email
	json.NewDecoder(r.Body).Decode(&email)
	fmt.Println(email)
	if err := database.Database.Insert(constants.SIGN_UP_COLL, context.Background(), email); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
