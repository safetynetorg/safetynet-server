package helpers

import (
	"net/smtp"
	"os"
	"safetynet/internal/constants"
)

var auth smtp.Auth

func AuthEmail() {
	from := constants.SAFETYNET_EMAIL
	pass := os.Getenv("EMAIL_PASS")

	auth = smtp.PlainAuth("", from, pass, "smtp.gmail.com")
}
