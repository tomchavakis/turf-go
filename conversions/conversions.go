package conversions

import (
	"errors"
	"math"

	"github.com/tomchavakis/turf-go/constants"
)

var factors = map[string]float64{
	constants.UnitMiles:         constants.EarthRadius / 1609.344,
	constants.UnitNauticalMiles: constants.EarthRadius / 1852.0,
	constants.UnitDegrees:       constants.EarthRadius / 111325.0,
	constants.UnitRadians:       1.0,
	constants.UnitInches:        constants.EarthRadius * 39.37,
	constants.UnitYards:         constants.EarthRadius / 1.0936,
	constants.UnitMeters:        constants.EarthRadius,
	constants.UnitCentimeters:   constants.EarthRadius * 100.0,
	constants.UnitKilometers:    constants.EarthRadius / 1000.0,
	constants.UnitFeet:          constants.EarthRadius * 3.28084,
	constants.UnitCentimetres:   constants.EarthRadius * 100.0,
	constants.UnitMetres:        constants.EarthRadius,
	constants.UnitKimometres:    constants.EarthRadius / 1000.0,
}

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

// LengthToDegrees convert a distance measurement (assuming a spherical Earth) from a real-world unit into degrees
// Valid units: miles, nauticalmiles, inches, yards, meters, metres, centimeters, kilometres, feet
func LengthToDegrees(distance float64, units string) (float64, error) {
	ltr, err := LengthToRadians(distance, units)
	if err != nil {
		return 0.0, err
	}

	return RadiansToDegrees(ltr), nil
}

// LengthToRadians convert a distance measurement (assuming a spherical Earth) from a real-world unit into radians.
func LengthToRadians(distance float64, units string) (float64, error) {
	if units == "" {
		units = constants.UnitDefault
	}
	if !validateUnit(units) {
		return 0.0, errors.New("invalid units")
	}

	return distance / factors[units], nil
}

// RadiansToLength convert a distance measurement (assuming a spherical Earth) from radians to a more friendly unit.
func RadiansToLength(radians float64, units string) (float64, error) {
	if units == "" {
		units = constants.UnitDefault
	}

	if !validateUnit(units) {
		return 0.0, errors.New("invalid unit")
	}

	return radians * factors[units], nil
}

// ConvertLength converts a distance to a different unit specified.
func ConvertLength(distance float64, originalUnits string, finalUnits string) (float64, error) {
	if finalUnits == "" {
		finalUnits = constants.UnitDefault
	}

	ltr, err := LengthToRadians(distance, originalUnits)

	if err != nil {
		return 0, err
	}
	return RadiansToLength(ltr, finalUnits)
}

func validateUnit(units string) bool {
	_, ok := factors[units]
	return ok
}
