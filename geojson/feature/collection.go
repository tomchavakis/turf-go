package feature

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
)

// Collection represents a feature collection which holds a list of Fetures
type Collection struct {
	Type     geojson.OBjectType `json:"type"`
	Features []Feature          `json:"features"`
}

// NewFeatureCollection initializes a new instance of FeatureCollection
func NewFeatureCollection(features []Feature) (*Collection, error) {
	return &Collection{Features: features}, nil
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
