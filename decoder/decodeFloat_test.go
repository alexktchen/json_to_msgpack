package decoder

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFloat64FromBytes(t *testing.T) {
	data := []byte{0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // Byte representation of float64(1.0)

	expectedValue := []byte(fmt.Sprintf("%f", float64(1.0)))

	decoder := &MsgpackDecoder{}
	decoder.setData(data)

	value, _ := decoder.float64FromBytes()

	if !bytes.Equal(value, expectedValue) {
		t.Errorf("Decoded value mismatch.\nExpected: %v\nGot: %v", expectedValue, value)
	}
}
