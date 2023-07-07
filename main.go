package main

import (
	"fmt"
	"log"

	"github.com/alexktchen/json_to_msgpack/builders"
)

// Example usage
func main() {
	jsonData := []byte(`{
		"name": "John",
		"age": 30,
		"city": "New York",
		"married": true,
		"interests": ["programming", "music", "sports"],
		"languages": {
			"go": true,
			"python": true,
			"java": false
		}
	}`)

	msgpackData, err := builders.ConvertJSONtoMessagePack(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", msgpackData)
}
