package projection

import (
	"errors"
	"math"

	"github.com/tomchavakis/turf-go/geojson/geometry"
	meta "github.com/tomchavakis/turf-go/meta/coordEach"
)

const a = 6378137.0

// ConvertToMercator converts a WGS84 GeoJSON object to Mercator (EPSG:3857 sometimes knows as EPSG:900913) projection
// https://spatialreference.org/ref/sr-org/epsg3857-wgs84-web-mercator-auxiliary-sphere/
func ConvertToMercator(lonlat []float64) []float64 {
	rad := math.Pi / 180.0
	maxExtend := 2 * math.Pi * a / 2.0 // 20037508.342789244

	// longitudes passing the 180th meridian
	var adjusted float64
	if math.Abs(lonlat[0]) <= 180.0 {
		adjusted = lonlat[0]
	} else {
		adjusted = lonlat[0] - float64(sign(lonlat[0]))*360.0
	}

	xy := []float64{
		a * adjusted * rad,
		a * math.Log(math.Tan(math.Pi*0.25+0.5*lonlat[1]*rad)),
	}

	if xy[0] > maxExtend {
		xy[0] = maxExtend
	}
	if xy[0] < -maxExtend {
		xy[0] = -maxExtend
	}
	if xy[1] > maxExtend {
		xy[1] = maxExtend
	}
	if xy[1] < -maxExtend {
		xy[1] = -maxExtend
	}
	return xy
}

// ConvertToWgs84 convert 900913 x/y values to lon/lat
func ConvertToWgs84(p []float64) []float64 {
	dgs := 180.0 / math.Pi

	return []float64{
		p[0] * dgs / a,
		(math.Pi*0.5 - 2.0*math.Atan(math.Exp(-p[1]/a))) * dgs,
	}
}

// ToMercator converts a WGS84 GeoJSON object into Mercator (EPSG:900913) projection
func ToMercator(geojson interface{}) (interface{}, error) {
	return Convert(geojson, "mercator")
}

// ToWgs84 Converts a Mercator (EPSG:900913) GeoJSON object into WGS84 projection
func ToWgs84(geojson interface{}) (interface{}, error) {
	return Convert(geojson, "wgs84")
}

// Convert converts a GeoJSON coordinates to the defined projection
// gjson is GeoJSON Feature or geometry
// projection defines the projection system to convert the coordinates to
func Convert(geojson interface{}, projection string) (interface{}, error) {
	// Validation
	if geojson == nil {
		return nil, errors.New("geojson is required")
	}

	_, err := meta.CoordEach(geojson, func(p geometry.Point) geometry.Point {
		if projection == "mercator" {
			res := ConvertToMercator([]float64{p.Lng, p.Lat})
			p.Lng = res[0]
			p.Lat = res[1]
		} else {
			res := ConvertToWgs84([]float64{p.Lng, p.Lat})
			p.Lng = res[0]
			p.Lat = res[1]
		}

		return p
	}, nil)

	if err != nil {
		return nil, err
	}

	return geojson, nil
}

func sign(x float64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
