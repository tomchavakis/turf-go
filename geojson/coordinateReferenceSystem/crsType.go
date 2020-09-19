package coordinateReferenceSystem

// CRSType defines the GeoJSON CRS types as defines the geojson.org v1.0 spec
// http://geojson.org/geojson-spec.html#coordinate-reference-system-objects
type CRSType string

const(
	// Unspecified defines a CRS Type where the CRS cannot be assumed
	Unspecified CRSType = "unspecified"
	// Name defines the Named CRS type
	Name CRSType = "name"
	// Link defines the Linked CRS Type
	Link CRSType = "link"
)
