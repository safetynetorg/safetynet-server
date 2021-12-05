package alert

import (
	"log"
	"os"

	"github.com/edganiukov/fcm"
)

var Client *fcm.Client

func InitClient() {
	client, err := fcm.NewClient(os.Getenv("SERVER_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
