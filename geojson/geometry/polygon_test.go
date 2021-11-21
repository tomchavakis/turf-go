package geometry

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewPolygon(t *testing.T) {
	type args struct {
		pts []Point
	}
	tests := map[string]struct {
		args    args
		want    *Polygon
		wantErr bool
		err     error
	}{
		"error polygon with less than 4 coordinates": {
			args: args{
				pts: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(35.55, 23.44)},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("a polygon must have at least 4 positions"),
		},
		"error polygon closed": {
			args: args{
				pts: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.50, 23.45), *NewPoint(35.55, 23.45)},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("all elements of a polygon must be closed linestrings"),
		},
		"new polygon": {
			args: args{
				pts: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)},
			},
			want: &Polygon{Coordinates: []LineString{
				{
					Coordinates: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)},
				},
			}},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			poly, err := NewPolygon([]LineString{
				{
					Coordinates: tt.args.pts,
				},
			})

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("NewPolygon() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := poly; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}
