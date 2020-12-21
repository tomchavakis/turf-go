package measurement

import (
	"testing"

	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"

	"github.com/stretchr/testify/assert"
)

const LineDistanceRouteOne = "../test-data/route1.json"
const LineDistanceRouteTwo = "../test-data/route2.json"
const LineDistancePolygon = "../test-data/polygon.json"
const LineDistanceMultiLineString = "../test-data/multiLineString.json"
const AreaPolygon = "../test-data/area-polygon.json"
const AreaMultiPolygon = "../test-data/area-multipolygon.json"
const AreaGeomPolygon = "../test-data/area-geom-polygon.json"
const AreaGeomMultiPolygon = "../test-data/area-geom-multipolgon.json"
const AreaFeatureCollection = "../test-data/area-feature-collection.json"
const BBoxPoint = "../test-data/bbox-point.json"
const BBoxMultiPoint = "../test-data/bbox-multipoint.json"
const BBoxLineString = "../test-data/bbox-linestring.json"
const BBoxPolygon = "../test-data/bbox-polygon.json"
const BBoxMultiLineString = "../test-data/bbox-multilinestring.json"
const BBoxMultiPolygon = "../test-data/bbox-multipolygon.json"
const BBoxGeometryMultiPolygon = "../test-data/bbox-geometry-multipolygon.json"
const AlongDCLine = "../test-data/along-dc-line.json"

func TestDistance(t *testing.T) {
	d := Distance(-77.03653, 38.89768, -77.05173, 38.8973)
	assert.Equal(t, d, 1.317556974720262, "error calculating the distance")
}

func TestPointDistance(t *testing.T) {
	p1 := geometry.Point{Lng: -77.03653, Lat: 38.89768}
	p2 := geometry.Point{Lng: -77.05173, Lat: 38.8973}
	d := PointDistance(p1, p2)
	assert.Equal(t, d, 1.317556974720262, "error calculating the distance")
}

func TestBearing(t *testing.T) {
	b := Bearing(-77.03653, 38.89768, -77.05173, 38.8973)
	assert.Equal(t, b, 268.16492117999513, "error calculating the bearing")
}

func TestPointBearing(t *testing.T) {
	p1 := geometry.Point{Lng: -77.03653, Lat: 38.89768}
	p2 := geometry.Point{Lng: -77.05173, Lat: 38.8973}
	b := PointBearing(p1, p2)
	assert.Equal(t, b, 268.16492117999513, "error calculating the point bearing")
}

func TestMidPoint(t *testing.T) {

	type args struct {
		p1 geometry.Point
		p2 geometry.Point
	}

	tests := map[string]struct {
		args    args
		wantErr bool
		want    geometry.Point
	}{
		"happy path: same lng": {
			args: args{
				p1: geometry.Point{Lat: 23.38, Lng: 43.16},
				p2: geometry.Point{Lat: 26.38, Lng: 43.16},
			},
			wantErr: false,
			want:    geometry.Point{Lat: 24.88, Lng: 43.16},
		},
		"happy path: same lat": {
			args: args{
				p1: geometry.Point{Lat: 23.38, Lng: 43.20},
				p2: geometry.Point{Lat: 23.38, Lng: 44.20},
			},
			wantErr: false,
			want:    geometry.Point{Lat: 23.38079468036304, Lng: 43.699999999999996},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := MidPoint(tt.args.p1, tt.args.p2)
			if tt.want != m {
				t.Errorf("error calculating the midpoint")
				return
			}
		})
	}

}

func TestDestinationPoint(t *testing.T) {
	p := geometry.Point{Lat: 23.34, Lng: 43.25}
	d := Destination(p, 10, 230)
	e := geometry.Point{Lat: 23.28223959663299, Lng: 43.175084627817945}

	if e.Lat != d.Lat && e.Lng != d.Lng {
		t.Errorf("error calculating the destination point")
	}

}

func TestLineDistanceWhenRouteIsPoint(t *testing.T) {
	p1 := geometry.Point{
		Lat: 1.0,
		Lng: 1.0,
	}
	p2 := geometry.Point{
		Lat: 1.0,
		Lng: 1.0,
	}
	coords := []geometry.Point{}
	coords = append(coords, p1, p2)

	ln, err := geometry.NewLineString(coords)
	assert.NoError(t, err, "error initializing the lineString")
	d := Length(*ln)

	assert.Equal(t, d, 0.0)
}

func TestLineDistanceWithGeometries(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistanceRouteOne)
	assert.NoError(t, err, "cannot load geojson")
	gjson2, err := utils.LoadJSONFixture(LineDistanceRouteTwo)
	assert.NoError(t, err, "cannot load geojson")

	feature1, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error decoding geojson to feature")
	feature2, err := feature.FromJSON(gjson2)
	assert.NoError(t, err, "error decoding geojson to feature")

	props := map[string]interface{}{
		"name":       nil,
		"cmt":        nil,
		"desc":       nil,
		"src":        nil,
		"link1_href": nil,
		"link1_text": nil,
		"link1_type": nil,
		"link2_href": nil,
		"link2_text": nil,
		"link2_type": nil,
		"number":     nil,
		"type":       nil,
	}
	assert.Equal(t, feature1.Properties, props, "invalid properties")
	assert.Equal(t, feature1.Geometry.GeoJSONType, geojson.LineString, "invalid geojson type")
	assert.Equal(t, feature2.Properties, props, "invalid properties")
	assert.Equal(t, feature2.Geometry.GeoJSONType, geojson.LineString, "invalid geojson type")

	ls1, err := feature1.ToLineString()
	assert.NoError(t, err, "error converting feature to LineString")

	ls2, err := feature2.ToLineString()
	assert.NoError(t, err, "error converting feature to LineString")

	l1 := Length(*ls1)
	l2 := Length(*ls2)

	assert.Equal(t, l1, 326.10170358450773)
	assert.Equal(t, l2, 742.3766554982323)

}

func TestLineDistancePolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistancePolygon)
	assert.NoError(t, err, "cannot load polygon geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error decoding geojson to feature")

	polygon, err := feature.ToPolygon()
	assert.NoError(t, err, "error converting feature to polygon")
	l := Length(*polygon)
	assert.Equal(t, l, 5.603584981972479, "invalid length value")
}

func TestLineDistanceMultiLineString(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistanceMultiLineString)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error decoding geojson to feature")

	mls, err := feature.ToMultiLineString()
	assert.NoError(t, err, "error converting feature to multiLineString")
	l := Length(*mls)
	assert.Equal(t, l, 4.709104188828164, "invalid length value")
}

func TestAreaPolygonAsFeature(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaPolygon)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")
	area, err := Area(feature)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 7766240.997209013, "invalid area value")
}

func TestAreaMultiPolygonAsFeature(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaMultiPolygon)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")
	area, err := Area(feature)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 24771.477332558756, "invalid area value")
}

func TestAreaPolygonAsGeometry(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomPolygon)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	geom, err := geometry.FromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")

	area, err := Area(geom)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 11.017976596496059, "invalid area value")

}

func TestAreaPolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomPolygon)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	geom, err := geometry.FromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")

	poly, err := geom.ToPolygon()
	assert.NoError(t, err, "error while converting geometry to polygon")

	area, err := Area(poly)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 11.017976596496059, "invalid area value")
}

func TestAreaMultiPolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomMultiPolygon)
	assert.NoError(t, err, "cannot load multiLineString geojson")

	geometry, err := geometry.FromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")

	multiPoly, err := geometry.ToMultiPolygon()
	assert.Nil(t, err, "multiPolygon cannot be nil")

	area, err := Area(multiPoly)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 24771.477332558756, "invalid area value")
}

func TestAreaFeatureCollection(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaFeatureCollection)
	assert.NoError(t, err, "cannot load feature collection geojson")

	collection, err := feature.CollectionFromJSON(gjson1)
	assert.NoError(t, err, "error while decoding geojson to feature")

	area, err := Area(collection)
	assert.NoError(t, err, "error while computing geojson to feature")

	assert.Equal(t, area, 294852.3713607366, "invalid area value")
}

func TestBBoxPoint(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxPoint)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	p, err := f.ToPoint()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(p)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 102.0)
	assert.Equal(t, bbox[1], 0.5)
	assert.Equal(t, bbox[2], 102.0)
	assert.Equal(t, bbox[3], 0.5)
}

func TestBBoxLineString(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxLineString)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	l, err := f.ToLineString()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(l)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 102.0)
	assert.Equal(t, bbox[1], -10.0)
	assert.Equal(t, bbox[2], 130.0)
	assert.Equal(t, bbox[3], 4.0)
}

func TestBBoxPolygon(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxPolygon)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	p, err := f.ToPolygon()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(p)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 101.0)
	assert.Equal(t, bbox[3], 1.0)
}

func TestMultiLineString(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxMultiLineString)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	ml, err := f.ToMultiLineString()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(ml)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestMultiPolygon(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxMultiPolygon)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	mpoly, err := f.ToMultiPolygon()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(mpoly)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestGeometry(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxGeometryMultiPolygon)
	assert.NoError(t, err, "cannot load geojson")

	g, err := geometry.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	mpoly, err := g.ToMultiPolygon()
	assert.NoError(t, err, "error while converting feature")
	bbox, err := BBox(mpoly)
	assert.NoError(t, err, "bbox error")

	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestGeometryCollection(t *testing.T) {
	geometries := []geometry.Geometry{}

	// Point
	gsonPoint, err := utils.LoadJSONFixture(BBoxPoint)
	assert.NoError(t, err, "cannot load geojson")

	point, err := feature.FromJSON(gsonPoint)
	assert.NoError(t, err, "error while decoding geojson")

	// MultiPoint
	gsonMultiPoint, err := utils.LoadJSONFixture(BBoxMultiPoint)
	assert.NoError(t, err, "cannot load geojson")

	multiPoint, err := feature.FromJSON(gsonMultiPoint)
	assert.NoError(t, err, "error while decoding geojson")

	// LineString
	gsonLineString, err := utils.LoadJSONFixture(BBoxLineString)
	assert.NoError(t, err, "cannot load geojson")

	linestring, err := feature.FromJSON(gsonLineString)
	assert.NoError(t, err, "error while decoding geojson")

	// MultiLineString
	gson, err := utils.LoadJSONFixture(BBoxMultiLineString)
	assert.NoError(t, err, "cannot load geojson")

	multiLineString, err := feature.FromJSON(gson)
	assert.NoError(t, err, "error while decoding geojson")

	// Polygon
	gsonPolygon, err := utils.LoadJSONFixture(BBoxPolygon)
	assert.NoError(t, err, "cannot load geojson")

	poly, err := feature.FromJSON(gsonPolygon)
	assert.NoError(t, err, "error while decoding geojson")

	// MultiPolygon
	gsonMultiPolygon, err := utils.LoadJSONFixture(BBoxMultiPolygon)
	assert.NoError(t, err, "cannot load geojson")

	multiPoly, err := feature.FromJSON(gsonMultiPolygon)
	assert.NoError(t, err, "error while decoding geojson")

	// geometries

	geometries = append(geometries, point.Geometry)
	geometries = append(geometries, multiPoint.Geometry)
	geometries = append(geometries, linestring.Geometry)
	geometries = append(geometries, multiLineString.Geometry)
	geometries = append(geometries, poly.Geometry)
	geometries = append(geometries, multiPoly.Geometry)

	gc, err := geometry.NewGeometryCollection(geometries)
	assert.NoError(t, err, "cannot create a new geometry collection")

	bbox, err := BBox(gc)
	assert.NoError(t, err, "bbox error")
	assert.Equal(t, len(bbox), 4, "invalid bbox length")

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], -10.0)
	assert.Equal(t, bbox[2], 130.0)
	assert.Equal(t, bbox[3], 4.0)
}

func TestAlong(t *testing.T) {

	gjson, err := utils.LoadJSONFixture(AlongDCLine)
	assert.NoError(t, err, "cannot load geojson")

	f, err := feature.FromJSON(gjson)
	assert.NoError(t, err, "error loading geojson")

	ln, err := f.ToLineString()
	assert.NoError(t, err, "error converting to linestring")

	p1 := Along(*ln, 1.0)
	p2 := Along(*ln, 1.2)
	p3 := Along(*ln, 1.4)
	p4 := Along(*ln, 1.6)
	p5 := Along(*ln, 1.8)
	p6 := Along(*ln, 2.0)
	p7 := Along(*ln, 100)
	p8 := Along(*ln, 0.0)

	fc := feature.Collection{
		Type: geojson.FeatureCollection,
		Features: []feature.Feature{
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p1.Lng, p1.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p2.Lng, p2.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p3.Lng, p3.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p4.Lng, p4.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p5.Lng, p5.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p6.Lng, p6.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p7.Lng, p7.Lat,
					},
				},
			},
			{
				Type: geojson.Feature,
				Geometry: geometry.Geometry{
					GeoJSONType: geojson.Point,
					Coordinates: []float64{
						p8.Lng, p8.Lat,
					},
				},
			},
		},
	}

	assert.Equal(t, len(fc.Features), 8, "features error")
	p7f, err := fc.Features[7].Geometry.ToPoint()
	assert.NoError(t, err)
	assert.Equal(t, p7f.Lng, p8.Lng)
	assert.Equal(t, p7f.Lat, p8.Lat)
}
