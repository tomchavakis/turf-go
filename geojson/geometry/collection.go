package geometry

import "github.com/tomchavakis/turf-go/geojson"

// Collection type
// https://tools.ietf.org/html/rfc7946#section-3.1.8
type Collection struct {
	Type       geojson.OBjectType `json:"type"`
	Geometries []Geometry         `json:"geometries"`
}

// NewGeometryCollection initializes a new instance of GeometryCollection
func NewGeometryCollection(geometries []Geometry) (*Collection, error) {
	return &Collection{Geometries: geometries, Type: geojson.GeometryCollection}, nil
}
