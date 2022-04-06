package asn1

import "testing"

func TestPrintableString(t *testing.T) {
	got, err := makePrintableString("foobar")
	assertNoError(t, err)
	assertEqual(t, "foobar", got)

	_, err = makePrintableString("***")
	assertError(t, err)
}

func TestMakeIA5String(t *testing.T) {
	got, err := makeIA5String("foobar")
	assertNoError(t, err)
	assertEqual(t, "foobar", got)

	_, err = makeIA5String("€")
	assertError(t, err)
}

func TestMakeNumericString(t *testing.T) {
	got, err := makeNumericString("1 2 3")
	assertNoError(t, err)
	assertEqual(t, "1 2 3", got)

	_, err = makeNumericString("foobar")
	assertError(t, err)
}

// func TestPrintableStringEncode(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		expected    []byte
// 		value       string
// 		errExpected bool
// 	}{
// 		{
// 			"Test encode string",
// 			[]byte{0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72},
// 			"foobar",
// 			false,
// 		},
// 		{
// 			"Test encode string",
// 			[]byte{},
// 			"***",
// 			true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			printable, err := makePrintableString(tt.value)
// 			encoding, err := stringEncoder(printable).encode()
// 			if tt.errExpected {
// 				assertError(t, err)
// 			} else {
// 				assertNoError(t, err)
// 				assertBytesEqual(t, tt.expected, encoding)
// 			}
// 		})
// 	}
// }
