package feature

import (
	"errors"
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"
)

func TestNewGeometryCollection(t *testing.T) {
	const MercatorPoint = "../../test-data/mercator.point.geojson"
	const MercatorLineString = "../../test-data/mercator.linestring.geojson"

	p, err := utils.LoadJSONFixture(MercatorPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	lts, err := utils.LoadJSONFixture(MercatorLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	lns, err := FromJSON(lts)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	fts := []Feature{*f, *lns}

	c, err := NewFeatureCollection(fts)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, len(c.Features), 2)
	assert.Equal(t, string(c.Type), "FeatureCollection")

	assert.Equal(t, string(c.Features[0].Geometry.GeoJSONType), "Point")
	assert.Equal(t, string(c.Features[1].Geometry.GeoJSONType), "LineString")
}

func TestCollectionFromJSONEmptyGeoJSON(t *testing.T) {
	c, err := CollectionFromJSON("")
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("input cannot be empty"))
}

func TestCollectionFromJSONInvalidGeoJSON(t *testing.T) {
	c, err := CollectionFromJSON("{ \"TYPE\": \"GeometryColl, \"geometries\":[] }")
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("cannot decode the input value"))
}

func TestCollectionFromJSON(t *testing.T) {
	const FeatureCollection = "../../test-data/wgs84.featurecollection.geojson"
	fc, err := utils.LoadJSONFixture(FeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	collection, err := CollectionFromJSON(fc)
	assert.Nil(t, err)

	assert.Equal(t, string(collection.Type), "FeatureCollection")
	assert.Equal(t, len(collection.Features), 3)
	assert.Equal(t, string(collection.Features[0].Geometry.GeoJSONType), "Polygon")
	pt, err := collection.Features[0].ToPolygon()
	assert.Nil(t, err)
	assert.Equal(t, *pt, geometry.Polygon{
		Coordinates: []geometry.LineString{
			{
				Coordinates: []geometry.Point{
					{
						Lat: 58,
						Lng: -116,
					}, {
						Lat: 58,
						Lng: -90,
					}, {
						Lat: 66,
						Lng: -90,
					}, {
						Lat: 66,
						Lng: -116,
					},
					{
						Lat: 58,
						Lng: -116,
					},
				},
			},
		},
	})

	assert.Equal(t, string(collection.Features[1].Geometry.GeoJSONType), "LineString")
	lt, err := collection.Features[1].ToLineString()
	assert.Nil(t, err)
	assert.Equal(t, *lt, geometry.LineString{
		Coordinates: []geometry.Point{
			{
				Lat: 33.72434,
				Lng: -20.390625,
			},
			{
				Lat: 47.517201,
				Lng: -3.515625,
			},
			{
				Lat: 16.972741,
				Lng: 14.0625,
			},
		},
	})

	assert.Equal(t, string(collection.Features[2].Geometry.GeoJSONType), "Point")
	ptt, err := collection.Features[2].ToPoint()
	assert.Nil(t, err)
	assert.Equal(t, *ptt, geometry.Point{
		Lat: 3.864255,
		Lng: -76.289062,
	})
}
