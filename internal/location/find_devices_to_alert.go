package location

import (
	"context"
	"fmt"
	"os"
	"safetynet/internal/alert"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"sync"

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

	client, _ := fcm.NewClient(os.Getenv("SERVER_KEY"))
	// TODO handle error

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

				// make sure that only one goroutine can mutate the [devices_alerted] variable at a time
				// prevents race condition
				mutex.Lock()
				devices_alerted++
				mutex.Unlock()

				_ = alert.PushNotif(device.Id, fmt.Sprintf("Lat: %f, Lon: %f", pair.LatSrc, pair.LonSrc), client)
				// TODO handle error

			}
		}(*cursor)
	}

	// wait for goroutines in waitgroup to complete
	wg.Wait()
	return devices_alerted, nil
}
