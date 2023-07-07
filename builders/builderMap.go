package builders

// withMap adds a map value to the MessagePack representation
func (b *Builder) WithMap(values map[string]interface{}) *Builder {
	mapBuilder := &Builder{}
	for key, value := range values {
		mapBuilder.withValue(key)
		mapBuilder.withValue(value)
	}
	mapData := mapBuilder.build()

	length := len(values)
	if length <= 15 {
		b.data = append(b.data, 0x80|byte(length))
	} else if length <= 65535 {
		b.data = append(b.data, 0xde, byte(length>>8), byte(length))
	} else {
		b.data = append(b.data, 0xdf, byte(length>>24), byte(length>>16), byte(length>>8), byte(length))
	}
	b.data = append(b.data, mapData...)
	return b
}
