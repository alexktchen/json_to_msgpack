package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alexktchen/json_to_msgpack/builders"
)

// Converter defines the interface for converting JSON to MessagePack
type Converter interface {
	Convert(data []byte) ([]byte, error)
}

// JSONConverter implements the Converter interface for JSON conversion
type JSONConverter struct{}

// Convert converts JSON to MessagePack format
func (c *JSONConverter) Convert(data []byte) ([]byte, error) {

	msgpackData, err := builders.ConvertJSONtoMessagePack(data)
	if err != nil {
		log.Fatal(err)
	}
	return msgpackData, nil
}

// MessagePackConverter implements the Converter interface for MessagePack conversion
type MessagePackConverter struct{}

func (c *MessagePackConverter) Convert(data []byte) ([]byte, error) {

	value, err := convertValue(data)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func convertValue(data []byte) ([]byte, error) {
	switch data[0] >> 4 {
	case 0x0:
		return convertPositiveFixNum(data)
	case 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7:
		return convertNegativeFixNum(data)
	case 0x8:
		return convertFixMap(data)
	case 0x9:
		return convertFixArray(data)
	case 0xa:
		return convertFixStr(data)
	case 0xb:
		return convertStr8(data)
	case 0xc:
		return convertStr16(data)
	case 0xd:
		return convertStr32(data)
	case 0xe:
		return convertArray16(data)
	case 0xf:
		return convertArray32(data)
	default:
		log.Printf("Unsupported MessagePack data type: %x", data[0])
		return nil, errors.New("unsupported MessagePack data type")
	}

	//return nil, fmt.Errorf("unsupported MessagePack type")
}

func convertPositiveFixNum(data []byte) ([]byte, error) {
	return json.Marshal(int(data[0] & 0x7f))
}

func convertNegativeFixNum(data []byte) ([]byte, error) {
	return json.Marshal(int(int8(data[0])))
}

func convertFixMap(data []byte) ([]byte, error) {
	count := int(data[0] & 0xf)
	result := make(map[string]interface{})
	offset := 1

	for i := 0; i < count; i++ {
		key, err := convertValue(data[offset:])
		if err != nil {
			return nil, err
		}
		offset += len(key)

		value, err := convertValue(data[offset:])
		if err != nil {
			return nil, err
		}
		offset += len(value)

		result[string(key)] = value
	}

	return json.Marshal(result)
}

func convertFixArray(data []byte) ([]byte, error) {
	count := int(data[0] & 0xf)
	result := make([]interface{}, count)
	offset := 1

	for i := 0; i < count; i++ {
		value, err := convertValue(data[offset:])
		if err != nil {
			return nil, err
		}
		offset += len(value)

		result[i] = value
	}

	return json.Marshal(result)
}

func convertFixStr(data []byte) ([]byte, error) {
	length := int(data[0] & 0x1f)
	return convertString(data[1:], length)
}

func convertStr8(data []byte) ([]byte, error) {
	length := int(data[1])
	return convertString(data[2:], length)
}

func convertStr16(data []byte) ([]byte, error) {
	length := int(data[1])<<8 | int(data[2])
	return convertString(data[3:], length)
}

func convertStr32(data []byte) ([]byte, error) {
	length := int(data[1])<<24 | int(data[2])<<16 | int(data[3])<<8 | int(data[4])
	return convertString(data[5:], length)
}

func convertArray16(data []byte) ([]byte, error) {
	length := int(data[1])<<8 | int(data[2])
	result := make([]interface{}, length)
	offset := 3

	for i := 0; i < length; i++ {
		value, err := convertValue(data[offset:])
		if err != nil {
			return nil, err
		}
		offset += len(value)

		result[i] = value
	}

	return json.Marshal(result)
}

func convertArray32(data []byte) ([]byte, error) {
	length := int(data[1])<<24 | int(data[2])<<16 | int(data[3])<<8 | int(data[4])
	result := make([]interface{}, length)
	offset := 5

	for i := 0; i < length; i++ {
		value, err := convertValue(data[offset:])
		if err != nil {
			return nil, err
		}
		offset += len(value)

		result[i] = value
	}

	return json.Marshal(result)
}

func convertString(data []byte, length int) ([]byte, error) {
	return json.Marshal(string(data[:length]))
}

// ConverterFactory creates the appropriate converter based on the given format
type ConverterFactory struct{}

// CreateConverter creates a converter based on the given format
func (f *ConverterFactory) CreateConverter(format string) Converter {
	switch format {
	case "json":
		return &JSONConverter{}
	case "msgpack":

		fmt.Println("results: Not yet implement")
		return nil
	default:
		return nil
	}
}

// Example usage
func main() {

	// Check the length of os.Args to ensure the required argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing required argument. e.g. type json or msgpack.")
		return
	}

	// Access the required argument
	msgType := os.Args[1]
	fmt.Println("type:", msgType)

	arg2 := os.Args[2]
	fmt.Println("json data:", arg2)

	jsonData := []byte(arg2)

	converterFactory := &ConverterFactory{}
	converter := converterFactory.CreateConverter(msgType)
	if converter == nil {
		log.Fatal("Invalid converter format")
	}

	msgpackData, err := converter.Convert(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	if msgpackData != nil {
		fmt.Printf("msgpackData: %x\n", msgpackData)
	}
}
