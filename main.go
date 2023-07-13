package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/alexktchen/json_to_msgpack/decoder"
	"github.com/alexktchen/json_to_msgpack/encoder"
)

// Converter defines the interface for converting JSON to MessagePack
type Converter interface {
	Convert(data []byte) ([]byte, error)
}

// JSONConverter implements the Converter interface for JSON conversion
type JSONConverter struct{}

// Convert converts JSON to MessagePack format
func (c *JSONConverter) Convert(data []byte) ([]byte, error) {

	msgpackData, err := encoder.ConvertJSONtoMessagePack(data)
	if err != nil {
		log.Fatal(err)
	}
	return msgpackData, nil
}

// MessagePackConverter implements the Converter interface for MessagePack conversion
type MessagePackConverter struct{}

func (c *MessagePackConverter) Convert(data []byte) ([]byte, error) {

	result, err := decoder.ConvertMessagePackToJson(data)

	if err != nil {
		fmt.Println("Error:", err)

	}

	return result, nil
}

// ConverterFactory creates the appropriate converter based on the given format
type ConverterFactory struct{}

// CreateConverter creates a converter based on the given format
func (f *ConverterFactory) CreateConverter(format string) Converter {
	switch format {
	case "json":
		return &JSONConverter{}
	case "msgpack":
		return &MessagePackConverter{}
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
	fmt.Println("data:", arg2)

	converterFactory := &ConverterFactory{}
	converter := converterFactory.CreateConverter(msgType)
	if converter == nil {
		log.Fatal("Invalid converter format")
	}

	if len(arg2) == 0 {
		log.Fatal("Invalid input format")
	}

	var input []byte

	if msgType == "json" {
		input = []byte(arg2)

	} else {
		input, _ = hex.DecodeString(arg2)
	}

	res, err := converter.Convert(input)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		if msgType == "json" {
			fmt.Printf("results: %x\n", res)
		} else {
			fmt.Printf("results : %s\n", string(res))
		}
	}
}
