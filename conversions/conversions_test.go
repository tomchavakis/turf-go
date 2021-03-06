package conversions

import (
	"math"
	"testing"
)

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
