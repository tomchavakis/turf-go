package feature

// Collection type
type Collection struct {
	features []Feature
}

// NewFeatureCollection initializes a new instance of Collection
func NewFeatureCollection(geometries []Feature) (*Collection, error) {
	return &Collection{}, nil
}
