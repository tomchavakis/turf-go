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
