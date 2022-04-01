package asn1

import (
	"reflect"
	"testing"
)

var stringType = reflect.TypeOf("hello")

func TestUnsupportedTypeError(t *testing.T) {
	err := UnsupportedTypeError{stringType}
	assertEqual(t, "asn1: unsupported type: string", err.Error())
}
