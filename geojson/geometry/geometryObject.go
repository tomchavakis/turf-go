package geometry

import "github.com/tomchavakis/turf-go/geojson"

// GeometryObject is the base interface for GeometryObject types
type GeometryObject struct {
	Type geojson.GeoJSONObjectType
}
