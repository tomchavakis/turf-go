package meta

import (
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestCoordEach(t *testing.T) {
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
		t.Errorf("CoordAll err %v", err)
	}

	assert.Equal(t, len(pts), 1)
	assert.Equal(t, pts[0].Lat, 54.0)
	assert.Equal(t, pts[0].Lng, 23.0)
}





