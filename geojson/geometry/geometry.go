package geometry

import (
	"encoding/json"
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
)

// Geometry type
// https://tools.ietf.org/html/rfc7946#section-3
type Geometry struct {
	// GeoJSONType describes the type of GeoJSON Geometry, Feature or FeatureCollection this object is.
	GeoJSONType geojson.OBjectType `json:"type"`
	Coordinates interface{}        `json:"coordinates"`
}

// FromJSON returns a new Geometry by passing in a valid JSON string.
func FromJSON(gjson string) (*Geometry, error) {

	if gjson == "" {
		return nil, errors.New("input cannot be empty")
	}

	var geometry Geometry
	err := json.Unmarshal([]byte(gjson), &geometry)
	if err != nil {
		return nil, errors.New("cannot decode the input value")
	}

	return &geometry, nil

}

// ToPolygon convert the Geometry to Polygon
func (g *Geometry) ToPolygon() (*Polygon, error) {

	if g.GeoJSONType == geojson.Polygon {

		var coords = []LineString{}

		var polygonCoordinates [][][]float64
		ccc, err := json.Marshal(g.Coordinates)
		if err != nil {
			return nil, errors.New("cannot marshal object")
		}
		err = json.Unmarshal(ccc, &polygonCoordinates)
		if err != nil {
			return nil, errors.New("cannot unmarshal object")
		}

		for i := 0; i < len(polygonCoordinates); i++ {
			var posArray = []Point{}
			for j := 0; j < len(polygonCoordinates[i]); j++ {
				pos := Point{
					Lng: polygonCoordinates[i][j][0],
					Lat: polygonCoordinates[i][j][1],
				}
				posArray = append(posArray, pos)
			}
			ln := LineString{
				Coordinates: posArray,
			}
			coords = append(coords, ln)
		}
		poly, err := NewPolygon(coords)
		if err != nil {
			return nil, errors.New("cannot creat a new polygon")
		}
		return poly, nil
	}
	return nil, nil
}

// ToMultiPolygon converts a MultiPolygon Feature to MultiPolygon geometry.
func (g *Geometry) ToMultiPolygon() (*MultiPolygon, error) {
	if g.GeoJSONType != geojson.MultiPolygon {
		return nil, errors.New("the feature must be a multiPolygon")
	}
	var multiPolygonCoordinates [][][][]float64
	ccc, err := json.Marshal(g.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &multiPolygonCoordinates)

	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var polys = []Polygon{}
	for k := 0; k < len(multiPolygonCoordinates); k++ {
		var coords = []LineString{}
		for i := 0; i < len(multiPolygonCoordinates[k]); i++ {
			var posArray = []Point{}
			for j := 0; j < len(multiPolygonCoordinates[k][i]); j++ {
				pos := Point{
					Lng: multiPolygonCoordinates[k][i][j][0],
					Lat: multiPolygonCoordinates[k][i][j][1],
				}
				posArray = append(posArray, pos)
			}
			ln := LineString{
				Coordinates: posArray,
			}
			coords = append(coords, ln)
		}
		poly := Polygon{
			Coordinates: coords,
		}
		polys = append(polys, poly)
	}

	poly, err := NewMultiPolygon(polys)
	if err != nil {
		return nil, errors.New("cannot creat a new polygon")
	}
	return poly, nil
}
