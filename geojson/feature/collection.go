package feature

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
)

// Collection represents a feature collection which holds a list of Fetures
type Collection struct {
	// A Feature object has a "Type" member with the value "Feature".
	Type     geojson.OBjectType `json:"type"`
	Features []Feature          `json:"features"`
}

// NewFeatureCollection initializes a new instance of Collection
func NewFeatureCollection(features []Feature) (*Collection, error) {
	return &Collection{Features: features}, nil
}

// CollectionFromJSON returns a new Feature by passing in a valid JSON string.
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
