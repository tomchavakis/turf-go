package geometry

import "errors"

// Polygon defines a polygon type
// https://tools.ietf.org/html/rfc7946#section-3.1.6
type Polygon struct {
	Coordinates []LineString
}

// NewPolygon initializes a new instance of a Polygon
// The coordinates of a polygon must be an array of linear ring coordinate arrays. For Polygons with more than one of these rings, the first MUST be
// the exterior ring, and any others MUST be interior rings. The exterior ring bounds the surface, and the interior ring bound holes within the surface.
func NewPolygon(coordinates []LineString) (*Polygon, error) {
	if len(coordinates) < 4 {
		return nil, errors.New("a polygon must have at least 4 positions")
	}

	for _, c := range coordinates {
		if !c.IsLinearRing() {
			return nil, errors.New("all elements of a polygon must be closed linestrings")
		}
	}

	return &Polygon{Coordinates: coordinates}, nil
}
