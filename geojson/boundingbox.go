package geojson

// BBOX defines a Bounding Box
// A GeoJSON object MAY have a member named "bbox" to include information on the coordinate range for its Geometries, Features, or FeatureCollections.
// The value of the bbox member MUST be an array of length 2*n where n is the number of dimensions represented in the contained geometries,
// with all axes of the most southwesterly point followed by all axes of the more northeasterly point.
// The axes order of a bbox follows the axes order of geometries.
// https://tools.ietf.org/html/rfc7946#section-5
type BBOX struct {
	West  float64
	South float64
	East  float64
	North float64
}

// NewBBox initializes a new Bounding Box
func NewBBox(west float64, south float64, east float64, north float64) *BBOX {
	return &BBOX{
		West:  west,
		South: south,
		East:  east,
		North: north,
	}
}
