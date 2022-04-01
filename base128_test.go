package asn1

import (
	"testing"
)

func TestBase128Encode(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    uint
	}{
		{
			"Test encode base-128",
			[]byte{0x7f},
			127,
		},
		{
			"Test encode base-128",
			[]byte{0x81, 0x00},
			128,
		},
		{
			"Test encode base-128",
			[]byte{0xc0, 0x00},
			8192,
		},
		{
			"Test encode base-128",
			[]byte{0xff, 0x7f},
			16383,
		},
		{
			"Test encode base-128",
			[]byte{0x81, 0x80, 0x00},
			16384,
		},
		{
			"Test encode base-128",
			[]byte{0xff, 0xff, 0x7f},
			2097151,
		},
		{
			"Test encode base-128",
			[]byte{0x81, 0x80, 0x80, 0x00},
			2097152,
		},
		{
			"Test encode base-128",
			[]byte{0xc0, 0x80, 0x80, 0x00},
			134217728,
		},
		{
			"Test encode base-128",
			[]byte{0xff, 0xff, 0xff, 0x7f},
			268435455,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoding := encodeBase128(tt.value)
			assertBytesEqual(t, tt.expected, encoding)
		})
	}
}
