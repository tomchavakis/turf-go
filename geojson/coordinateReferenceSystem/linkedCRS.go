package coordinateReferenceSystem

import "errors"

// LinkedCRS defines a Linked CRS Type
// http://geojson.org/geojson-spec.html#named-crs
type LinkedCRS struct {
	Href string
	Type string
}

// NewLinkedCRS initializes a new instance of the LinkedCRS
// href must be a URI string
func NewLinkedCRS(href string, tp string) (*CRSBase,error){
	if href == "" || tp == ""{
		return nil, errors.New("href or type can't be empty")
	}
	return &CRSBase{
		Properties: map[string]string{"href":href},
		Type:       Link,
	}, nil
}
