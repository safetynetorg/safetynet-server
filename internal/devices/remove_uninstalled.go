package devices

import (
	"context"
	"sync"
	"time"

	"safetynet/internal/constants"
	"safetynet/internal/database"
	"safetynet/internal/firebase"
	"safetynet/internal/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RemoveUninstalledDevices() {
	ticker := time.NewTicker(30 * time.Minute)
	ctx := context.Background()
	var wg sync.WaitGroup

	for range ticker.C {
		devicesColl := database.Database.Safetynet.Collection(constants.DEVICES_COLL)

		cursor, err := devicesColl.Find(ctx, bson.D{{}})
		if err != nil {
			continue
		}

		for cursor.Next(ctx) {
			wg.Add(1)
			go checkAndRemoveDevice(*cursor, ctx, &wg)
		}

		wg.Wait()
		cursor.Close(ctx)
	}
}

func checkAndRemoveDevice(c mongo.Cursor, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var device database.SafetynetDevice

	if err := c.Decode(&device); err != nil {
		return
	}

	if firebase.VerifyToken(device.Id, ctx) {
		return
	}

	if err := database.Database.Delete(constants.DEVICES_COLL, ctx, device.Id); err != nil {
		if err = helpers.Retry(
			func() error {
				return database.Database.Delete(constants.DEVICES_COLL, ctx, device.Id)
			}, 2*time.Second, 2); err != nil {
			return
		}
	}
}
