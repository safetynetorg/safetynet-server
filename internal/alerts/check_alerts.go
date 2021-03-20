package alerts

import (
	"context"
	"safetynet/internal/constants"
	"safetynet/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAlert(id primitive.ObjectID) (*database.SafetynetDevice, bool, error) {
	var device *database.SafetynetDevice
	doc, err := database.Database.FindDeviceById(constants.ALERT_COLL, context.Background(), id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, false, nil
		}
		return nil, false, err
	}

	bsonBytes, _ := bson.Marshal(doc)
	bson.Unmarshal(bsonBytes, &device)

	return device, true, err
}
