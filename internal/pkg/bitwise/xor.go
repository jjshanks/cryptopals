package bitwise

import (
	"errors"
)

type ByteBuffer interface {
	Len() int
	ReadByte() (byte, error)
}

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
