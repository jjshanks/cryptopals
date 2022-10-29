package cryptopals

import (
	xbytes "bytes"
	"testing"

	"github.com/jjshanks/cryptopals/internal/pkg/base64"
	"github.com/jjshanks/cryptopals/internal/pkg/bitwise"
	"github.com/jjshanks/cryptopals/internal/pkg/bytes"
	"github.com/jjshanks/cryptopals/internal/pkg/hex"
	"github.com/jjshanks/cryptopals/internal/pkg/scoring"
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
	leftBuffer := xbytes.NewBuffer(leftBytes)
	require.NoError(t, err)
	right := "686974207468652062756c6c277320657965"
	rightBytes, err := hex.Decode(right)
	rightBuffer := xbytes.NewBuffer(rightBytes)
	require.NoError(t, err)
	expected := "746865206b696420646f6e277420706c6179"
	expectedBytes, err := hex.Decode(expected)
	require.NoError(t, err)
	actual, err := bitwise.FixedXOR(leftBuffer, rightBuffer)
	require.NoError(t, err)
	assert.Equal(t, expectedBytes, actual)
}

func TestChallenge3(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputBytes, err := hex.Decode(input)
	require.NoError(t, err)
	bestScore := float64(1 << 31)
	bestAns := ""
	for _, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		xor, err := bitwise.FixedXOR(xbytes.NewBuffer(inputBytes), &bytes.FixedBuffer{Fix: byte(c), Length: len(inputBytes)})
		require.NoError(t, err)
		score := scoring.EnglishScore(string(xor))
		if score < float64(bestScore) {
			bestScore = score
			bestAns = string(xor)
		}
	}
	assert.Equal(t, "Cooking MC's like a pound of bacon", bestAns)
}
