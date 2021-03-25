package geometry

import "github.com/tomchavakis/turf-go/geojson"

// Object is the base interface for GeometryObject types
type Object struct {
	Type geojson.OBjectType
}
