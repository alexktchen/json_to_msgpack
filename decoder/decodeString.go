package decoder

import "fmt"

func (d *MsgpackDecoder) decodeString(length int) ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.data[:length])), nil
}
