package alert

import (
	"log"
	"safetynet/internal/keys"

	"github.com/edganiukov/fcm"
)

var Client *fcm.Client

func InitClient() {
	client, err := fcm.NewClient(keys.SERVER_KEY)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
