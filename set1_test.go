package cryptopals

import (
	"bufio"
	xbytes "bytes"
	"os"
	"testing"

	"github.com/jjshanks/cryptopals/internal/pkg/base64"
	"github.com/jjshanks/cryptopals/internal/pkg/bitwise"
	"github.com/jjshanks/cryptopals/internal/pkg/cryptanalysis"
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
	solution, err := cryptanalysis.SingleCharXOR(inputBytes)
	require.NoError(t, err)
	assert.Equal(t, "Cooking MC's like a pound of bacon", solution.Text)
}

func TestChallenge4(t *testing.T) {
	file, err := os.Open("data/challenge4.txt")
	require.NoError(t, err)
	defer file.Close()

	bestSolution := cryptanalysis.XORCryptanalysisSolution{
		Score: 0,
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputBytes, err := hex.Decode(line)
		require.NoError(t, err)
		solution, err := cryptanalysis.SingleCharXOR(inputBytes)
		require.NoError(t, err)
		if bestSolution.Score < solution.Score {
			bestSolution = solution
		}
	}
	require.NoError(t, scanner.Err())
	assert.Equal(t, "Now that the party is jumping\n", bestSolution.Text)
}
