package geometry

import "errors"

// MultiLineString type
//https://tools.ietf.org/html/rfc7946#section-3.1.5
type MultiLineString struct {
	coordinates []LineString
}

// NewMultiLineString initializes a new MultiLineString
func NewMultiLineString(coordinates []LineString) (*MultiLineString, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("according to the GeoJSON v1.0 spec a MultiLineString must have at least two or more positions")
	}

	return &MultiLineString{coordinates: coordinates}, nil
}
