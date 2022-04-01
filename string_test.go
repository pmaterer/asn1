package asn1

import (
	"testing"
)

func TestPrintableStringEncode(t *testing.T) {
	tests := []struct {
		name        string
		expected    []byte
		value       string
		errExpected bool
	}{
		{
			"Test encode string",
			[]byte{0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72},
			"foobar",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := stringEncoder(tt.value).encode()
			if tt.errExpected {
				assertError(t, err)
			} else {
				assertNoError(t, err)
				assertBytesEqual(t, tt.expected, encoding)
			}
		})
	}
}
