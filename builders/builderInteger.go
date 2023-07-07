package builders

// withInt adds an integer value to the MessagePack representation
func (b *Builder) WithInt(value int64) *Builder {
	if value >= 0 {
		if value <= 0x7f {
			b.data = append(b.data, byte(value)) // positive fixint
		} else if value <= 0xff {
			b.data = append(b.data, 0xcc, byte(value)) // uint 8
		} else if value <= 0xffff {
			b.data = append(b.data, 0xcd, byte(value>>8), byte(value)) // uint 16
		} else if value <= 0xffffffff {
			b.data = append(b.data, 0xce, byte(value>>24), byte(value>>16), byte(value>>8), byte(value)) // uint 32
		} else {
			b.data = append(b.data, 0xcf, byte(value>>56), byte(value>>48), byte(value>>40), byte(value>>32),
				byte(value>>24), byte(value>>16), byte(value>>8), byte(value)) // uint 64
		}
	} else {
		if value >= -32 {
			b.data = append(b.data, byte(value)) // negative fixint
		} else if value >= -128 {
			b.data = append(b.data, 0xd0, byte(value)) // int 8
		} else if value >= -32768 {
			b.data = append(b.data, 0xd1, byte(value>>8), byte(value)) // int 16
		} else if value >= -2147483648 {
			b.data = append(b.data, 0xd2, byte(value>>24), byte(value>>16), byte(value>>8), byte(value)) // int 32
		} else {
			b.data = append(b.data, 0xd3, byte(value>>56), byte(value>>48), byte(value>>40), byte(value>>32),
				byte(value>>24), byte(value>>16), byte(value>>8), byte(value)) // int 64
		}
	}
	return b
}
