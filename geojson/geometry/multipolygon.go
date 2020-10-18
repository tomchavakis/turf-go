package geometry

// MultiPolygon defines the MultiPolygon type
//https://tools.ietf.org/html/rfc7946#section-3.1.7
type MultiPolygon struct {
	Coordinates []Polygon
}

// NewMultiPolygon initializes a new MultiLineString
func NewMultiPolygon(coordinates []Polygon) (*MultiPolygon, error) {

	// if len(coordinates) < 2 {
	// 	return nil, errors.New("according to the GeoJSON v1.0 spec a MultiPolygon must have at least two or more positions")
	// }

	return &MultiPolygon{Coordinates: coordinates}, nil
}
