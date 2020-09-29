package crs

// Type defines the GeoJSON CRS types as defines the geojson.org v1.0 spec
// http://geojson.org/geojson-spec.html#coordinate-reference-system-objects
type Type string

const (
	// Unspecified defines a CRS Type where the CRS cannot be assumed
	Unspecified Type = "unspecified"
	// Name defines the Named CRS type
	Name Type = "name"
	// Link defines the Linked CRS Type
	Link Type = "link"
)
