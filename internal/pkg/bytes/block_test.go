package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		chunkSize   int
		expected    [][]byte
		errExpected bool
	}{
		{"simple", []byte{0x1, 0x2, 0x3}, 2, [][]byte{{0x1, 0x2}, {0x3}}, false},
		{"larger chunk", []byte{0x1}, 5, [][]byte{{0x1}}, false},
		{"multi chunks", []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}, 3, [][]byte{{0x1, 0x2, 0x3}, {0x4, 0x5, 0x6}, {0x7, 0x8}}, false},
		{"err chunk", []byte{0x1}, -1, nil, true},
	}

	for _, tc := range tests {
		actual, err := Chunk(tc.input, tc.chunkSize)
		if tc.errExpected {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expected, actual, tc.name)
		}
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		name        string
		input       [][]byte
		expected    [][]byte
		expectedErr bool
	}{
		{"simple", [][]byte{{0x1, 0x2}, {0x3, 0x4}}, [][]byte{{0x1, 0x3}, {0x2, 0x4}}, false},
		{"odd", [][]byte{{0x1, 0x2, 0xff, 0xaa}, {0x3, 0x4, 0xfe}}, [][]byte{{0x1, 0x3}, {0x2, 0x4}, {0xff, 0xfe}, {0xaa}}, false},
		{"nil input", nil, nil, true},
		{"empty input", [][]byte{}, nil, true},
	}
	for _, tc := range tests {
		actual, err := Transpose(tc.input)
		if tc.expectedErr {
			assert.Error(t, err, tc.name)
		} else {
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expected, actual, tc.name)
		}
	}
}
