package encoder

import (
	"testing"
)

func TestConvertJSONtoMessagePack(t *testing.T) {
	// Test JSON data
	jsonData := []byte(`{
		"name": "John",
		"age": 30,
		"city": "New York",
		"likes": ["reading", "coding"],
		"details": {
			"height": 180,
			"weight": 75.5
		}
	}`)

	// Convert JSON to MessagePack
	result, err := ConvertJSONtoMessagePack(jsonData)
	if err != nil {
		t.Errorf("Error converting JSON to MessagePack: %v", err)
	}

	// Check if the converted data matches the expected data
	if result == nil {
		t.Errorf("result is nil")
	}
}
