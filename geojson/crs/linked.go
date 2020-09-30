package crs

import "errors"

// Linked defines a Linked CRS Type
// http://geojson.org/geojson-spec.html#named-crs
type Linked struct {
}

// New initializes a new instance of the Linked CRS
// href must be a URI string
func (l *Linked) New(href string, tp string) (*Base, error) {
	if href == "" || tp == "" {
		return nil, errors.New("href or type can't be empty")
	}
	return &Base{
		Properties: map[string]string{"href": href},
		Type:       LinkedCRS,
	}, nil
}
