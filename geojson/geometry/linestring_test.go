package geometry

import (
	"reflect"
	"testing"
)

func TestNewLineString(t *testing.T) {
	type args struct {
		coordinates []Position
	}
	coords := []Position{
		{
			Altitude:  nil,
			Latitude:  34.44,
			Longitude: 23.53,
		},
		{
			Altitude:  nil,
			Latitude:  34.44,
			Longitude: 23.53,
		}}

	tests := map[string]struct {
		args    args
		want    *LineString
		wantErr bool
	}{
		"coordinates less than 2": {
			args: args{
				coordinates: []Position{
					{
						Altitude:  nil,
						Latitude:  34.44,
						Longitude: 23.53,
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
		coordinates []Position
	}
	closedCoords := []Position{
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
		{
			Altitude:  nil,
			Latitude:  52.3711451105601,
			Longitude: 4.895267486572266,
		},
		{
			Altitude:  nil,
			Latitude:  52.36931095278263,
			Longitude: 4.892091751098633,
		},
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
	}
	openCoords := []Position{
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
		{
			Altitude:  nil,
			Latitude:  52.3711451105601,
			Longitude: 4.895267486572266,
		},
		{
			Altitude:  nil,
			Latitude:  52.36931095278263,
			Longitude: 4.892091751098633,
		},
		{
			Altitude:  nil,
			Latitude:  52.3707258812113,
			Longitude: 4.8892593383789,
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
		coordinates []Position
	}
	closedCoords := []Position{
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
		{
			Altitude:  nil,
			Latitude:  52.3711451105601,
			Longitude: 4.895267486572266,
		},
		{
			Altitude:  nil,
			Latitude:  52.36931095278263,
			Longitude: 4.892091751098633,
		},
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
	}
	nonLinearCoords := []Position{
		{
			Altitude:  nil,
			Latitude:  52.370725881211314,
			Longitude: 4.889259338378906,
		},
		{
			Altitude:  nil,
			Latitude:  52.3711451105601,
			Longitude: 4.895267486572266,
		},
		{
			Altitude:  nil,
			Latitude:  52.36931095278263,
			Longitude: 4.892091751098633,
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
