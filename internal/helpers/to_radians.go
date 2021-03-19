package helpers

import "math"

func AsRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
