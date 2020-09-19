package geometry

import "errors"

// NewMultiPoint type
//https://tools.ietf.org/html/rfc7946#section-3.1.7
type MultiPolygon struct {
	coordinates []Polygon
}

// NewMultiPoint initializes a new MultiLineString
func NewMultiPolygon(coordinates []Polygon)  (*MultiPolygon, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("according to the GeoJSON v1.0 spec a MultiLineString must have at least two or more positions")
	}

	return &MultiPolygon{coordinates: coordinates}, nil
}
