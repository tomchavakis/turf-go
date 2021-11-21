package geojson

import (
	"reflect"
	"testing"
)

func TestNewBoundingBox(t *testing.T) {
	type args struct {
		west  float64
		south float64
		east  float64
		north float64
	}
	tests := map[string]struct {
		args    args
		want    *BBOX
		wantErr bool
		err     error
	}{
		"new position - no altitude": {
			args: args{
				north: -10.0,
				west:  -10.0,
				east:  10.0,
				south: 10.0,
			},
			want: &BBOX{
				North: -10.0,
				West:  -10.0,
				East:  10.0,
				South: 10.0,
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			bbox := NewBBox(tt.args.west, tt.args.south, tt.args.east, tt.args.north)

			if got := bbox; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
