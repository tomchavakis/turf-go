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

var areaFactors = map[string]float64{
	constants.UnitAcres:       0.000247105,
	constants.UnitCentimeters: 10000.0,
	constants.UnitCentimetres: 10000.0,
	constants.UnitFeet:        10.763910417,
	constants.UnitHectares:    0.0001,
	constants.UnitInches:      1550.003100006,
	constants.UnitKilometers:  0.000001,
	constants.UnitKimometres:  0.000001,
	constants.UnitMeters:      1.0,
	constants.UnitMetres:      1.0,
	constants.UnitMiles:       3.86e-7,
	constants.UnitMillimeters: 1000000.0,
	constants.UnitMillimetres: 1000000.0,
	constants.UnitYards:       1.195990046,
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
	if units == "" {
		units = constants.UnitDefault
	}

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
	if originalUnits == "" {
		originalUnits = constants.UnitMeters
	}

	if finalUnits == "" {
		finalUnits = constants.UnitDefault
	}

	ltr, err := LengthToRadians(distance, originalUnits)

	if err != nil {
		return 0, err
	}
	return RadiansToLength(ltr, finalUnits)
}

// ConvertArea converts an area to the requested unit
func ConvertArea(area float64, originalUnits string, finalUnits string) (float64, error) {
	if originalUnits == "" {
		originalUnits = constants.UnitMeters
	}

	if finalUnits == "" {
		finalUnits = constants.UnitKilometers
	}
	if area < 0 {
		return 0.0, errors.New("area must be a positive number")
	}

	if !validateAreaUnit(originalUnits) {
		return 0.0, errors.New("invalid original units")
	}

	if !validateAreaUnit(finalUnits) {
		return 0.0, errors.New("invalid finalUnits units")
	}
	startFactor := areaFactors[originalUnits]
	finalFactor := areaFactors[finalUnits]
	return (area / startFactor) * finalFactor, nil
}

func validateAreaUnit(units string) bool {
	_, ok := areaFactors[units]
	return ok
}

func validateUnit(units string) bool {
	_, ok := factors[units]
	return ok
}
