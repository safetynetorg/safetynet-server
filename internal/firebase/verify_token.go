package firebase

import (
	"context"
	"net/http"
	"os"
	"safetynet/internal/alert"
	"safetynet/internal/constants"
)

func VerifyToken(fcmToken string, ctx context.Context) bool {
	url := constants.TOKEN_VERIFY + fcmToken

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "key="+os.Getenv("SERVER_KEY"))

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode == 404 {
		return false
	}

	if err != nil || res.StatusCode != 200 {
		msg := "Something is wrong with your Safetynet installation. Please reinstall the app."
		alert.PushNotif(fcmToken, msg, alert.Client)
	}

	return true
}
