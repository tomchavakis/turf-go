package crs

// Type defines the GeoJSON CRS types as defines the geojson.org v1.0 spec
// http://geojson.org/geojson-spec.html#coordinate-reference-system-objects
type Type string

const (
	// UnspecifiedCRS defines a CRS Type where the CRS cannot be assumed
	UnspecifiedCRS Type = "unspecified"
	// NamedCRS defines the Named CRS type
	NamedCRS Type = "name"
	// LinkedCRS defines the Linked CRS Type
	LinkedCRS Type = "link"
)
