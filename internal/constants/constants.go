package constants

import "time"

const (
	PORT            = "8080"
	EARTH_RADIUS    = 6371
	ALERT_RADIUS    = 0.2
	DEVICES_COLL    = "devices"
	ALERT_COLL      = "alert-ids"
	SIGN_UP_COLL    = "sign-up"
	CONTACT_COLL    = "contact"
	DATABASE        = "safetynet"
	ALERT_LIFE_TIME = 30 * time.Second
	NO_DOC_FOUND    = "mongo: no documents in result"
)
