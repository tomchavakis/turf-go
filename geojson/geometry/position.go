package geometry

// Position is the fundamental geometry construct, consisting of Latitude, Longtitude and Altitude
type Position struct {
	Altitude  *float64
	Latitude  float64
	Longitude float64
}

// NewPosition initializes a new instance of the Position
func NewPosition(altitude *float64, latitude float64, longitude float64) *Position {
	return &Position{Altitude: altitude, Latitude: latitude, Longitude: longitude}
}

// ToPoint converts a Position to Point
func (p *Position) ToPoint() Point {
	return Point{
		Lat: p.Latitude,
		Lng: p.Longitude,
	}
}
