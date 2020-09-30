package geometry

import (
	"github.com/tomchavakis/turf-go/geojson/object"
)

// Object is the base interface for GeometryObject types
type Object struct {
	Type object.Type
}
