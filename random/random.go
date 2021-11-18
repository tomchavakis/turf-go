package random

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// LineStringOptions object
type LineStringOptions struct {
	// BBox is the bounding box of which geometries are placed.
	BBox geojson.BBOX
	// NumVertices is how many coordinates each LineString will contain. 10 is the default value
	NumVertices *int
	// MaxLength os the maximum number of decimal degrees that a vertex can be from its predecessor. 0.0001 is the default value
	MaxLength *float64
	// MaxRotation is the maximum number of radians that a line segment can turn from the previous segment.  math.Pi / 8 is the default value
	MaxRotation *float64
}

// Position returns a random position within a bounding box
func Position(bbox geojson.BBOX) geometry.Position {
	pos := coordInBBox(bbox)
	res := geometry.NewPosition(nil, pos[0], pos[1])
	return *res
}

// Point returns a a GeoJSON FeatureCollection of random Points within a bounding box
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

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

// LineString returns a GeoJSON FeatureCollection of random LineStrings within a bounding box
func LineString(count int, options LineStringOptions) (*feature.Collection, error) {
	if count == 0 {
		count = 1
	}
	dv := 2
	if options.NumVertices == nil || *options.NumVertices < dv {
		v := 10
		options.NumVertices = &v
	}

	if options.MaxLength == nil {
		l := 0.0001
		options.MaxLength = &l
	}

	if options.MaxRotation == nil {
		r := math.Pi / 8.0
		options.MaxRotation = &r
	}

	fc := []feature.Feature{}

	for i := 0; i < count; i++ {

		startingPoint := Position(options.BBox)
		vertices := [][]float64{}
		vertices = append(vertices, []float64{startingPoint.ToPoint().Lng, startingPoint.ToPoint().Lat})

		for j := 0; j < *options.NumVertices-1; j++ {
			var priorAngle float64
			if j == 0 {
				priorAngle = rand.Float64() * 2 * math.Pi
			} else {
				priorAngle = math.Tan((vertices[j][1] - vertices[j-1][1]) / (vertices[j][0] - vertices[j-1][0]))
			}
			angle := priorAngle + (rand.Float64()-0.5)*(*options.MaxRotation)*2
			distance := rand.Float64() * (*options.MaxLength)
			vv := []float64{vertices[j][0] + distance*math.Cos(angle), vertices[j][1] + distance*math.Sin(angle)}
			vertices = append(vertices, vv)
		}

		g := geometry.Geometry{
			GeoJSONType: geojson.LineString,
			Coordinates: vertices,
		}
		f, err := feature.New(g, []float64{options.BBox.North, options.BBox.West, options.BBox.East, options.BBox.South}, nil, "")
		if err != nil {
			return nil, fmt.Errorf("cannot create a new Feature with error: %v", err)
		}
		fc = append(fc, *f)
	}

	if len(fc) > 0 {
		return feature.NewFeatureCollection(fc)
	}

	return nil, errors.New("can't generate a random LineString")

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
	// res[0] = rand.Float64()*(bbox.East-bbox.North) + bbox.North
	// res[1] = rand.Float64()*(bbox.West-bbox.South) + bbox.South

	res[0] = rand.Float64()*(bbox.East-bbox.South) + bbox.South
	res[1] = rand.Float64()*(bbox.West-bbox.North) + bbox.North

	return res
}
