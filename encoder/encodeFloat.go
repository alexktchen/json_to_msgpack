package encoder

import (
	"math"
)

// withFloat adds a floating-point value to the MessagePack representation
func (b *Builder) WithFloat(value float64) *Builder {
	b.data = append(b.data, 0xcb)
	bits := math.Float64bits(value)
	b.data = append(b.data,
		byte(bits>>56),
		byte(bits>>48),
		byte(bits>>40),
		byte(bits>>32),
		byte(bits>>24),
		byte(bits>>16),
		byte(bits>>8),
		byte(bits),
	)
	return b
}
