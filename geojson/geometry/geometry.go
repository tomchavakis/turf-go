package geometry

import (
	"github.com/tomchavakis/turf-go/geojson"
)

// Geometry type
// https://tools.ietf.org/html/rfc7946#section-3
type Geometry struct {
	// GeoJSONType describes the type of GeoJSON Geometry, Feature or FeatureCollection this object is.
	GeoJSONType geojson.OBjectType `json:"type"`
	Coordinates interface{}        `json:"coordinates"`
}
