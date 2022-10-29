package cryptopals

import (
	"testing"

	"github.com/jjshanks/cryptopals/internal/pkg/base64"
	"github.com/jjshanks/cryptopals/internal/pkg/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChallenge1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	byteInput, err := hex.Decode(input)
	require.NoError(t, err)
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	assert.Equal(t, expectedOutput, base64.Encode(byteInput))
}
