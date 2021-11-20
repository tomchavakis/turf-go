package geometry

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewMultiLineString(t *testing.T) {
	type args struct {
		lns []LineString
	}
	tests := map[string]struct {
		args    args
		want    *MultiLineString
		wantErr bool
		err     error
	}{
		"error multilinestring with less than 2 coordinates": {
			args: args{
				lns: []LineString{
					{
						[]Point{*NewPoint(35.55, 23.44)},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("a MultiLineString must have at least two or more linestrings"),
		},
		"new multilinestring": {
			args: args{
				lns: []LineString{
					{
						Coordinates: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)},
					},
					{
						Coordinates: []Point{*NewPoint(38.55, 23.44), *NewPoint(38.60, 23.60), *NewPoint(39.00, 24.04), *NewPoint(39.50, 24.67), *NewPoint(40.55, 23.44)},
					},
				},
			},
			want: &MultiLineString{
				[]LineString{
					{
						Coordinates: []Point{*NewPoint(35.55, 23.44), *NewPoint(36.60, 23.60), *NewPoint(37.00, 24.04), *NewPoint(37.50, 24.67), *NewPoint(35.55, 23.44)},
					},
					{
						Coordinates: []Point{*NewPoint(38.55, 23.44), *NewPoint(38.60, 23.60), *NewPoint(39.00, 24.04), *NewPoint(39.50, 24.67), *NewPoint(40.55, 23.44)},
					},
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			geo, err := NewMultiLineString(tt.args.lns)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestNewMultiLineString() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewMultiLineString() = %v, want %v", got, tt.want)
			}
		})
	}
}
