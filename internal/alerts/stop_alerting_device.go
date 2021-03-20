package alerts

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// stop alerting a device after [constants.ALERT_LIFE_TIME] of time
func StopAlertingDevice(id primitive.ObjectID) {
	const duration = constants.ALERT_LIFE_TIME

	deadline := time.Now().Add(duration)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// wait for the context to be finished
	<-ctx.Done()
	database.Database.Delete(constants.ALERT_COLL, context.Background(), id)
}
