package decoder

import (
	"encoding/hex"
	"testing"
)

func TestConvertMessagePackToJson(t *testing.T) {

	input, _ := hex.DecodeString("82a3616765a131a46e616d65a644616e69656c")

	jsonData, err := ConvertMessagePackToJson(input)
	if err != nil {
		t.Errorf("Error converting MessagePack to JSON: %v", err)
	}
	expected := `{"age":"1","name":"Daniel"}`
	res := string(jsonData)
	if res != expected {
		t.Errorf("Unexpected result.\nExpected: %s\nGot: %s", expected, jsonData)
	}
}
