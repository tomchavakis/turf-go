package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestCoordAllPont(t *testing.T) {
	json := "{ \"type\": \"Point\", \"coordinates\": [23.0, 54.0]}"
	geometry, err := geometry.FromJSON(json)
	assert.Nil(t, err, "geometry can't be nil")

	pt, err := geometry.ToPoint()
	assert.Nil(t, err, "convert to Point error")

	pts, err := CoordAll(pt, nil)
	assert.Nil(t, err, "coord all err")

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 54.0)
	assert.Equal(t, pts[0].Lng, 23.0)
}

func TestCoordAllLineString(t *testing.T) {
	json := "{ \"type\": \"LineString\", \"coordinates\": [[0.0, 0.0], [1.0, 1.0]]}"
	geometry, err := geometry.FromJSON(json)
	assert.Nil(t, err, "geometry can't be nil")

	ln, err := geometry.ToLineString()
	assert.Nil(t, err, "convert to Linestring error")

	pts, err := CoordAll(ln, nil)
	assert.Nil(t, err, "coord all err")
	assert.Equal(t, len(pts), 2)
	assert.Equal(t, pts[0].Lat, 0.0)
	assert.Equal(t, pts[0].Lng, 0.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 1.0)
}

func TestCoordAllPolygon(t *testing.T) {
	json := "{ \"type\": \"Polygon\", \"coordinates\": [[[0.0, 0.0], [1.0, 1.0], [0.0, 1.0], [0.0, 0.0]]]}"
	geometry, err := geometry.FromJSON(json)
	assert.Nil(t, err, "geometry can't be nil")

	poly, err := geometry.ToPolygon()
	assert.Nil(t, err, "convert to Polygon error")

	f := false
	pts, err := CoordAll(poly, &f)
	assert.Nil(t, err, "coord all err")

	assert.Equal(t, len(pts), 4)

	assert.Equal(t, pts[0].Lat, 0.0)
	assert.Equal(t, pts[0].Lng, 0.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 1.0)

	assert.Equal(t, pts[2].Lat, 1.0)
	assert.Equal(t, pts[2].Lng, 0.0)

	assert.Equal(t, pts[3].Lat, 0.0)
	assert.Equal(t, pts[3].Lng, 0.0)
}

func TestCoordExclueWrapCoord(t *testing.T) {
	json := "{ \"type\": \"Polygon\", \"coordinates\": [[[0.0, 0.0], [1.0, 1.0], [0.0, 1.0], [0.0, 0.0]]]}"
	geometry, err := geometry.FromJSON(json)
	assert.Nil(t, err, "geometry can't be nil")

	poly, err := geometry.ToPolygon()
	assert.Nil(t, err, "convert to Polygon error")

	f := true
	pts, err := CoordAll(poly, &f)
	assert.Nil(t, err, "coord all err")

	assert.Equal(t, len(pts), 3)

	assert.Equal(t, pts[0].Lat, 0.0)
	assert.Equal(t, pts[0].Lng, 0.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 1.0)

	assert.Equal(t, pts[2].Lat, 1.0)
	assert.Equal(t, pts[2].Lng, 0.0)
}

func TestCoordMultiPolygon(t *testing.T) {
	json := "{ \"type\": \"MultiPolygon\", \"coordinates\": [[[[0, 0], [1, 1], [0, 1], [0, 0]]]]}"
	geometry, err := geometry.FromJSON(json)
	assert.Nil(t, err, "geometry can't be nil")

	multiPoly, err := geometry.ToMultiPolygon()
	assert.Nil(t, err, "convert to MultiPolygon error")

	f := false
	pts, err := CoordAll(multiPoly, &f)
	assert.Nil(t, err, "coord all err")

	assert.Equal(t, len(pts), 4)

	assert.Equal(t, pts[0].Lat, 0.0)
	assert.Equal(t, pts[0].Lng, 0.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 1.0)

	assert.Equal(t, pts[2].Lat, 1.0)
	assert.Equal(t, pts[2].Lng, 0.0)

	assert.Equal(t, pts[3].Lat, 0.0)
	assert.Equal(t, pts[3].Lng, 0.0)

}

func TestInvariantGetCoord(t *testing.T) {
	json := "{ \"type\": \"Feature\", \"geometry\": { \"type\":\"Point\", \"coordinates\":[1,2]}}"
	feature, err := feature.FromJSON(json)
	assert.Nil(t, err, "feature can't be nil")

	g, err := GetCoord(*feature)
	assert.Nil(t, err, "feature can't be nil")

	assert.Equal(t, g.Lat, 2.0)
	assert.Equal(t, g.Lng, 1.0)
}

func TestCoordAllFeatureCollection(t *testing.T) {
	json := "{\"type\": \"FeatureCollection\", \"features\": [{\"type\": \"Feature\",\"properties\": {\"population\": 200},\"geometry\": {\"type\": \"Point\",\"coordinates\": [-112.0372, 46.608058]}}]}"
	c, err := feature.CollectionFromJSON(json)
	assert.Nil(t, err, "point can't be nil")
	assert.NotNil(t, c, "feature collection can't be nil")
	exclude := true
	pts, err := CoordAll(c, &exclude)
	assert.Nil(t, err, "coord all err")

	assert.Equal(t, len(pts), 1)

	assert.Equal(t, pts[0].Lat, 46.608058)
	assert.Equal(t, pts[0].Lng, -112.0372)
}

// TODO: Write Test after BBOX
func TestCoordAllGeometryCollection(t *testing.T) {

}
