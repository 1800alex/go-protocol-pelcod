package pelcod

// xorChecksum defines a custom xorChecksum object.
type xorChecksum struct {
	sum byte
}

// Reset clears the xorChecksum value to zero
func (xor *xorChecksum) Reset() *xorChecksum {
	xor.sum = 0
	return xor
}

// PushByte pushes a single byte into the xorChecksum
func (xor *xorChecksum) PushByte(b byte) *xorChecksum {
	xor.sum ^= b
	return xor
}

// PushBytes pushes an array of bytes into the xorChecksum
func (xor *xorChecksum) PushBytes(data []byte) *xorChecksum {
	var b byte
	for _, b = range data {
		xor.sum ^= b
	}
	return xor
}

// Value returns the current calculated xorChecksum
func (xor *xorChecksum) Value() byte {
	return xor.sum
}
