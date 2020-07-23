package turf

import (
	"math"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	r := DegreesToRadians(180)
	if r != math.Pi {
		t.Errorf("error converting degrees to radians")
	}
}

func TestRadiansToDegrees(t *testing.T) {
	r := RadiansToDegrees(math.Pi)
	if r != 180 {
		t.Errorf("error converting radians to degrees")
	}
}
