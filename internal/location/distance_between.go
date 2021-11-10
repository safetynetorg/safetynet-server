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
func checkInDistance(coords coordPair) bool {
	distanceBetween := distanceBetweenCoords(coords)

	return distanceBetween <= constants.ALERT_RADIUS
}

// find the distance between two coordinated (in km)
func distanceBetweenCoords(coords coordPair) float64 {
	deltaLat := helpers.AsRadians(coords.LatRecv - coords.LatSrc)
	deltaLon := helpers.AsRadians(coords.LonRecv - coords.LonSrc)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(helpers.AsRadians(coords.LatSrc))*
			math.Cos(helpers.AsRadians(coords.LatRecv))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EARTH_RADIUS * c

	return d
}
