package asn1

type intEncoder int

func (e intEncoder) length() int {
	length := 1

	for e > 127 {
		length++
		e >>= 8
	}

	for e < -128 {
		length++
		e >>= 8
	}

	return length
}

func (e intEncoder) encode() ([]byte, error) {
	length := e.length()
	buf := make([]byte, length)
	for j := 0; j < length; j++ {
		shift := uint((length - 1 - j) * 8)
		buf[j] = byte(e >> shift)
	}

	return buf, nil
}
