package invariant

import (
	"errors"
	"reflect"
	"testing"

	"github.com/tomchavakis/turf-go/geojson"
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

func TestGetCoords(t *testing.T) {
	type args struct {
		coords interface{}
	}
	tests := map[string]struct {
		args    args
		want    interface{}
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
		"error array - single array ": {
			args: args{
				[]float64{44.34},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("coord must be GeoJSON Point or an Array of numbers"),
		},
		"error feature  - polygon": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.Polygon,
						Coordinates: [][][]float64{
							{{2, 1}, {4, 3}}, {{6, 5}, {8, 7}},
						},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("cannot create a new polygon a polygon must have at least 4 positions"),
		},
		"error - feature  - point": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.Point,
						Coordinates: [][]float64{{23.33, 33, 33}},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("cannot unmarshal object"),
		},
		"error - feature - multiPoint": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiPoint,
						Coordinates: []float64{
							102,
						},
					},
				},
			},
			wantErr: true,
			err:     errors.New("cannot unmarshal object"),
			want:    nil,
		},
		"error - feature  - lineString": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.LineString,
						Coordinates: []float64{
							44.34, 23.52,
						},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("cannot marshal object"),
		},
		"error - feature  - multiLineString": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiLineString,
						Coordinates: [][][]float64{
							{{44.34, 23.52}, {33.33, 44.44}},
						},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("can't create a new multiLineString"),
		},
		"error feature  - multiPolygon": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiPolygon,
						Coordinates: [][][]float64{
							{
								{44.34}, {23.52}, {33.33}, {44.44},
							},
						},
					},
				},
			},
			wantErr: true,
			want:    nil,
			err:     errors.New("cannot marshal object"),
		},
		"feature  - point": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.Point,
						Coordinates: []float64{44.34, 23.52},
					},
				},
			},
			wantErr: false,
			want:    []float64{44.34, 23.52},
			err:     nil,
		},
		"feature - multiPoint": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiPoint,
						Coordinates: [][]float64{
							{102, -10},
							{103, 1},
							{104, 0},
							{130, 4},
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
			want: [][]float64{
				{102, -10},
				{103, 1},
				{104, 0},
				{130, 4},
			},
		},
		"feature  - lineString": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.LineString,
						Coordinates: [][]float64{
							{44.34, 23.52}, {33.33, 44.44},
						},
					},
				},
			},
			wantErr: false,
			want: [][]float64{
				{44.34, 23.52}, {33.33, 44.44},
			},
			err: nil,
		},
		"feature  - multiLineString": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiLineString,
						Coordinates: [][][]float64{
							{{44.34, 23.52}, {33.33, 44.44}},
							{{45.34, 23.52}, {35.33, 46.44}},
						},
					},
				},
			},
			wantErr: false,
			want: [][][]float64{
				{{44.34, 23.52}, {33.33, 44.44}},
				{{45.34, 23.52}, {35.33, 46.44}},
			},
			err: nil,
		},
		"feature  - multiPolygon": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.MultiPolygon,
						Coordinates: [][][][]float64{
							{
								{
									{44.34, 23.52}, {33.33, 44.44},
								},
								{
									{45.34, 23.52}, {33.33, 44.44},
								},
								{
									{46.34, 23.52}, {33.33, 44.44},
								},
								{
									{47.34, 23.52}, {33.33, 44.44},
								},
							},
							{
								{
									{48.34, 23.52}, {34.33, 44.44},
								},
								{
									{49.34, 23.52}, {35.33, 44.44},
								},
								{
									{50.34, 23.52}, {36.33, 44.44},
								},
								{
									{51.34, 23.52}, {37.33, 44.44},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			want: [][][][]float64{
				{
					{
						{44.34, 23.52}, {33.33, 44.44},
					},
					{
						{45.34, 23.52}, {33.33, 44.44},
					},
					{
						{46.34, 23.52}, {33.33, 44.44},
					},
					{
						{47.34, 23.52}, {33.33, 44.44},
					},
				},
				{
					{
						{48.34, 23.52}, {34.33, 44.44},
					},
					{
						{49.34, 23.52}, {35.33, 44.44},
					},
					{
						{50.34, 23.52}, {36.33, 44.44},
					},
					{
						{51.34, 23.52}, {37.33, 44.44},
					},
				},
			},
			err: nil,
		},
		"feature  - polygon": {
			args: args{
				coords: &feature.Feature{
					ID:         "",
					Type:       geojson.Feature,
					Properties: map[string]interface{}{},
					Bbox:       []float64{},
					Geometry: geometry.Geometry{
						GeoJSONType: geojson.Polygon,
						Coordinates: [][][]float64{
							{
								{101, 0},
								{101, 1},
								{100, 1},
								{100, 0},
								{101, 0},
							},
						},
					},
				},
			},
			wantErr: false,
			want: [][][]float64{
				{
					{101, 0},
					{101, 1},
					{100, 1},
					{100, 0},
					{101, 0},
				},
			},
			err: nil,
		},
		"geometry - polygon": {
			args: args{
				coords: &geometry.Polygon{
					Coordinates: []geometry.LineString{
						{
							Coordinates: []geometry.Point{
								{
									Lat: 1.0,
									Lng: 2.0,
								},
								{
									Lat: 3.0,
									Lng: 4.0,
								},
							},
						},
						{
							Coordinates: []geometry.Point{
								{
									Lat: 5.0,
									Lng: 6.0,
								},
								{
									Lat: 7.0,
									Lng: 8.0,
								},
							},
						},
					},
				},
			},
			wantErr: false,
			want: [][][]float64{
				{{2, 1}, {4, 3}}, {{6, 5}, {8, 7}},
			},
			err: nil,
		},
		"geometry - lineString": {
			args: args{
				coords: &geometry.LineString{
					Coordinates: []geometry.Point{
						{
							Lat: 1.0,
							Lng: 2.0,
						},
						{
							Lat: 3.0,
							Lng: 4.0,
						},
					},
				},
			},
			wantErr: false,
			want: [][]float64{
				{2, 1}, {4, 3},
			},
			err: nil,
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
		"geometry - multiPoint": {
			args: args{
				coords: &geometry.MultiPoint{
					Coordinates: []geometry.Point{
						{
							Lat: 23.44,
							Lng: 43.33,
						},
						{
							Lat: 25.44,
							Lng: 44.33,
						},
						{
							Lat: 26.46,
							Lng: 45.33,
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
			want: [][]float64{
				{43.33, 23.44},
				{44.33, 25.44},
				{45.33, 26.46},
			},
		},
		"geometry - multiLineString": {
			args: args{
				coords: &geometry.MultiLineString{
					Coordinates: []geometry.LineString{
						{
							Coordinates: []geometry.Point{
								{
									Lat: 23.44,
									Lng: 43.33,
								},
								{
									Lat: 25.44,
									Lng: 44.33,
								},
								{
									Lat: 26.46,
									Lng: 45.33,
								},
							},
						},
						{
							Coordinates: []geometry.Point{
								{
									Lat: 29.44,
									Lng: 48.33,
								},
								{
									Lat: 36.46,
									Lng: 55.33,
								},
							},
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
			want: [][][]float64{
				{
					{43.33, 23.44}, {44.33, 25.44}, {45.33, 26.46},
				},
				{
					{48.33, 29.44}, {55.33, 36.46},
				},
			},
		},
		"geometry - multiPolygon": {
			args: args{
				coords: &geometry.MultiPolygon{
					Coordinates: []geometry.Polygon{
						{
							Coordinates: []geometry.LineString{
								{
									Coordinates: []geometry.Point{
										{
											Lat: 23.55,
											Lng: 43.66,
										},
										{
											Lat: 24.55,
											Lng: 44.67,
										},
										{
											Lat: 25.55,
											Lng: 45.68,
										},
									},
								},
								{
									Coordinates: []geometry.Point{
										{
											Lat: 25.55,
											Lng: 43.66,
										},
										{
											Lat: 26.55,
											Lng: 44.67,
										},
										{
											Lat: 27.55,
											Lng: 45.68,
										},
									},
								},
							},
						},
						{
							Coordinates: []geometry.LineString{
								{
									Coordinates: []geometry.Point{
										{
											Lat: 30.55,
											Lng: 43.66,
										},
										{
											Lat: 34.55,
											Lng: 44.67,
										},
										{
											Lat: 35.55,
											Lng: 45.68,
										},
									},
								},
								{
									Coordinates: []geometry.Point{
										{
											Lat: 45.55,
											Lng: 43.66,
										},
										{
											Lat: 46.55,
											Lng: 44.67,
										},
										{
											Lat: 47.55,
											Lng: 45.68,
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
			err:     nil,
			want: [][][][]float64{
				{
					{
						{43.66, 23.55}, {44.67, 24.55}, {45.68, 25.55},
					},
					{
						{43.66, 25.55}, {44.67, 26.55}, {45.68, 27.55},
					},
				},
				{
					{
						{43.66, 30.55}, {44.67, 34.55}, {45.68, 35.55},
					},
					{
						{43.66, 45.55}, {44.67, 46.55}, {45.68, 47.55},
					},
				},
			},
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
			geo, err := GetCoords(tt.args.coords)

			if (err != nil) && tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("TestGetCoords() error = %v, wantErr %v", err.Error(), tt.err.Error())
					return
				}
			}

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestGetCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	type args struct {
		geojson interface{}
	}
	fp, _ := feature.FromJSON("{ \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Point\", \"coordinates\": [102, 0.5] } }")
	g, _ := geometry.FromJSON("{ \"type\": \"MultiPolygon\", \"coordinates\": [ [ [ [-116, -45], [-90, -45], [-90, -56], [-116, -56], [-116, -45] ] ], [ [ [-90.351563, 9.102097], [-77.695312, -3.513421], [-65.039063, 12.21118], [-65.742188, 21.616579], [-84.023437, 24.527135], [-90.351563, 9.102097] ] ] ] }")
	fcl, _ := feature.CollectionFromJSON("{ \"type\": \"FeatureCollection\", \"features\": [ { \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Polygon\", \"coordinates\": [ [ [-12913060.93202, 7967317.535016], [-10018754.171395, 7967317.535016], [-10018754.171395, 9876845.895795], [-12913060.93202, 9876845.895795], [-12913060.93202, 7967317.535016] ] ] } }, { \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"LineString\", \"coordinates\": [ [-2269873.991957, 3991847.410438], [-391357.58482, 6026906.856034], [1565430.33928, 1917652.163291] ] } }, { \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Point\", \"coordinates\": [-8492459.534936, 430493.386177] } } ] }")
	gcl, _ := geometry.CollectionFromJSON("{ \"TYPE\": \"GeometryCollection\", \"geometries\": [ { \"TYPE\": \"Point\", \"coordinates\": [-71.0, 40.99999999999998] }, { \"TYPE\": \"LineString\", \"coordinates\": [ [-20.39062500000365, 33.72434000000235], [-3.5156249999990803, 47.51720099999992], [14.062499999996321, 16.97274100000141] ] } ] }")

	tests := map[string]struct {
		args args
		want string
	}{
		"error - required coords": {
			args: args{
				geojson: "",
			},
			want: "invalid",
		},
		"feature - point": {
			args: args{
				geojson: fp,
			},
			want: string(geojson.Point),
		},
		"feature - collection": {
			args: args{
				geojson: fcl,
			},
			want: string(geojson.FeatureCollection),
		},
		"geometry - collection": {
			args: args{
				geojson: gcl,
			},
			want: string(geojson.GeometryCollection),
		},
		"geometry - multipolygon": {
			args: args{
				geojson: g,
			},
			want: string(geojson.MultiPolygon),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			geo := GetType(tt.args.geojson)

			if got := geo; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestGetCoord() = %v, want %v", got, tt.want)
			}
		})
	}
}
