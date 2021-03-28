package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type contact struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Question string `bson:"question,omitempty"`
}

// Adding contact questions into contact collection
func Contact(w http.ResponseWriter, r *http.Request) {
	c := new(contact)
	json.NewDecoder(r.Body).Decode(c)
	if err := send(c); err != nil {
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
