package asn1

import (
	"bytes"
	"encoding/hex"
	"testing"
	"time"
)

type intStruct struct {
	A int
}

type twoIntStruct struct {
	A int
	B int
}

type nestedStruct struct {
	A intStruct
}

type implicitTagTest struct {
	A int `asn1:"tag:5"`
}

type explicitTagTest struct {
	A int `asn1:"explicit,tag:5"`
}

// type generalizedTimeTest struct {
// 	A time.Time `asn1:"generalized"`
// }

type ia5StringTest struct {
	A string `asn1:"ia5"`
}

type printableStringTest struct {
	A string `asn1:"printable"`
}

// type genericStringTest struct {
// 	A string
// }

// type nullApplicationTest struct {
// 	A Null `asn1:"tag:2,application"`
// }

// type contextSpecificTest struct {
// 	A int `asn1:"tag:5"`
// }

// type applicationTest struct {
// 	A int `asn1:"application,tag:0`
// 	B int `asn1:"application,tag:1`
// }

type optionalTest struct {
	X int `asn1:"optional"`
	Y int `asn1:"optional"`
}

type optionalTaggedTest struct {
	X int `asn1:"optional,tag:0"`
	Y int `asn1:"optional,tag:1"`
}

// type testSet []int

type marshalTest struct {
	in  interface{}
	out string
}

var marshalTests = []marshalTest{
	// booleans
	{true, "0101ff"},
	{false, "010100"},
	// strings
	{"😎", "0c04f09f988e"},

	// object identifiers
	{ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}, "06092a864886f70d01010b"},
	{ObjectIdentifier{1, 2, 840, 113549}, "06062a864886f70d"},
	{ObjectIdentifier{1, 2, 3, 4}, "06032a0304"},
	{ObjectIdentifier{1, 2, 840, 133549, 1, 1, 5}, "06092a864888932d010105"},
	{ObjectIdentifier{2, 100, 3}, "0603813403"},
	// octet strings
	{[]byte{0x03, 0x02, 0x06, 0xa0}, "0404030206a0"},
	{[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, "04080123456789abcdef"},
	{[]byte("Hello!"), "040648656c6c6f21"},
	{[]byte{}, "0400"},
	{[]byte{0x00}, "040100"},
	{[]byte{0x01, 0x02, 0x03}, "0403010203"},
	{[]byte{1, 2, 3}, "0403010203"},
	{[]byte{0x03, 0x02, 0x06, 0xa0}, "0404030206A0"},
	// integers
	{-12345, "0202cfc7"},
	{-1000, "0202fc18"},
	{-1, "0201ff"},
	{0, "020100"},
	{1, "020101"},
	{50, "020132"},
	{127, "02017f"},
	{128, "02020080"},
	{255, "020200ff"},
	{256, "02020100"},
	{-128, "020180"},
	{-129, "0202ff7f"},
	{1000, "020203e8"},
	{50000, "020300c350"},
	// null
	{Null{}, "0500"},
	// structs
	{intStruct{64}, "3003020140"},
	{twoIntStruct{64, 65}, "3006020140020141"},
	{nestedStruct{intStruct{127}}, "3005300302017f"},
	{ia5StringTest{"test"}, "3006160474657374"},
	{printableStringTest{"test"}, "3006130474657374"},

	{optionalTest{X: 9}, "3003020109"},
	{optionalTaggedTest{X: 9}, "3003800109"},
	{optionalTaggedTest{Y: 9}, "3003810109"},
	{optionalTaggedTest{X: 9, Y: 9}, "3006800109810109"},
	// {nullApplicationTest{Null{}}, "4200"},
	// implicit/explicit
	{implicitTagTest{64}, "3003850140"},
	{explicitTagTest{64}, "3005a503020140"},
	// timestamp
	{time.Date(1991, time.May, 6, 23, 45, 40, 0, time.UTC), "17113931303530363233343534302b30303030"},
	// bit string
	{BitString{[]byte{0x80}, 7}, "03020780"},
	{BitString{[]byte{0x81, 0xf0}, 4}, "03030481f0"},
	{BitString{[]byte{0b01101110, 0b01011101, 0b11000000}, 6}, "0304066e5dc0"},

	// set
	// {testSet([]int{10}), "310302010a"},
}

func TestMarshal(t *testing.T) {
	for i, test := range marshalTests {
		data, err := Marshal(test.in)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		}
		out, _ := hex.DecodeString(test.out)
		if !bytes.Equal(out, data) {
			t.Errorf("#%d got: %x want %x\n\t%q\n\t%q", i, data, out, data, out)
		}
	}
}

type marshalWithOptionsTest struct {
	in      interface{}
	out     string
	options string
}

var marshalWithOptionsTests = []marshalWithOptionsTest{
	// strings
	{"hi", "13026869", "printable"},
	{"Test User 1", "130b5465737420557365722031", "printable"},
	{"test1@rsa.com", "160d7465737431407273612e636f6d", "ia5"},
	{"hi", "16026869", "ia5"},
	{"hi", "85026869", "tag:5"},
	{"hi", "A5040C026869", "tag:5,explicit"},
	{
		"" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // 127 times 'x'
		"137f" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"78787878787878787878787878787878787878787878787878787878787878",
		"printable",
	},
	{
		"" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" +
			"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // 128 times 'x'
		"138180" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"7878787878787878787878787878787878787878787878787878787878787878" +
			"7878787878787878787878787878787878787878787878787878787878787878",
		"printable",
	},
	{"test", "130474657374", "printable"},

	// time
	{time.Date(2019, time.December, 15, 19, 0o2, 10, 0, time.FixedZone("UTC-8", -8*60*60)), "17113139313231353139303231302d30383030", "utc"},
	{time.Date(1991, time.May, 6, 16, 45, 40, 0, time.FixedZone("UTC-8", -7*60*60)), "17113931303530363136343534302D30373030", "utc"},

	// enumerated
	{0, "0a0100", "enumerated"},
}

func TestMarshalWithOptions(t *testing.T) {
	for i, test := range marshalWithOptionsTests {
		data, err := MarshalWithOptions(test.in, test.options)
		if err != nil {
			t.Errorf("#%d failed: %s", i, err)
		}
		out, _ := hex.DecodeString(test.out)
		if !bytes.Equal(out, data) {
			t.Errorf("#%d got: %x want %x\n\t%q\n\t%q", i, data, out, data, out)
		}
	}
}

func TestEncodeLength(t *testing.T) {
	shortForm1 := encodeLength(5)
	assertBytesEqual(t, []byte{0x05}, shortForm1)

	shortForm2 := encodeLength(123)
	assertBytesEqual(t, []byte{0x7b}, shortForm2)

	longForm1 := encodeLength(500)
	assertBytesEqual(t, []byte{0x82, 0x01, 0xf4}, longForm1)

	longForm2 := encodeLength(1234)
	assertBytesEqual(t, []byte{0x82, 0x04, 0xd2}, longForm2)

	longForm3 := encodeLength(201)
	assertBytesEqual(t, []byte{0x81, 0xc9}, longForm3)
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, test := range marshalTests {
			Marshal(test.in)
		}
	}
}
