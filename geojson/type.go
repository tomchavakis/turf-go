package geojson

// OBjectType is the base class for all GeometryObjects
type OBjectType string

const (
	// Point Defines a Point Type https://tools.ietf.org/html/rfc7946#section-3.1.2
	Point OBjectType = "Point"
	// MultiPoint Defines a MultiPoint Type https://tools.ietf.org/html/rfc7946#section-3.1.3
	MultiPoint OBjectType = "MultiPoint"
	// LineString Defines a LineString Type https://tools.ietf.org/html/rfc7946#section-3.1.4
	LineString OBjectType = "LineString"
	// MultiLineString Defines a MultiLineString Type https://tools.ietf.org/html/rfc7946#section-3.1.5
	MultiLineString OBjectType = "MultiLineString"
	// Polygon Defines a Polygon Type https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon OBjectType = "Polygon"
	// MultiPolygon Defines a MultiPolygon Type https://tools.ietf.org/html/rfc7946#section-3.1.7
	MultiPolygon OBjectType = "MultiPolygon"
	// GeometryCollection Defines a GeometryCollection Type https://tools.ietf.org/html/rfc7946#section-3.1.8
	GeometryCollection OBjectType = "GeometryCollection"
	// Feature Defines a Feature Type https://tools.ietf.org/html/rfc7946#section-3.2
	Feature OBjectType = "Feature"
	// FeatureCollection Defines a FeatureCollection Type https://tools.ietf.org/html/rfc7946#section-3.3
	FeatureCollection OBjectType = "FeatureCollection"
)
