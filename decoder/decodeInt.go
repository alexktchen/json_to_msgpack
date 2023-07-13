package decoder

import "fmt"

func (d *MsgpackDecoder) decodeInt(size int) ([]byte, error) {
	d.data = d.data[1:]
	var value int64
	if d.data[0]&(1<<7) != 0 {
		value = -1 << (size*8 - 1)
	}

	for i := 0; i < size; i++ {
		value = (value << 8) | int64(d.data[i])
	}
	return []byte(fmt.Sprintf("%d", value)), nil

}
