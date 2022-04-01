package asn1

import (
	"math"
	"testing"
)

func TestRealEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    float64
	}{
		{
			"Test encode real",
			[]byte{0x02, 0x33, 0x2e, 0x31, 0x34, 0x35, 0x39},
			3.1459,
		},
		{
			"Test encode real - infinity",
			[]byte{0x40},
			math.Inf(1),
		},
		{
			"Test encode real - negative infinity",
			[]byte{0x41},
			math.Inf(-1),
		},
		{
			"Test encode real - NaN",
			[]byte{0x42},
			math.Log(-1.0),
		},
		{
			"Test encode real - negative 0",
			[]byte{0x43},
			math.Copysign(0, -1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := realEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}
