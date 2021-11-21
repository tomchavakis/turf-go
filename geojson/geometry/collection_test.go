package geometry

import (
	"errors"
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/utils"
)

func TestNewGeometryCollection(t *testing.T) {

	p, err := FromJSON("{ \"type\" : \"Point\", \"coordinates\": [-71, 41] }")
	assert.NotNil(t, p)
	assert.Nil(t, err)

	lns, err := FromJSON("{ \"type\" : \"LineString\", \"coordinates\":  [ [102, -10],[103, 1],[104, 0],[130, 4] ]}")
	assert.Nil(t, err)
	assert.NotNil(t, lns)

	geoms := []Geometry{*p, *lns}

	c, err := NewGeometryCollection(geoms)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, len(c.Geometries), 2)
	assert.Equal(t, string(c.Type), "GeometryCollection")

	assert.Equal(t, string(c.Geometries[0].GeoJSONType), "Point")
	assert.Equal(t, string(c.Geometries[1].GeoJSONType), "LineString")
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
	const WGS84GeometryCollection = "../../test-data/wgs84.geometrycollection.geojson"
	p, err := utils.LoadJSONFixture(WGS84GeometryCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	collection, err := CollectionFromJSON(p)
	assert.Nil(t, err)

	assert.Equal(t, string(collection.Type), "GeometryCollection")
	assert.Equal(t, len(collection.Geometries), 2)
	assert.Equal(t, string(collection.Geometries[0].GeoJSONType), "Point")
	pt, err := collection.Geometries[0].ToPoint()
	assert.Nil(t, err)
	assert.Equal(t, *pt, Point{
		Lat: 40.99999999999998,
		Lng: -71.0,
	})

	assert.Equal(t, string(collection.Geometries[1].GeoJSONType), "LineString")
	lst, err := collection.Geometries[1].ToLineString()
	assert.Nil(t, err)
	assert.Equal(t, *lst, LineString{
		Coordinates: []Point{
			{
				Lat: 33.72434000000235,
				Lng: -20.39062500000365,
			},
			{
				Lat: 47.51720099999992,
				Lng: -3.5156249999990803,
			},
			{
				Lat: 16.97274100000141,
				Lng: 14.062499999996321,
			},
		},
	})
}
