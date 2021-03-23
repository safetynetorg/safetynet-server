package location

import (
	"context"
	"fmt"
	"os"
	"safetynet/internal/alert"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"time"

	"github.com/edganiukov/fcm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// find devices to alert when someone is in dange
func FindDevicesToAlert(ctx context.Context, src *database.SafetynetDevice) {

	devicesColl := database.Database.Safetynet.Collection(constants.DEVICES_COLL)

	cursor, err := devicesColl.Find(ctx, bson.D{{}})
	if err != nil {
		return

	}
	defer cursor.Close(ctx)

	client, err := fcm.NewClient(os.Getenv("SERVER_KEY"))
	if err != nil {
		if client = retyConnect(2*time.Second, 2); client == nil {
			return
		}
	}

	for cursor.Next(ctx) {

		go func(c mongo.Cursor) {

			var device database.SafetynetDevice

			if err := c.Decode(&device); err != nil || device.Id == src.Id {
				return
			}

			pair := &coordPair{
				LatSrc:  src.Lat,
				LonSrc:  src.Lon,
				LatRecv: device.Lat,
				LonRecv: device.Lon,
			}

			// check if the receiver device is in range of the alert
			if checkInDistance(pair) {

				if err = alert.PushNotif(device.Id, fmt.Sprintf("Lat: %f, Lon: %f", pair.LatSrc, pair.LonSrc), client); err != nil {

					if err = helpers.Rety(func() error {

						return alert.PushNotif(device.Id, fmt.Sprintf("Lat: %f, Lon: %f", pair.LatSrc, pair.LonSrc), client)

					}, 1*time.Second, 2); err != nil {

						return

					}

				}

			}
		}(*cursor)
	}
}

func retyConnect(sleep time.Duration, attempts int) *fcm.Client {
	for i := 0; i < attempts; i++ {
		time.Sleep(sleep)
		client, err := fcm.NewClient(os.Getenv("SERVER_KEY"))
		if err != nil {
			return client
		}
	}
	return nil
}
