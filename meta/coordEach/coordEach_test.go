package meta

import (
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestCoordEachPoint(t *testing.T) {
	json := "{ \"type\": \"Point\", \"coordinates\": [23.0, 54.0]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	pt, err := geom.ToPoint()
	if err != nil {
		t.Errorf("convert to Point error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}

	pts, err := CoordEach(pt, fnc, nil)

	if err != nil {
		t.Errorf("CoordEach err %v", err)
	}

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 54.0)
	assert.Equal(t, pts[0].Lng, 23.0)
}

func TestCoordEachMultiPoint(t *testing.T) {
	json := "{ \"type\": \"MultiPoint\", \"coordinates\": [[102, -10],[103, 1],[104, 0],[130, 4]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	pt, err := geom.ToMultiPoint()
	if err != nil {
		t.Errorf("convert to MultiPoint error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}

	pts, err := CoordEach(pt, fnc, nil)

	if err != nil {
		t.Errorf("CoordEach err %v", err)
	}

	assert.Equal(t, len(pts), 4)
	assert.Equal(t, pts[0].Lat, -10.0)
	assert.Equal(t, pts[0].Lng, 102.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 103.0)

	assert.Equal(t, pts[2].Lat, 0.0)
	assert.Equal(t, pts[2].Lng, 104.0)

	assert.Equal(t, pts[3].Lat, 4.0)
	assert.Equal(t, pts[3].Lng, 130.0)
}

func TestCoordEachLineString(t *testing.T) {
	json := "{ \"type\": \"LineString\", \"coordinates\": [[102, -10],[103, 1],[104, 0],[130, 4]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	pt, err := geom.ToLineString()
	if err != nil {
		t.Errorf("convert to LineString error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}

	pts, err := CoordEach(pt, fnc, nil)

	if err != nil {
		t.Errorf("CoordEach err %v", err)
	}

	assert.Equal(t, len(pts), 4)
	assert.Equal(t, pts[0].Lat, -10.0)
	assert.Equal(t, pts[0].Lng, 102.0)

	assert.Equal(t, pts[1].Lat, 1.0)
	assert.Equal(t, pts[1].Lng, 103.0)

	assert.Equal(t, pts[2].Lat, 0.0)
	assert.Equal(t, pts[2].Lng, 104.0)

	assert.Equal(t, pts[3].Lat, 4.0)
	assert.Equal(t, pts[3].Lng, 130.0)
}

func TestCoordEachMultiLineString(t *testing.T) {
	json := "{ \"type\": \"MultiLineString\", \"coordinates\": [[[100, 0],[101, 1]],[[102, 2],[103, 3]]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	pt, err := geom.ToMultiLineString()
	if err != nil {
		t.Errorf("convert to MultiLineString error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		p.Lat = p.Lat + 1
		p.Lng = p.Lng + 1
		return p
	}

	pts, err := CoordEach(pt, fnc, nil)

	if err != nil {
		t.Errorf("CoordEach err %v", err)
	}

	assert.Equal(t, len(pts), 4)
	assert.Equal(t, pts[0].Lat, 1.0)
	assert.Equal(t, pts[0].Lng, 101.0)

	assert.Equal(t, pts[1].Lat, 2.0)
	assert.Equal(t, pts[1].Lng, 102.0)

	assert.Equal(t, pts[2].Lat, 3.0)
	assert.Equal(t, pts[2].Lng, 103.0)

	assert.Equal(t, pts[3].Lat, 4.0)
	assert.Equal(t, pts[3].Lng, 104.0)
}

func TestCoordEachPolygon(t *testing.T) {
	json := "{ \"type\": \"Polygon\", \"coordinates\": [[[125.0,-15.0],[113.0,-22.0],[117.0,-37.0],[130.0,-33.0],[148.0,-39.0],[154.0,-27.0],[144.0,-15.0],[125.0,-15.0]]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	poly, err := geom.ToPolygon()
	if err != nil {
		t.Errorf("convert to Polygon error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}
	extractWrapPolygon := false
	pts, err := CoordEach(poly, fnc, &extractWrapPolygon)

	if err != nil {
		t.Errorf("CoordEach err %v", err)
	}

	assert.Equal(t, len(pts), 8)
	assert.Equal(t, pts[0].Lat, -15.0)
	assert.Equal(t, pts[0].Lng, 125.0)

	assert.Equal(t, pts[1].Lat, -22.0)
	assert.Equal(t, pts[1].Lng, 113.0)

	assert.Equal(t, pts[2].Lat, -37.0)
	assert.Equal(t, pts[2].Lng, 117.0)

	assert.Equal(t, pts[3].Lat, -33.0)
	assert.Equal(t, pts[3].Lng, 130.0)

	assert.Equal(t, pts[4].Lat, -39.0)
	assert.Equal(t, pts[4].Lng, 148.0)

	assert.Equal(t, pts[5].Lat, -27.0)
	assert.Equal(t, pts[5].Lng, 154.0)

	assert.Equal(t, pts[6].Lat, -15.0)
	assert.Equal(t, pts[6].Lng, 144.0)

	assert.Equal(t, pts[7].Lat, -15.0)
	assert.Equal(t, pts[7].Lng, 125.0)
}

func TestCoordEachPolygonExcludeWrapCoord(t *testing.T) {
	json := "{ \"type\": \"Polygon\", \"coordinates\": [[[125.0,-15.0],[113.0,-22.0],[117.0,-37.0],[130.0,-33.0],[148.0,-39.0],[154.0,-27.0],[144.0,-15.0],[125.0,-15.0]]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	poly, err := geom.ToPolygon()
	if err != nil {
		t.Errorf("convert to Polygon error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}
	extractWrapPolygon := true
	pts, err := CoordEach(poly, fnc, &extractWrapPolygon)

	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}

	assert.Equal(t, len(pts), 7)
	assert.Equal(t, pts[0].Lat, -15.0)
	assert.Equal(t, pts[0].Lng, 125.0)

	assert.Equal(t, pts[1].Lat, -22.0)
	assert.Equal(t, pts[1].Lng, 113.0)

	assert.Equal(t, pts[2].Lat, -37.0)
	assert.Equal(t, pts[2].Lng, 117.0)

	assert.Equal(t, pts[3].Lat, -33.0)
	assert.Equal(t, pts[3].Lng, 130.0)

	assert.Equal(t, pts[4].Lat, -39.0)
	assert.Equal(t, pts[4].Lng, 148.0)

	assert.Equal(t, pts[5].Lat, -27.0)
	assert.Equal(t, pts[5].Lng, 154.0)

	assert.Equal(t, pts[6].Lat, -15.0)
	assert.Equal(t, pts[6].Lng, 144.0)

}

func TestCoordEachMultiPolygon(t *testing.T) {
	json := "{ \"type\": \"MultiPolygon\", \"coordinates\": [[[[102, 2],[103, 2],[103, 3],[102, 3],[102, 2]]],[[[100, 0],[101, 0],[101, 1],[100, 1],[100, 0]]]]}"
	geom, err := geometry.FromJSON(json)
	if err != nil {
		t.Errorf("geometry error %v", err)
	}

	poly, err := geom.ToMultiPolygon()
	if err != nil {
		t.Errorf("convert to MultiPolygon error %v", err)
	}

	fnc := func(p geometry.Point) geometry.Point {
		return p
	}
	extractWrapPolygon := false
	pts, err := CoordEach(poly, fnc, &extractWrapPolygon)

	if err != nil {
		t.Errorf("CoordAll err %v", err)
	}

	assert.Equal(t, len(pts), 10)
	assert.Equal(t, pts[0].Lat, 2.0)
	assert.Equal(t, pts[0].Lng, 102.0)

	assert.Equal(t, pts[1].Lat, 2.0)
	assert.Equal(t, pts[1].Lng, 103.0)

	assert.Equal(t, pts[2].Lat, 3.0)
	assert.Equal(t, pts[2].Lng, 103.0)

	assert.Equal(t, pts[3].Lat, 3.0)
	assert.Equal(t, pts[3].Lng, 102.0)

	assert.Equal(t, pts[4].Lat, 2.0)
	assert.Equal(t, pts[4].Lng, 102.0)

	assert.Equal(t, pts[5].Lat, 0.0)
	assert.Equal(t, pts[5].Lng, 100.0)

	assert.Equal(t, pts[6].Lat, 0.0)
	assert.Equal(t, pts[6].Lng, 101.0)

	assert.Equal(t, pts[7].Lat, 1.0)
	assert.Equal(t, pts[7].Lng, 101.0)

	assert.Equal(t, pts[8].Lat, 1.0)
	assert.Equal(t, pts[8].Lng, 100.0)

	assert.Equal(t, pts[9].Lat, 0.0)
	assert.Equal(t, pts[9].Lng, 100.0)
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

	fnc := func(p geometry.Point) geometry.Point {
		p.Lat = p.Lat + 0.1
		p.Lng = p.Lng + 0.1
		return p
	}

	exclude := true
	pts, err := CoordEach(c, fnc, &exclude)
	if err != nil {
		t.Errorf("CoordEach error %v", err)
	}

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 46.708058)
	assert.Equal(t, pts[0].Lng, -111.9372)
}
