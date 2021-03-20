package alerts

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StopAlertingDevice(id primitive.ObjectID) {
	const duration = constants.ALERTLIFETIME * time.Second

	deadline := time.Now().Add(duration)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	<-ctx.Done()
	database.Database.Delete(constants.ALERT_COLL, context.Background(), id)
}
