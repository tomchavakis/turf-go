package geometry


// GeometryCollection type
// https://tools.ietf.org/html/rfc7946#section-3.1.8
type GeometryCollection struct {
	geometries []GeometryObject
}

// NewGeometryCollection initializes a new instance of GeometryCollection
func NewGeometryCollection(geometries []GeometryObject) (*GeometryCollection, error) {
	return &GeometryCollection{},nil
}