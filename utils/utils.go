package utils

import (
	"io/ioutil"
	"reflect"
)

// LoadJSONFixture loads a testing file
func LoadJSONFixture(filename string) (string, error) {
	filepath := filename

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// IsArray returns true of the interface is an array.
func IsArray(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice
}
