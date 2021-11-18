package random

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// Position returns a random position within a bounding box
func Position(bbox geojson.BBOX) geometry.Position {
	pos := coordInBBox(bbox)
	res := geometry.NewPosition(nil, pos[0], pos[1])
	return *res
}

// Point returns a random point within a bounding box
// count is how many geometries will be generated. default = 1
func Point(count int, bbox geojson.BBOX) (*feature.Collection, error) {
	if count == 0 {
		count = 1
	}
	fc := []feature.Feature{}

	for i := 0; i < count; i++ {
		p := coordInBBox(bbox)
		coords := []float64{p[0], p[1]}
		g := geometry.Geometry{
			GeoJSONType: geojson.Point,
			Coordinates: coords,
		}
		f, err := feature.New(g, []float64{bbox.North, bbox.West, bbox.East, bbox.South}, nil, "")
		if err != nil {
			return nil, fmt.Errorf("cannot create a new Feature with error: %v", err)
		}
		fc = append(fc, *f)
	}

	if len(fc) > 0 {
		return feature.NewFeatureCollection(fc)
	}

	return nil, errors.New("can't generate a random point")
}

// func rnd() float64 {
// 	return rand.Float64() - 0.5
// }

// func lon() float64 {
// 	return rnd() * 360
// }

// func lat() float64 {
// 	return rnd() * 180
// }

func coordInBBox(bbox geojson.BBOX) []float64 {
	res := make([]float64, 2)
	res[0] = rand.Float64()*(bbox.East-bbox.North) + bbox.North
	res[1] = rand.Float64()*(bbox.West-bbox.South) + bbox.South

	return res
}
