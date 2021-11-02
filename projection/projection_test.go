package projection

import (
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"
)

const MercatorPoint = "../test-data/mercator.point.geojson"
const MercatorPolygon = "../test-data/mercator.polygon.geojson"
const MercatorMultiPolygon = "../test-data/mercator.multipolygon.geojson"

func TestConvertToMercatorPoint(t *testing.T) {
	p := geometry.Point{
		Lng: -71.0,
		Lat: 41.0,
	}
	xy := ConvertToMercator([]float64{p.Lng, p.Lat})
	assert.Equal(t, xy[0], -7903683.846322424)
	assert.Equal(t, xy[1], 5012341.663847514)
}

func TestConvertToWgs84(t *testing.T) {
	p := []float64{-7903683.846322424, 5012341.663847514}
	wgs84Point := ConvertToWgs84(p)
	assert.Equal(t, wgs84Point[0], -71.0)
	assert.Equal(t, wgs84Point[1], 40.99999999999998) //=41.0
}

func TestProjectionPoint(t *testing.T) {
	p := geometry.Point{
		Lng: 10.0,
		Lat: 40.0,
	}
	mercator := ConvertToMercator([]float64{p.Lng, p.Lat})
	wgs84 := ConvertToWgs84(mercator)
	assert.Equal(t, p.Lng, wgs84[0])
	assert.Equal(t, p.Lat, wgs84[1])
}

func TestConvertToWGS84EmptyGeoJSON(t *testing.T) {
	geojson, err := ToWgs84(nil)
	assert.Nil(t, geojson)
	assert.Equal(t, err.Error(), "geojson is required")
}

func TestConvertToMercatorEmptyGeoJSON(t *testing.T) {
	geojson, err := ToMercator(nil)
	assert.Nil(t, geojson)
	assert.Equal(t, err.Error(), "geojson is required")
}

func TestConvertToWGS84PointFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToWgs84(f)
	assert.Nil(t, err)

	k, ok := r.(*feature.Feature)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	coords := k.Geometry.Coordinates
	coord, ok := coords.(geometry.Point)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}
	assert.Equal(t, coord.Lat, 40.99999999999998)
	assert.Equal(t, coord.Lng, -71.0)
}

func TestConvertToWGS84PolygonFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToWgs84(f)
	assert.Nil(t, err)

	k, ok := r.(*feature.Feature)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	coords := k.Geometry.Coordinates
	coord, ok := coords.([]geometry.LineString)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord[0].Coordinates[0].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -116.00000000000237)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 66.00000000000074)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[3].Lat, 66.00000000000074)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -116.00000000000237)

	assert.Equal(t, coord[0].Coordinates[4].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -116.00000000000237)
}

//TODO: Fix MultiPolygon Test
func TestConvertToWGS84MultiPolygonFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToWgs84(f)
	assert.Nil(t, err)

	k, ok := r.(*feature.Feature)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	coords := k.Geometry.Coordinates
	coord, ok := coords.([]geometry.Polygon)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[0].Lat, -44.999999999999595)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[0].Lng, -116.00000000000237)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[1].Lat, -44.999999999999595)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[1].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[2].Lat, -56.000000000001094)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[2].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[3].Lat, -56.000000000001094)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[3].Lng, -116.00000000000237)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[4].Lat, -44.999999999999595)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[4].Lng, -116.00000000000237)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[0].Lat, 9.102096999999901)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[0].Lng, -90.35156200000102)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[1].Lat, -3.5134210000008026)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[1].Lng, -77.69531200000432)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[2].Lat, 12.211180000002315)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[2].Lng, -65.03906199999867)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[3].Lat, 21.61657899999976)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[3].Lng, -65.74218699999848)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[4].Lat, 24.52713500000161)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[4].Lng, -84.02343700000269)
}

//TODO: Add Fiji Test

//TODO: Add MultiLine Test
//TODO: Add Feature Collection Test
