package helpers

import (
	"net/smtp"
	"safetynet/internal/constants"
	"safetynet/internal/keys"
)

func SendEmail(msg string, to string) error {
	from := constants.SAFETYNET_EMAIL
	pass := keys.EMAIL_PASS

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		from, []string{to}, []byte(msg))

	return err
}
