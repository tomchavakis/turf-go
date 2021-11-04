package geometry

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
)

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

// CollectionFromJSON returns a new Collection by passing in a valid JSON string.
func CollectionFromJSON(gjson string) (*Collection, error) {

	if gjson == "" {
		return nil, errors.New("input cannot be empty")
	}

	var collection Collection
	err := json.Unmarshal([]byte(gjson), &collection)
	if err != nil {
		return nil, errors.New("cannot decode the input value")
	}

	return &collection, nil

}
