package asn1

import (
	"bytes"
	"testing"
)

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: \n"+
			"%+v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func assertBytesEqual(t *testing.T, expected []byte, actual []byte) {
	t.Helper()
	n := bytes.Compare(expected, actual)
	if n != 0 {
		t.Fatalf("Not equal: \n"+
			"expected: % x\n"+
			"actual: % x", expected, actual)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Not equal: \n"+
			"expected: % x\n"+
			"actual: % x", expected, actual)
	}
}
