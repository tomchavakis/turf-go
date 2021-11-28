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
