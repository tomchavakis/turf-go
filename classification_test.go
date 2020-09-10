package turf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestNearestPoint(t *testing.T) {

	p1 := geometry.Point{Lng: -75.33, Lat: 39.44}
	p2 := geometry.Point{Lng: -75.33, Lat: 39.45}
	p3 := geometry.Point{Lng: -75.31, Lat: 39.46}
	p4 := geometry.Point{Lng: -75.30, Lat: 39.46}

	var points []geometry.Point
	points = append(points, p1, p2, p3, p4)

	refPoint := geometry.Point{Lat: 39.50, Lng: -75.33}

	r := NearestPoint(refPoint, points)

	assert.Equal(t, r, p3, "error computing the Nearest Point")
}
