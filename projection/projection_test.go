package projection

import (
	"testing"

	"github.com/tomchavakis/turf-go/assert"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

func TestConvertToMercatorPoint(t *testing.T) {
	p := geometry.Point{
		Lng: -71.0,
		Lat: 41.0,
	}
	xy := ConvertToMercator(p)
	assert.Equal(t, xy[0], -7903683.846322424)
	assert.Equal(t, xy[1], 5012341.663847514)
}

func TestConvertToWgs84(t *testing.T) {
	p := []float64{-7903683.846322424, 5012341.663847514}
	wgs84Point := ConvertToWgs84(p)
	assert.Equal(t, wgs84Point.Lng, -71.0)
	assert.Equal(t, wgs84Point.Lat, 40.99999999999998) //=41.0
}

func TestProjectionPoint(t *testing.T) {
	p := geometry.Point{
		Lat: 40.0,
		Lng: 10.0,
	}
	mercator := ConvertToMercator(p)
	wgs84 := ConvertToWgs84(mercator)
	assert.Equal(t, p, wgs84)
}
