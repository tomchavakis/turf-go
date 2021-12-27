package measurement

import (
	"errors"
	"reflect"
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/constants"
	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/internal/common"
	"github.com/tomchavakis/turf-go/utils"
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
const ImbalancedPolygon = "../test-data/imbalanced-polygon.json"
const BBoxPoint = "../test-data/bbox-point.json"
const BBoxMultiPoint = "../test-data/bbox-multipoint.json"
const BBoxLineString = "../test-data/bbox-linestring.json"
const BBoxPolygonLineString = "../test-data/bbox-polygon-linestring.json"
const BBoxPoly = "../test-data/bbox-polygon.json"
const BBoxMultiLineString = "../test-data/bbox-multilinestring.json"
const BBoxMultiPolygon = "../test-data/bbox-multipolygon.json"
const BBoxGeometryMultiPolygon = "../test-data/bbox-geometry-multipolygon.json"
const AlongDCLine = "../test-data/along-dc-line.json"

func TestDistance(t *testing.T) {
	d, err := Distance(-75.343, 39.984, -75.534, 39.123, constants.UnitMiles)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 60.35329997171416)

	d, err = Distance(-75.343, 39.984, -75.534, 39.123, constants.UnitNauticalMiles)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 52.445583795722655)

	d, err = Distance(-75.343, 39.984, -75.534, 39.123, constants.UnitKilometers)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 97.12922118967835)

	d, err = Distance(-75.343, 39.984, -75.534, 39.123, constants.UnitRadians)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 0.015245501024842149)

	d, err = Distance(-75.343, 39.984, -75.534, 39.123, constants.UnitDegrees)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 0.8724834600465156)
}

func TestPointDistance(t *testing.T) {
	p1 := geometry.Point{Lng: -75.343, Lat: 39.984}
	p2 := geometry.Point{Lng: -75.534, Lat: 39.123}
	d, err := PointDistance(p1, p2, constants.UnitDefault)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, d, 97.12922118967835)
}

func TestBearing(t *testing.T) {
	b := Bearing(-77.03653, 38.89768, -77.05173, 38.8973)
	assert.Equal(t, b, 268.16492117999513)
}

func TestPointBearing(t *testing.T) {
	p1 := geometry.Point{Lng: -77.03653, Lat: 38.89768}
	p2 := geometry.Point{Lng: -77.05173, Lat: 38.8973}
	b := PointBearing(p1, p2)
	assert.Equal(t, b, 268.16492117999513)
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
	d, err := Destination(p, 10, 230, constants.UnitDefault)
	if err != nil {
		t.Errorf("Destination error %v", err)
	}
	e := geometry.Point{Lat: 23.282174951509955, Lng: 43.17500084522403}

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
	assert.Equal(t, err, nil)

	d, err := Length(*ln, constants.UnitDefault)
	if err != nil {
		t.Errorf("Length error %v", err)
	}
	assert.Equal(t, d, 0.0)
}

func TestLineDistanceWithGeometries(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistanceRouteOne)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}
	gjson2, err := utils.LoadJSONFixture(LineDistanceRouteTwo)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	feature1, err := feature.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	feature2, err := feature.FromJSON(gjson2)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

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

	if !reflect.DeepEqual(feature1.Properties, props) {
		t.Errorf("invalid properties")
	}

	assert.Equal(t, feature1.Geometry.GeoJSONType, geojson.LineString)
	if !reflect.DeepEqual(feature2.Properties, props) {
		t.Errorf("invalid properties")
	}
	assert.Equal(t, feature2.Geometry.GeoJSONType, geojson.LineString)

	ls1, err := feature1.ToLineString()
	if err != nil {
		t.Errorf("ToLineString error: %v", err)
	}

	ls2, err := feature2.ToLineString()
	if err != nil {
		t.Errorf("ToLineString error: %v", err)
	}

	l1, err := Length(*ls1, constants.UnitDefault)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	l2, err := Length(*ls2, constants.UnitDefault)
	if err != nil {
		t.Errorf("distance error %v", err)
	}

	assert.Equal(t, l1, 325.737252622811)
	assert.Equal(t, l2, 741.5469760360743)
}

func TestLineDistancePolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistancePolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	feature, err := feature.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	polygon, err := feature.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}
	l, err := Length(*polygon, constants.UnitDefault)
	if err != nil {
		t.Errorf("Length error %v", err)
	}
	assert.Equal(t, l, 5.597322420589979)
}

func TestLineDistanceMultiLineString(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistanceMultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	fs, err := feature.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	mls, err := fs.ToMultiLineString()
	if err != nil {
		t.Errorf("ToMultiLineString error: %v", err)
	}
	l, err := Length(*mls, constants.UnitDefault)
	if err != nil {
		t.Errorf("distance error %v", err)
	}
	assert.Equal(t, l, 4.703841298351085)
}

func TestAreaPolygonAsFeature(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	feature, err := feature.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}
	area, err := Area(feature)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 7748891609977)
}

func TestAreaMultiPolygonAsFeature(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	feature, err := feature.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	area, err := Area(feature)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 24716139112)
}

func TestAreaPolygonAsGeometry(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	geom, err := geometry.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	area, err := Area(geom)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 10993362)

}

func TestAreaPolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	geom, err := geometry.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	poly, err := geom.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}

	area, err := Area(poly)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 10993362)
}

func TestAreaMultiPolygon(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaGeomMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	geometry, err := geometry.FromJSON(gjson1)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	multiPoly, err := geometry.ToMultiPolygon()
	if err != nil {
		t.Errorf("ToMultiPolygon error: %v", err)
	}

	area, err := Area(multiPoly)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 24716139112)
}

func TestAreaFeatureCollection(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(AreaFeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	collection, err := feature.CollectionFromJSON(gjson1)
	if err != nil {
		t.Errorf("CollectionFromJSON error: %v", err)
	}

	area, err := Area(collection)
	if err != nil {
		t.Errorf("Area error: %v", err)
	}

	assert.Equal(t, int(area), 294193686165)
}

func TestBBoxPoint(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	p, err := f.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	bbox, err := BBox(p)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)
	assert.Equal(t, bbox[0], 102.0)
	assert.Equal(t, bbox[1], 0.5)
	assert.Equal(t, bbox[2], 102.0)
	assert.Equal(t, bbox[3], 0.5)
}

func TestBBoxLineString(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	l, err := f.ToLineString()
	if err != nil {
		t.Errorf("ToLineString error: %v", err)
	}

	bbox, err := BBox(l)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 102.0)
	assert.Equal(t, bbox[1], -10.0)
	assert.Equal(t, bbox[2], 130.0)
	assert.Equal(t, bbox[3], 4.0)
}

func TestBBoxPoly(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxPoly)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	p, err := f.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}

	bbox, err := BBox(p)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 101.0)
	assert.Equal(t, bbox[3], 1.0)
}

func TestMultiLineString(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxMultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	ml, err := f.ToMultiLineString()
	if err != nil {
		t.Errorf("ToMultiLineString error: %v", err)
	}

	bbox, err := BBox(ml)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestMultiPolygon(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	mpoly, err := f.ToMultiPolygon()
	if err != nil {
		t.Errorf("ToMultiPolygon error: %v", err)
	}

	bbox, err := BBox(mpoly)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestBBoxPolygonFromLineString(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxPolygonLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	g, err := geometry.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	ln, err := g.ToLineString()
	if err != nil {
		t.Errorf("ToLineString error: %v", err)
	}

	// Use the lineString object to calculate its bounding area
	bbox, err := BBox(ln)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	// Use the boundingBox coordinates to create an actual BoundingBox object
	boudingBox := geojson.BBOX{
		West:  bbox[1],
		South: bbox[0],
		East:  bbox[3],
		North: bbox[2],
	}
	f, err := BBoxPolygon(boudingBox, "")
	if err != nil {
		t.Errorf("BBoxPolygon error: %v", err)
	}

	if f == nil {
		t.Error("bboxPolygon is nil")
	}

	poly, err := f.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}

	if poly == nil {
		t.Error("feature to polygon error")
	}
	if poly != nil {
		assert.Equal(t, len(poly.Coordinates[0].Coordinates), 5)
		assert.Equal(t, poly.Coordinates[0].Coordinates[0], geometry.Point{
			Lat: -10,
			Lng: 102,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[1], geometry.Point{
			Lat: 4,
			Lng: 102,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[2], geometry.Point{
			Lat: 4,
			Lng: 130,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[3], geometry.Point{
			Lat: -10,
			Lng: 130,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[4], geometry.Point{
			Lat: -10,
			Lng: 102,
		})
	}
}

func TestGeometry(t *testing.T) {
	gson, err := utils.LoadJSONFixture(BBoxGeometryMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	g, err := geometry.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	mpoly, err := g.ToMultiPolygon()
	if err != nil {
		t.Errorf("ToMultiPolygon error: %v", err)
	}

	bbox, err := BBox(mpoly)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], 0.0)
	assert.Equal(t, bbox[2], 103.0)
	assert.Equal(t, bbox[3], 3.0)
}

func TestGeometryCollection(t *testing.T) {
	geometries := []geometry.Geometry{}

	// Point
	gsonPoint, err := utils.LoadJSONFixture(BBoxPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	point, err := feature.FromJSON(gsonPoint)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// MultiPoint
	gsonMultiPoint, err := utils.LoadJSONFixture(BBoxMultiPoint)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	multiPoint, err := feature.FromJSON(gsonMultiPoint)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// LineString
	gsonLineString, err := utils.LoadJSONFixture(BBoxLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	linestring, err := feature.FromJSON(gsonLineString)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// MultiLineString
	gson, err := utils.LoadJSONFixture(BBoxMultiLineString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	multiLineString, err := feature.FromJSON(gson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// Polygon
	gsonPolygon, err := utils.LoadJSONFixture(BBoxPoly)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	poly, err := feature.FromJSON(gsonPolygon)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// MultiPolygon
	gsonMultiPolygon, err := utils.LoadJSONFixture(BBoxMultiPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	multiPoly, err := feature.FromJSON(gsonMultiPolygon)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	// geometries

	geometries = append(geometries, point.Geometry)
	geometries = append(geometries, multiPoint.Geometry)
	geometries = append(geometries, linestring.Geometry)
	geometries = append(geometries, multiLineString.Geometry)
	geometries = append(geometries, poly.Geometry)
	geometries = append(geometries, multiPoly.Geometry)

	gc, err := geometry.NewGeometryCollection(geometries)
	if err != nil {
		t.Errorf("NewGeometryCollection error: %v", err)
	}

	bbox, err := BBox(gc)
	if err != nil {
		t.Errorf("BBox error: %v", err)
	}

	assert.Equal(t, len(bbox), 4)

	assert.Equal(t, bbox[0], 100.0)
	assert.Equal(t, bbox[1], -10.0)
	assert.Equal(t, bbox[2], 130.0)
	assert.Equal(t, bbox[3], 4.0)
}

func TestAlong(t *testing.T) {

	gjson, err := utils.LoadJSONFixture(AlongDCLine)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	ln, err := f.ToLineString()
	if err != nil {
		t.Errorf("ToLineString error: %v", err)
	}

	p1, err := Along(*ln, 1.0, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p2, err := Along(*ln, 1.2, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p3, err := Along(*ln, 1.4, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p4, err := Along(*ln, 1.6, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p5, err := Along(*ln, 1.8, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p6, err := Along(*ln, 2.0, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p7, err := Along(*ln, 100, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}
	p8, err := Along(*ln, 0.0, constants.UnitDefault)
	if err != nil {
		t.Errorf("Along error %v", err)
	}

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

	assert.Equal(t, len(fc.Features), 8)
	p7f, err := fc.Features[7].Geometry.ToPoint()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	assert.Equal(t, p7f.Lng, p8.Lng)
	assert.Equal(t, p7f.Lat, p8.Lat)
}

func TestCenterFeature(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	ifs, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	coords := []float64{133.5, -27.0}
	g := geometry.Geometry{
		GeoJSONType: geojson.Point,
		Coordinates: coords,
	}

	ef, err := feature.New(g, nil, nil, "")
	ef.Bbox = []float64{
		113, -39, 154, -15,
	}
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cfs, err := CenterFeature(*ifs, nil, "")
	if err != nil {
		t.Errorf("CenterFeature error: %v", err)
	}

	if !reflect.DeepEqual(ef, cfs) {
		t.Error("Center Feature error")
	}
}

func TestCenterFeatureWithId(t *testing.T) {
	id := "testId"
	gjson, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	ifs, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}
	cfs, err := CenterFeature(*ifs, nil, id)
	if err != nil {
		t.Errorf("CenterFeature error: %v", err)
	}

	p, err := cfs.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	assert.Equal(t, p.Lng, 133.5)
	assert.Equal(t, p.Lat, -27.0)

}

func TestCenterFeatureWithProperties(t *testing.T) {
	properties := make(map[string]interface{})
	properties["key"] = "value"
	gjson, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	ifs, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	cfs, err := CenterFeature(*ifs, properties, "")
	if err != nil {
		t.Errorf("CenterFeature error: %v", err)
	}

	p, err := cfs.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	assert.Equal(t, p.Lng, 133.5)
	assert.Equal(t, p.Lat, -27.0)
	if cfs.Properties == nil {
		t.Errorf("properties cannot be empty")
	}
	if !reflect.DeepEqual(cfs.Properties, properties) {
		t.Error("properties are not equal")
	}
}

func TestCenterFeatureCollection(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(AreaFeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	fc, err := feature.CollectionFromJSON(gjson)
	if err != nil {
		t.Errorf("CollectionFromJSON error: %v", err)
	}
	cf, err := CenterFeatureCollection(*fc, nil, "")
	if err != nil {
		t.Errorf("CenterFeatureCollection error: %v", err)
	}

	p, err := cf.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	if p == nil {
		t.Error("point cannot be empty")
	}
	if p != nil {
		assert.Equal(t, p.Lng, 4.1748046875)
		assert.Equal(t, p.Lat, 47.214224817196836)
	}
}

func TestEnvelope(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(AreaFeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	fc, err := feature.CollectionFromJSON(gjson)
	if err != nil {
		t.Errorf("CollectionFromJSON error: %v", err)
	}
	f, err := Envelope(*fc)

	if err != nil {
		t.Errorf("Envelope error: %v", err)
	}

	if f == nil {
		t.Error("Polygon is nil")
	}

	poly, err := f.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}

	if poly == nil {
		t.Error("feature to polygon error")
	}
	if poly != nil {
		assert.Equal(t, len(poly.Coordinates[0].Coordinates), 5)
		assert.Equal(t, poly.Coordinates[0].Coordinates[0], geometry.Point{
			Lat: -3.515625,
			Lng: 49.83798245308484,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[1], geometry.Point{
			Lat: -3.515625,
			Lng: 44.59046718130883,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[2], geometry.Point{
			Lat: 11.865234375,
			Lng: 44.59046718130883,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[3], geometry.Point{
			Lat: 11.865234375,
			Lng: 49.83798245308484,
		})

		assert.Equal(t, poly.Coordinates[0].Coordinates[4], geometry.Point{
			Lat: -3.515625,
			Lng: 49.83798245308484,
		})
	}
}

func TestCentroidFeature(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FeatureFromJSON error: %v", err)
	}
	cf, err := CentroidFeature(*f, nil, "")
	if err != nil {
		t.Errorf("CentroidFeature error: %v", err)
	}

	p, err := cf.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	if p == nil {
		t.Error("point cannot be empty")
	}
	if p != nil {
		assert.Equal(t, p.Lng, 133.0)
		assert.Equal(t, p.Lat, -26.857142857142858)
	}
}

func TestImbalancedPolygonFeature(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(ImbalancedPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FeatureFromJSON error: %v", err)
	}
	cf, err := CentroidFeature(*f, nil, "")
	if err != nil {
		t.Errorf("CentroidFeature error: %v", err)
	}

	p, err := cf.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	if p == nil {
		t.Error("point cannot be empty")
	}
	if p != nil {
		assert.Equal(t, p.Lng, 4.851791984156558)
		assert.Equal(t, p.Lat, 45.78143055383553)
	}
}

func TestCentroidFeatureWithProperties(t *testing.T) {
	properties := make(map[string]interface{})
	properties["key"] = "value"
	gjson, err := utils.LoadJSONFixture(AreaPolygon)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	ifs, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}

	cfs, err := CentroidFeature(*ifs, properties, "")
	if err != nil {
		t.Errorf("CentroidFeature error: %v", err)
	}

	p, err := cfs.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	assert.Equal(t, p.Lng, 133.0)
	assert.Equal(t, p.Lat, -26.857142857142858)
	if cfs.Properties == nil {
		t.Errorf("properties cannot be empty")
	}
	if !reflect.DeepEqual(cfs.Properties, properties) {
		t.Error("properties are not equal")
	}
}

func TestCentroidFeatureCollection(t *testing.T) {
	gjson, err := utils.LoadJSONFixture(AreaFeatureCollection)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	fc, err := feature.CollectionFromJSON(gjson)
	if err != nil {
		t.Errorf("CollectionFromJSON error: %v", err)
	}
	cf, err := CentroidFeatureCollection(*fc, nil, "")
	if err != nil {
		t.Errorf("CentroidFeatureCollection error: %v", err)
	}

	p, err := cf.Geometry.ToPoint()
	if err != nil {
		t.Errorf("ToPoint error: %v", err)
	}

	if p == nil {
		t.Error("point cannot be empty")
	}
	if p != nil {
		assert.Equal(t, p.Lng, 6.774169921875)
		assert.Equal(t, p.Lat, 47.486422855836416)
	}
}

func TestRhumbBearing(t *testing.T) {
	type RhumbObj struct {
		start geometry.Point
		end   geometry.Point
		final bool
	}
	type args struct {
		rhumbObj RhumbObj
	}
	tests := map[string]struct {
		args    args
		want    *float64
		wantErr bool
		err     error
	}{
		"point - start Bearing": {
			args: args{
				rhumbObj: RhumbObj{
					start: geometry.Point{
						Lat: 45.0,
						Lng: -75.0,
					},
					end: geometry.Point{
						Lat: 60.0,
						Lng: 20.0,
					},
					final: false,
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(75.28061364784332),
			err:     nil,
		},
		"point - final Bearing": {
			args: args{
				rhumbObj: RhumbObj{
					start: geometry.Point{
						Lat: 45.0,
						Lng: -75.0,
					},
					end: geometry.Point{
						Lat: 60.0,
						Lng: 20.0,
					},
					final: true,
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(-104.7193863521567),
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			geo, err := RhumbBearing(tt.args.rhumbObj.start, tt.args.rhumbObj.end, tt.args.rhumbObj.final)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestRhumbBearing() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestRhumbBearing() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestRhumbDestination(t *testing.T) {
	type RhumbObj struct {
		geojson  string
		bearing  float64
		distance float64
		units    string
	}
	type args struct {
		rhumbObj RhumbObj
	}
	tests := map[string]struct {
		args    args
		want    *feature.Feature
		wantErr bool
		err     error
	}{
		"error - invalid units": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  -90.0,
					distance: 100.0,
					units:    "inv",
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": -90, \"dist\": 100 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-539.5, -16.5] } }",
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("invalid units"),
		},
		"point - negative distance": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  -90.0,
					distance: -1.0,
					units:    "",
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": -90, \"dist\": -1 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-539.5, -16.5] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{-539.499990620548, -16.5},
				},
			},
			err: nil,
		},
		"point - fiji east west 539 lng": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  -90.0,
					distance: 100.0,
					units:    "",
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": -90, \"dist\": 100 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-539.5, -16.5] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{-539.5009379451956, -16.5},
				},
			},
			err: nil,
		},
		"point - fiji east west": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  -90.0,
					distance: 100.0,
					units:    "",
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": -90, \"dist\": 100 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-179.5, -16.5] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{-179.50093794519552, -16.5},
				},
			},
			err: nil,
		},
		"point - fiji west east": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  120.0,
					distance: 150.0,
					units:    "",
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": 120, \"dist\": 150 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [179.5, -16.5] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{179.5012184286744, -16.500674490272797},
				},
			},
			err: nil,
		},
		"point - no bearing - no distance": {
			args: args{
				rhumbObj: RhumbObj{
					bearing: 0,
					units:   "",
					geojson: "{ \"type\": \"Feature\", \"properties\": { \"bearing\": 0 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-75, 38.10096062273525] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{-75, 38.10096062273525},
				},
			},
			err: nil,
		},
		"point - vertical bearing no distance": {
			args: args{
				rhumbObj: RhumbObj{
					bearing: 90.0,
					units:   "",
					geojson: "{ \"type\": \"Feature\", \"properties\": { \"bearing\": 90 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-75, 39] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{-75, 39},
				},
			},
			err: nil,
		},
		"point - big distance": {
			args: args{
				rhumbObj: RhumbObj{
					bearing:  90.0,
					distance: 5000,
					units:    constants.UnitMiles,
					geojson:  "{ \"type\": \"Feature\", \"properties\": { \"bearing\": 90, \"distance\": 5000 }, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-75, 39] } }",
				},
			},
			wantErr: false,
			want: &feature.Feature{
				ID:         "",
				Type:       "Feature",
				Properties: nil,
				Bbox:       []float64{},
				Geometry: geometry.Geometry{
					GeoJSONType: "Point",
					Coordinates: []float64{18.117374548567227, 39},
				},
			},
			err: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := feature.FromJSON(tt.args.rhumbObj.geojson)
			assert.NotNil(t, err)
			p, err := g.ToPoint()
			assert.NotNil(t, err)
			assert.NotNil(t, p)

			geo, err := RhumbDestination(*p, tt.args.rhumbObj.distance, tt.args.rhumbObj.bearing, tt.args.rhumbObj.units, nil)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestRhumbDestination() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestRhumbDestination() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestRhumbDistance(t *testing.T) {
	type RhumbObj struct {
		from geometry.Point
		to   geometry.Point
	}
	type args struct {
		rhumbObj RhumbObj
	}
	tests := map[string]struct {
		args    args
		want    *float64
		wantErr bool
		err     error
	}{
		"rhumbdistance - fiji 539": {
			args: args{
				rhumbObj: RhumbObj{
					from: geometry.Point{
						Lat: -16.5,
						Lng: -539.5,
					},
					to: geometry.Point{
						Lat: -18.5,
						Lng: -541.5,
					},
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(307.30629278030636),
			err:     nil,
		},
		"rhumbdistance - fiji": {
			args: args{
				rhumbObj: RhumbObj{
					from: geometry.Point{
						Lat: -16.5,
						Lng: -179.5,
					},
					to: geometry.Point{
						Lat: -16.5,
						Lng: 178.5,
					},
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(213.23207469632695),
			err:     nil,
		},
		"rhumbdistance - fiji 1": {
			args: args{
				rhumbObj: RhumbObj{
					from: geometry.Point{
						Lat: 39.984,
						Lng: -75.343,
					},
					to: geometry.Point{
						Lat: 39.123,
						Lng: -75.534,
					},
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(97.12923942772163),
			err:     nil,
		},
		"rhumbdistance - fiji 2": {
			args: args{
				rhumbObj: RhumbObj{
					from: geometry.Point{
						Lat: 35.60371874069731,
						Lng: -119.17968749999999,
					},
					to: geometry.Point{
						Lat: 46.92025531537451,
						Lng: -67.5,
					},
				},
			},
			wantErr: false,
			want:    common.Float64Ptr(4482.044244192197),
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			geo, err := RhumbDistance(tt.args.rhumbObj.from, tt.args.rhumbObj.to, "")

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestRhumbDistance() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestRhumbDistance() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestInBBox(t *testing.T)  {
	tests := map[string]struct {
		point   geometry.Point
		bbox    []float64
		want    bool
	}{
		"point in bbox": {
			point:	geometry.Point{
				Lng: 116.0,
				Lat: -20.0,
			},
			bbox: []float64{
				113.0, -39.0, 154.0, -15.0,
			},
			want: true,
		},
		"point on bbox": {
			point:	geometry.Point{
				Lng: 116.0,
				Lat: -20.0,
			},
			bbox: []float64{
				116.0, -20.0, 154.0, -15.0,
			},
			want: true,
		},
		"point not in bbox": {
			point:	geometry.Point{
				Lng: -20.0,
				Lat: 116.0,
			},
			bbox: []float64{
				113.0, -39.0, 154.0, -15.0,
			},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, inBBox(tt.point, tt.bbox), tt.want)
		})
	}
}


func TestInRing(t *testing.T) {
	tests := map[string]struct {
		point   		geometry.Point
		ring    		[]geometry.Point
		ignoreBoundary 	bool
		want    		bool
	} {
		"point in ring": {
			point: geometry.Point{
				Lng: 50.0,
				Lat: 50.0,
			},
			ring: []geometry.Point {
				{
					Lng: 0.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 0.0,
				},
			},
			ignoreBoundary: false,
			want: true,
		},
		"point not in ring": {
			point: geometry.Point{
				Lng: -50.0,
				Lat: -50.0,
			},
			ring: []geometry.Point {
				{
					Lng: 0.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 0.0,
				},
			},
			ignoreBoundary: false,
			want: false,
		},
		"point on ring": {
			point: geometry.Point{
				Lng: 0.0,
				Lat: 0.0,
			},
			ring: []geometry.Point {
				{
					Lng: 0.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 0.0,
				},
			},
			ignoreBoundary: false,
			want: true,
		},
		"point on ring, ignore boundary": {
			point: geometry.Point{
				Lng: 0.0,
				Lat: 0.0,
			},
			ring: []geometry.Point {
				{
					Lng: 0.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 100.0,
				},
				{
					Lng: 100.0,
					Lat: 0.0,
				},
				{
					Lng: 0.0,
					Lat: 0.0,
				},
			},
			ignoreBoundary: true,
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, inRing(tt.point, tt.ring, tt.ignoreBoundary), tt.want)
		})
	}
}


func TestInPolygon(t *testing.T) {
	fs, err := createFeature(t, AreaPolygon)
	if err != nil {
		t.Errorf("createFeature error = %v", err)
	}
	tests := map[string]struct {
		point   geometry.Point
		polygon    geometry.Geometry
		want    bool
	} {
		"point in polygon": {
			point: geometry.Point{
				Lng: 133.0,
				Lat: -26.857142857142858,
			},
			polygon: fs.Geometry,
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, pointOnPolygon(tt.point, tt.polygon), tt.want)
		})
	}
}


func TestPointOnFeature(t *testing.T) {
	fs, err := createFeature(t, AreaPolygon)
	if err != nil {
		t.Errorf("createFeature error = %v", err)
	}

	p, err := PointOnFeature(*fs, nil, "")
	if err != nil {
		t.Errorf("PointOnFeature error = %v", err)
	} else {
		if p == nil {
			t.Error("point cannot be empty")
		}
		if p != nil {
			assert.Equal(t, p.Lng, 133.0)
			assert.Equal(t, p.Lat, -26.857142857142858)
		}
		return
	}
}

func TestPointOnFeatureOfImbalancedPolygon(t *testing.T) {
	fs, err := createFeature(t, ImbalancedPolygon)
	if err != nil {
		t.Errorf("createFeature error = %v", err)
	}

	p, err := PointOnFeature(*fs, nil, "")
	if err != nil {
		t.Errorf("PointOnFeature error = %v", err)
	} else {
		if p == nil {
			t.Error("point cannot be empty")
		}
		if p != nil {
			assert.Equal(t, p.Lng, 4.851791984156558)
			assert.Equal(t, p.Lat, 45.78143055383553)
		}
		return
	}
}

func TestPointOnFeatureCollection(t *testing.T) {
	fc, err := createFeatureCollection(t, AreaFeatureCollection)
	if err != nil {
		t.Errorf("createFeatureCollection error: %v", err)
	}

	p, err := PointOnFeatureCollection(*fc, nil, "")
	if err != nil {
		t.Errorf("PointOnFeature error = %v", err)
	} else {
		if p == nil {
			t.Error("point cannot be empty")
		}
		if p != nil {
			assert.Equal(t, p.Lng, 6.774169921875)
			assert.Equal(t, p.Lat, 47.486422855836416)
		}
		return
	}
}


func createFeature(t *testing.T, geometryString string) (*feature.Feature, error) {
	gjson, err := utils.LoadJSONFixture(geometryString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}
	fs, err := feature.FromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}
	return fs, err
}

func createFeatureCollection(t *testing.T, geometryString string) (*feature.Collection, error) {
	gjson, err := utils.LoadJSONFixture(geometryString)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}
	fs, err := feature.CollectionFromJSON(gjson)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}
	return fs, err
}
