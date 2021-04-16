package projection

import (
	"math"

	"github.com/tomchavakis/turf-go/geojson/geometry"
)

const a = 6378137.0

// ConvertToMercator converts a WGS84 GeoJSON object to Mercator (EPSG:3857 sometimes knows as EPSG:900913) projection
// https://spatialreference.org/ref/sr-org/epsg3857-wgs84-web-mercator-auxiliary-sphere/
func ConvertToMercator(p geometry.Point) []float64 {
	rad := math.Pi / 180.0
	maxExtend := 2 * math.Pi * a / 2.0 // 20037508.342789244

	// longitudes passing the 180th meridian
	var adjusted float64
	if math.Abs(p.Lng) <= 180.0 {
		adjusted = p.Lng
	} else {
		adjusted = p.Lng - float64(sign(p.Lng))*360.0
	}

	xy := []float64{
		a * adjusted * rad,
		a * math.Log(math.Tan(math.Pi*0.25+0.5*p.Lat*rad)),
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
func ConvertToWgs84(p []float64) geometry.Point {
	dgs := 180.0 / math.Pi

	return geometry.Point{
		Lng: p[0] * dgs / a,
		Lat: (math.Pi*0.5 - 2.0*math.Atan(math.Exp(-p[1]/a))) * dgs,
	}
}

//TODO: Implement Convert - needs coordEach

func sign(x float64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
