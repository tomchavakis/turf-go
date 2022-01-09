package invariant

import (
	"errors"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"
)

// GetCoord unwrap a coordinate from a point feature, geometry or a single coordinate.
// coord is a GeoJSON point or an Array of numbers.
//
//  Examples:
//
// Point Example:
//
//  p := &geometry.Point{
//	   Lat: 34.6,
//	   Lng: 23.5,
//  }
//  coord, err := GetCoord(p)
//
//  Array Example
//
//  p := []float64{
// 	  23.5,
// 	  34.6,
//  }
//  coord, err := GetCoord(p)
//
//  Feature Example
//
//  p := &feature.Feature{
// 	   Type:       "Feature",
// 	   Properties: map[string]interface{}{},
// 	   Bbox:       []float64{},
// 	   Geometry: geometry.Geometry{
// 		   GeoJSONType: "Point",
// 		   Coordinates: []float64{44.34, 23.52},
// 	   },
//  }
//  coord, err := GetCoord(p)
//
func GetCoord(coord interface{}) ([]float64, error) {
	if coord == nil {
		return nil, errors.New("coord is required")
	}
	result := []float64{}
	switch gtp := coord.(type) {
	case *feature.Feature:
		if gtp.Type == geojson.Feature && gtp.Geometry.GeoJSONType == "Point" {
			p, err := gtp.Geometry.ToPoint()
			if err != nil {
				return nil, err
			}
			result = append(result, p.Lng, p.Lat)
			return result, nil
		}
	case *geometry.Point:
		result = append(result, gtp.Lng, gtp.Lat)
		return result, nil
	case []float64:
		if utils.IsArray(gtp) && len(gtp) >= 2 && !utils.IsArray(gtp[0]) && !utils.IsArray(gtp[1]) {
			return gtp, nil
		}
	}

	return nil, errors.New("coord must be GeoJSON Point or an Array of numbers")
}

// GetCoords unwrap coordinates from a Feature, Geometry, Object or an Array
//
// Example:
//
// coords: &geometry.Point{
//	  Lat: 23.52,
//	  Lng: 44.34,
//  },
//
// coords,err := GetCoords(coords)
// = []float64{44.34, 23.52}
//
func GetCoords(coords interface{}) (interface{}, error) {
	if coords == nil {
		return nil, errors.New("coord is required")
	}
	switch gtp := coords.(type) {
	case *feature.Feature: // Feature
		if gtp.Type == geojson.Feature {
			switch gtp.Geometry.GeoJSONType {
			case geojson.Point:
				_, err := gtp.Geometry.ToPoint()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			case geojson.MultiPoint:
				_, err := gtp.Geometry.ToMultiPoint()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			case geojson.LineString:
				_, err := gtp.Geometry.ToLineString()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			case geojson.Polygon:
				_, err := gtp.Geometry.ToPolygon()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			case geojson.MultiLineString:
				_, err := gtp.Geometry.ToMultiLineString()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			case geojson.MultiPolygon:
				_, err := gtp.Geometry.ToMultiPolygon()
				if err != nil {
					return nil, err
				}
				return gtp.Geometry.Coordinates, nil
			}
		}
	// Geometry
	case *geometry.Polygon:
		result := [][][]float64{}
		for i := 0; i < len(gtp.Coordinates); i++ {
			coords := [][]float64{}
			for j := 0; j < len(gtp.Coordinates[i].Coordinates); j++ {
				coords = append(coords, []float64{gtp.Coordinates[i].Coordinates[j].Lng, gtp.Coordinates[i].Coordinates[j].Lat})
			}
			result = append(result, coords)
		}
		return result, nil
	case *geometry.LineString:
		result := [][]float64{}
		for i := 0; i < len(gtp.Coordinates); i++ {
			result = append(result, []float64{gtp.Coordinates[i].Lng, gtp.Coordinates[i].Lat})
		}
		return result, nil
	case *geometry.MultiLineString:
		result := [][][]float64{}
		for i := 0; i < len(gtp.Coordinates); i++ {
			tmp := [][]float64{}
			for j := 0; j < len(gtp.Coordinates[i].Coordinates); j++ {
				tmp = append(tmp, []float64{gtp.Coordinates[i].Coordinates[j].Lng, gtp.Coordinates[i].Coordinates[j].Lat})
			}
			result = append(result, tmp)
		}
		return result, nil
	case *geometry.Point:
		result := []float64{}
		result = append(result, gtp.Lng, gtp.Lat)
		return result, nil
	case *geometry.MultiPoint:
		result := [][]float64{}
		for i := 0; i < len(gtp.Coordinates); i++ {
			tmp := []float64{}
			tmp = append(tmp, gtp.Coordinates[i].Lng, gtp.Coordinates[i].Lat)
			result = append(result, tmp)
		}
		return result, nil
	case *geometry.MultiPolygon:
		result := [][][][]float64{}
		for i := 0; i < len(gtp.Coordinates); i++ {
			tmp := [][][]float64{}
			for j := 0; j < len(gtp.Coordinates[i].Coordinates); j++ {
				tmPoly := [][]float64{}
				for k := 0; k < len(gtp.Coordinates[i].Coordinates[j].Coordinates); k++ {
					tmPoly = append(tmPoly, []float64{gtp.Coordinates[i].Coordinates[j].Coordinates[k].Lng, gtp.Coordinates[i].Coordinates[j].Coordinates[k].Lat})
				}
				tmp = append(tmp, tmPoly)
			}
			result = append(result, tmp)
		}
		return result, nil
	case []float64:
		if utils.IsArray(gtp) && len(gtp) >= 2 {
			return gtp, nil
		}
	}
	return nil, errors.New("coord must be GeoJSON Point or an Array of numbers")
}

// GetType returns the GeoJSON object's type
//
// Examples:
//
// fp, err := feature.FromJSON("{ \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Point\", \"coordinates\": [102, 0.5] } }")
// result := GetType(fp)
// ="Point"
//
func GetType(geojson interface{}) string {
	switch gtp := geojson.(type) {
	case *feature.Feature:
		return string(gtp.Geometry.GeoJSONType)
	case *feature.Collection:
		return string(gtp.Type)
	case *geometry.Collection:
		return string(gtp.Type)
	case *geometry.Geometry:
		return string(gtp.GeoJSONType)
	}
	return "invalid"
}

// GetGeom returns the Geometry from Feature or Geometry Object
func GetGeom(geojson interface{}) *geometry.Geometry {
	switch gtp := geojson.(type) {
	case *feature.Feature:
		return &gtp.Geometry
	case *geometry.Geometry:
		return gtp
	}
	return nil
}
