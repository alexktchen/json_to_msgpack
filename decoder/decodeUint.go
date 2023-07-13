package decoder

import "fmt"

func (d *MsgpackDecoder) decodeUint(size int) ([]byte, error) {
	d.data = d.data[1:]
	var value uint64
	for i := 0; i < size; i++ {
		value = (value << 8) | uint64(d.data[i])
	}
	d.data = d.data[size:]
	return []byte(fmt.Sprintf("%d", value)), nil
}
