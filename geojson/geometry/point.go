package geometry

// Point represents a geolocation using ESPG-900913/(ESPG-3875) Projection
type Point struct {
	Lat float64
	Lng float64
}

// NewPoint initializes a new Point
func NewPoint(lat float64, lng float64) *Point {
	return &Point{
		Lat: lat,
		Lng: lng,
	}
}
