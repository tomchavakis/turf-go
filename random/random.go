package random

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/internal/common"
)

// LineStringOptions ...
type LineStringOptions struct {
	// BBox is the bounding box of which geometries are placed.
	BBox geojson.BBOX
	// NumVertices is how many coordinates each LineString will contain. 10 is the default value
	NumVertices *int
	// MaxLength is the maximum number of decimal degrees that a vertex can be from its predecessor. 0.0001 is the default value
	MaxLength *float64
	// MaxRotation is the maximum number of radians that a line segment can turn from the previous segment.  math.Pi / 8 is the default value
	MaxRotation *float64
}

// PolygonOptions ...
type PolygonOptions struct {
	// BBox is the bounding box of which geometries are placed.
	BBox geojson.BBOX
	// NumVertices is how many coordinates each Polygon will contain. 10 is the default value
	NumVertices *int
	// MaxRadialLength  is the maximum number of decimal degrees latitude or longitude that a vertex can reach out of the center of a Polygon.
	MaxRadialLength *float64
}

// Position returns a random position within a bounding box
//
// Examples:
// 	random.Position(
//		geojson.BBOX{
//			West:  -180,
//			South: -90,
//			East:  180,
//			North: 90,
//		})
func Position(bbox geojson.BBOX) geometry.Position {
	pos := coordInBBox(bbox)
	res := geometry.NewPosition(nil, pos[0], pos[1])
	return *res
}

// Point returns a GeoJSON FeatureCollection of random Point within a bounding box.
//
// count is how many geometries will be generated. default = 1
//
// Examples:
// 	random.Point(0,
//		geojson.BBOX{
//			West:  -180,
//			South: -90,
//			East:  180,
//			North: 90,
//		})
//
// 	random.Point(10,
//		geojson.BBOX{
//			West:  -180,
//			South: -90,
//			East:  180,
//			North: 90,
//		})
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

// LineString returns a GeoJSON FeatureCollection of random LineString within a bounding box.
//
// count=1 how many geometries will be generated
//
// Examples:
//   random.LineString(10,LineStringOptions{
//		BBox:geojson.BBOX{
//			West:  -180,
//			South: -90,
//			East:  180,
//			North: 90,
//		},
//		NumVertices: nil,
//		MaxLength:   nil,
//		MaxRotation: nil,
//	})
func LineString(count int, options LineStringOptions) (*feature.Collection, error) {
	if count == 0 {
		count = 1
	}

	if options.NumVertices == nil || *options.NumVertices < 2 {
		options.NumVertices = common.IntPtr(10)
	}

	if options.MaxLength == nil {
		options.MaxLength = common.Float64Ptr(0.0001)
	}

	if options.MaxRotation == nil {
		options.MaxRotation = common.Float64Ptr(math.Pi / 8.0)
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

// Polygon returns a GeoJSON FeatureCollection of random Polygon.
//
// count=1 how many geometries will be generated
//
// Examples:
// 	random.Polygon(10, PolygonOptions{
//		BBox:geojson.BBOX{
//			West:  -180,
//			South: -90,
//			East:  180,
//			North: 90,
//		},
//		NumVertices:     nil,
//		MaxRadialLength: nil,
//	})
func Polygon(count int, options PolygonOptions) (*feature.Collection, error) {

	if count == 0 {
		count = 1
	}

	if options.NumVertices == nil {
		options.NumVertices = common.IntPtr(10)
	}

	if options.MaxRadialLength == nil {
		options.MaxRadialLength = common.Float64Ptr(10.0)
	}

	fc := []feature.Feature{}

	for i := 0; i < count; i++ {

		vts := [][][]float64{}
		vertices := [][]float64{}
		circleOffsets := make([]float64, *options.NumVertices+1)

		// sum Offsets
		for i := 0; i < len(circleOffsets); i++ {
			if i == 0 {
				circleOffsets[i] = rand.Float64()
			} else {
				circleOffsets[i] = rand.Float64() + circleOffsets[i-1]
			}
		}

		// scale offsets
		for j := 0; j < len(circleOffsets); j++ {
			cur := (circleOffsets[j] * 2 * math.Pi) / circleOffsets[len(circleOffsets)-1]
			radialScaler := rand.Float64()
			vertices = append(vertices, []float64{radialScaler * (*options.MaxRadialLength) * math.Sin(cur), radialScaler * (*options.MaxRadialLength) * math.Cos(cur)})
		}

		// close the ring
		vertices[len(vertices)-1] = vertices[0]

		// center the polygon around something
		res := vertexToCoordinate(vertices, options.BBox)
		vts = append(vts, res)

		g := geometry.Geometry{
			GeoJSONType: geojson.Polygon,
			Coordinates: vts,
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

func vertexToCoordinate(vtc [][]float64, bbox geojson.BBOX) [][]float64 {
	res := [][]float64{}
	p := Position(bbox)
	for i := 0; i < len(vtc); i++ {
		tmp := []float64{}
		tmp = append(tmp, vtc[i][0]+p.Longitude, vtc[i][1]+p.Latitude)
		res = append(res, tmp)
	}
	return res
}

func coordInBBox(bbox geojson.BBOX) []float64 {
	res := make([]float64, 2)
	res[0] = rand.Float64()*(bbox.East-bbox.South) + bbox.South
	res[1] = rand.Float64()*(bbox.West-bbox.North) + bbox.North

	return res
}
