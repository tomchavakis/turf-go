package feature

// FeatureCollection type
type FeatureCollection struct {
	geometries []Feature
}

// NewGeometryCollection initializes a new instance of FeatureCollection
func NewFeatureCollection(geometries []Feature) (*FeatureCollection, error) {
	return &FeatureCollection{},nil
}