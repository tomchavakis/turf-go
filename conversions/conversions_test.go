package conversions

import (
	"math"
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/constants"
)

func TestRadiansToDistance(t *testing.T) {
	rtl, err := RadiansToLength(1, constants.UnitRadians)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.0)

	rtl, err = RadiansToLength(1, constants.UnitKilometers)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, constants.EarthRadius/1000.0)

	rtl, err = RadiansToLength(1, constants.UnitMiles)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, constants.EarthRadius/1609.344)
}

func TestDistanceToRadians(t *testing.T) {
	rtl, err := LengthToRadians(1.0, constants.UnitRadians)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.0)

	rtl, err = LengthToRadians(constants.EarthRadius/1000, constants.UnitKilometers)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.0)

	rtl, err = LengthToRadians(constants.EarthRadius/1609.344, constants.UnitMiles)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.0)
}

func TestDistanceToDegrees(t *testing.T) {
	rtl, err := LengthToDegrees(1.0, constants.UnitRadians)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 57.29577951308232)

	rtl, err = LengthToDegrees(100.0, constants.UnitKilometers)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 0.899320363724538)

	rtl, err = LengthToDegrees(10.0, constants.UnitMiles)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 0.14473158314379025)
}

func TestConvertLength(t *testing.T) {
	rtl, err := ConvertLength(1000.0, constants.UnitMeters, "")
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.0)

	rtl, err = ConvertLength(1.0, constants.UnitKilometers, constants.UnitMiles)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 0.6213711922373341)

	rtl, err = ConvertLength(1.0, constants.UnitMiles, constants.UnitKilometers)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.6093439999999997)

	rtl, err = ConvertLength(1.0, constants.UnitNauticalMiles, "")
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 1.8519999999999999)

	rtl, err = ConvertLength(1.0, constants.UnitMeters, constants.UnitCentimeters)
	if err != nil {
		t.Errorf("RadiansToLength error: %v", err)
	}

	assert.Equal(t, rtl, 100.00000000000001)

}

func TestConvertArea(t *testing.T) {
	a, err := ConvertArea(1000.0, constants.UnitMetres, constants.UnitKilometers)
	assert.Equal(t, err, nil)
	assert.Equal(t, a, 0.001)

	b, err := ConvertArea(1, constants.UnitKilometers, constants.UnitMiles)
	assert.Equal(t, err, nil)
	assert.Equal(t, b, 0.386)

	c, err := ConvertArea(1, constants.UnitMiles, constants.UnitKilometers)
	assert.Equal(t, err, nil)
	assert.Equal(t, c, 2.5906735751295336)

	d, err := ConvertArea(1, constants.UnitMeters, constants.UnitCentimeters)
	assert.Equal(t, err, nil)
	assert.Equal(t, d, 10000.0)

	k, err := ConvertArea(1, constants.UnitMeters, constants.UnitCentimetres)
	assert.Equal(t, err, nil)
	assert.Equal(t, k, 10000.0)

	f, err := ConvertArea(100, constants.UnitMeters, constants.UnitAcres)
	assert.Equal(t, err, nil)
	assert.Equal(t, f, 0.0247105)

	g, err := ConvertArea(100, "", constants.UnitYards)
	assert.Equal(t, err, nil)
	assert.Equal(t, g, 119.59900459999999)

	h, err := ConvertArea(100, constants.UnitMeters, constants.UnitFeet)
	assert.Equal(t, err, nil)
	assert.Equal(t, h, 1076.3910417)

	i, err := ConvertArea(100000, constants.UnitFeet, "")
	assert.Equal(t, err, nil)
	assert.Equal(t, i, 0.009290303999749462)

	j, err := ConvertArea(1, constants.UnitMeters, constants.UnitHectares)
	assert.Equal(t, err, nil)
	assert.Equal(t, j, 0.0001)

	_, err = ConvertArea(-1, constants.UnitMeters, constants.UnitMillimeters)
	assert.Equal(t, err.Error(), "area must be a positive number")

	_, err = ConvertArea(1, "foo", "bar")
	assert.Equal(t, err.Error(), "invalid original units")

	_, err = ConvertArea(1, constants.UnitMeters, "bar")
	assert.Equal(t, err.Error(), "invalid finalUnits units")
}

func TestDegreesToRadians(t *testing.T) {
	r := DegreesToRadians(180)

	if r != math.Pi {
		t.Errorf("degrees to radians = %f; want %f", r, math.Pi)
	}
}

func TestRadiansToDegrees(t *testing.T) {
	r := RadiansToDegrees(math.Pi)

	if r != float64(180) {
		t.Error("error converting radians to degrees")
	}
}
