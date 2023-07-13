package decoder

import (
	"fmt"

	"github.com/alexktchen/json_to_msgpack/unitily"
)

func ConvertMessagePackToJson(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty data")
	}
	return decodeValue(data)
}

func decodeValue(value []byte) ([]byte, error) {

	decoder := &MsgpackDecoder{}
	decoder.setData(value)
	return decoder.withValue()
}

func (d *MsgpackDecoder) setData(value []byte) {
	d.data = value
}

func (decoder *MsgpackDecoder) withValue() ([]byte, error) {

	if len(decoder.data) == 0 {
		return nil, fmt.Errorf("The decoder not contains any data")
	}

	code := decoder.data[0]

	switch {
	case code <= unitily.PositiveFixNum: // Positive FixNum
		v := []byte(fmt.Sprintf("%d", code))
		decoder.data = decoder.data[1:]
		return v, nil

	case code >= unitily.FixMap && code <= unitily.FixMap+0xf: // FixMap
		decoder.data = decoder.data[1:]
		return decoder.decodeMap(code)

	case code >= unitily.FixArray && code <= unitily.FixArray+0xf: // FixArray
		decoder.data = decoder.data[1:]
		return decoder.decodeArray(code)

	case code >= unitily.FixStr && code <= unitily.FixStr+0x1f: // FixStr
		decoder.data = decoder.data[1:]
		key, err := decoder.decodeString(int(code - unitily.FixStr))
		decoder.data = decoder.data[len(key)-2:]
		return key, err
	case code == unitily.Float32Code: // float32
		return decoder.decodeFloat()

	case code == unitily.Float64Code: // float64
		decoder.data = decoder.data[1:]
		return decoder.float64FromBytes()
	case code == unitily.NilCode: // Nil
		return []byte("null"), nil

	case code == unitily.FalseCode: // False
		return []byte("false"), nil

	case code == unitily.TrueCode: // True
		return []byte("true"), nil

	case code == unitily.Uint8Code: // uint8
		return decoder.decodeUint(1)

	case code == unitily.Uint16Code: // uint16
		return decoder.decodeUint(2)

	case code == unitily.Uint32Code: // uint32
		return decoder.decodeUint(4)

	case code == unitily.Uint64Code: // uint64
		return decoder.decodeUint(8)

	case code == unitily.Int8Code: // int8
		return decoder.decodeInt(1)

	case code == unitily.Int16Code: // int16
		return decoder.decodeInt(2)

	case code == unitily.Int32Code: // int32
		return decoder.decodeInt(4)

	case code == unitily.Int64Code: // int64
		return decoder.decodeInt(8)

	case code >= unitily.FixRaw && code <= unitily.FixRaw+0x1f: // FixRaw
		return decoder.data[1 : code-unitily.FixRaw+1], nil

	case code == unitily.Raw16Code: // Raw16
		length := uint16(decoder.data[1])<<8 | uint16(decoder.data[2])
		return decoder.data[3 : 3+length], nil

	case code == unitily.Raw32Code: // Raw32
		length := uint32(decoder.data[1])<<24 | uint32(decoder.data[2])<<16 |
			uint32(decoder.data[3])<<8 | uint32(decoder.data[4])
		return decoder.data[5 : 5+length], nil

	case code >= 0xe0 && code <= 0xff: // Negative FixNum
		return []byte(fmt.Sprintf("%d", int8(code))), nil

	default:
		return nil, fmt.Errorf("Unknown data code: %x", code)
	}
}
