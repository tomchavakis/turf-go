package invariant

import (
	"errors"
	"reflect"
	"testing"

	"github.com/tomchavakis/turf-go/geojson/feature"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestGetCoord(t *testing.T) {
	type args struct {
		coords interface{}
	}
	tests := map[string]struct {
		args    args
		want    []float64
		wantErr bool
		err     error
	}{
		"error - required coords": {
			args: args{
				coords: nil,
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("coord is required"),
		},
		"error  - polygon interface": {
			args: args{
				coords: feature.Feature{
					ID:         "",
					Type:       "Feature",
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: "Polygon",
						Coordinates: nil,
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("coord must be GeoJSON Point or an Array of numbers"),
		},
		"error array - single array ": {
			args: args{
				[]float64{44.34},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("coord must be GeoJSON Point or an Array of numbers"),
		},
		"feature  - point": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       "Feature",
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: "Point",
						Coordinates: []float64{44.34, 23.52},
					},
				},
			},
			wantErr: false,
			want:    []float64{44.34, 23.52},
			err:     nil,
		},
		"geometry - point": {
			args: args{
				coords: &geometry.Point{
					Lat: 23.52,
					Lng: 44.34,
				},
			},
			wantErr: false,
			want:    []float64{44.34, 23.52},
			err:     nil,
		},
		"array - point": {
			args: args{
				[]float64{44.34, 23.52},
			},
			wantErr: false,
			want:    []float64{44.34, 23.52},
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			geo, err := GetCoord(tt.args.coords)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestGetCoord() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestGetCoord() = %v, want %v", got, tt.want)
			}
		})
	}
}
