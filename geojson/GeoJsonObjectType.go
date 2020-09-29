package geojson

// ObjectType is the base class for all GeometryObjects
type ObjectType string

const (
	// Point Defines a Point Type https://tools.ietf.org/html/rfc7946#section-3.1.2
	Point ObjectType = "Point"
	// MultiPoint Defines a MultiPoint Type https://tools.ietf.org/html/rfc7946#section-3.1.3
	MultiPoint ObjectType = "MultiPoint"
	// LineString Defines a LineString Type https://tools.ietf.org/html/rfc7946#section-3.1.4
	LineString ObjectType = "LineString"
	// MiltiLineString Defines a MultiLineString Type https://tools.ietf.org/html/rfc7946#section-3.1.5
	MiltiLineString ObjectType = "MultiLineString"
	// Polygon Defines a Polygon Type https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon ObjectType = "Polygon"
	// MultiPolygon Defines a MultiPolygon Type https://tools.ietf.org/html/rfc7946#section-3.1.7
	MultiPolygon ObjectType = "MultiPolygon"
	// GeometryCollection Defines a GeometryCollection Type https://tools.ietf.org/html/rfc7946#section-3.1.8
	GeometryCollection ObjectType = "GeometryCollection"
	// Feature Defines a Feature Type https://tools.ietf.org/html/rfc7946#section-3.2
	Feature ObjectType = "Feature"
	// FeatureCollection Defines a FeatureCollection Type https://tools.ietf.org/html/rfc7946#section-3.3
	FeatureCollection ObjectType = "FeatureCollection"
)
