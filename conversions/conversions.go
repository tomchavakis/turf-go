package conversions

import (
	"math"
)

// DegreesToRadians converts an angle in degrees to radians.
// degrees angle between 0 and 360
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// RadiansToDegrees converts radians to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// ToKilometersPerHour converts knots to km/h
func ToKilometersPerHour(knots float64) float64 {
	return knots * 1.852
}
