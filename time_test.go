package asn1

import (
	"testing"
	"time"
)

var datetime = time.Date(1988, time.December, 9, 3, 0, 0, 0, time.Local)

func TestTimeUTCEncode(t *testing.T) {
	expected := "881209030000-0600"
	assertEqual(t, expected, makeUTCTime(datetime))
}

func TestTimeGeneralizedEncode(t *testing.T) {
	expected := "19881209030000Z"
	assertEqual(t, expected, makeGeneralizedTime(datetime))
}
