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
