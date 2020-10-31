package geometry

import (
	"errors"
)

// LineString defines the linestring type.
type LineString struct {
	Coordinates []Point
}

// NewLineString initializes a new LineString
func NewLineString(coordinates []Point) (*LineString, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("according to the GeoJSON v1.0 spec a LineString must have at least two or more positions")
	}

	return &LineString{Coordinates: coordinates}, nil
}

// IsClosed determines if the Linestring is closed which means that has its first and last coordinate at the same position
func (l *LineString) IsClosed() bool {
	first := l.Coordinates[0]
	end := l.Coordinates[len(l.Coordinates)-1]

	return first.Lng == end.Lng &&
		first.Lat == end.Lat
}

// IsLinearRing returns true if it is a closed LineString with four or more positions
// https://tools.ietf.org/html/rfc7946#section-3.1.1
func (l *LineString) IsLinearRing() bool {
	return len(l.Coordinates) >= 4 && l.IsClosed()
}
