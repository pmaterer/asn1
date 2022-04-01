package asn1

type boolEncoder bool

func (e boolEncoder) encode() ([]byte, error) {
	if e {
		return []byte{0xff}, nil
	}
	return []byte{0x00}, nil
}
