package location

import (
	"math"
	"safetynet/internal/constants"
	"safetynet/internal/helpers"
)

type coordPair struct {
	LatSrc  float64
	LonSrc  float64
	LatRecv float64
	LonRecv float64
}

// check if the distance between two coordinates is within [constants.ALERT_RADIUS]
func checkInDistance(coords *coordPair) bool {
	distance_between := distanceBetweenCoords(coords)

	return distance_between <= constants.ALERT_RADIUS
}

// find the distance between two coordinated (in km)
func distanceBetweenCoords(coords *coordPair) float64 {
	delta_lat := helpers.AsRadians(coords.LatRecv - coords.LatSrc)
	delta_lon := helpers.AsRadians(coords.LonRecv - coords.LonSrc)

	a := math.Sin(delta_lat/2)*math.Sin(delta_lat/2) +
		math.Cos(helpers.AsRadians(coords.LatSrc))*
			math.Cos(helpers.AsRadians(coords.LatRecv))*
			math.Sin(delta_lon/2)*math.Sin(delta_lon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EARTH_RADIUS * c

	return d
}
