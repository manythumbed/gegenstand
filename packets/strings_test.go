package packets

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	given := "A\U0002A6D4"
	expected := []byte{0x0, 0x5, 0x41, 0xF0, 0xAA, 0x9B, 0x94}

	actual, err := encode(given)

	if err != nil {
		t.Errorf("Unexpected error with encode(%s) - %v", given, err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("encode(%s) = %v, expected = %v", given, actual, expected)
	}
}
