package builders

// withArray adds an array value to the MessagePack representation
func (b *Builder) WithArray(values []interface{}) *Builder {
	arrayBuilder := &Builder{}
	for _, value := range values {
		arrayBuilder.withValue(value)
	}
	arrayData := arrayBuilder.build()

	length := len(values)
	if length <= 15 {
		b.data = append(b.data, 0x90|byte(length))
	} else if length <= 65535 {
		b.data = append(b.data, 0xdc, byte(length>>8), byte(length))
	} else {
		b.data = append(b.data, 0xdd, byte(length>>24), byte(length>>16), byte(length>>8), byte(length))
	}
	b.data = append(b.data, arrayData...)
	return b
}
