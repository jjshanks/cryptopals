package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		output string
	}{
		{"happy small", []byte{77, 97, 110}, "TWFu"},
		{"happy long", []byte{72, 101, 108, 108, 111, 33}, "SGVsbG8h"},
		{"small pad", []byte{72, 101, 108, 108, 111}, "SGVsbG8="},
		{"large pad", []byte{71, 111, 111, 100}, "R29vZA=="},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, Encode(test.input), test.name)
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		output      []byte
		errExpected bool
	}{
		{"simple", "ABCD", []byte{0x0, 0x10, 0x83}, false},
		{"padding", "R29vZA==", []byte{71, 111, 111, 100}, false},
		{"padding2", "SGVsbG8=", []byte{72, 101, 108, 108, 111}, false},
		{"bad char", "abc&", nil, true},
		{"bad length", "a", nil, true},
	}

	for _, test := range tests {
		actual, err := Decode(test.input)
		if test.errExpected {
			assert.Error(t, err, test.name)
		} else {
			assert.NoError(t, err, test.name)
			assert.Equal(t, test.output, actual, test.name)
		}

	}
}
