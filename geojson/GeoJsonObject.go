package geojson

import "github.com/tomchavakis/turf-go/geojson/crs"

// Object type
type Object struct {
	// BoundingBoxes The value of the bbox member MUST be an array of
	//   length 2*n where n is the number of dimensions represented in the
	//   contained geometries, with all axes of the most southwesterly point
	//   followed by all axes of the more northeasterly point.  The axes order
	//   of a bbox follows the axes order of geometries.
	// https://tools.ietf.org/html/rfc7946#section-5
	BBox []float64
	CRS  crs.Object
	Type ObjectType
}
