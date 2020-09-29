package feature

// Feature defines a new feature type
// https://tools.ietf.org/html/rfc7946#section-3.2
// TODO: geometry type is a valid Geometry Type
type Feature struct {
	Geometry   interface{}
	Properties interface{}
	ID         *string
}
