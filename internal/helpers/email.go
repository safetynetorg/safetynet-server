package helpers

import (
	"net/smtp"
	"safetynet/internal/constants"
)

func SendEmail(msg, to string) error {
	from := constants.SAFETYNET_EMAIL

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		from, []string{to}, []byte(msg))

	return err
}
