package builders

// withBool adds a boolean value to the MessagePack representation
func (b *Builder) WithBool(value bool) *Builder {
	if value {
		b.data = append(b.data, 0xc3) // true
	} else {
		b.data = append(b.data, 0xc2) // false
	}
	return b
}
