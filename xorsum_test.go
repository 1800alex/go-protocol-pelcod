package pelcod

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var createXORChecksumTests = []struct {
	name   string
	buffer []byte
	result byte
}{
	{
		name:   "3 bytes",
		buffer: []byte{0x00, 0x01, 0x09},
		result: 0x08,
	},
	{
		name:   "7 bytes",
		buffer: []byte{0x00, 0x01, 0x02, 0x03, 0x06, 0x1B, 0x09},
		result: 0x14,
	},
}

func TestXORChecksum(t *testing.T) {
	for _, tt := range createXORChecksumTests {
		t.Run(tt.name, func(t *testing.T) {

			xor := xorChecksum{}
			xor.Reset()
			xor.PushByte(0x00)

			for i := 0; i < len(tt.buffer); i++ {
				xor.PushByte(tt.buffer[i])
			}

			result := xor.Value()
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestXORChecksum_Multiple(t *testing.T) {
	for _, tt := range createXORChecksumTests {
		t.Run(tt.name, func(t *testing.T) {

			xor := xorChecksum{}
			xor.Reset()
			xor.PushBytes(tt.buffer)

			result := xor.Value()
			assert.Equal(t, tt.result, result)
		})
	}
}
