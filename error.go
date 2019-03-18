package pelcod

import "errors"

var (
	// ErrParseInvalidLength is the error returned when attempting to parse
	// a buffer that is too short.
	ErrParseInvalidLength = errors.New("invalid length")

	// ErrParseNoSTX is the error returned when attempting to parse
	// a buffer without an STX or ACK.
	ErrParseNoSTX = errors.New("no STX or ACK")

	// ErrParseNoSTX is the error returned when attempting to parse
	// a buffer without an ETX.
	ErrParseNoETX = errors.New("no ETX")

	// ErrParseNoSTX is the error returned when attempting to parse
	// a buffer with back to back ESC bytes.
	ErrParseDuplicateESC = errors.New("duplicate ESC")

	// ErrParseInvalidEscaped is the error returned when attempting to parse
	// a buffer with an ESC byte followed by an invalid byte (0x80 unset).
	ErrParseInvalidEscaped = errors.New("invalid escaped byte")

	// ErrParseInvalidChecksum is the error returned when attempting to parse
	// a buffer with a bad checksum.
	ErrParseInvalidChecksum = errors.New("invalid checksum")
)
