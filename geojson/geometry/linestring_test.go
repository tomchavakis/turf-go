package geometry

import (
	"reflect"
	"testing"
)

func TestNewLineString(t *testing.T) {
	type args struct {
		coordinates []Point
	}
	coords := []Point{
		{
			Lat: 34.44,
			Lng: 23.53,
		},
		{
			Lat: 34.44,
			Lng: 23.53,
		}}

	tests := map[string]struct {
		args    args
		want    *LineString
		wantErr bool
	}{
		"coordinates less than 2": {
			args: args{
				coordinates: []Point{
					{
						Lat: 34.44,
						Lng: 23.53,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		"happy path": {
			args: args{
				coordinates: coords,
			},
			want:    &LineString{Coordinates: coords},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewLineString(tt.args.coordinates)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLineString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineString_IsClosed(t *testing.T) {
	type fields struct {
		coordinates []Point
	}
	closedCoords := []Point{
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
		{
			Lat: 52.3711451105601,
			Lng: 4.895267486572266,
		},
		{
			Lat: 52.36931095278263,
			Lng: 4.892091751098633,
		},
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
	}
	openCoords := []Point{
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
		{
			Lat: 52.3711451105601,
			Lng: 4.895267486572266,
		},
		{
			Lat: 52.36931095278263,
			Lng: 4.892091751098633,
		},
		{
			Lat: 52.3707258812113,
			Lng: 4.8892593383789,
		},
	}
	tests := map[string]struct {
		fields fields
		want   bool
	}{
		"open linestring": {
			fields: fields{
				coordinates: openCoords,
			},
			want: false,
		},
		"happy path": {
			fields: fields{
				coordinates: closedCoords,
			},
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := LineString{
				Coordinates: tt.fields.coordinates,
			}
			if got := l.IsClosed(); got != tt.want {
				t.Errorf("IsClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineString_IsLinearRing(t *testing.T) {
	type fields struct {
		coordinates []Point
	}
	closedCoords := []Point{
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
		{
			Lat: 52.3711451105601,
			Lng: 4.895267486572266,
		},
		{
			Lat: 52.36931095278263,
			Lng: 4.892091751098633,
		},
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
	}
	nonLinearCoords := []Point{
		{
			Lat: 52.370725881211314,
			Lng: 4.889259338378906,
		},
		{
			Lat: 52.3711451105601,
			Lng: 4.895267486572266,
		},
		{
			Lat: 52.36931095278263,
			Lng: 4.892091751098633,
		},
	}
	tests := map[string]struct {
		fields fields
		want   bool
	}{
		"non linear coords": {
			fields: fields{
				coordinates: nonLinearCoords,
			},
			want: false,
		},
		"happy path": {
			fields: fields{
				coordinates: closedCoords,
			},
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := LineString{
				Coordinates: tt.fields.coordinates,
			}
			if got := l.IsLinearRing(); got != tt.want {
				t.Errorf("IsLinearRing() = %v, want %v", got, tt.want)
			}
		})
	}
}
