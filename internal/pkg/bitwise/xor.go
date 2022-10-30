package bitwise

import (
	"errors"
)

type ByteBuffer interface {
	Len() int
	ReadByte() (byte, error)
}

// FixedXOR xor's two buffers of the same length together
func FixedXOR(left, right ByteBuffer) ([]byte, error) {
	if left.Len() != right.Len() {
		return nil, errors.New("inputs must be the same length")
	}
	result := make([]byte, left.Len())
	bytes := left.Len()
	for i := 0; i < bytes; i += 1 {
		lb, le := left.ReadByte()
		if le != nil {
			return nil, errors.New("unable to read input")
		}
		rb, re := right.ReadByte()
		if re != nil {
			return nil, errors.New("unable to read input")
		}
		result[i] = lb ^ rb
	}
	return result, nil
}

// RepeatingXOR applies the provided key to the input like
// result[i] = input[i] ^ key[i % len(key)]
func RepeatingXOR(input, key []byte) ([]byte, error) {
	if input == nil || key == nil {
		return nil, errors.New("input and key must not be nil")
	}
	result := make([]byte, len(input))
	for i := 0; i < len(input); i += 1 {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return result, nil
}
