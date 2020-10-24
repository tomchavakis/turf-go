package geojson

// Object type
type Object struct {
	// BBox is the bounding box of the coordinate range of the object's Geometries, Features, or Feature Collections.
	// https://tools.ietf.org/html/rfc7946#section-5
	BBox *BBOX
	// GeoJSONType describes the type of GeoJSON Geometry, Feature or FeatureCollection this object is.
	GeoJSONType OBjectType
}
