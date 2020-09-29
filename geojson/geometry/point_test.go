package geometry

import (
	"reflect"
	"testing"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		lat float64
		lng float64
	}
	tests := map[string]struct {
		args args
		want *Point
	}{
		"new point": {
			args: args{
				lat: 35.55,
				lng: 23.44,
			},
			want: &Point{
				Lat: 35.55,
				Lng: 23.44,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := NewPoint(tt.args.lat, tt.args.lng); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
