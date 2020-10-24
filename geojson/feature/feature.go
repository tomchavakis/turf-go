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

// ToPolygon converts a Polygon Feature to Polygon geometry
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
		var posArray = []geometry.Position{}
		for j := 0; j < len(polygonCoordinates[i]); j++ {
			pos := geometry.Position{
				Altitude:  nil,
				Longitude: polygonCoordinates[i][j][0],
				Latitude:  polygonCoordinates[i][j][1],
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
		return nil, errors.New("cannot creat a new polygon")
	}
	return poly, nil

}

// ToMultiPolygon converts a MultiPolygon Feature to MultiPolygon geometry
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
			var posArray = []geometry.Position{}
			for j := 0; j < len(multiPolygonCoordinates[k][i]); j++ {
				pos := geometry.Position{
					Altitude:  nil,
					Longitude: multiPolygonCoordinates[k][i][j][0],
					Latitude:  multiPolygonCoordinates[k][i][j][1],
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

//TODO: Add ToPolyLine etc...
