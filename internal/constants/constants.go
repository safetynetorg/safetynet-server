package constants

import "time"

const (
	PORT            string        = "8080"
	EARTH_RADIUS    float64       = 6371
	ALERT_RADIUS    float64       = 0.2
	DEVICES_COLL    string        = "devices"
	ALERT_COLL      string        = "alert-ids"
	SIGN_UP_COLL    string        = "sign-up"
	CONTACT_COLL    string        = "contact"
	DATABASE        string        = "safetynet"
	ALERT_LIFE_TIME time.Duration = 30 * time.Second
	NO_DOC_FOUND    string        = "mongo: no documents in result"
	SAFETYNET_EMAIL string        = "help.safetynetorg@gmail.com"
	TOKEN_VERIFY    string        = "https://iid.googleapis.com/iid/info/"
)
