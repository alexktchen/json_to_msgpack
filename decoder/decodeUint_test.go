package decoder

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDecodeUint(t *testing.T) {
	testCases := []struct {
		data              []byte
		expectedValue     uint64
		expectedBytesRead int
	}{
		{[]byte{0xcc, 0x7f}, 127, 1},                        // uint 8
		{[]byte{0xCD, 0xff, 0xff}, 65535, 2},                // uint 16
		{[]byte{0xce, 0x02, 0x0f, 0x5C, 0x2B}, 34561067, 4}, // uint 32
	}

	for _, tc := range testCases {

		decoder := &MsgpackDecoder{}
		decoder.setData(tc.data)
		value, err := decoder.withValue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !bytes.Equal(value, []byte(fmt.Sprintf("%d", tc.expectedValue))) {
			t.Errorf("Decoded value mismatch.\nExpected: %v\nGot: %v", tc.expectedValue, value)
		}
	}
}
