package geometry

// MultiPolygon defines the MultiPolygon type
// For type "MultiPolygon", the "coordinates" member is an array of Polygon coordinate arrays.
//https://tools.ietf.org/html/rfc7946#section-3.1.7
type MultiPolygon struct {
	Coordinates []Polygon
}

// NewMultiPolygon initialize a new MultiPolygon
func NewMultiPolygon(coordinates []Polygon) (*MultiPolygon, error) {

	return &MultiPolygon{Coordinates: coordinates}, nil
}
