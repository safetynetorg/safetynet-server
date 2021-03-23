package location

import (
	"context"
	"fmt"
	"os"
	"safetynet/internal/alert"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"sync"
	"time"

	"github.com/edganiukov/fcm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mutex = &sync.Mutex{}

// find devices to alert when someone is in dange
func FindDevicesToAlert(ctx context.Context, src *database.SafetynetDevice) (int, error) {

	var devices_alerted int
	var wg sync.WaitGroup

	devicesColl := database.Database.Safetynet.Collection(constants.DEVICES_COLL)

	cursor, err := devicesColl.Find(ctx, bson.D{{}})
	if err != nil {
		return 0, err

	}
	defer cursor.Close(ctx)

	client, err := fcm.NewClient(os.Getenv("SERVER_KEY"))
	if err != nil {
		if client = retyConnect(2*time.Second, 2); client == nil {
			return 0, err
		}
	}

	for cursor.Next(ctx) {

		wg.Add(1)

		go func(c mongo.Cursor) {
			defer wg.Done()

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

				// make sure that only one goroutine can mutate the [devices_alerted] variable at a time
				// prevents race condition
				mutex.Lock()
				devices_alerted++
				mutex.Unlock()
			}
		}(*cursor)
	}

	// wait for goroutines in waitgroup to complete
	wg.Wait()
	return devices_alerted, nil
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
