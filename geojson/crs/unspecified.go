package crs

// Unspecified defines the Named CRS Type
type Unspecified struct {
}

// NewUnspecified initializes a new instance of the Unspecified CRS
func (u *Unspecified) NewUnspecified() (*Base, error) {
	return &Base{
		Properties: map[string]string{},
		Type:       UnspecifiedCRS,
	}, nil
}
