package geometry

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewMultiPoint(t *testing.T) {
	type args struct {
		pts []Point
	}
	tests := map[string]struct {
		args    args
		want    *MultiPoint
		wantErr bool
		err     error
	}{
		"error multipoint with less than 2 coordinates": {
			args: args{
				pts: []Point{*NewPoint(35.55, 23.44)},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("a MultiPoint must have at least two or more positions"),
		},
		"new multipoint": {
			args: args{
				pts: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)},
			},
			want:    &MultiPoint{Coordinates: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)}},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			mpt, err := NewMultiPoint(tt.args.pts)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestNewMultiPoint() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := mpt; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewMultiPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
