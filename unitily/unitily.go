package unitily

import (
	"fmt"
)

// Msgpack type codes
const (
	PositiveFixNum = 0x09 // 0xxxxxxx
	FixMap         = 0x80 // 1000xxxx
	FixArray       = 0x90 // 1001xxxx
	FixStr         = 0xa0 // 101xxxxx
	NilCode        = 0xc0 // 11000000
	FalseCode      = 0xc2 // 11000010
	TrueCode       = 0xc3 // 11000011
	Uint8Code      = 0xcc // 11001100
	Uint16Code     = 0xcd // 11001101
	Uint32Code     = 0xce // 11001110
	Uint64Code     = 0xcf // 11001111
	Int8Code       = 0xd0 // 11010000
	Int16Code      = 0xd1 // 11010001
	Int32Code      = 0xd2 // 11010010
	Int64Code      = 0xd3 // 11010011
	Float32Code    = 0xca // 11001010
	Float64Code    = 0xcb // 11001011
	FixRaw         = 0xa0 // 101xxxxx
	Raw16Code      = 0xda // 11011010
	Raw32Code      = 0xdb // 11011011
)

func ReadMapLength(data []byte) (int, int) {

	if data[0] >= FixMap && data[0] <= FixMap+0xf {
		return int(data[0] - FixMap), 1
	}

	switch data[0] {
	case 0xde: // Map16
		return int(data[1])<<8 | int(data[2]), 3
	case 0xdf: // Map32
		return int(data[1])<<24 | int(data[2])<<16 | int(data[3])<<8 | int(data[4]), 5
	default:
		panic(fmt.Errorf("Invalid map length prefix: %x", data[0]))
	}
}

func ReadArrayLength(data []byte) (int, int) {
	if data[0] >= FixArray && data[0] <= FixArray+0xf {
		return int(data[0] - FixArray), 1
	}

	switch data[0] {
	case 0xdc: // Array16
		return int(data[1])<<8 | int(data[2]), 3
	case 0xdd: // Array32
		return int(data[1])<<24 | int(data[2])<<16 | int(data[3])<<8 | int(data[4]), 5
	default:
		panic(fmt.Errorf("Invalid array length prefix: %x", data[0]))
	}
}

func ReadStringLength(data []byte) (int, int) {
	if data[0] >= FixStr && data[0] <= FixStr+0x1f {
		return int(data[0] - FixStr), 1
	}

	switch data[0] {
	case 0xd9: // Str8
		return int(data[1]), 2
	case 0xda: // Str16
		return int(data[1])<<8 | int(data[2]), 3
	case 0xdb: // Str32
		return int(data[1])<<24 | int(data[2])<<16 | int(data[3])<<8 | int(data[4]), 5
	default:
		panic(fmt.Errorf("Invalid string length prefix: %x", data[0]))
	}
}

func CalculateDataLen(data []byte, code byte) (d []byte) {

	switch {

	case code == Float32Code: // float32
		data = data[6:]
	case code == Float64Code: // float64
		data = data[9:]
	case code == Uint16Code: // uint16
		data = data[3:]
	case code == TrueCode: // True
		data = data[1:]
	case code >= FixStr && code <= FixStr+0x1f: // FixStr
		data = data[1+int(code-FixStr):]
	case code >= FixArray && code <= FixArray+0xf: // FixArray
		data = data[1+int(code-FixStr):]
	default:
		data = data[len(data)-1:]

	}
	return data
}
