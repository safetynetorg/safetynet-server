package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type contact struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Question string `json:"question,omitempty"`
}

// Adding contact questions into contact collection
func Contact(w http.ResponseWriter, r *http.Request) {
	c := new(contact)
	json.NewDecoder(r.Body).Decode(c)
	if err := send(c); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func send(c *contact) error {
	body := fmt.Sprintf("Email: %s\nName: %s\nQuestion: %s\n", c.Email, c.Name, c.Question)
	msg := "From: " + c.Email + "\n" +
		"To: " + constants.SAFETYNET_EMAIL + "\n" +
		"Subject: Question\n\n" +
		body

	err := helpers.SendEmail(msg, constants.SAFETYNET_EMAIL)

	return err
}
