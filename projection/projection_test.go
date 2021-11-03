package projection

import (
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"
)

const MercatorPoint = "../test-data/mercator.point.geojson"
const MercatorMultiPoint = "../test-data/mercator.multipoint.geojson"

const MercatorPolygon = "../test-data/mercator.polygon.geojson"
const MercatorMultiPolygon = "../test-data/mercator.multipolygon.geojson"

const MercatorMultiLineString = "../test-data/mercator.multilinestring.geojson"
const MercatorLineString = "../test-data/mercator.linestring.geojson"

const MercatorFeatureCollection = "../test-data/mercator.featurecollection.geojson"

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

func TestConvertToWGS84MultiPointFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorMultiPoint)
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
	coord, ok := coords.([]geometry.Point)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord[0].Lat, 33.72434000000235)
	assert.Equal(t, coord[0].Lng, -20.39062500000365)

	assert.Equal(t, coord[1].Lat, 47.51720099999992)
	assert.Equal(t, coord[1].Lng, -3.5156249999990803)

	assert.Equal(t, coord[2].Lat, 16.97274100000141)
	assert.Equal(t, coord[2].Lng, 14.062499999996321)
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

func TestConvertToWGS84LineStringFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorLineString)
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
	coord, ok := coords.([]geometry.Point)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord[0].Lat, 33.72434000000235)
	assert.Equal(t, coord[0].Lng, -20.39062500000365)

	assert.Equal(t, coord[1].Lat, 47.51720099999992)
	assert.Equal(t, coord[1].Lng, -3.5156249999990803)

	assert.Equal(t, coord[2].Lat, 16.97274100000141)
	assert.Equal(t, coord[2].Lng, 14.062499999996321)
}

func TestConvertToWGS84MultilineFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorMultiLineString)
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
	assert.Equal(t, coord[0].Coordinates[0].Lat, -1.4061090000014986)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -127.96875000000244)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 33.43144100000098)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -91.05468700000083)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 57.51582299999881)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -71.71874999999919)

	assert.Equal(t, coord[0].Coordinates[3].Lat, 65.658275)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -29.531250000001258)

	assert.Equal(t, coord[1].Coordinates[0].Lat, 29.535230000001896)
	assert.Equal(t, coord[1].Coordinates[0].Lng, -72.77343699999663)

	assert.Equal(t, coord[1].Coordinates[1].Lat, 13.23994499999653)
	assert.Equal(t, coord[1].Coordinates[1].Lng, -82.61718700000304)

	assert.Equal(t, coord[1].Coordinates[2].Lat, 9.79567800000412)
	assert.Equal(t, coord[1].Coordinates[2].Lng, -49.57031200000271)

	assert.Equal(t, coord[1].Coordinates[3].Lat, -11.178402000004093)
	assert.Equal(t, coord[1].Coordinates[3].Lng, -84.3749999999959)

	assert.Equal(t, coord[1].Coordinates[4].Lat, -55.57834500000183)
	assert.Equal(t, coord[1].Coordinates[4].Lng, -29.531250000001258)
}

func TestConvertToWGS84FeatureCollection(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorFeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	fc, err := feature.CollectionFromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToWgs84(fc)
	assert.Nil(t, err)

	k, ok := r.(*feature.Collection)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	coords := k.Features

	coord, ok := coords[0].Geometry.Coordinates.([]geometry.LineString)

	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord[0].Coordinates[0].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -116.00000000000237)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 66.00000000000074)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -90.0000000000034)

	assert.Equal(t, coord[0].Coordinates[4].Lat, 58.00000000000044)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -116.00000000000237)

	coord2, ok2 := coords[1].Geometry.Coordinates.([]geometry.Point)

	if !ok2 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord2[0].Lat, 33.72434000000235)
	assert.Equal(t, coord2[0].Lng, -20.39062500000365)

	assert.Equal(t, coord2[1].Lat, 47.51720099999992)
	assert.Equal(t, coord2[1].Lng, -3.5156249999990803)

	assert.Equal(t, coord2[2].Lat, 16.97274100000141)
	assert.Equal(t, coord2[2].Lng, 14.062499999996321)

	coord3, ok3 := coords[2].Geometry.Coordinates.(geometry.Point)

	if !ok3 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord3.Lat, 3.864254999997437)
	assert.Equal(t, coord3.Lng, -76.28906199999572)
}

//TODO: Passed 180th meridian

//TODO: Add Fiji Test
