package geometry

import (
	"errors"
	"reflect"
	"testing"
)

func TestPointFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *Point
		wantErr bool
		err     error
	}{
		"empty geojson": {
			args: args{
				geojson: "",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("input cannot be empty"),
		},
		"err geojson error": {
			args: args{
				geojson: "{ \"type\" : \"Point\", \"coordinates\": [-71, 41}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("cannot decode the input value"),
		},
		"error Point invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"MultiPoint\", \"coordinates\": [-71, 41]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"point geojson": {
			args: args{
				geojson: "{ \"type\" : \"Point\", \"coordinates\": [-71, 41] }",
			},
			want: &Point{
				Lat: 41.0,
				Lng: -71.0,
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestPointFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToPoint()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestPointFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestPointFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPointFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *MultiPoint
		wantErr bool
		err     error
	}{
		"error MultiPoint invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"LineString\", \"coordinates\":  [[ [102, -10],[103, 1],[104, 0],[130, 4] ]]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"MultiPoint geojson": {
			args: args{
				geojson: "{ \"type\" : \"MultiPoint\", \"coordinates\":  [[102, -10],[103, 1],[104, 0],[130, 4]]}",
			},
			want: &MultiPoint{
				Coordinates: []Point{
					{
						Lat: -10.0,
						Lng: 102.0,
					},
					{
						Lat: 1.0,
						Lng: 103.0,
					},
					{
						Lat: 0.0,
						Lng: 104.0,
					},
					{
						Lat: 4.0,
						Lng: 130.0,
					},
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiPointFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToMultiPoint()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiPointFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestMultiPointFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygonFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *Polygon
		wantErr bool
		err     error
	}{
		"error Polygon invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"LineString\", \"coordinates\":  [[ [102, -10],[103, 1],[104, 0],[130, 4] ]]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"Polygon non closed geojson": {
			args: args{
				geojson: "{ \"type\" : \"Polygon\", \"coordinates\":  [[ [102, -10],[103, 1],[104, 0],[130, 4] ]]}",
			},
			want: &Polygon{
				Coordinates: []LineString{{
					Coordinates: []Point{
						{
							Lat: -10.0,
							Lng: 102.0,
						},
						{
							Lat: 1.0,
							Lng: 103.0,
						},
						{
							Lat: 0.0,
							Lng: 104.0,
						},
						{
							Lat: 4.0,
							Lng: 130.0,
						},
					},
				}},
			},
			wantErr: true,
			err:     errors.New("cannot create a new polygon all elements of a polygon must be closed linestrings"),
		},
		"Polygon - geojson": {
			args: args{
				geojson: "{ \"type\" : \"Polygon\", \"coordinates\":  [[ [102, -10],[103, 1],[104, 0],[130, 4], [102, -10] ]]}",
			},
			want: &Polygon{
				Coordinates: []LineString{{
					Coordinates: []Point{
						{
							Lat: -10.0,
							Lng: 102.0,
						},
						{
							Lat: 1.0,
							Lng: 103.0,
						},
						{
							Lat: 0.0,
							Lng: 104.0,
						},
						{
							Lat: 4.0,
							Lng: 130.0,
						},
						{
							Lat: -10.0,
							Lng: 102.0,
						},
					},
				}},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestPolygonFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToPolygon()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestPolygonFromJson() error = %v, wantErr %v", err.Error(), tt.err.Error())
					return
				}
				return
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestPolygonFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygonFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *MultiPolygon
		wantErr bool
		err     error
	}{
		"error MultiPolygon invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"Polygon\", \"coordinates\":  [[[ [102, 2],[103, 2],[103, 3],[102, 3],[102, 2]]],[[[100.2, 0.2],[100.8, 0.2],[100.8, 0.8],[100.2, 0.8],[100.2, 0.2]]] ]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"MultiPolygon geojson": {
			args: args{
				geojson: "{ \"type\" : \"MultiPolygon\", \"coordinates\":  [[[ [102, 2],[103, 2],[103, 3],[102, 3],[102, 2]]],[[[100.2, 0.2],[100.8, 0.2],[100.8, 0.8],[100.2, 0.8],[100.2, 0.2]]] ]}",
			},
			want: &MultiPolygon{
				Coordinates: []Polygon{
					{
						[]LineString{
							{
								Coordinates: []Point{
									{
										Lat: 2.0,
										Lng: 102.0,
									},
									{
										Lat: 2.0,
										Lng: 103.0,
									},
									{
										Lat: 3.0,
										Lng: 103.0,
									},
									{
										Lat: 3.0,
										Lng: 102.0,
									},
									{
										Lat: 2.0,
										Lng: 102.0,
									},
								},
							},
						},
					},
					{
						[]LineString{
							{
								Coordinates: []Point{
									{
										Lat: 0.2,
										Lng: 100.2,
									},
									{
										Lat: 0.2,
										Lng: 100.8,
									},
									{
										Lat: 0.8,
										Lng: 100.8,
									},
									{
										Lat: 0.8,
										Lng: 100.2,
									},
									{
										Lat: 0.2,
										Lng: 100.2,
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiPolygonFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToMultiPolygon()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiPolygonFromJson() error = %v, wantErr %v", err.Error(), tt.err.Error())
					return
				}
				return
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestMultiPolygonFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineStringFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *LineString
		wantErr bool
		err     error
	}{
		"error LineString invalid geojson": {
			args: args{
				geojson: "{ \"type\" : \"LineString\", \"coordinates\": [ [ [102, -10],[103, 1],[104, 0],[130, 4] ]]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("cannot marshal object"),
		},
		"error LineString invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"Polygon\", \"coordinates\": [ [ [102, -10],[103, 1],[104, 0],[130, 4] ]]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"LineString - geojson": {
			args: args{
				geojson: "{ \"type\" : \"LineString\", \"coordinates\":  [ [102, -10],[103, 1],[104, 0],[130, 4] ]}",
			},
			want: &LineString{
				Coordinates: []Point{
					{
						Lat: -10.0,
						Lng: 102.0,
					},
					{
						Lat: 1.0,
						Lng: 103.0,
					},
					{
						Lat: 0.0,
						Lng: 104.0,
					},
					{
						Lat: 4.0,
						Lng: 130.0,
					},
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestLineStringFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToLineString()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestLineStringFromJson() error = %v, wantErr %v", err.Error(), tt.err.Error())
					return
				}
				return
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestLineStringFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiLineStringFromJson(t *testing.T) {
	type args struct {
		geojson string
	}
	tests := map[string]struct {
		args    args
		want    *MultiLineString
		wantErr bool
		err     error
	}{
		"error MultiLineString invalid type geojson": {
			args: args{
				geojson: "{ \"type\" : \"Polygon\", \"coordinates\":  [[[ [102, 2],[103, 2],[103, 3],[102, 3],[102, 2]]],[[[100.2, 0.2],[100.8, 0.2],[100.8, 0.8],[100.2, 0.8],[100.2, 0.2]]] ]}",
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("invalid geometry type"),
		},
		"MultiLineString geojson": {
			args: args{
				geojson: "{ \"type\" : \"MultiLineString\", \"coordinates\":  [[ [102, 2],[103, 2],[103, 3],[102, 3],[102, 2]], [[100.2, 0.2],[100.8, 0.2],[100.8, 0.8],[100.2, 0.8],[100.2, 0.2]]]}",
			},
			want: &MultiLineString{
				[]LineString{
					{
						Coordinates: []Point{
							{
								Lat: 2.0,
								Lng: 102.0,
							},
							{
								Lat: 2.0,
								Lng: 103.0,
							},
							{
								Lat: 3.0,
								Lng: 103.0,
							},
							{
								Lat: 3.0,
								Lng: 102.0,
							},
							{
								Lat: 2.0,
								Lng: 102.0,
							},
						},
					},
					{
						Coordinates: []Point{
							{
								Lat: 0.2,
								Lng: 100.2,
							},
							{
								Lat: 0.2,
								Lng: 100.8,
							},
							{
								Lat: 0.8,
								Lng: 100.8,
							},
							{
								Lat: 0.8,
								Lng: 100.2,
							},
							{
								Lat: 0.2,
								Lng: 100.2,
							},
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			g, err := FromJSON(tt.args.geojson)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiLineStringFromJson() error = %v, wantErr %v", err, tt.err.Error())
					return
				}
				return
			}

			p, err := g.ToMultiLineString()

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestMultiLineStringFromJson() error = %v, wantErr %v", err.Error(), tt.err.Error())
					return
				}
				return
			}

			if got := p; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestMultiLineStringFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
