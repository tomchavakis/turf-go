package geometry

import (
	"encoding/json"
	"errors"
	"fmt"

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

// ToPoint converts the Geometry to Point
func (g *Geometry) ToPoint() (*Point, error) {
	if g.GeoJSONType == geojson.Point {

		var coords []float64
		ccc, err := json.Marshal(g.Coordinates)
		if err != nil {
			return nil, errors.New("cannot marshal object")
		}
		err = json.Unmarshal(ccc, &coords)
		if err != nil {
			return nil, errors.New("cannot unmarshal object")
		}
		var pos = Point{}
		pos.Lat = coords[1]
		pos.Lng = coords[0]

		return &pos, nil
	}
	return nil, errors.New("invalid geometry type")
}

// ToMultiPoint converts the Geometry to MultiPoint
func (g *Geometry) ToMultiPoint() (*MultiPoint, error) {
	if g.GeoJSONType == geojson.MultiPoint {
		var m MultiPoint
		var coords [][]float64
		ccc, err := json.Marshal(g.Coordinates)
		if err != nil {
			return nil, errors.New("cannot marshal object")
		}
		err = json.Unmarshal(ccc, &coords)
		if err != nil {
			return nil, errors.New("cannot unmarshal object")
		}
		for i := 0; i < len(coords); i++ {
			p := NewPoint(coords[i][1], coords[i][0])
			m.Coordinates = append(m.Coordinates, *p)
		}
		return &m, nil
	}

	return nil, errors.New("invalid geometry type")
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
			return nil, fmt.Errorf("cannot create a new polygon %v", err.Error())
		}
		return poly, nil
	}
	return nil, errors.New("invalid geometry type")
}

// ToMultiPolygon converts a MultiPolygon Feature to MultiPolygon geometry.
func (g *Geometry) ToMultiPolygon() (*MultiPolygon, error) {
	if g.GeoJSONType != geojson.MultiPolygon {
		return nil, errors.New("invalid geometry type")
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
		return nil, errors.New("cannot creat a new multipolygon")
	}
	return poly, nil
}

// ToLineString converts a ToLineString Geometry to ToLineString geometry.
func (g *Geometry) ToLineString() (*LineString, error) {
	if g.GeoJSONType != geojson.LineString {
		return nil, errors.New("invalid geometry type")
	}

	var coords [][]float64
	ccc, err := json.Marshal(g.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}
	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var coordinates []Point
	for _, coord := range coords {
		p := Point{
			Lat: coord[1],
			Lng: coord[0],
		}
		coordinates = append(coordinates, p)
	}

	lineString, err := NewLineString(coordinates)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new linestring %v", err.Error())
	}
	return lineString, nil
}

// ToMultiLineString converts a MultiLineString faeture to MultiLineString geometry.
func (g *Geometry) ToMultiLineString() (*MultiLineString, error) {
	if g.GeoJSONType != geojson.MiltiLineString {
		return nil, errors.New("invalid geometry type")
	}
	var coords [][][]float64
	ccc, err := json.Marshal(g.Coordinates)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	err = json.Unmarshal(ccc, &coords)
	if err != nil {
		return nil, errors.New("cannot marshal object")
	}

	var coordinates []LineString
	for i := 0; i < len(coords); i++ {
		var ls LineString
		var points []Point
		for j := 0; j < len(coords[i]); j++ {
			p := Point{
				Lat: coords[i][j][1],
				Lng: coords[i][j][0],
			}
			points = append(points, p)
		}
		ls.Coordinates = points
		coordinates = append(coordinates, ls)
	}

	ml, err := NewMultiLineString(coordinates)
	if err != nil {
		return nil, errors.New("can't create a new multiLineString")
	}
	return ml, nil
}
