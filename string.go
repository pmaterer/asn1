package asn1

import (
	"fmt"
	"unicode"
)

type stringEncoder string

func (e stringEncoder) encode() ([]byte, error) {
	return []byte(e), nil
}

// https://en.wikipedia.org/wiki/PrintableString
func isPrintable(b byte) bool {
	for _, c := range string(b) {
		switch {
		case c >= 'a' && c <= 'z':
		case c >= 'A' && c <= 'Z':
		case c >= '0' && c <= '9':
		default:
			switch c {
			case ' ', '\'', '(', ')', '+', ',', '-', '.', '/', ':', '=', '?':
			default:
				return false
			}
		}
	}
	return true
}

func makePrintableString(s string) (string, error) {
	stringBytes := []byte(s)
	for i := 0; i < len(stringBytes); i++ {
		if !isPrintable(stringBytes[i]) {
			return "", fmt.Errorf("PrintableString contains invalid character: '%s'", string(stringBytes[i]))
		}
	}
	return s, nil
}

func makeIA5String(s string) (string, error) {
	stringBytes := []byte(s)
	for i := 0; i < len(stringBytes); i++ {
		if stringBytes[i] > 127 {
			return "", fmt.Errorf("IA5String contains invalid character: '%s'", string(stringBytes[i]))
		}
	}
	return s, nil
}

func makeNumericString(s string) (string, error) {
	stringBytes := []byte(s)
	for i := 0; i < len(stringBytes); i++ {
		if !unicode.IsDigit(rune(stringBytes[i])) || stringBytes[i] != ' ' {
			return "", fmt.Errorf("NumericString contains invalid character: '%s'", string(stringBytes[i]))
		}
	}
	return s, nil
}
