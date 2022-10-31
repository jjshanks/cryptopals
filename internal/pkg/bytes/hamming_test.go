package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		name        string
		left        []byte
		right       []byte
		expected    int
		errExpected bool
	}{
		{"simple", []byte{0x1}, []byte{0x3}, 1, false},
		{"left nil", nil, []byte{0x1}, 0, true},
		{"right nil", []byte{0x1}, nil, 0, true},
		{"diff lengths", []byte{0x1}, []byte{0x1, 0x1}, 0, true},
		{"long", []byte{0xff, 0x0, 0b10101010}, []byte{0x00, 0xff, 0b01010101}, 8 * 3, false},
	}

	for _, tc := range tests {
		actual, err := HammingDistance(tc.left, tc.right)
		if tc.errExpected {
			assert.Error(t, err, tc.name)
			continue
		}
		assert.NoError(t, err, tc.name)
		assert.Equal(t, tc.expected, actual, tc.name)
	}
}
