package projection

import (
	"testing"

	"github.com/tomchavakis/geojson/feature"
	"github.com/tomchavakis/geojson/geometry"
	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/utils"
)

const MercatorPoint = "../test-data/mercator.point.geojson"
const MercatorMultiPoint = "../test-data/mercator.multipoint.geojson"
const MercatorPolygon = "../test-data/mercator.polygon.geojson"
const MercatorMultiPolygon = "../test-data/mercator.multipolygon.geojson"
const MercatorLineString = "../test-data/mercator.linestring.geojson"
const MercatorMultiLineString = "../test-data/mercator.multilinestring.geojson"
const MercatorPassedMeridian = "../test-data/mercator.passed180thmeridian.geojson"
const MercatorPassedMeridian2 = "../test-data/mercator.passed180thmeridian2.geojson"
const MercatorGeometryPoint = "../test-data/mercator.geometry.point.geojson"
const MercatorGeometryMultiPoint = "../test-data/mercator.geometry.multipoint.geojson"
const MercatorGeometryLineString = "../test-data/mercator.geometry.linestring.geojson"
const MercatorGeometryMultiLineString = "../test-data/mercator.geometry.multilinestring.geojson"
const MercatorGeometryPolygon = "../test-data/mercator.geometry.polygon.geojson"
const MercatorGeometryMultiPolygon = "../test-data/mercator.geometry.multipolygon.geojson"
const MercatorFeatureCollection = "../test-data/mercator.featurecollection.geojson"
const MercatorGeometryCollection = "../test-data/mercator.geometrycollection.geojson"

const WGS84Point = "../test-data/wgs84.point.geojson"
const WGS84MultiPoint = "../test-data/wgs84.multipoint.geojson"
const WGS84Polygon = "../test-data/wgs84.polygon.geojson"
const WGS84MultiPolygon = "../test-data/wgs84.multipolygon.geojson"
const WGS84LineString = "../test-data/wgs84.linestring.geojson"
const WGS84MultiLineString = "../test-data/wgs84.multilinestring.geojson"
const WGS84PassedMeridian = "../test-data/wgs84.passed180thmeridian.geojson"
const WGS84PassedMeridian2 = "../test-data/wgs84.passed180thmeridian2.geojson"

const WGS84GeometryPoint = "../test-data/wgs84.geometry.point.geojson"
const WGS84GeometryMultiPoint = "../test-data/wgs84.geometry.multipoint.geojson"
const WGS84GeometryPolygon = "../test-data/wgs84.geometry.polygon.geojson"
const WGS84GeometryMultiPolygon = "../test-data/wgs84.geometry.multipolygon.geojson"
const WGS84GeometryLineString = "../test-data/wgs84.geometry.linestring.geojson"
const WGS84GeometryMultiLineString = "../test-data/wgs84.geometry.multilinestring.geojson"

const WGS84FeatureCollection = "../test-data/wgs84.featurecollection.geojson"
const WGS84GeometryCollection = "../test-data/wgs84.geometrycollection.geojson"

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

func TestConvertToWGS84PassedMeridian(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorPassedMeridian)
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

	assert.Equal(t, coord[0].Coordinates[0].Lat, -23.56398700000191)
	assert.Equal(t, coord[0].Coordinates[0].Lng, 113.20312499999733)

	assert.Equal(t, coord[0].Coordinates[1].Lat, -34.8138030000016)
	assert.Equal(t, coord[0].Coordinates[1].Lng, 116.7187499999964)

	assert.Equal(t, coord[0].Coordinates[2].Lat, -31.05293399999659)
	assert.Equal(t, coord[0].Coordinates[2].Lng, 131.92382799999635)

	assert.Equal(t, coord[0].Coordinates[3].Lat, -38.41055799999684)
	assert.Equal(t, coord[0].Coordinates[3].Lng, 141.41601599999615)

	assert.Equal(t, coord[0].Coordinates[4].Lat, -38.34165599999897)
	assert.Equal(t, coord[0].Coordinates[4].Lng, 148.5351559999959)

	assert.Equal(t, coord[0].Coordinates[5].Lat, -27.371766999999128)
	assert.Equal(t, coord[0].Coordinates[5].Lng, 153.98437499999565)

	assert.Equal(t, coord[0].Coordinates[6].Lat, -10.746968999999027)
	assert.Equal(t, coord[0].Coordinates[6].Lng, 142.03124999999878)

	assert.Equal(t, coord[0].Coordinates[7].Lat, -17.560247000003802)
	assert.Equal(t, coord[0].Coordinates[7].Lng, 140.27343800000153)

	assert.Equal(t, coord[0].Coordinates[8].Lat, -15.029686000000622)
	assert.Equal(t, coord[0].Coordinates[8].Lng, 135.61523400000323)

	assert.Equal(t, coord[0].Coordinates[9].Lat, -11.953348999996189)
	assert.Equal(t, coord[0].Coordinates[9].Lng, 136.66992199999626)

	assert.Equal(t, coord[0].Coordinates[10].Lat, -11.350797000000282)
	assert.Equal(t, coord[0].Coordinates[10].Lng, 131.30859400000273)

	assert.Equal(t, coord[0].Coordinates[11].Lat, -16.9727410000014)
	assert.Equal(t, coord[0].Coordinates[11].Lng, 122.25585899999774)

	assert.Equal(t, coord[0].Coordinates[12].Lat, -19.394068000000363)
	assert.Equal(t, coord[0].Coordinates[12].Lng, 121.46484399999632)

	assert.Equal(t, coord[0].Coordinates[13].Lat, -23.56398700000191)
	assert.Equal(t, coord[0].Coordinates[13].Lng, 113.20312499999733)
}

func TestConvertToWGS84PassedMeridian2(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorPassedMeridian2)
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

	assert.Equal(t, coord[0].Coordinates[0].Lat, 11.350797000000295)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -73.47656199999646)

	assert.Equal(t, coord[0].Coordinates[1].Lat, -3.6888550000002938)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -81.73828100000442)

	assert.Equal(t, coord[0].Coordinates[2].Lat, -14.94478499999696)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -75.41015599999709)

	assert.Equal(t, coord[0].Coordinates[3].Lat, -19.145167999997195)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -70.31249999999957)

	assert.Equal(t, coord[0].Coordinates[4].Lat, -50.513426999998934)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -76.28906199999572)

	assert.Equal(t, coord[0].Coordinates[5].Lat, -55.677583999997594)
	assert.Equal(t, coord[0].Coordinates[5].Lng, -68.73046900000112)

	assert.Equal(t, coord[0].Coordinates[6].Lat, -54.673831000001236)
	assert.Equal(t, coord[0].Coordinates[6].Lng, -65.39062500000085)

	assert.Equal(t, coord[0].Coordinates[7].Lat, -50.513426999998934)
	assert.Equal(t, coord[0].Coordinates[7].Lng, -68.90624999999993)

	assert.Equal(t, coord[0].Coordinates[8].Lat, -39.50404099999976)
	assert.Equal(t, coord[0].Coordinates[8].Lng, -62.753906000000406)

	assert.Equal(t, coord[0].Coordinates[9].Lat, -38.41055799999684)
	assert.Equal(t, coord[0].Coordinates[9].Lng, -58.007812000000506)

	assert.Equal(t, coord[0].Coordinates[10].Lat, -28.459033000000893)
	assert.Equal(t, coord[0].Coordinates[10].Lng, -48.8671870000029)

	assert.Equal(t, coord[0].Coordinates[11].Lat, -25.64152600000281)
	assert.Equal(t, coord[0].Coordinates[11].Lng, -48.69140600000408)

	assert.Equal(t, coord[0].Coordinates[12].Lat, -22.105998999999827)
	assert.Equal(t, coord[0].Coordinates[12].Lng, -41.660155999996945)

	assert.Equal(t, coord[0].Coordinates[13].Lat, -6.839170000003316)
	assert.Equal(t, coord[0].Coordinates[13].Lng, -34.80468699999759)

	assert.Equal(t, coord[0].Coordinates[14].Lat, 1.05462800000055)
	assert.Equal(t, coord[0].Coordinates[14].Lng, -50.62499999999574)

	assert.Equal(t, coord[0].Coordinates[15].Lat, 4.56547399999967)
	assert.Equal(t, coord[0].Coordinates[15].Lng, -52.031250000004356)

	assert.Equal(t, coord[0].Coordinates[16].Lat, 10.833305999997641)
	assert.Equal(t, coord[0].Coordinates[16].Lng, -63.98437500000122)

	assert.Equal(t, coord[0].Coordinates[17].Lat, 11.350797000000295)
	assert.Equal(t, coord[0].Coordinates[17].Lng, -73.47656199999646)
}

func TestConvertToWGS84GeometryPoint(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	pt, err := f.ToPoint()
	if err != nil {
		t.Errorf("error %v", err)
	}

	r, err := ToWgs84(pt)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Point)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}
	assert.Equal(t, k.Lat, 40.99999999999998)
	assert.Equal(t, k.Lng, -71.0)
}

func TestConvertToWGS84GeometryMultiPoint(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryMultiPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	mpt, err := f.ToMultiPoint()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToWgs84(mpt)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiPoint)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Lat, 33.72434000000235)
	assert.Equal(t, k.Coordinates[0].Lng, -20.39062500000365)

	assert.Equal(t, k.Coordinates[1].Lat, 47.51720099999992)
	assert.Equal(t, k.Coordinates[1].Lng, -3.5156249999990803)

	assert.Equal(t, k.Coordinates[2].Lat, 16.97274100000141)
	assert.Equal(t, k.Coordinates[2].Lng, 14.062499999996321)
}

func TestConvertToWGS84GeometryPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	poly, err := f.ToPolygon()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToWgs84(poly)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Polygon)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lat, 58.00000000000044)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lng, -116.00000000000237)

	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lat, 58.00000000000044)
	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lng, -90.0000000000034)

	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lat, 66.00000000000074)
	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lng, -90.0000000000034)

	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lat, 66.00000000000074)
	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lng, -116.00000000000237)

	assert.Equal(t, k.Coordinates[0].Coordinates[4].Lat, 58.00000000000044)
	assert.Equal(t, k.Coordinates[0].Coordinates[4].Lng, -116.00000000000237)
}

func TestConvertToWGS84GeometryMultiPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	mpl, err := f.ToMultiPolygon()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToWgs84(mpl)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiPolygon)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[0].Lat, -44.999999999999595)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[0].Lng, -116.00000000000237)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[1].Lat, -44.999999999999595)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[1].Lng, -90.0000000000034)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[2].Lat, -56.000000000001094)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[2].Lng, -90.0000000000034)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[3].Lat, -56.000000000001094)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[3].Lng, -116.00000000000237)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[4].Lat, -44.999999999999595)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[4].Lng, -116.00000000000237)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[0].Lat, 9.102096999999901)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[0].Lng, -90.35156200000102)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[1].Lat, -3.5134210000008026)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[1].Lng, -77.69531200000432)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[2].Lat, 12.211180000002315)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[2].Lng, -65.03906199999867)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[3].Lat, 21.61657899999976)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[3].Lng, -65.74218699999848)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[4].Lat, 24.52713500000161)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[4].Lng, -84.02343700000269)

}

func TestConvertToWGS84GeometryLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	lns, err := f.ToLineString()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToWgs84(lns)
	assert.Nil(t, err)

	k, ok := r.(*geometry.LineString)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Lat, -1.4061090000014986)
	assert.Equal(t, k.Coordinates[0].Lng, -127.96875000000244)

	assert.Equal(t, k.Coordinates[1].Lat, 33.43144100000098)
	assert.Equal(t, k.Coordinates[1].Lng, -91.05468700000083)

	assert.Equal(t, k.Coordinates[2].Lat, 57.51582299999881)
	assert.Equal(t, k.Coordinates[2].Lng, -71.71874999999919)
}

func TestConvertToWGS84GeometryMultiLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryMultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	mlns, err := f.ToMultiLineString()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToWgs84(mlns)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiLineString)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lat, -1.4061090000014986)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lng, -127.96875000000244)

	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lat, 33.43144100000098)
	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lng, -91.05468700000083)

	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lat, 57.51582299999881)
	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lng, -71.71874999999919)

	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lat, 65.658275)
	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lng, -29.531250000001258)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Lat, 29.535230000001896)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Lng, -72.77343699999663)

	assert.Equal(t, k.Coordinates[1].Coordinates[1].Lat, 13.23994499999653)
	assert.Equal(t, k.Coordinates[1].Coordinates[1].Lng, -82.61718700000304)

	assert.Equal(t, k.Coordinates[1].Coordinates[2].Lat, 9.79567800000412)
	assert.Equal(t, k.Coordinates[1].Coordinates[2].Lng, -49.57031200000271)

	assert.Equal(t, k.Coordinates[1].Coordinates[3].Lat, -11.178402000004093)
	assert.Equal(t, k.Coordinates[1].Coordinates[3].Lng, -84.3749999999959)

	assert.Equal(t, k.Coordinates[1].Coordinates[4].Lat, -55.57834500000183)
	assert.Equal(t, k.Coordinates[1].Coordinates[4].Lng, -29.531250000001258)

}

func TestConvertToWGS84GeometryCollection(t *testing.T) {
	p, err := utils.LoadJSONFixture(MercatorGeometryCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	collection, err := geometry.CollectionFromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToWgs84(collection)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Collection)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	coords := k.Geometries

	coord, ok := coords[0].Coordinates.(geometry.Point)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord.Lat, 40.99999999999998)
	assert.Equal(t, coord.Lng, -71.0)

	coord2, ok2 := coords[1].Coordinates.([]geometry.Point)

	if !ok2 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord2[0].Lat, 33.72434000000235)
	assert.Equal(t, coord2[0].Lng, -20.39062500000365)

	assert.Equal(t, coord2[1].Lat, 47.51720099999992)
	assert.Equal(t, coord2[1].Lng, -3.5156249999990803)

	assert.Equal(t, coord2[2].Lat, 16.97274100000141)
	assert.Equal(t, coord2[2].Lng, 14.062499999996321)
}

func TestConvertToMercatorPoint(t *testing.T) {
	p := geometry.Point{
		Lng: -71.0,
		Lat: 41.0,
	}
	xy := ConvertToMercator([]float64{p.Lng, p.Lat})
	assert.Equal(t, xy[0], -7903683.846322424)
	assert.Equal(t, xy[1], 5012341.663847514)
}

func TestConvertToMercatorEmptyGeoJSON(t *testing.T) {
	geojson, err := ToMercator(nil)
	assert.Nil(t, geojson)
	assert.Equal(t, err.Error(), "geojson is required")
}

func TestConvertToMercatorPointFeature(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84Point)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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
	assert.Equal(t, coord.Lat, 5012341.663847514)
	assert.Equal(t, coord.Lng, -7903683.846322424)
}

func TestConvertToMercatorMultiPoint(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84MultiPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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
	assert.Equal(t, coord[0].Lat, 3991847.4104379974)
	assert.Equal(t, coord[0].Lng, -2269873.991957)

	assert.Equal(t, coord[1].Lat, 6026906.856034)
	assert.Equal(t, coord[1].Lng, -391357.58482)

	assert.Equal(t, coord[2].Lat, 1917652.1632909998)
	assert.Equal(t, coord[2].Lng, 1565430.33928)
}

func TestConvertToMercatorPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84Polygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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

	assert.Equal(t, coord[0].Coordinates[0].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -12913060.932019735)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 9876845.895794801)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[3].Lat, 9876845.895794801)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -12913060.932019735)

	assert.Equal(t, coord[0].Coordinates[4].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -12913060.932019735)
}

func TestConvertToMercatorMultiPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84MultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[0].Lat, -5621521.486192067)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[0].Lng, -12913060.932019735)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[1].Lat, -5621521.486192067)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[1].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[2].Lat, -7558415.656081783)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[2].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[3].Lat, -7558415.656081783)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[3].Lng, -12913060.932019735)

	assert.Equal(t, coord[0].Coordinates[0].Coordinates[4].Lat, -5621521.486192067)
	assert.Equal(t, coord[0].Coordinates[0].Coordinates[4].Lng, -12913060.932019735)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[0].Lat, 1017529.7499880107)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[0].Lng, -10057889.985536376)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[1].Lat, -391357.57972990983)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[1].Lng, -8649002.568864517)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[2].Lat, 1369751.5250587354)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[2].Lng, -7240115.374831639)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[3].Lat, 2465552.7440450294)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[3].Lng, -7318386.89179566)

	assert.Equal(t, coord[1].Coordinates[0].Coordinates[4].Lat, 2817774.6324118003)
	assert.Equal(t, coord[1].Coordinates[0].Coordinates[4].Lng, -9353446.221540703)

}

func TestConvertToMercatorLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84LineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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

	assert.Equal(t, coord[0].Lat, 3991847.4104379974)
	assert.Equal(t, coord[0].Lng, -2269873.991957)

	assert.Equal(t, coord[1].Lat, 6026906.856034)
	assert.Equal(t, coord[1].Lng, -391357.58482)

	assert.Equal(t, coord[2].Lat, 1917652.1632909998)
	assert.Equal(t, coord[2].Lng, 1565430.33928)

}

func TestConvertToMercatorMultiLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84MultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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
	assert.Equal(t, coord[0].Coordinates[0].Lat, -156543.0522530012)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -14245416.087452)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 3952711.5619209986)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -10136161.391181)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 7866287.4827539995)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -7983694.730329999)

	assert.Equal(t, coord[0].Coordinates[3].Lat, 9783939.750186)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -3287403.712489)

	assert.Equal(t, coord[1].Coordinates[0].Lat, 3443946.8023359994)
	assert.Equal(t, coord[1].Coordinates[0].Lng, -8101101.950116)

	assert.Equal(t, coord[1].Coordinates[1].Lat, 1487158.7652179995)
	assert.Equal(t, coord[1].Coordinates[1].Lng, -9196903.187613)

	assert.Equal(t, coord[1].Coordinates[2].Lat, 1095801.2846229984)
	assert.Equal(t, coord[1].Coordinates[2].Lng, -5518141.890304)

	assert.Equal(t, coord[1].Coordinates[3].Lat, -1252344.285755001)
	assert.Equal(t, coord[1].Coordinates[3].Lng, -9392582.035682)

	assert.Equal(t, coord[1].Coordinates[4].Lat, -7474929.9346210025)
	assert.Equal(t, coord[1].Coordinates[4].Lng, -3287403.712489)

}

func TestConvertToMercatorPassedMeridian(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84PassedMeridian)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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
	assert.Equal(t, coord[0].Coordinates[0].Lat, -2700367.3196590007)
	assert.Equal(t, coord[0].Coordinates[0].Lng, 12601714.231207)

	assert.Equal(t, coord[0].Coordinates[1].Lat, -4138606.4164760034)
	assert.Equal(t, coord[0].Coordinates[1].Lng, 12993071.816027)

	assert.Equal(t, coord[0].Coordinates[2].Lat, -3639625.540684)
	assert.Equal(t, coord[0].Coordinates[2].Lng, 14685693.356459)

	assert.Equal(t, coord[0].Coordinates[3].Lat, -4637587.344467)
	assert.Equal(t, coord[0].Coordinates[3].Lng, 15742358.891133)

	assert.Equal(t, coord[0].Coordinates[4].Lat, -4627803.413133998)
	assert.Equal(t, coord[0].Coordinates[4].Lng, 16534857.930818997)

	assert.Equal(t, coord[0].Coordinates[5].Lat, -3169996.399371)
	assert.Equal(t, coord[0].Coordinates[5].Lng, 17141462.21512)

	assert.Equal(t, coord[0].Coordinates[6].Lat, -1203424.5372380016)
	assert.Equal(t, coord[0].Coordinates[6].Lng, 15810846.426732002)

	assert.Equal(t, coord[0].Coordinates[7].Lat, -1986139.800958002)
	assert.Equal(t, coord[0].Coordinates[7].Lng, 15615167.689982003)

	assert.Equal(t, coord[0].Coordinates[8].Lat, -1692621.5824070002)
	assert.Equal(t, coord[0].Coordinates[8].Lng, 15096618.792691)

	assert.Equal(t, coord[0].Coordinates[9].Lat, -1340399.6832170007)
	assert.Equal(t, coord[0].Coordinates[9].Lng, 15214026.123796)

	assert.Equal(t, coord[0].Coordinates[10].Lat, -1271912.1821860003)
	assert.Equal(t, coord[0].Coordinates[10].Lng, 14617205.820861)

	assert.Equal(t, coord[0].Coordinates[11].Lat, -1917652.163291)
	assert.Equal(t, coord[0].Coordinates[11].Lng, 13609459.970374001)

	assert.Equal(t, coord[0].Coordinates[12].Lat, -2201386.426958)
	assert.Equal(t, coord[0].Coordinates[12].Lng, 13521404.583364)

	assert.Equal(t, coord[0].Coordinates[13].Lat, -2700367.3196590007)
	assert.Equal(t, coord[0].Coordinates[13].Lng, 12601714.231207)
}

func TestConvertToMercatorPassedMeridian2(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84PassedMeridian2)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(f)
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

	assert.Equal(t, coord[0].Coordinates[0].Lat, 1271912.1821859663)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -8179373.467080395)

	assert.Equal(t, coord[0].Coordinates[1].Lat, -410925.44809296844)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -9099063.819237508)

	assert.Equal(t, coord[0].Coordinates[2].Lat, -1682837.6291183494)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -8394620.166561324)

	assert.Equal(t, coord[0].Coordinates[3].Lat, -2172034.5726313316)
	assert.Equal(t, coord[0].Coordinates[3].Lng, -7827151.696402048)

	assert.Equal(t, coord[0].Coordinates[4].Lat, -6535671.749414185)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -8492459.534936476)

	assert.Equal(t, coord[0].Coordinates[5].Lat, -7494497.668144476)
	assert.Equal(t, coord[0].Coordinates[5].Lng, -7651040.811062874)

	assert.Equal(t, coord[0].Coordinates[6].Lat, -7298818.963453764)
	assert.Equal(t, coord[0].Coordinates[6].Lng, -7279251.077653904)

	assert.Equal(t, coord[0].Coordinates[7].Lat, -6535671.749414185)
	assert.Equal(t, coord[0].Coordinates[7].Lng, -7670608.662474007)

	assert.Equal(t, coord[0].Coordinates[8].Lat, -4794130.456523035)
	assert.Equal(t, coord[0].Coordinates[8].Lng, -6985732.861208956)

	assert.Equal(t, coord[0].Coordinates[9].Lat, -4637587.344467449)
	assert.Equal(t, coord[0].Coordinates[9].Lng, -6457400.205191435)

	assert.Equal(t, coord[0].Coordinates[10].Lat, -3306971.5892318888)
	assert.Equal(t, coord[0].Coordinates[10].Lng, -5439870.373339678)

	assert.Equal(t, coord[0].Coordinates[11].Lat, -2954749.7193256505)
	assert.Equal(t, coord[0].Coordinates[11].Lng, -5420302.521928546)

	assert.Equal(t, coord[0].Coordinates[12].Lat, -2524256.4461500226)
	assert.Equal(t, coord[0].Coordinates[12].Lng, -4637587.352288341)

	assert.Equal(t, coord[0].Coordinates[13].Lat, -763147.3322926272)
	assert.Equal(t, coord[0].Coordinates[13].Lng, -3874440.0340592684)

	assert.Equal(t, coord[0].Coordinates[14].Lat, 117407.2818729388)
	assert.Equal(t, coord[0].Coordinates[14].Lng, -5635549.2214094745)

	assert.Equal(t, coord[0].Coordinates[15].Lat, 508764.9104400351)
	assert.Equal(t, coord[0].Coordinates[15].Lng, -5792092.255337516)

	assert.Equal(t, coord[0].Coordinates[16].Lat, 1213208.5147962673)
	assert.Equal(t, coord[0].Coordinates[16].Lng, -7122708.043725864)

	assert.Equal(t, coord[0].Coordinates[17].Lat, 1271912.1821859663)
	assert.Equal(t, coord[0].Coordinates[17].Lng, -8179373.467080395)
}

func TestConvertToMercatorFeatureCollection(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84FeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	fc, err := feature.CollectionFromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(fc)
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

	assert.Equal(t, coord[0].Coordinates[0].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[0].Lng, -12913060.932019735)

	assert.Equal(t, coord[0].Coordinates[1].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[1].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[2].Lat, 9876845.895794801)
	assert.Equal(t, coord[0].Coordinates[2].Lng, -10018754.171394622)

	assert.Equal(t, coord[0].Coordinates[4].Lat, 7967317.535015907)
	assert.Equal(t, coord[0].Coordinates[4].Lng, -12913060.932019735)

	coord2, ok2 := coords[1].Geometry.Coordinates.([]geometry.Point)

	if !ok2 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord2[0].Lat, 3991847.410437685)
	assert.Equal(t, coord2[0].Lng, -2269873.991956594)

	assert.Equal(t, coord2[1].Lat, 6026906.856034012)
	assert.Equal(t, coord2[1].Lng, -391357.5848201024)

	assert.Equal(t, coord2[2].Lat, 1917652.1632908355)
	assert.Equal(t, coord2[2].Lng, 1565430.3392804095)

	coord3, ok3 := coords[2].Geometry.Coordinates.(geometry.Point)

	if !ok3 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord3.Lat, 430493.3861772847)
	assert.Equal(t, coord3.Lng, -8492459.534936476)
}

func TestConvertToMercatorGeometryPoint(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	pt, err := f.ToPoint()
	if err != nil {
		t.Errorf("error %v", err)
	}

	r, err := ToMercator(pt)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Point)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}
	assert.Equal(t, k.Lat, 5012341.663847514)
	assert.Equal(t, k.Lng, -7903683.846322424)
}

func TestConvertToMercatorGeometryMultiPoint(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryMultiPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	mpt, err := f.ToMultiPoint()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToMercator(mpt)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiPoint)
	if !ok {
		t.
			Errorf("invalid geometry %v", err)
	}
	assert.Equal(t, k.Coordinates[0].Lat, 3991847.4104379974)
	assert.Equal(t, k.Coordinates[0].Lng, -2269873.991957)

	assert.Equal(t, k.Coordinates[1].Lat, 6026906.856034)
	assert.Equal(t, k.Coordinates[1].Lng, -391357.58482)

	assert.Equal(t, k.Coordinates[2].Lat, 1917652.1632909998)
	assert.Equal(t, k.Coordinates[2].Lng, 1565430.33928)
}

func TestConvertToMercatorGeometryPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	poly, err := f.ToPolygon()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToMercator(poly)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Polygon)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lat, 7967317.535015907)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lng, -12913060.932019735)

	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lat, 7967317.535015907)
	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lng, -10018754.171394622)

	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lat, 9876845.895794801)
	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lng, -10018754.171394622)

	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lat, 9876845.895794801)
	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lng, -12913060.932019735)

	assert.Equal(t, k.Coordinates[0].Coordinates[4].Lat, 7967317.535015907)
	assert.Equal(t, k.Coordinates[0].Coordinates[4].Lng, -12913060.932019735)
}

func TestConvertToMercatorGeometryMultiPolygon(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	poly, err := f.ToMultiPolygon()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToMercator(poly)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiPolygon)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[0].Lat, -5621521.486192067)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[0].Lng, -12913060.932019735)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[1].Lat, -5621521.486192067)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[1].Lng, -10018754.171394622)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[2].Lat, -7558415.656081783)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[2].Lng, -10018754.171394622)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[3].Lat, -7558415.656081783)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[3].Lng, -12913060.932019735)

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[4].Lat, -5621521.486192067)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Coordinates[4].Lng, -12913060.932019735)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[0].Lat, 1017529.7499880107)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[0].Lng, -10057889.985536376)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[1].Lat, -391357.57972990983)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[1].Lng, -8649002.568864517)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[2].Lat, 1369751.5250587354)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[2].Lng, -7240115.374831639)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[3].Lat, 2465552.7440450294)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[3].Lng, -7318386.89179566)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[4].Lat, 2817774.6324118003)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Coordinates[4].Lng, -9353446.221540703)
}

func TestConvertToMercatorGeometryLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	lns, err := f.ToLineString()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToMercator(lns)
	assert.Nil(t, err)

	k, ok := r.(*geometry.LineString)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Lat, 3991847.4104379974)
	assert.Equal(t, k.Coordinates[0].Lng, -2269873.991957)

	assert.Equal(t, k.Coordinates[1].Lat, 6026906.856034)
	assert.Equal(t, k.Coordinates[1].Lng, -391357.58482)

	assert.Equal(t, k.Coordinates[2].Lat, 1917652.1632909998)
	assert.Equal(t, k.Coordinates[2].Lng, 1565430.33928)

}

func TestConvertToMercatorGeometryMultiLineString(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryMultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := geometry.FromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	mlns, err := f.ToMultiLineString()
	if err != nil {
		t.Errorf("conversion error %v", err)
	}

	r, err := ToMercator(mlns)
	assert.Nil(t, err)

	k, ok := r.(*geometry.MultiLineString)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lat, -156543.0522530012)
	assert.Equal(t, k.Coordinates[0].Coordinates[0].Lng, -14245416.087452)

	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lat, 3952711.5619209986)
	assert.Equal(t, k.Coordinates[0].Coordinates[1].Lng, -10136161.391181)

	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lat, 7866287.4827539995)
	assert.Equal(t, k.Coordinates[0].Coordinates[2].Lng, -7983694.730329999)

	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lat, 9783939.750186)
	assert.Equal(t, k.Coordinates[0].Coordinates[3].Lng, -3287403.712489)

	assert.Equal(t, k.Coordinates[1].Coordinates[0].Lat, 3443946.8023359994)
	assert.Equal(t, k.Coordinates[1].Coordinates[0].Lng, -8101101.950116)

	assert.Equal(t, k.Coordinates[1].Coordinates[1].Lat, 1487158.7652179995)
	assert.Equal(t, k.Coordinates[1].Coordinates[1].Lng, -9196903.187613)

	assert.Equal(t, k.Coordinates[1].Coordinates[2].Lat, 1095801.2846229984)
	assert.Equal(t, k.Coordinates[1].Coordinates[2].Lng, -5518141.890304)

	assert.Equal(t, k.Coordinates[1].Coordinates[3].Lat, -1252344.285755001)
	assert.Equal(t, k.Coordinates[1].Coordinates[3].Lng, -9392582.035682)

	assert.Equal(t, k.Coordinates[1].Coordinates[4].Lat, -7474929.9346210025)
	assert.Equal(t, k.Coordinates[1].Coordinates[4].Lng, -3287403.712489)

}

func TestConvertToMercatorGeometryCollection(t *testing.T) {
	p, err := utils.LoadJSONFixture(WGS84GeometryCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	collection, err := geometry.CollectionFromJSON(p)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}

	r, err := ToMercator(collection)
	assert.Nil(t, err)

	k, ok := r.(*geometry.Collection)
	if !ok {
		t.Errorf("invalid geometry %v", err)
	}

	coords := k.Geometries

	coord, ok := coords[0].Coordinates.(geometry.Point)
	if !ok {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord.Lat, 5012341.663847514)
	assert.Equal(t, coord.Lng, -7903683.846322424)

	coord2, ok2 := coords[1].Coordinates.([]geometry.Point)

	if !ok2 {
		t.Errorf("invalid feature %v", err)
	}

	assert.Equal(t, coord2[0].Lat, 3991847.4104379974)
	assert.Equal(t, coord2[0].Lng, -2269873.991957)

	assert.Equal(t, coord2[1].Lat, 6026906.856034)
	assert.Equal(t, coord2[1].Lng, -391357.58482)

	assert.Equal(t, coord2[2].Lat, 1917652.1632909998)
	assert.Equal(t, coord2[2].Lng, 1565430.33928)
}
