package encoder

// withString adds a string value to the MessagePack representation
func (b *Builder) WithString(value string) *Builder {
	strBytes := []byte(value)
	length := len(strBytes)
	if length <= 31 {
		b.data = append(b.data, 0xa0|byte(length))
	} else if length <= 255 {
		b.data = append(b.data, 0xd9, byte(length))
	} else {
		b.data = append(b.data, 0xda, byte(length>>8), byte(length))
	}
	b.data = append(b.data, strBytes...)
	return b
}
