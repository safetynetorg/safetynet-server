package location

import (
	"math"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type latLonPair struct {
	LatSrc  float64
	LonSrc  float64
	LatDest float64
	LonDest float64
}

func checkInDistance(lat_lon *latLonPair) bool {
	distance_between := distanceBetweenLatLon(lat_lon)

	return distance_between <= constants.ALERT_RADIUS
}

// in km
func distanceBetweenLatLon(lat_lon *latLonPair) float64 {
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
