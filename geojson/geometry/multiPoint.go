package geometry
import "errors"

// NewMultiPoint type
//https://tools.ietf.org/html/rfc7946#section-3.1.3
type MultiPoint struct {
	coordinates []Point
}

// NewMultiPoint initializes a new MultiLineString
func NewMultiPoint(coordinates []Point)  (*MultiPoint, error) {

	if len(coordinates) < 2 {
		return nil, errors.New("according to the GeoJSON v1.0 spec a MultiLineString must have at least two or more positions")
	}

	return &MultiPoint{coordinates: coordinates}, nil
}
