package bitwise

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedXOR(t *testing.T) {
	tests := []struct {
		name        string
		left        ByteBuffer
		right       ByteBuffer
		expected    []byte
		expectedErr bool
	}{
		{"simple", bytes.NewBuffer([]byte{0x3}), bytes.NewBuffer([]byte{0x2}), []byte{0x1}, false},
		{"long simple", bytes.NewBuffer([]byte{0x00, 0x01, 0x2, 0x3}), bytes.NewBuffer([]byte{0xff, 0xfe, 0xfd, 0xfc}), []byte{0xff, 0xff, 0xff, 0xff}, false},
		{"diff lengths", bytes.NewBuffer([]byte{0x33, 0x33}), bytes.NewBuffer([]byte{0x22}), nil, true},
		{"left err", mockError([]byte{0x1}), bytes.NewBuffer([]byte{0x1, 0x2}), nil, true},
		{"right err", bytes.NewBuffer([]byte{0x1, 0x2}), mockError([]byte{0x1}), nil, true},
	}
	for _, tc := range tests {
		actual, err := FixedXOR(tc.left, tc.right)
		if tc.expectedErr {
			assert.Error(t, err, tc.name)
		} else {
			assert.Equal(t, tc.expected, actual, tc.name)
		}
	}
}

func TestRepeatingXOR(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		key         []byte
		expected    []byte
		expectedErr bool
	}{
		{"simple", []byte{0x3}, []byte{0x2}, []byte{0x1}, false},
		{"repeated key", []byte{0x00, 0x01, 0x02, 0x03}, []byte{0xff}, []byte{0xff, 0xfe, 0xfd, 0xfc}, false},
		{"input nil", nil, []byte{0x1}, nil, true},
		{"key nil", []byte{0x1}, nil, nil, true},
	}
	for _, tc := range tests {
		actual, err := RepeatingXOR(tc.input, tc.key)
		if tc.expectedErr {
			assert.Error(t, err, tc.name)
		} else {
			assert.Equal(t, tc.expected, actual, tc.name)
		}
	}
}

func mockError(bytesBeforeErr []byte) *MockBuffer {
	orgBuffer := bytes.NewBuffer(bytesBeforeErr)
	return &MockBuffer{orgBuffer}
}

type MockBuffer struct {
	b *bytes.Buffer
}

func (m *MockBuffer) Len() int {
	return m.b.Len() + 1
}

func (m *MockBuffer) ReadByte() (byte, error) {
	return m.b.ReadByte()
}
