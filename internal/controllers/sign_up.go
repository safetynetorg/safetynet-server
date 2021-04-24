package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
)

// Adding email into sign up collection
func SignUp(w http.ResponseWriter, r *http.Request) {
	var email database.Email
	json.NewDecoder(r.Body).Decode(&email)

	exists, err := database.Database.SignupEmailExists(context.Background(), email.Email)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if exists {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := database.Database.Insert(constants.SIGN_UP_COLL, context.Background(), email); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := sendConfirm(email.Email); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func sendConfirm(email string) error {
	body := "Thank you for signing up for the Safetynet newsletter!"

	msg := "From: " + constants.SAFETYNET_EMAIL + "\n" +
		"To: " + email + "\n" +
		"Subject: Thank you!\n\n" +
		body

	err := helpers.SendEmail(msg, email)

	return err
}
