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
