package geometry

// BBOX defines a Bounding Box
type BBOX struct {
	West  float64
	South float64
	East  float64
	North float64
}

// NewBBox initializes a new Bounding Box
func NewBBox(west float64, south float64, east float64, north float64) *BBOX {
	return &BBOX{
		West:  west,
		South: south,
		East:  east,
		North: north,
	}
}
