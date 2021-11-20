package geometry

import (
	"reflect"
	"testing"

	"github.com/tomchavakis/turf-go/internal/common"
)

func TestNewPosition(t *testing.T) {
	type args struct {
		altitude  *float64
		latitude  float64
		longitude float64
	}
	tests := map[string]struct {
		args    args
		want    *Position
		wantErr bool
		err     error
	}{
		"new position - no altitude": {
			args: args{
				altitude:  nil,
				latitude:  34.7,
				longitude: 2.0,
			},
			want: &Position{
				Altitude:  nil,
				Latitude:  34.7,
				Longitude: 2.0,
			},
			wantErr: false,
			err:     nil,
		},
		"new position": {
			args: args{
				altitude:  common.Float64Ptr(40.5),
				latitude:  34.7,
				longitude: 2.0,
			},
			want: &Position{
				Altitude:  common.Float64Ptr(40.5),
				Latitude:  34.7,
				Longitude: 2.0,
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			pos := NewPosition(tt.args.altitude, tt.args.latitude, tt.args.longitude)

			if got := pos; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewPosition() = %v, want %v", got, tt.want)
			}

			p := pos.ToPoint()
			if p.Lat != tt.args.latitude {
				t.Errorf("Latitude point = %v, want %v", p.Lat, tt.args.latitude)
			}

			if p.Lng != tt.args.longitude {
				t.Errorf("longitude point = %v, want %v", p.Lng, tt.args.longitude)
			}
		})
	}
}
