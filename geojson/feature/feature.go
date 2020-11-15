package feature

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// Feature defines a new feature type
// A Feature object represents a spatially bounded thing. Every object is a GeoJSON object no matter where it
// occurs in a GeoJSON text.
// https://tools.ietf.org/html/rfc7946#section-3.2
type Feature struct {
	// A Feature object has a "Type" member with the value "Feature".
	Type geojson.OBjectType `json:"type"`
	// A Feature object has a member with the name "properties". The
	// value of the properties member is an object (any JSON object or a
	// JSON null value).
	Properties map[string]interface{} `json:"properties"`
	// Bbox is the bounding box of the feature.
	Bbox []float64 `json:"bbox"`
	// A Feature object has a member with the name "Geometry".  The value
	// of the geometry member SHALL be either a Geometry object as
	// defined above or, in the case that the Feature is unlocated, a
	// JSON null value.
	Geometry geometry.Geometry `json:"geometry"`
}

// New initializes a new Feature
func New(geometry geometry.Geometry, bbox []float64, properties map[string]interface{}) (*Feature, error) {
	return &Feature{
		Geometry:   geometry,
		Properties: properties,
		Type:       geojson.Feature,
		Bbox:       bbox,
	}, nil
}

// FromJSON returns a new Feature by passing in a valid JSON string.
func FromJSON(gjson string) (*Feature, error) {

	if gjson == "" {
		return nil, errors.New("input cannot be empty")
	}

	var feature Feature
	err := json.Unmarshal([]byte(gjson), &feature)
	if err != nil {
		return nil, errors.New("cannot decode the input value")
	}

	return &feature, nil

}

// ToPoint converts the Feature to Point.
func (f *Feature) ToPoint() (*geometry.Point, error) {
	if f.Geometry.GeoJSONType != geojson.Point {
		return nil, errors.New("the feature must be a point")
	}

	var coords []float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot unmarshal object")
	}
	var pos = geometry.Point{}
	pos.Lat = coords[1]
	pos.Lng = coords[0]

	return &pos, nil
}

// ToMultiPoint converts the Feature to MultiPoint type.
func (f *Feature) ToMultiPoint() (*geometry.MultiPoint, error) {
	if f.Geometry.GeoJSONType != geojson.MultiPoint {
		return nil, errors.New("the feature must be a MultiPoint")
	}

	var m geometry.MultiPoint
	var coords [][]float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot unmarshal object")
	}
	for i := 0; i < len(coords); i++ {
		p := geometry.NewPoint(coords[i][1], coords[i][0])
		m.Coordinates = append(m.Coordinates, *p)
	}
	return &m, nil
}

// ToPolygon converts a Polygon Feature to Polygon geometry.
func (f *Feature) ToPolygon() (*geometry.Polygon, error) {
	if f.Geometry.GeoJSONType != geojson.Polygon {
		return nil, errors.New("the feature must be a polygon")
	}
	var coords = []geometry.LineString{}

	var polygonCoordinates [][][]float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &polygonCoordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	for i := 0; i < len(polygonCoordinates); i++ {
		var posArray = []geometry.Point{}
		for j := 0; j < len(polygonCoordinates[i]); j++ {
			pos := geometry.Point{
				Lng: polygonCoordinates[i][j][0],
				Lat: polygonCoordinates[i][j][1],
			}
			posArray = append(posArray, pos)
		}
		ln := geometry.LineString{
			Coordinates: posArray,
		}
		coords = append(coords, ln)
	}
	poly, err := geometry.NewPolygon(coords)
	if err != nil {
		return nil, errors.New("cannot create a new polygon")
	}
	return poly, nil

}

// ToMultiPolygon converts a MultiPolygon Feature to MultiPolygon geometry.
func (f *Feature) ToMultiPolygon() (*geometry.MultiPolygon, error) {
	if f.Geometry.GeoJSONType != geojson.MultiPolygon {
		return nil, errors.New("the feature must be a multiPolygon")
	}
	var multiPolygonCoordinates [][][][]float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &multiPolygonCoordinates)

	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var polys = []geometry.Polygon{}
	for k := 0; k < len(multiPolygonCoordinates); k++ {
		var coords = []geometry.LineString{}
		for i := 0; i < len(multiPolygonCoordinates[k]); i++ {
			var posArray = []geometry.Point{}
			for j := 0; j < len(multiPolygonCoordinates[k][i]); j++ {
				pos := geometry.Point{
					Lng: multiPolygonCoordinates[k][i][j][0],
					Lat: multiPolygonCoordinates[k][i][j][1],
				}
				posArray = append(posArray, pos)
			}
			ln := geometry.LineString{
				Coordinates: posArray,
			}
			coords = append(coords, ln)
		}
		poly := geometry.Polygon{
			Coordinates: coords,
		}
		polys = append(polys, poly)
	}

	poly, err := geometry.NewMultiPolygon(polys)
	if err != nil {
		return nil, errors.New("cannot creat a new polygon")
	}
	return poly, nil
}

// ToLineString converts a ToLineString Feature to ToLineString geometry.
func (f *Feature) ToLineString() (*geometry.LineString, error) {
	if f.Geometry.GeoJSONType != geojson.LineString {
		return nil, errors.New("the feature must be a linestring")
	}

	var coords [][]float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var coordinates []geometry.Point
	for _, coord := range coords {
		p := geometry.Point{
			Lat: coord[1],
			Lng: coord[0],
		}
		coordinates = append(coordinates, p)
	}

	lineString, err := geometry.NewLineString(coordinates)
	if err != nil {
		return nil, errors.New("cannot creat a new polygon")
	}
	return lineString, nil
}

// ToMultiLineString converts a MultiLineString faeture to MultiLineString geometry.
func (f *Feature) ToMultiLineString() (*geometry.MultiLineString, error) {
	if f.Geometry.GeoJSONType != geojson.MiltiLineString {
		return nil, errors.New("the feature must be a multiLineString")
	}
	var coords [][][]float64
	ccc, err := json.Marshal(f.Geometry.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var coordinates []geometry.LineString
	for i := 0; i < len(coords); i++ {
		var ls geometry.LineString
		var points []geometry.Point
		for j := 0; j < len(coords[i]); j++ {
			p := geometry.Point{
				Lat: coords[i][j][1],
				Lng: coords[i][j][0],
			}
			points = append(points, p)
		}
		ls.Coordinates = points
		coordinates = append(coordinates, ls)
	}

	ml, err := geometry.NewMultiLineString(coordinates)
	if err != nil {
		return nil, errors.New("can't create a new multiLineString")
	}
	return ml, nil
}
