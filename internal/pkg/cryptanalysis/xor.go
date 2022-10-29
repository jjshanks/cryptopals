package cryptanalysis

import (
	xbytes "bytes"

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
