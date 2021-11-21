package geometry

import "errors"

// MultiPoint defines the MultiPoint type
//https://tools.ietf.org/html/rfc7946#section-3.1.3
type MultiPoint struct {
	Coordinates []Point
}

// NewMultiPoint initializes a new MultiLineString
func NewMultiPoint(coordinates []Point) (*MultiPoint, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("a MultiPoint must have at least two or more positions")
	}

	return &MultiPoint{Coordinates: coordinates}, nil
}
