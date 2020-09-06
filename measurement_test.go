package turf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	d := Distance(-77.03653, 38.89768, -77.05173, 38.8973)
	assert.Equal(t, d, 1.317556974720262, "error calculating the distance")
}

func TestDistancePoint(t *testing.T) {
	p1 := Point{Lng: -77.03653, Lat: 38.89768}
	p2 := Point{Lng: -77.05173, Lat: 38.8973}
	d := DistancePoint(p1, p2)
	assert.Equal(t, d, 1.317556974720262, "error calculating the distance")
}
