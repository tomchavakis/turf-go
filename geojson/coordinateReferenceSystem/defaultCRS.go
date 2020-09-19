package coordinateReferenceSystem

// DefaultCRS CRS is a geographic coordinate reference system using the WGS84 datum.
// https://tools.ietf.org/html/rfc7946#section-4
// http://geojson.org/geojson-spec.html#coordinate-reference-system-objects
type DefaultCRS struct {

}

// Get the Instance of the DefaultCRS
func (d *DefaultCRS) Instance() string{
	return "urn:ogc:def:crs:OGC::CRS84"
}
