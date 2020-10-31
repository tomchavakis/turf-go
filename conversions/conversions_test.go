package conversions

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegreesToRadians(t *testing.T) {
	r := DegreesToRadians(180)
	assert.Equal(t, r, math.Pi, "error converting degrees to radians")
}

func TestRadiansToDegrees(t *testing.T) {
	r := RadiansToDegrees(math.Pi)
	assert.Equal(t, r, float64(180), "error converting radians to degrees")
}
