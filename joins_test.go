package turf

import (
	"reflect"
	"testing"

	"github.com/tomchavakis/geojson"
	"github.com/tomchavakis/geojson/feature"
	"github.com/tomchavakis/geojson/geometry"
	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/utils"
)

const PolyWithHoleFixture = "test-data/poly-with-hole.json"
const MultiPolyWithHoleFixture = "test-data/multipoly-with-hole.json"

func TestPointInPolygon(t *testing.T) {
	type args struct {
		point   geometry.Point
		polygon geometry.Polygon
	}

	poly := geometry.Polygon{
		Coordinates: []geometry.LineString{
			{
				Coordinates: []geometry.Point{
					{
						Lat: 36.171278341935434,
						Lng: -86.76624298095703,
					},
					{
						Lat: 36.170862616662134,
						Lng: -86.74238204956055,
					},
					{
						Lat: 36.19607929145354,
						Lng: -86.74100875854492,
					},
					{
						Lat: 36.2014818084173,
						Lng: -86.77362442016602,
					},
					{
						Lat: 36.171278341935434,
						Lng: -86.76624298095703,
					},
				},
			},
		},
	}

	tests := map[string]struct {
		args    args
		want    bool
		wantErr bool
	}{
		"point in Polygon": {
			args: args{
				point: geometry.Point{
					Lat: 36.185411688981105,
					Lng: -86.76074981689453,
				},
				polygon: poly,
			},
			want:    true,
			wantErr: false,
		},
		"point in Polygon 2": {
			args: args{
				point: geometry.Point{
					Lat: 36.19393203374786,
					Lng: -86.75946235656737,
				},
				polygon: poly,
			},
			want:    true,
			wantErr: false,
		},
		"point out of Polygon": {
			args: args{
				point: geometry.Point{
					Lat: 36.18416473150645,
					Lng: -86.73036575317383,
				},
				polygon: poly,
			},
			want:    false,
			wantErr: false,
		},
		"point out of Polygon - really close to polygon": {
			args: args{
				point: geometry.Point{
					Lat: 36.18200632243299,
					Lng: -86.74175441265106,
				},
				polygon: poly,
			},
			want:    false,
			wantErr: false,
		},
		"point in Polygon - on boundary": {
			args: args{
				point: geometry.Point{
					Lat: 36.171278341935434,
					Lng: -86.76624298095703,
				},
				polygon: poly,
			},
			want:    true,
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := PointInPolygon(tt.args.point, tt.args.polygon)
			if (err != nil) != tt.wantErr {
				t.Errorf("PointInPolygon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PointInPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFeatureCollection(t *testing.T) {
	// test for a simple Polygon
	coords := []geometry.Point{
		{
			Lat: 0,
			Lng: 0,
		},
		{
			Lat: 0,
			Lng: 100,
		},
		{
			Lat: 100,
			Lng: 100,
		},
		{
			Lat: 100,
			Lng: 0,
		},
		{
			Lat: 0,
			Lng: 0,
		},
	}

	ml := []geometry.LineString{}
	ln, err := geometry.NewLineString(coords)
	if err != nil {
		t.Errorf("NewLineString error %v", err)
	}

	ml = append(ml, *ln)

	poly, err := geometry.NewPolygon(ml)
	if err != nil {
		t.Errorf("NewPolygon error %v", err)
	}

	ptIn := geometry.Point{
		Lat: 50,
		Lng: 50,
	}

	ptOut := geometry.Point{
		Lat: 140,
		Lng: 150,
	}

	pip, err := PointInPolygon(ptIn, *poly)
	if err != nil {
		t.Errorf("PointInPolygon error %v", err)
	}

	if !pip {
		t.Error("Point is not in Polygon")
	}

	pop, err := PointInPolygon(ptOut, *poly)
	if err != nil {
		t.Errorf("PointInPolygon error %v", err)
	}
	if pop {
		t.Error("Point is not in Polygon")
	}
}

func TestPolyWithHole(t *testing.T) {
	ptInHole := geometry.Point{
		Lat: 36.20373274711739,
		Lng: -86.69208526611328,
	}
	ptInPoly := geometry.Point{
		Lat: 36.20258997094334,
		Lng: -86.72229766845702,
	}
	ptOutsidePoly := geometry.Point{
		Lat: 36.18527313913089,
		Lng: -86.75079345703125,
	}

	fix, err := utils.LoadJSONFixture(PolyWithHoleFixture)
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(fix)
	if err != nil {
		t.Errorf("FromJSON error %v", err)
	}
	if f == nil {
		t.Error("feature cannot be nil")
	}

	if f != nil {
		assert.Equal(t, f.Type, geojson.Feature)
	}
	props := map[string]interface{}{
		"name":     "Poly with Hole",
		"value":    float64(3),
		"filename": "poly-with-hole.json",
	}
	if !reflect.DeepEqual(f.Properties, props) {
		t.Error("Properties are not equal")
	}
	if !reflect.DeepEqual(f.Bbox, []float64{-86.73980712890625, 36.173495506147, -86.67303085327148, 36.23084281427824}) {
		t.Error("BBOX error")
	}

	assert.Equal(t, f.Geometry.GeoJSONType, geojson.Polygon)

	poly, err := f.ToPolygon()
	if err != nil {
		t.Errorf("ToPolygon error: %v", err)
	}

	pih, err := PointInPolygon(ptInHole, *poly)
	if err != nil {
		t.Errorf("PointInPolygon error: %v", err)
	}
	if pih {
		t.Error("Point in hole is not in Polygon")
	}

	pip, err := PointInPolygon(ptInPoly, *poly)
	if err != nil {
		t.Errorf("PointInPolygon error: %v", err)
	}
	if !pip {
		t.Error("Point in poly is not in Polygon")
	}

	pop, err := PointInPolygon(ptOutsidePoly, *poly)
	if err != nil {
		t.Errorf("PointInPolygon error: %v", err)
	}
	if pop {
		t.Error("Point is not in Polygon")
	}
}

func TestMultiPolyWithHole(t *testing.T) {
	ptInHole := geometry.Point{
		Lat: 36.20373274711739,
		Lng: -86.69208526611328,
	}
	ptInPoly := geometry.Point{
		Lat: 36.20258997094334,
		Lng: -86.72229766845702,
	}
	ptInPoly2 := geometry.Point{
		Lat: 36.18527313913089,
		Lng: -86.75079345703125,
	}
	ptOutsidePoly := geometry.Point{
		Lat: 36.23015046460186,
		Lng: -86.75302505493164,
	}

	fixture, err := utils.LoadJSONFixture(MultiPolyWithHoleFixture)
	if err != nil {
		t.Errorf("LoadJSONFixture error: %v", err)
	}

	f, err := feature.FromJSON(fixture)
	if err != nil {
		t.Errorf("FromJSON error: %v", err)
	}
	if f == nil {
		t.Error("Feature cannot be nil")
	}
	if f != nil {
		assert.Equal(t, f.Type, geojson.Feature)
		props := map[string]interface{}{
			"name":     "Poly with Hole",
			"value":    float64(3),
			"filename": "poly-with-hole.json",
		}
		if !reflect.DeepEqual(f.Properties, props) {
			t.Error("Properties are not equal")
		}
		if !reflect.DeepEqual(f.Bbox, []float64{-86.77362442016602, 36.170862616662134, -86.67303085327148, 36.23084281427824}) {
			t.Error("BBOX error")
		}
		assert.Equal(t, f.Geometry.GeoJSONType, geojson.MultiPolygon)

		poly, err := f.ToMultiPolygon()
		if err != nil {
			t.Errorf("ToMultiPolygon error: %v", err)
		}

		pih := PointInMultiPolygon(ptInHole, *poly)
		if pih {
			t.Error("Point in hole is not in MultiPolygon")
		}

		pip := PointInMultiPolygon(ptInPoly, *poly)
		if !pip {
			t.Error("Point in poly is not in MultiPolygon")
		}

		pip2 := PointInMultiPolygon(ptInPoly2, *poly)
		if !pip2 {
			t.Error("Point in poly is not in MultiPolygon")
		}

		pop := PointInMultiPolygon(ptOutsidePoly, *poly)
		if pop {
			t.Error("Point in not in MultiPolygon")
		}
	}
}
