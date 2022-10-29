package cryptopals

import (
	"bytes"
	"testing"

	"github.com/jjshanks/cryptopals/internal/pkg/base64"
	"github.com/jjshanks/cryptopals/internal/pkg/bitwise"
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

func TestChallenge2(t *testing.T) {
	left := "1c0111001f010100061a024b53535009181c"
	leftBytes, err := hex.Decode(left)
	leftBuffer := bytes.NewBuffer(leftBytes)
	require.NoError(t, err)
	right := "686974207468652062756c6c277320657965"
	rightBytes, err := hex.Decode(right)
	rightBuffer := bytes.NewBuffer(rightBytes)
	require.NoError(t, err)
	expected := "746865206b696420646f6e277420706c6179"
	expectedBytes, err := hex.Decode(expected)
	require.NoError(t, err)
	actual, err := bitwise.FixedXOR(leftBuffer, rightBuffer)
	require.NoError(t, err)
	assert.Equal(t, expectedBytes, actual)
}
