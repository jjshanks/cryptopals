package cryptanalysis

import (
	xbytes "bytes"
	"io"

	"github.com/jjshanks/cryptopals/internal/pkg/bitwise"
	"github.com/jjshanks/cryptopals/internal/pkg/bytes"
	"github.com/jjshanks/cryptopals/internal/pkg/scoring"
)

// SingleCharXOR scores the input against ever alpha key a single char xor
// returns the best string result, key, and score
func SingleCharXOR(input []byte) (XORCryptanalysisSolution, error) {
	bestScore := 0.0
	bestAns := ""
	bestKey := byte(0)
	for key := 0; key < 256; key += 1 {
		xor, err := bitwise.FixedXOR(xbytes.NewBuffer(input), &bytes.FixedBuffer{Fix: byte(key), Length: len(input)})
		if err != nil {
			return XORCryptanalysisSolution{}, err
		}
		score := scoring.EnglishScore(string(xor))
		if score > bestScore {
			bestScore = score
			bestAns = string(xor)
			bestKey = byte(key)
		}
	}
	return XORCryptanalysisSolution{bestAns, bestScore, bestKey}, nil
}

type XORCryptanalysisSolution struct {
	Text  string
	Score float64
	Key   byte
}

type SeekableReader interface {
	Read([]byte) (int, error)
	Seek(int64, int) (int64, error)
}

func RepeatingXORKeySize(input SeekableReader, min, max int) (int, error) {
	bestNorm := float64(1 << 31)
	bestKeySize := max + 1
	for i := min; i <= max; i += 1 {
		first := make([]byte, i)
		second := make([]byte, i)
		_, err := io.ReadFull(input, first)
		if err != nil {
			return 0, err
		}
		_, err = io.ReadFull(input, second)
		if err != nil {
			return 0, err
		}
		actual, err := bytes.HammingDistance(first, second)
		if err != nil {
			return 0, err
		}
		norm := float64(actual) / float64(i)
		if norm < bestNorm {
			bestNorm = norm
			bestKeySize = i
		}
		input.Seek(0, 0)
	}
	return bestKeySize, nil
}
