package location

import (
	"math"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type LatLonPair struct {
	LatSrc  float64
	LonSrc  float64
	LatDest float64
	LonDest float64
}

func CheckInDistance(lat_lon *LatLonPair) bool {
	distance_between := distance_between_lat_lon(lat_lon)

	return distance_between <= constants.ALERT_RADIUS
}

// in km
func distance_between_lat_lon(lat_lon *LatLonPair) float64 {
	delta_lat := helpers.AsRadians(lat_lon.LatDest - lat_lon.LatSrc)
	delta_lon := helpers.AsRadians(lat_lon.LonDest - lat_lon.LonSrc)

	a := math.Sin(delta_lat/2)*math.Sin(delta_lat/2) +
		math.Cos(helpers.AsRadians(lat_lon.LatSrc))*
			math.Cos(helpers.AsRadians(lat_lon.LatDest))*
			math.Sin(delta_lon/2)*math.Sin(delta_lon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EARTH_RADIUS * c

	return d
}
