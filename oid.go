package asn1

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

type ObjectIdentifier []uint

func isValidRootNode(root uint) bool {
	switch root {
	case 0, 1, 2:
		return true
	default:
		return false
	}
}

func NewObjectIdentifier(root uint, node []uint) (ObjectIdentifier, error) {
	if !isValidRootNode(root) {
		return nil, fmt.Errorf("error creating object identifier: invalid root node %d", root)
	}
	if len(node) == 0 {
		return nil, fmt.Errorf("error creating object identifier: empty node")
	}

	switch root {
	case 0, 1:
		if node[0] >= 40 {
			return nil, fmt.Errorf("invalid node")
		}
	}

	oid := []uint{root}
	oid = append(oid, node...)

	return ObjectIdentifier(oid), nil
}

var objectIdentifierType = reflect.TypeOf(ObjectIdentifier{})

type objectIdentifierEncoder ObjectIdentifier

func (e objectIdentifierEncoder) encode() ([]byte, error) {
	b := new(bytes.Buffer)

	b.Write(encodeBase128((e[0]*40 + e[1])))

	for i := 2; i < len(e); i++ {
		b.Write(encodeBase128(e[i]))
	}

	return b.Bytes(), nil
}

func (o ObjectIdentifier) ToString() string {
	oidString := strconv.Itoa(int(o[0])) + "."
	for j, i := range o[1:] {
		oidString += strconv.Itoa(int(i))
		if j+1 < len(o[1:]) {
			oidString += "."
		}
	}
	return oidString
}
