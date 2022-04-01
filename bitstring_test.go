package asn1

import (
	"testing"
)

func TestBitStringEncode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    BitString
	}{
		{
			"Test encode bitstring",
			[]byte{0x04, 0x0a, 0x3b, 0x5f, 0x29, 0x1c, 0xd0},
			BitString{Bytes: []byte{0x0a, 0x3b, 0x5f, 0x29, 0x1c, 0xd0}, PaddingBits: 4},
		},
		{
			"Test encode bitstring",
			[]byte{0x06, 0x6e, 0x5d, 0xc0},
			BitString{Bytes: []byte{0b01101110, 0b01011101, 0b11000000}, PaddingBits: 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding, err := bitStringEncoder(tt.value).encode()
			assertNoError(t, err)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}

func TestNewBitString(t *testing.T) {
	// ok
	_, err := NewBitString([]byte{0x6e, 0x5d, 0xc0}, 6)
	assertNoError(t, err)

	// too many padding bits
	_, err = NewBitString([]byte{0x6e}, 9)
	assertError(t, err)

	// empty with padding
	_, err = NewBitString([]byte{}, 4)
	assertError(t, err)
}
