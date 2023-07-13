package encoder

// withNil adds a nil value to the MessagePack representation
func (b *Builder) WithNil() *Builder {
	b.data = append(b.data, 0xc0) // nil
	return b
}
