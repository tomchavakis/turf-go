package feature

// TODO: geometry type is a valid Geometry Type
// Feature type
// https://tools.ietf.org/html/rfc7946#section-3.2
type Feature struct {
	geometry interface{}
	properties interface{}
	Id *string
}