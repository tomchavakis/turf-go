package coordinateReferenceSystem

import "errors"

// NamedCRS defines the Named CRS Type
type NamedCRS struct {
	name string
}

// NewNamedCRS initializes a new instance of the NamedCRS
// name must be a a name identifying a coordinate system with the OGS CRS URNs like 'urn:ogc:def:crs:OGC:1.3:CRS84'
func NewNamedCRS(name string) (*CRSBase, error){
	if name == "" {
		return nil, errors.New("name must be specified")
	}

	return &CRSBase{
		Properties: map[string]string{"name":name},
		Type:       Name,
	}, nil
}
