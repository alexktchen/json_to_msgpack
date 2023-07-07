package builders

import (
	"encoding/json"
	"log"
	"reflect"
)

// ConvertJSONtoMessagePack converts a JSON byte slice to MessagePack format
func ConvertJSONtoMessagePack(jsonData []byte) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return encodeValue(data), nil
}

// encodeValue recursively encodes a value into MessagePack format
func encodeValue(value interface{}) []byte {
	builder := &Builder{}
	builder.withValue(value)
	return builder.build()
}

// build constructs the final MessagePack representation
func (b *Builder) build() []byte {
	return b.data
}

// WithValue adds a value to the MessagePack representation based on its type
func (b *Builder) withValue(value interface{}) *Builder {
	switch val := value.(type) {
	case nil:
		return b.WithNil()
	case bool:
		return b.WithBool(val)
	case int64:
		return b.WithInt(val)
	case float64:
		return b.WithFloat(val)
	case string:
		return b.WithString(val)
	case []interface{}:
		return b.WithArray(val)
	case map[string]interface{}:
		return b.WithMap(val)
	default:
		log.Printf("Unsupported data type: %v", reflect.TypeOf(value))
	}
	return b
}
