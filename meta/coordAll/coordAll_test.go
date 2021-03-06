package meta

import (
	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"testing"
)

func TestCoordAllPont(t *testing.T) {
	json := "{ \"type\": \"Point\", \"coordinates\": [23.0, 54.0]}"
	geometry, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	pt, err := geometry.ToPoint()
	if err != nil {
		t.Errorf("convert to Point error %v", err)
	}

	pts, err := CoordAll(pt, nil)
	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 54.0)
	assert.Equal(t, pts[0].Lng, 23.0)
}

func TestCoordAllLineString(t *testing.T) {
	json := "{ \"type\": \"LineString\", \"coordinates\": [[0.0, 0.0], [1.0, 1.0]]}"
	geometry, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	ln, err := geometry.ToLineString()
	if err != nil {
		t.Errorf("convert to LineString error %v", err)
	}

	pts, err := CoordAll(ln, nil)
	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}
	assert.Equal(t, len(pts), 2)

	assert.Equal(t, pts[0].Lat, 0.0)
	assert.Equal(t, pts[0].Lng, 0.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 1.0)
}

func TestCoordAllPolygon(t *testing.T) {
	json := "{ \"type\": \"Polygon\", \"coordinates\": [[[0.0, 0.0], [1.0, 1.0], [0.0, 1.0], [0.0, 0.0]]]}"
	geometry, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	poly, err := geometry.ToPolygon()
	if err != nil {
		t.Errorf("convert to Polygon error %v", err)
	}

	f := false
	pts, err := CoordAll(poly, &f)
	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}
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
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	poly, err := geometry.ToPolygon()
	if err != nil {
		t.Errorf("convert to Polygon error %v", err)
	}

	f := true
	pts, err := CoordAll(poly, &f)
	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}

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
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	multiPoly, err := geometry.ToMultiPolygon()
	if err != nil {
		t.Errorf("convert to MultiPolygon error %v", err)
	}

	f := false
	pts, err := CoordAll(multiPoly, &f)
	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}

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
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	g, err := GetCoord(*feature)
	if err != nil {
		t.Errorf("GetCoord error %v", err)
	}

	assert.Equal(t, g.Lat, 2.0)
	assert.Equal(t, g.Lng, 1.0)
}

func TestCoordAllFeatureCollection(t *testing.T) {
	json := "{\"type\": \"FeatureCollection\", \"features\": [{\"type\": \"Feature\",\"properties\": {\"population\": 200},\"geometry\": {\"type\": \"Point\",\"coordinates\": [-112.0372, 46.608058]}}]}"
	c, err := feature.CollectionFromJSON(json)
	if err != nil {
		t.Errorf("CollectionFromJSON error %v", err)
	}

	if c == nil {
		t.Error("feature collection can't be nil")
	}

	exclude := true
	pts, err := CoordAll(c, &exclude)
	if err != nil {
		t.Errorf("CoordAll error %v", err)
	}

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 46.608058)
	assert.Equal(t, pts[0].Lng, -112.0372)
}

// TODO: Write Test after BBOX
func TestCoordAllGeometryCollection(t *testing.T) {

}
