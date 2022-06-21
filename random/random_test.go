package random

import (
	"testing"

	"github.com/tomchavakis/geojson"
	"github.com/tomchavakis/geojson/feature"
	"github.com/tomchavakis/turf-go"
	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/internal/common"
)

func TestRandomPosition(t *testing.T) {
	bbox := geojson.NewBBox(-10.0, -10.0, 10.0, 10.0)
	rndPos := Position(*bbox)
	assert.True(t, turf.InBBOX(rndPos.ToPoint(), *bbox), "point is not within the Bounding Box")
}

func TestRandomPoint0(t *testing.T) {
	fc, err := Point(0, *geojson.NewBBox(0, 0, 0, 0))
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 1)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Point")
	assert.Equal(t, fc.Features[0].Geometry.Coordinates, interface{}([]float64{0.0, 0.0}))
}

func TestRandomPoint(t *testing.T) {
	fc, err := Point(10, *geojson.NewBBox(0, 0, 0, 0))
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Point")
	assert.Equal(t, fc.Features[0].Geometry.Coordinates, interface{}([]float64{0.0, 0.0}))
}

func TestRandomPointInBBox(t *testing.T) {
	bbox := geojson.NewBBox(-10.0, -10.0, 10.0, 10.0)
	fc, err := Point(10, *bbox)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Point")
	p, err := fc.Features[0].Geometry.ToPoint()
	assert.Nil(t, err)
	assert.True(t, turf.InBBOX(*p, *bbox))
}

func TestRandomLineString(t *testing.T) {
	bbox := geojson.NewBBox(-10.0, -10.0, 10.0, 10.0)
	options := LineStringOptions{
		BBox:        *bbox,
		NumVertices: nil,
		MaxLength:   nil,
		MaxRotation: nil,
	}
	fc, err := LineString(10, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "LineString")
	checkFeaturesInBBox(t, bbox, fc.Features)
}

func TestRandomLineString_Vertex_1(t *testing.T) {
	bbox := geojson.NewBBox(-10.0, -10.0, 10.0, 10.0)
	options := LineStringOptions{
		BBox:        *bbox,
		NumVertices: common.IntPtr(1),
		MaxLength:   nil,
		MaxRotation: nil,
	}
	fc, err := LineString(10, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "LineString")
	checkFeaturesInBBox(t, bbox, fc.Features)
}

func TestRandomLineString_Count_0(t *testing.T) {
	bbox := geojson.NewBBox(-10.0, -10.0, 10.0, 10.0)
	options := LineStringOptions{
		BBox:        *bbox,
		NumVertices: common.IntPtr(1),
		MaxLength:   nil,
		MaxRotation: nil,
	}
	fc, err := LineString(0, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 1)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "LineString")
	checkFeaturesInBBox(t, bbox, fc.Features)
}

func TestRandomPolygon(t *testing.T) {
	bbox := geojson.NewBBox(-20.0, -20.0, 20.0, 20.0)
	options := PolygonOptions{
		BBox:            *bbox,
		NumVertices:     nil,
		MaxRadialLength: nil,
	}
	fc, err := Polygon(1, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 1)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Polygon")
}

func TestRandomPolygon_Features_ShouldBe_10(t *testing.T) {
	bbox := geojson.NewBBox(-20.0, -20.0, 20.0, 20.0)
	options := PolygonOptions{
		BBox:            *bbox,
		NumVertices:     nil,
		MaxRadialLength: nil,
	}
	fc, err := Polygon(10, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Polygon")
}

func TestRandomPolygon_Vertices_ShouldBe_NumVerticesPlusOne(t *testing.T) {
	bbox := geojson.NewBBox(-20.0, -20.0, 20.0, 20.0)
	options := PolygonOptions{
		BBox:            *bbox,
		NumVertices:     common.IntPtr(20),
		MaxRadialLength: nil,
	}
	fc, err := Polygon(10, options)
	assert.Nil(t, err)
	assert.Equal(t, string(fc.Type), "FeatureCollection")
	assert.NotNil(t, fc.Features)
	assert.Equal(t, len(fc.Features), 10)
	assert.Equal(t, string(fc.Features[0].Geometry.GeoJSONType), "Polygon")
	poly, err := fc.Features[0].Geometry.ToPolygon()
	assert.Nil(t, err)
	assert.Equal(t, len(poly.Coordinates[0].Coordinates), *options.NumVertices+1)
}

func checkFeaturesInBBox(t *testing.T, bbox *geojson.BBOX, fc []feature.Feature) {
	for i := 0; i < len(fc); i++ {
		ln, err := fc[i].Geometry.ToLineString()
		assert.Nil(t, err)
		for _, v := range ln.Coordinates {
			assert.True(t, turf.InBBOX(v, *bbox))
		}
	}
}
