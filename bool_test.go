package asn1

import (
	"testing"
)

func TestBoolEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    bool
	}{
		{
			"Test encode bool",
			[]byte{0xff},
			true,
		},
		{
			"Test encode bool",
			[]byte{0x00},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := boolEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}
