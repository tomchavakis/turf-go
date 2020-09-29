package crs

import "errors"

// Named defines the Named CRS Type
type Named struct {
}

// New initializes a new instance of the Named CRS
// name must be a a name identifying a coordinate system with the OGS CRS URNs like 'urn:ogc:def:crs:OGC:1.3:CRS84'
func (n *Named)New(name string) (*Base, error) {
	if name == "" {
		return nil, errors.New("name must be specified")
	}

	return &Base{
		Properties: map[string]string{"name": name},
		Type:       Name,
	}, nil
}
