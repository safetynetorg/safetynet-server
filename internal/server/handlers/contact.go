package handlers

import (
	"encoding/json"
	"net/http"
	"safetynet/internal/database"
	"net/smtp"
	"log"
)

// Adding contact questions into contact collection
func Contact(w http.ResponseWriter, r *http.Request) {
	var contact database.Contact
	json.NewDecoder(r.Body).Decode(&contact)
	send("Email: " + contact.Email + "\nName: " + contact.Name + "\nQuestion: " + contact.Question)
	w.WriteHeader(http.StatusOK)
}

func send(body string) {
	from := "arcanederp@gmail.com"
	pass := "Gamingmaster4224"
	to := "help.safetynetorg@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Contact\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("sent")
}