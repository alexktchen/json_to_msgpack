package decoder

import "github.com/alexktchen/json_to_msgpack/unitily"

func (d *MsgpackDecoder) decodeArray(code byte) ([]byte, error) {

	length := int(code - unitily.FixArray)

	jsonData := []byte("[")
	for i := 0; i < length; i++ {

		value, err := d.withValue()
		if err != nil {
			return nil, err
		}

		jsonData = append(jsonData, value...)

		if i < length-1 {
			jsonData = append(jsonData, ',')
		}
	}
	jsonData = append(jsonData, ']')

	return jsonData, nil
}
