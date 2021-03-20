package constants

import "time"

const (
	PORT          string = "8080"
	EARTH_RADIUS         = 6371
	ALERT_RADIUS         = 0.5
	DEVICES_COLL         = "devices"
	ALERT_COLL           = "alert-ids"
	DATABASE             = "safetynet"
	ALERTLIFETIME        = 30 * time.Second
)
