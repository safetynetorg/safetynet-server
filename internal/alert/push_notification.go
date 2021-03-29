package alert

import (
	"github.com/edganiukov/fcm"
)

func PushNotif(token, data string, client *fcm.Client) error {
	notif := &fcm.Notification{
		Body: data,
	}

	msg := &fcm.Message{
		Token:        token,
		Notification: notif,
	}

	if _, err := client.Send(msg); err != nil {
		return err
	}

	return nil
}
