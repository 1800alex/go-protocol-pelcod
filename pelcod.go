package pelcod

import ()

func build(start byte, data []byte) (result []byte) {
	result = make([]byte, len(data)+3)

	//var b byte

	offset := 0
	result[0] = start
	offset++

	i := 0
	xor := xorChecksum{}
	xor.Reset()

	for i = 0; i < len(data); i++ {

		if (data[i] == STX) || (data[i] == ACK) || (data[i] == ETX) || (data[i] == ESC) {
			result = append(result, 0)

			result[i+offset] = ESC
			offset++
			result[i+offset] = data[i] | 0x80
		} else {
			result[i+offset] = data[i]
		}

		xor.PushByte(data[i])
	}

	xorsum := xor.Value()

	if (xorsum == STX) || (xorsum == ACK) || (xorsum == ETX) || (xorsum == ESC) {
		result = append(result, 0)

		result[i+offset] = ESC
		offset++
		result[i+offset] = xorsum | 0x80
	} else {
		result[i+offset] = xorsum
	}

	offset++
	result[i+offset] = ETX

	return
}

// BuildSTX returns a PelcoD encoded byte array
// from the incoming byte array beginning with the
// STX byte.
func BuildSTX(data []byte) []byte {
	return build(STX, data)
}

// BuildACK returns a PelcoD encoded byte array
// from the incoming byte array beginning with the
// ACK byte.
func BuildACK(data []byte) []byte {
	return build(ACK, data)
}

// Parse returns a byte array from the incoming PelcoD
// encoded byte array.
func Parse(data []byte) (output []byte, err error) {
	err = nil

	dataLen := len(data)

	if dataLen <= 3 {
		err = ErrParseInvalidLength
		return
	}

	result := make([]byte, dataLen-3)
	resultLen := 0

	if data[0] == STX {
		//is STX
	} else if data[0] == ACK {
		//is ACK
	} else {
		err = ErrParseNoSTX
		return
	}

	if data[dataLen-1] == ETX {
		//is ETX
	} else {
		err = ErrParseNoETX
		return
	}

	i := 0
	xor := xorChecksum{}
	xor.Reset()
	escapeNext := false

	var dataxor byte

	for i = 1; i < dataLen-1; i++ {

		if data[i] == ESC {
			if escapeNext == true {
				err = ErrParseDuplicateESC
				return
			}

			escapeNext = true
		} else if true == escapeNext {
			if (data[i] & 0x80) == 0 {
				err = ErrParseInvalidEscaped
				return
			}

			escapeNext = false

			if i < dataLen-2 {
				result[resultLen] = data[i] & 0x7F
				xor.PushByte(result[resultLen])
				resultLen++
			} else {
				dataxor = data[i] & 0x7F
			}
		} else {

			if i < dataLen-2 {
				result[resultLen] = data[i]
				xor.PushByte(result[resultLen])
				resultLen++
			} else {
				dataxor = data[i]
			}

		}
	}

	xorsum := xor.Value()

	if xorsum != dataxor {
		err = ErrParseInvalidChecksum
		return
	}

	//todo this is inefficient
	output = make([]byte, resultLen)

	for i = 0; i < resultLen; i++ {
		output[i] = result[i]
	}

	return output, nil
}
