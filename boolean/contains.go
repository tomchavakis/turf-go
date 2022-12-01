package boolean

import (
    "github.com/tomchavakis/geojson/feature"
    "github.com/tomchavakis/geojson/geometry"
)


type ContainTypes interface {
    feature.Feature | geometry.Geometry
}

// Boolean-contains returns true if the second geometry is completely contained by the first geometry.
// The interiors of both geometries must intersect and, the interior and boundary of the secondary (geometry b)
// must not intersect the exterior of the primary (geometry a).
// Boolean-contains returns the exact opposite result if the boolean-within
func Contains[T ContainTypes](f1 T, f2 T){
    // requirement: invariant -- getGeom    
}

