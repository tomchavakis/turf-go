package turf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/turf-go/geojson"
	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"github.com/tomchavakis/turf-go/utils"
)

const PolyWithHoleFixture = "poly-with-hole.json"

func TestPointInPolygon(t *testing.T) {
	type args struct {
		point   geometry.Point
		polygon geometry.Polygon
	}

	poly := geometry.Polygon{
		Coordinates: []geometry.LineString{
			{
				Coordinates: []geometry.Position{
					{
						Latitude:  36.171278341935434,
						Longitude: -86.76624298095703,
					},
					{
						Latitude:  36.170862616662134,
						Longitude: -86.74238204956055,
					},
					{
						Latitude:  36.19607929145354,
						Longitude: -86.74100875854492,
					},
					{
						Latitude:  36.2014818084173,
						Longitude: -86.77362442016602,
					},
					{
						Latitude:  36.171278341935434,
						Longitude: -86.76624298095703,
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
	coords := []geometry.Position{
		{
			Altitude:  nil,
			Latitude:  0,
			Longitude: 0,
		},
		{
			Altitude:  nil,
			Latitude:  0,
			Longitude: 100,
		},
		{
			Altitude:  nil,
			Latitude:  100,
			Longitude: 100,
		},
		{
			Altitude:  nil,
			Latitude:  100,
			Longitude: 0,
		},
		{
			Altitude:  nil,
			Latitude:  0,
			Longitude: 0,
		},
	}

	ml := []geometry.LineString{}
	ln, err := geometry.NewLineString(coords)
	assert.Nil(t, err, "error message %s", err)
	ml = append(ml, *ln)

	poly, err := geometry.NewPolygon(ml)
	assert.Nil(t, err, "error message %s", err)

	ptIn := geometry.Point{
		Lat: 50,
		Lng: 50,
	}

	ptOut := geometry.Point{
		Lat: 140,
		Lng: 150,
	}

	pip, err := PointInPolygon(ptIn, *poly)
	assert.Nil(t, err, "error message %s", err)
	assert.True(t, pip, "Point in not in Polygon")

	pop, err := PointInPolygon(ptOut, *poly)
	assert.Nil(t, err, "error message %s", err)
	assert.False(t, pop, "Point in not in Polygon")
}

func TestPolyWithHole(t *testing.T) {
	// ptInHole := geometry.Point{
	// 	Lat: 36.20373274711739,
	// 	Lng: -86.69208526611328,
	// }
	// ptInPoly := geometry.Point{
	// 	Lat: 36.20258997094334,
	// 	Lng: -86.72229766845702,
	// }
	// ptOutsidePoly := geometry.Point{
	// 	Lat: 36.18527313913089,
	// 	Lng: -86.75079345703125,
	// }

	fix, err := utils.LoadJSONFixture(PolyWithHoleFixture)
	assert.NoError(t, err, "error loading fixture")

	f, err := feature.FromJSON(fix)
	assert.NoError(t, err, "error decoding json to feature")
	assert.NotNil(t, f, "feature is nil")

	assert.Equal(t, f.Type, geojson.Feature, "invalid base type")
	props := map[string]interface{}{
		"name":     "Poly with Hole",
		"value":    float64(3),
		"filename": "poly-with-hole.json",
	}
	assert.Equal(t, f.Properties, props, "invalid properties")
	assert.Equal(t, f.Bbox, []float64{-86.73980712890625, 36.173495506147, -86.67303085327148, 36.23084281427824}, "invalid properties object")
	assert.Equal(t, f.Geometry.GeoJSONType, geojson.Polygon, "invalid geojson type")
	assert.Equal(t, len(f.Geometry.Coordinates), 2, "invalid coordinates")

	// PointInPolygon(ptInHole, polyHole.Geometry)
	// assert.False()

}
