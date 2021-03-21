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
	time.Sleep(constants.ALERT_LIFE_TIME)

	database.Database.Delete(constants.ALERT_COLL, context.Background(), id)
}
