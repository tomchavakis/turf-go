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
	assert.NoError(t, err, "can't load geojson")
	gjson2, err := utils.LoadJSONFixture(LineDistanceRouteTwo)
	assert.NoError(t, err, "can't load geojson")

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
	assert.NoError(t, err, "can't load polygon geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error decoding geojson to feature")

	polygon, err := feature.ToPolygon()
	assert.NoError(t, err, "error converting feature to polygon")
	l := Length(*polygon)
	assert.Equal(t, l, 5.603584981972479, "invalid length value")
}

func TestLineDistanceMultiLineString(t *testing.T) {
	gjson1, err := utils.LoadJSONFixture(LineDistanceMultiLineString)
	assert.NoError(t, err, "can't load multiLineString geojson")

	feature, err := feature.FromJSON(gjson1)
	assert.NoError(t, err, "error decoding geojson to feature")

	mls, err := feature.ToMultiLineString()
	assert.NoError(t, err, "error converting feature to multiLineString")
	l := Length(*mls)
	assert.Equal(t, l, 4.709104188828164, "invalid length value")
}
