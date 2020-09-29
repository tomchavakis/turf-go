package turf

import (
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"math"
)

// Distance calculates the distance between two points in kilometers. This uses the Haversine formula
func Distance(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {

	dLat := DegreesToRadians(lat2 - lat1)
	dLng := DegreesToRadians(lon2 - lon1)
	lat1R := DegreesToRadians(lat1)
	lat2R := DegreesToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLng/2), 2)*math.Cos(lat1R)*math.Cos(lat2R)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := EarthRadius * c
	return d
}

// PointDistance calculates the distance between two points
func PointDistance(p1 geometry.Point, p2 geometry.Point) float64 {
	return Distance(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}

// Bearing finds the geographic bearing between two given points.
func Bearing(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {
	dLng := DegreesToRadians(lon2 - lon1)
	lat1R := DegreesToRadians(lat1)
	lat2R := DegreesToRadians(lat2)
	y := math.Sin(dLng) * math.Cos(lat2R)
	x := math.Cos(lat1R)*math.Sin(lat2R) - math.Sin(lat1R)*math.Cos(lat2R)*math.Cos(dLng)

	// convert to degrees
	bd := RadiansToDegrees(math.Atan2(y, x))

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
	dLon := DegreesToRadians(p2.Lng - p1.Lng)
	lat1R := DegreesToRadians(p1.Lat)
	lon1R := DegreesToRadians(p1.Lng)
	lat2R := DegreesToRadians(p2.Lat)
	Bx := math.Cos(lat2R) * math.Cos(dLon)
	By := math.Cos(lat2R) * math.Sin(dLon)
	midLat := math.Atan2(math.Sin(lat1R)+math.Sin(lat2R), math.Sqrt((math.Cos(lat1R)+Bx)*(math.Cos(lat1R)+Bx)+By*By))
	midLng := lon1R + math.Atan2(By, math.Cos(lat1R)+Bx)

	return geometry.Point{Lat: RadiansToDegrees(midLat), Lng: RadiansToDegrees(midLng)}
}

// Destination returns a destination point according to a reference point, a distance in km and a bearing in degrees from True North.
func Destination(p1 geometry.Point, distance float64, bearing float64) geometry.Point {
	lonR := DegreesToRadians(p1.Lng)
	latR := DegreesToRadians(p1.Lat)
	bR := DegreesToRadians(bearing)
	dLat := math.Asin(math.Sin(latR)*math.Cos(distance/EarthRadius) + math.Cos(latR)*math.Sin(distance/EarthRadius)*math.Cos(bR))
	dLng := lonR + math.Atan2(math.Sin(bR)*math.Sin(distance/EarthRadius)*math.Cos(latR), math.Cos(distance/EarthRadius)-math.Sin(latR)*math.Sin(dLat))

	return geometry.Point{Lat: RadiansToDegrees(dLat), Lng: RadiansToDegrees(dLng)}
}
