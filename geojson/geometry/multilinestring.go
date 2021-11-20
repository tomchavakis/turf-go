package geometry

import "errors"

// MultiLineString type
//https://tools.ietf.org/html/rfc7946#section-3.1.5
type MultiLineString struct {
	Coordinates []LineString
}

// NewMultiLineString initializes a new MultiLineString
func NewMultiLineString(coordinates []LineString) (*MultiLineString, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("a MultiLineString must have at least two or more linestrings")
	}

	return &MultiLineString{Coordinates: coordinates}, nil
}
