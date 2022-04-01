package asn1

import (
	"testing"
)

func TestNullEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    Null
	}{
		{
			"Test encode integer",
			[]byte{},
			Null{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := nullEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}
