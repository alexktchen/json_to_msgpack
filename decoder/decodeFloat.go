package decoder

import (
	"fmt"
	"unsafe"
)

func (d *MsgpackDecoder) decodeFloat() ([]byte, error) {

	bytes := d.data[1:5]
	bits := uint32FromBytes(bytes)
	dd := *(*float32)(unsafe.Pointer(&bits))

	return []byte(fmt.Sprintf("%f", dd)), nil
}

func uint32FromBytes(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) | (uint32(bytes[1]) << 16) | (uint32(bytes[2]) << 8) | uint32(bytes[3])
}

func (d *MsgpackDecoder) float64FromBytes() ([]byte, error) {

	bits := uint64FromBytes(d.data)
	dd := *(*float64)(unsafe.Pointer(&bits))
	d.data = d.data[8:]
	return []byte(fmt.Sprintf("%f", dd)), nil
}

func uint64FromBytes(bytes []byte) uint64 {
	return (uint64(bytes[0]) << 56) | (uint64(bytes[1]) << 48) | (uint64(bytes[2]) << 40) |
		(uint64(bytes[3]) << 32) | (uint64(bytes[4]) << 24) | (uint64(bytes[5]) << 16) |
		(uint64(bytes[6]) << 8) | uint64(bytes[7])
}
