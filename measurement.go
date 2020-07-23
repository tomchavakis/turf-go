package turf

import "math"

// Distance calculates the distance between two points in degrees, radians, miles or kilometers. This uses the Haversine formula
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

// DistancePoint calculates the distance between two points
func DistancePoint(p1 Point, p2 Point) float64 {
	return Distance(p1.Lng, p1.Lat, p2.Lng, p2.Lat)
}
