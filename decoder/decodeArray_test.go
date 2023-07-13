package decoder

import (
	"reflect"
	"testing"
)

func TestDecodeArray(t *testing.T) {
	data := []byte{0x93, 0x01, 0x02, 0x03} // MessagePack array [1, 2, 3]
	expectedValue := "[1,2,3]"
	decoder := &MsgpackDecoder{}
	decoder.setData(data)
	value, err := decoder.withValue()
	res := string(value[:])
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, expectedValue) {
		t.Errorf("Decoded value mismatch.\nExpected: %v\nGot: %v", expectedValue, value)
	}

}
