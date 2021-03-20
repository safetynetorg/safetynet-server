package location

import (
	"math"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type latLonPair struct {
	LatSrc  float64
	LonSrc  float64
	LatRecv float64
	LonRecv float64
}

// check if the distance between two coordinates is within [constants.ALERT_RADIUS]
func checkInDistance(lat_lon *latLonPair) bool {
	distance_between := distanceBetweenLatLon(lat_lon)

	return distance_between <= constants.ALERT_RADIUS
}

// find the distance between two coordinated (in km)
func distanceBetweenLatLon(lat_lon *latLonPair) float64 {
	delta_lat := helpers.AsRadians(lat_lon.LatRecv - lat_lon.LatSrc)
	delta_lon := helpers.AsRadians(lat_lon.LonRecv - lat_lon.LonSrc)

	a := math.Sin(delta_lat/2)*math.Sin(delta_lat/2) +
		math.Cos(helpers.AsRadians(lat_lon.LatSrc))*
			math.Cos(helpers.AsRadians(lat_lon.LatRecv))*
			math.Sin(delta_lon/2)*math.Sin(delta_lon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EARTH_RADIUS * c

	return d
}
