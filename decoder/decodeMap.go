package decoder

import "github.com/alexktchen/json_to_msgpack/unitily"

func (d *MsgpackDecoder) decodeMap(code byte) ([]byte, error) {

	length := int(code - unitily.FixMap)

	jsonData := []byte("{")
	for i := 0; i < length; i++ {
		key, err := d.withValue()
		if err != nil {
			return nil, err
		}
		//d.data = d.data[len(key)-1:]

		value, err := d.withValue()
		if err != nil {
			return nil, err
		}
		//d.data = unitily.CalculateDataLen(d.data, code)

		jsonData = append(jsonData, key...)
		jsonData = append(jsonData, ':')
		jsonData = append(jsonData, value...)

		if i < length-1 {
			jsonData = append(jsonData, ',')
		}
	}
	jsonData = append(jsonData, '}')

	return jsonData, nil
}
