package feature

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// Feature defines a new feature type
// A Feature object represents a spatially bounded thing. Every object is a GeoJSON object no matter where it
// occurs in a GeoJSON text.
// https://tools.ietf.org/html/rfc7946#section-3.2
type Feature struct {
	// A Feature object has a "Type" member with the value "Feature".
	Type geojson.OBjectType `json:"type"`
	// A Feature object has a member with the name "properties". The
	// value of the properties member is an object (any JSON object or a
	// JSON null value).
	Properties map[string]interface{} `json:"properties"`
	// Bbox is the bounding box of the feature.
	Bbox []float64 `json:"bbox"`
	// A Feature object has a member with the name "Geometry".  The value
	// of the geometry member SHALL be either a Geometry object as
	// defined above or, in the case that the Feature is unlocated, a
	// JSON null value.
	Geometry geometry.Geometry `json:"geometry"`
}

// New initializes a new Feature
func New(geometry geometry.Geometry, bbox []float64, properties map[string]interface{}) (*Feature, error) {
	return &Feature{
		Geometry:   geometry,
		Properties: properties,
		Type:       geojson.Feature,
		Bbox:       bbox,
	}, nil
}

// FromJSON returns a new Feature by passing in a valid JSON string.
func FromJSON(gjson string) (*Feature, error) {

	if gjson == "" {
		return nil, errors.New("input cannot be empty")
	}

	var feature Feature
	err := json.Unmarshal([]byte(gjson), &feature)
	if err != nil {
		return nil, errors.New("cannot decode the input value")
	}

	return &feature, nil

}
