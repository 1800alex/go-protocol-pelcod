package pelcod

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var createBuildSTXAndACKTests = []struct {
	name   string
	buffer []byte
	result []byte
}{
	{
		name:   "no escaped data",
		buffer: []byte{0x00, 0x01, 0x09},
		result: []byte{0x02, 0x00, 0x01, 0x09, 0x08, 0x03},
	},
	{
		name:   "escaped data",
		buffer: []byte{0x00, 0x01, 0x02, 0x03, 0x06, 0x1B, 0x09},
		result: []byte{0x02, 0x00, 0x01, 0x1B, 0x82, 0x1B, 0x83, 0x1B, 0x86, 0x1B, 0x9B, 0x09, 0x14, 0x03},
	},
	{
		name:   "escaped checksum",
		buffer: []byte{0x00, 0x02},
		result: []byte{0x02, 0x00, 0x1B, 0x82, 0x1B, 0x82, 0x03},
	},
}

func TestBuildSTX(t *testing.T) {
	for _, tt := range createBuildSTXAndACKTests {
		t.Run(tt.name, func(t *testing.T) {

			expectedBuffer := []byte{0x02}

			for i := 1; i < len(tt.result); i++ {
				expectedBuffer = append(expectedBuffer, tt.result[i])
			}

			result := BuildSTX(tt.buffer)
			assert.Equal(t, expectedBuffer, result)
		})
	}
}

func TestBuildACK(t *testing.T) {
	for _, tt := range createBuildSTXAndACKTests {
		t.Run(tt.name, func(t *testing.T) {

			expectedBuffer := []byte{0x06}

			for i := 1; i < len(tt.result); i++ {
				expectedBuffer = append(expectedBuffer, tt.result[i])
			}

			result := BuildACK(tt.buffer)
			assert.Equal(t, expectedBuffer, result)
		})
	}
}

var createParseTests = []struct {
	name   string
	buffer []byte
	result []byte
	err    error
}{
	{
		name:   "no escaped data",
		buffer: []byte{0x02, 0x00, 0x01, 0x09, 0x08, 0x03},
		result: []byte{0x00, 0x01, 0x09},
	},
	{
		name:   "escaped data",
		buffer: []byte{0x02, 0x00, 0x01, 0x1B, 0x82, 0x1B, 0x83, 0x1B, 0x86, 0x1B, 0x9B, 0x09, 0x14, 0x03},
		result: []byte{0x00, 0x01, 0x02, 0x03, 0x06, 0x1B, 0x09},
	},
	{
		name:   "escaped checksum",
		buffer: []byte{0x02, 0x00, 0x1B, 0x82, 0x1B, 0x82, 0x03},
		result: []byte{0x00, 0x02},
	},
	{
		name:   "too short",
		buffer: []byte{0x02, 0x00, 0x03},
		err:    ErrParseInvalidLength,
	},
	{
		name:   "no stx or ack",
		buffer: []byte{0x01, 0x00, 0x01, 0x09, 0x08, 0x03},
		err:    ErrParseNoSTX,
	},
	{
		name:   "no etx",
		buffer: []byte{0x02, 0x00, 0x01, 0x09, 0x08, 0x04},
		err:    ErrParseNoETX,
	},
	{
		name:   "duplicate esc",
		buffer: []byte{0x02, 0x00, 0x1B, 0x1B, 0x1B, 0x82, 0x03},
		err:    ErrParseDuplicateESC,
	},
	{
		name:   "invalid escaped byte",
		buffer: []byte{0x02, 0x00, 0x1B, 0x01, 0x02, 0x82, 0x03},
		err:    ErrParseInvalidEscaped,
	},
	{
		name:   "invalid checksum",
		buffer: []byte{0x02, 0x00, 0x1B, 0x82, 0x02, 0x83, 0x03},
		err:    ErrParseInvalidChecksum,
	},
}

func TestParse_STX(t *testing.T) {
	for _, tt := range createParseTests {
		t.Run(tt.name, func(t *testing.T) {

			inputBuffer := []byte{}
			start := 0

			if tt.buffer[0] == STX {
				inputBuffer = append(inputBuffer, STX)
				start++
			}

			for i := start; i < len(tt.buffer); i++ {
				inputBuffer = append(inputBuffer, tt.buffer[i])
			}

			result, err := Parse(inputBuffer)

			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, result)
			}
		})
	}
}

func TestParse_ACK(t *testing.T) {
	for _, tt := range createParseTests {
		t.Run(tt.name, func(t *testing.T) {

			inputBuffer := []byte{}
			start := 0

			if tt.buffer[0] == STX {
				inputBuffer = append(inputBuffer, ACK)
				start++
			}

			for i := start; i < len(tt.buffer); i++ {
				inputBuffer = append(inputBuffer, tt.buffer[i])
			}

			result, err := Parse(inputBuffer)

			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, result)
			}
		})
	}
}

func ExampleBuildSTX() {
	result := BuildSTX([]byte{0x00, 0x01, 0x09})

	fmt.Printf("%x\n", result)
}

func ExampleBuildACK() {
	result := BuildACK([]byte{0x00, 0x01, 0x09})

	fmt.Printf("%x\n", result)
}

func ExampleParse() {
	result, err := Parse([]byte{0x02, 0x00, 0x01, 0x09, 0x08, 0x03})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%x\n", result)
}

func BenchmarkBuildSTX(b *testing.B) {
	var result []byte

	for n := 0; n < b.N; n++ {
		result = BuildSTX([]byte{
			0x80, 0xf2, 0xd8, 0x75, 0xfe, 0x6a, 0xd3,
			0x21, 0x24, 0x30, 0x8e, 0x3a, 0x12, 0x90,
			0x34, 0x86, 0x39, 0x2a, 0x97, 0x1a, 0x32,
			0xb0, 0x1b, 0xd8, 0xa1, 0xae, 0x76, 0x71,
			0x42, 0x52, 0x15, 0xc8, 0x16, 0x8d, 0xa8,
			0xf2, 0x81, 0x84, 0x73, 0x37, 0x81, 0x95,
			0x26, 0xdb, 0x71, 0xd2, 0x8a, 0xb0, 0x67})
	}

	if len(result) < 1 {
		b.Fatalf("something went wrong")
	}
}

func BenchmarkBuildACK(b *testing.B) {
	var result []byte

	for n := 0; n < b.N; n++ {
		result = BuildACK([]byte{
			0x80, 0xf2, 0xd8, 0x75, 0xfe, 0x6a, 0xd3,
			0x21, 0x24, 0x30, 0x8e, 0x3a, 0x12, 0x90,
			0x34, 0x86, 0x39, 0x2a, 0x97, 0x1a, 0x32,
			0xb0, 0x1b, 0xd8, 0xa1, 0xae, 0x76, 0x71,
			0x42, 0x52, 0x15, 0xc8, 0x16, 0x8d, 0xa8,
			0xf2, 0x81, 0x84, 0x73, 0x37, 0x81, 0x95,
			0x26, 0xdb, 0x71, 0xd2, 0x8a, 0xb0, 0x67})
	}

	if len(result) < 1 {
		b.Fatalf("something went wrong")
	}
}

func BenchmarkBuildParse(b *testing.B) {
	var result []byte
	var err error

	for n := 0; n < b.N; n++ {
		result, err = Parse([]byte{
			0x06, 0x80, 0xF2, 0xD8, 0x75, 0xFE, 0x6A,
			0xD3, 0x21, 0x24, 0x30, 0x8E, 0x3A, 0x12,
			0x90, 0x34, 0x86, 0x39, 0x2A, 0x97, 0x1A,
			0x32, 0xB0, 0x1B, 0x9B, 0xD8, 0xA1, 0xAE,
			0x76, 0x71, 0x42, 0x52, 0x15, 0xC8, 0x16,
			0x8D, 0xA8, 0xF2, 0x81, 0x84, 0x73, 0x37,
			0x81, 0x95, 0x26, 0xDB, 0x71, 0xD2, 0x8A,
			0xB0, 0x67, 0xA4, 0x03})
	}

	if err != nil {
		b.Fatalf(err.Error())
	}

	if len(result) < 1 {
		b.Fatalf("something went wrong")
	}
}
