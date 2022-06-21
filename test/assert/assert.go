package assert

import (
	"bytes"
	"reflect"
	"testing"
)

// Equal checks if values are equal
func Equal(t *testing.T, a, b interface{}) {
	if a == nil || b == nil {
		if a == b {
			return
		}
	}
	exp, ok := a.([]byte)
	if !ok {
		if reflect.DeepEqual(a, b) {
			return
		}
	}
	act, ok := b.([]byte)
	if !ok {
		t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
	if exp == nil || act == nil {
		if exp == nil && act == nil {
			return
		}
	}
	if bytes.Equal(exp, act) {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

// True asserts that the specified value is true
func True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	if !value {
		t.Errorf("Received %v, expected %v", value, true)
	}
	return value
}

func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)
	if isNilableKind && value.IsNil() {
		return true
	}
	return false
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
func Nil(t *testing.T, object interface{}) bool {
	return isNil(object)
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
func NotNil(t *testing.T, object interface{}) bool {
	return !isNil(object)
}
