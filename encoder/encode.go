package encoder

// Builder builds the MessagePack representation of a value
type Builder struct {
	data []byte
}

// Build constructs the final MessagePack representation
func (b *Builder) Build() []byte {
	return b.data
}
