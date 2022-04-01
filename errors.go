package asn1

import (
	"fmt"
	"reflect"
)

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("asn1: unsupported type: %s", e.Type.String())
}
