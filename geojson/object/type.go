package object

// Type is the base class for all GeometryObjects
type Type string

const (
	// Point Defines a Point Type https://tools.ietf.org/html/rfc7946#section-3.1.2
	Point Type = "Point"
	// MultiPoint Defines a MultiPoint Type https://tools.ietf.org/html/rfc7946#section-3.1.3
	MultiPoint Type = "MultiPoint"
	// LineString Defines a LineString Type https://tools.ietf.org/html/rfc7946#section-3.1.4
	LineString Type = "LineString"
	// MiltiLineString Defines a MultiLineString Type https://tools.ietf.org/html/rfc7946#section-3.1.5
	MiltiLineString Type = "MultiLineString"
	// Polygon Defines a Polygon Type https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon Type = "Polygon"
	// MultiPolygon Defines a MultiPolygon Type https://tools.ietf.org/html/rfc7946#section-3.1.7
	MultiPolygon Type = "MultiPolygon"
	// GeometryCollection Defines a GeometryCollection Type https://tools.ietf.org/html/rfc7946#section-3.1.8
	GeometryCollection Type = "GeometryCollection"
	// Feature Defines a Feature Type https://tools.ietf.org/html/rfc7946#section-3.2
	Feature Type = "Feature"
	// FeatureCollection Defines a FeatureCollection Type https://tools.ietf.org/html/rfc7946#section-3.3
	FeatureCollection Type = "FeatureCollection"
)
