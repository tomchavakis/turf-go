package turf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestPoint(t *testing.T) {
	p1 := Point{Lng: -75.33, Lat: 39.44}
	p2 := Point{Lng: -75.33, Lat: 39.45}
	p3 := Point{Lng: -75.31, Lat: 39.46}
	p4 := Point{Lng: -75.30, Lat: 39.46}

	points := []Point{
		p1,
		p2,
		p3,
		p4,
	}

	refPoint := Point{Lat: 39.50, Lng: -75.33}

	r := NearestPoint(refPoint, points)

	assert.Equal(t, r, p3, "error computing the Nearest Point")
}
