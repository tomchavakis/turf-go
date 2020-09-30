package geometry

// Collection type
// https://tools.ietf.org/html/rfc7946#section-3.1.8
type Collection struct {
	geometries []Object
}

// NewGeometryCollection initializes a new instance of GeometryCollection
func NewGeometryCollection(geometries []Object) (*Collection, error) {
	return &Collection{geometries: geometries}, nil
}
