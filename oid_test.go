package asn1

import (
	"testing"
)

func TestObjectIdentifierEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    ObjectIdentifier
	}{
		{
			"Test encode object identifier",
			[]byte{0x88, 0x37, 0x03},
			ObjectIdentifier{2, 999, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := objectIdentifierEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}

func TestNewObjectIdentifier(t *testing.T) {
	// valid
	_, err := NewObjectIdentifier(0, []uint{9, 2342, 19200300, 100, 1, 2})
	assertNoError(t, err)

	// bad root node
	_, err = NewObjectIdentifier(5, []uint{9, 2342, 19200300, 100, 1, 2})
	assertError(t, err)
	_, err = NewObjectIdentifier(50, []uint{9, 2342, 19200300, 100, 1, 2})
	assertError(t, err)

	// empty node
	_, err = NewObjectIdentifier(0, []uint{})
	assertError(t, err)
}

func TestObjectIdentifierToString(t *testing.T) {
	oid, _ := NewObjectIdentifier(0, []uint{9, 2342, 19200300, 100, 1, 1})
	oidStr := oid.ToString()
	assertEqual(t, "0.9.2342.19200300.100.1.1", oidStr)
}
