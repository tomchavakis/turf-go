package coordinateReferenceSystem

// UnspecifiedCRS defines the Named CRS Type
type UnspecifiedCRS struct {
	name string
}

// NewUnspecifiedCRS initializes a new instance of the UnspecifiedCRS
func NewUnspecifiedCRS() (*CRSBase, error){

	return &CRSBase{
		Properties: map[string]string{},
		Type:       Unspecified,
	}, nil
}