package encoder

import (
	"reflect"
	"testing"
)

func TestWithArray(t *testing.T) {
	// Create a new builder
	builder := &Builder{}

	// Define test array data
	arrayData := []interface{}{1, 2, 3, "four", true}

	// Add array data to the builder
	builder.WithArray(arrayData)

	// Get the resulting MessagePack data
	result := builder.Build()

	// Expected MessagePack data for the test array
	expectedData := []byte{0x95, 0x01, 0x02, 0x03, 0xa4, 0x66, 0x6f, 0x75, 0x72, 0xc3}

	// Check if the resulting MessagePack data matches the expected data
	if !reflect.DeepEqual(result, expectedData) {
		t.Errorf("Unexpected MessagePack data.\nExpected: %v\nActual: %v", expectedData, result)
	}
}
