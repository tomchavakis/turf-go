package crs

// UnspecifiedCRS defines the Named CRS Type
type UnspecifiedCRS struct {
}

// NewUnspecifiedCRS initializes a new instance of the UnspecifiedCRS
func (u *UnspecifiedCRS)NewUnspecifiedCRS() (*Base, error) {
	return &Base{
		Properties: map[string]string{},
		Type:       Unspecified,
	}, nil
}
