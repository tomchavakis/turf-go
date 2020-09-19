package geojson

// GeoJSONObjectType is the base class for all GeometryObjects
type GeoJSONObjectType string

const(
	// Defines a Point Type https://tools.ietf.org/html/rfc7946#section-3.1.2
	Point GeoJSONObjectType = "Point"
	// Defines a MultiPoint Type https://tools.ietf.org/html/rfc7946#section-3.1.3
	MultiPoint GeoJSONObjectType = "MultiPoint"
	// Defines a LineString Type https://tools.ietf.org/html/rfc7946#section-3.1.4
	LineString GeoJSONObjectType = "LineString"
	// Defines a MultiLineString Type https://tools.ietf.org/html/rfc7946#section-3.1.5
	MiltiLineString GeoJSONObjectType = "MultiLineString"
	// Defines a Polygon Type https://tools.ietf.org/html/rfc7946#section-3.1.6
	Polygon GeoJSONObjectType = "Polygon"
	// Defines a MultiPolygon Type https://tools.ietf.org/html/rfc7946#section-3.1.7
	MultiPolygon GeoJSONObjectType = "MultiPolygon"
	// Defines a GeometryCollection Type https://tools.ietf.org/html/rfc7946#section-3.1.8
	GeometryCollection GeoJSONObjectType = "GeometryCollection"
	// Defines a Feature Type https://tools.ietf.org/html/rfc7946#section-3.2
	Feature GeoJSONObjectType = "Feature"
	// Defines a FeatureCollection Type https://tools.ietf.org/html/rfc7946#section-3.3
	FeatureCollection GeoJSONObjectType = "FeatureCollection"
)
