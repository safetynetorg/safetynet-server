package location

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"safetynet/internal/alert"
	"safetynet/internal/constants"
	"safetynet/internal/database"
	"safetynet/internal/helpers"
	"strconv"
	"sync"
	"time"

	"github.com/edganiukov/fcm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type addressData struct {
	Data []addressLocation
}

type addressLocation struct {
	Distance      float64
	Name          string
	Neighbourhood string
}

// find devices to alert when someone is in dange
func FindDevicesToAlert(ctx context.Context, src database.SafetynetDevice) (int, error) {
	var wg sync.WaitGroup
	var alertedDevices int

	devicesColl := database.Database.Safetynet.Collection(constants.DEVICES_COLL)

	cursor, err := devicesColl.Find(ctx, bson.D{{}})
	if err != nil {
		return 0, err

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		wg.Add(1)
		go checkAndAlert(*cursor, &wg, src, &alertedDevices, alert.Client)
	}
	wg.Wait()
	return alertedDevices, nil
}

func checkAndAlert(c mongo.Cursor, wg *sync.WaitGroup, src database.SafetynetDevice, alertedDevices *int, client *fcm.Client) {
	defer wg.Done()
	var device database.SafetynetDevice

	if err := c.Decode(&device); err != nil || device.Id == src.Id {
		return
	}

	pair := coordPair{
		LatSrc:  src.Lat,
		LonSrc:  src.Lon,
		LatRecv: device.Lat,
		LonRecv: device.Lon,
	}

	// check if the receiver device is in range of the alert
	if checkInDistance(pair) {
		if err := alertDevice(device, pair, client); err == nil {
			*alertedDevices++
		}
	}
}

func alertDevice(device database.SafetynetDevice, pair coordPair, client *fcm.Client) error {
	address, err := getLocation(pair)
	var msg string

	if err != nil {
		msg = "Alert: \n* location not found *"
	} else {
		msg = "ALERT: \n" + address.Name + "\n" + address.Neighbourhood + "\n" + fmt.Sprint(address.Distance) + "m away"
	}

	if err := alert.PushNotif(device.Id, msg, client); err != nil {

		retry := func() error { return alert.PushNotif(device.Id, msg, client) }

		err = helpers.Retry(retry, 1*time.Second, 2)

		if err != nil {
			return err
		}

	}
	return nil
}

func getLocation(coords coordPair) (*addressLocation, error) {
	var address addressData
	baseURL, err := url.Parse("http://api.positionstack.com")
	if err != nil {
		return nil, err
	}

	baseURL.Path += "v1/reverse"

	params := url.Values{}

	params.Add("access_key", os.Getenv("GEO_KEY"))

	lon := strconv.FormatFloat(coords.LonSrc, 'E', -1, 64)

	lat := strconv.FormatFloat(coords.LatSrc, 'E', -1, 64)

	params.Add("query", fmt.Sprintf("%s,%s", lat, lon))

	params.Add("output", "json")

	params.Add("limit", "1")

	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&address); err != nil {
		return nil, err
	}

	return &address.Data[0], nil
}
