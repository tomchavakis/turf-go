package measurement

import (
	"math"

	"github.com/tomchavakis/turf-go/constants"
	"github.com/tomchavakis/turf-go/conversions"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// Distance calculates the distance between two points in kilometers. This uses the Haversine formula
func Distance(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {

	dLat := conversions.DegreesToRadians(lat2 - lat1)
	dLng := conversions.DegreesToRadians(lon2 - lon1)
	lat1R := conversions.DegreesToRadians(lat1)
	lat2R := conversions.DegreesToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1R)*math.Cos(lat2R)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := constants.EarthRadius * c
	return d
}

// PointDistance calculates the distance between two points
func PointDistance(p1 geometry.Point, p2 geometry.Point) float64 {
	return Distance(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}

// Bearing finds the geographic bearing between two given points.
func Bearing(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {
	dLng := conversions.DegreesToRadians(lon2 - lon1)
	lat1R := conversions.DegreesToRadians(lat1)
	lat2R := conversions.DegreesToRadians(lat2)
	y := math.Sin(dLng) * math.Cos(lat2R)
	x := math.Cos(lat1R)*math.Sin(lat2R) - math.Sin(lat1R)*math.Cos(lat2R)*math.Cos(dLng)

	// convert to degrees
	bd := conversions.RadiansToDegrees(math.Atan2(y, x))

	if bd < 0.0 {
		bd += 360.0
	}

	if bd >= 360.0 {
		bd -= 360.0

	}
	return bd

}

// PointBearing finds the geographic bearing between two points.
func PointBearing(p1 geometry.Point, p2 geometry.Point) float64 {
	return Bearing(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}

// MidPoint finds the point midway between them.
func MidPoint(p1 geometry.Point, p2 geometry.Point) geometry.Point {
	dLon := conversions.DegreesToRadians(p2.Lng - p1.Lng)
	lat1R := conversions.DegreesToRadians(p1.Lat)
	lon1R := conversions.DegreesToRadians(p1.Lng)
	lat2R := conversions.DegreesToRadians(p2.Lat)
	Bx := math.Cos(lat2R) * math.Cos(dLon)
	By := math.Cos(lat2R) * math.Sin(dLon)
	midLat := math.Atan2(math.Sin(lat1R)+math.Sin(lat2R), math.Sqrt((math.Cos(lat1R)+Bx)*(math.Cos(lat1R)+Bx)+By*By))
	midLng := lon1R + math.Atan2(By, math.Cos(lat1R)+Bx)

	return geometry.Point{Lat: conversions.RadiansToDegrees(midLat), Lng: conversions.RadiansToDegrees(midLng)}
}

// Destination returns a destination point according to a reference point, a distance in km and a bearing in degrees from True North.
func Destination(p1 geometry.Point, distance float64, bearing float64) geometry.Point {
	lonR := conversions.DegreesToRadians(p1.Lng)
	latR := conversions.DegreesToRadians(p1.Lat)
	bR := conversions.DegreesToRadians(bearing)
	dLat := math.Asin(math.Sin(latR)*math.Cos(distance/constants.EarthRadius) + math.Cos(latR)*math.Sin(distance/constants.EarthRadius)*math.Cos(bR))
	dLng := lonR + math.Atan2(math.Sin(bR)*math.Sin(distance/constants.EarthRadius)*math.Cos(latR), math.Cos(distance/constants.EarthRadius)-math.Sin(latR)*math.Sin(dLat))

	return geometry.Point{Lat: conversions.RadiansToDegrees(dLat), Lng: conversions.RadiansToDegrees(dLng)}
}

// Length measures the length of a set of points
// http://turfjs.org/docs/#linedistance
func Length(coords []geometry.Point) float64 {
	travelled := 0.0
	prevCoords := coords[0]
	var currentCoords geometry.Point
	for i := 1; i < len(coords); i++ {
		currentCoords = coords[i]
		travelled += PointDistance(prevCoords, currentCoords)
		prevCoords = currentCoords
	}
	return travelled
}

// LineStringLength measures the length of a Linestring
func LineStringLength(l geometry.LineString) float64 {
	return Length(l.Coordinates)
}

// MultiLineStringLength measures the length of a MultiLineString
func MultiLineStringLength(ml geometry.MultiLineString) float64 {
	len := 0.0
	coords := ml.Coordinates // []LineString
	for _, c := range coords {
		len += LineStringLength(c)
	}

	return len
}

// PolygonLength measures the length of a polygon
func PolygonLength(p geometry.Polygon) float64 {
	len := 0.0
	for _, c := range p.Coordinates {
		len += LineStringLength(c)
	}

	return len
}

// MultiPolygonLength measures the length of a MultiPolygon
func MultiPolygonLength(mp geometry.MultiPolygon) float64 {
	len := 0.0
	coords := mp.Coordinates
	for _, coord := range coords {
		for _, pl := range coord.Coordinates {
			len += LineStringLength(pl)
		}
	}
	return len
}
