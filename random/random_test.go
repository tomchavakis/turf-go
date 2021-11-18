package random

import (
	"testing"

	"github.com/tomchavakis/turf-go"
	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson"
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
