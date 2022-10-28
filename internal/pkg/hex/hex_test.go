package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		input       string
		expected    []byte
		errExpected bool
		name        string
	}{
		{"", nil, true, "blank"},
		{"0", nil, true, "odd"},
		{"0001", []byte{0, 1}, false, "simple happy path"},
		{"8af2", []byte{0x8a, 0xf2}, false, "less simple happy path"},
		{"8zf2", nil, true, "char out of range"},
	}

	for _, tc := range tests {
		actual, err := Decode(tc.input)
		if tc.errExpected {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tc.expected, actual)
		}
	}
}
