package location

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mutex = &sync.Mutex{}

func FindDevicesToAlert(ctx context.Context, src *database.SafetynetDevice) ([]*database.SafetynetDevice, error) {

	var devices []*database.SafetynetDevice
	var wg sync.WaitGroup

	devicesColl := database.Database.Safetynet.Collection(constants.DEVICES_COLL)

	cursor, err := devicesColl.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		wg.Add(1)

		go func(c mongo.Cursor) {
			defer wg.Done()

			var device database.SafetynetDevice

			if err := c.Decode(&device); err != nil || device.Id == src.Id {
				return
			}

			pair := &latLonPair{
				LatSrc:  src.Lat,
				LonSrc:  src.Lon,
				LatDest: device.Lat,
				LonDest: device.Lon,
			}

			if checkInDistance(pair) {
				mutex.Lock()
				devices = append(devices, &device)
				mutex.Unlock()
				database.Database.Insert(constants.ALERT_IDS_COLL, ctx, database.AlertThisId{Id: device.Id})
			}
		}(*cursor)
	}

	wg.Wait()
	return devices, nil
}
