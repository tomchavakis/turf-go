package turf

import (
	"testing"

	"github.com/tomchavakis/turf-go/geojson/geometry"
)

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

//TODO: Add Multiple Test Cases
