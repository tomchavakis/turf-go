package classification

import (
	"math"

	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/measurement"
)

// NearestPoint takes a reference point and a list of points and returns the point from the point list closest to the reference.
func NearestPoint(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
	if len(points) == 0 {
		return &refPoint, nil
	}

	p := points[0]
	minDist := math.MaxFloat64

	for _, point := range points {
		dist, err := measurement.PointDistance(refPoint, point, units)
		if err != nil {
			return nil, err
		}
		if dist < minDist {
			p = point
			minDist = dist
		}
	}

	return &p, nil
}
