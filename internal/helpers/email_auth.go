package helpers

import (
	"net/smtp"
	"safetynet/internal/constants"
	"safetynet/internal/keys"
)

var auth smtp.Auth

func AuthEmail() {
	from := constants.SAFETYNET_EMAIL
	pass := keys.EMAIL_PASS

	auth = smtp.PlainAuth("", from, pass, "smtp.gmail.com")
}
