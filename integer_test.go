package asn1

import (
	"testing"
)

func TestIntegerEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    int
	}{
		{
			"Test encode integer",
			[]byte{0x00},
			0,
		},
		{
			"Test encode integer",
			[]byte{0x48},
			72,
		},
		{
			"Test encode integer",
			[]byte{0x7f},
			127,
		},
		{
			"Test encode integer",
			[]byte{0x80},
			-128,
		},
		{
			"Test encode integer",
			[]byte{0x00, 0x80},
			128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := intEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}
