package database

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/location"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mutex = &sync.Mutex{}

func (db *db) FindWhoAlert(ctx context.Context, src *SafetynetDevice) ([]*SafetynetDevice, error) {

	var devices []*SafetynetDevice
	var wg sync.WaitGroup

	collection := db.safetynet.Collection(constants.DEVICES_COLL)

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		wg.Add(1)

		go func(c mongo.Cursor) {
			defer wg.Done()

			var device SafetynetDevice

			if err := c.Decode(&device); err != nil || device.Id == src.Id {
				return
			}

			pair := &location.LatLonPair{
				LatSrc:  src.Lat,
				LonSrc:  src.Lon,
				LatDest: device.Lat,
				LonDest: device.Lon,
			}

			if location.CheckInDistance(pair) {
				mutex.Lock()
				devices = append(devices, &device)
				mutex.Unlock()
			}
		}(*cursor)
	}

	wg.Wait()
	return devices, nil
}
