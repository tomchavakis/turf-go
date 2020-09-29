package crs

// Default CRS is a geographic coordinate reference system using the WGS84 datum.
// https://tools.ietf.org/html/rfc7946#section-4
// http://geojson.org/geojson-spec.html#coordinate-reference-system-objects
type Default struct {
}

// Instance returns the Instance of the Default CRS
func (d *Default) Instance() string {
	return "urn:ogc:def:crs:OGC::CRS84"
}
