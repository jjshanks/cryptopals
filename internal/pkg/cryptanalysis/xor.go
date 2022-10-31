package cryptanalysis

import (
	xbytes "bytes"
	"io"
	"sort"

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

func RepeatingXORKeySize(input SeekableReader, min, max, candidateCount int) ([]int, error) {
	candidates := make([]keySizeResult, 0)
	for i := min; i <= max; i += 1 {
		first := make([]byte, i)
		second := make([]byte, i)
		_, err := io.ReadFull(input, first)
		if err != nil {
			return nil, err
		}
		_, err = io.ReadFull(input, second)
		if err != nil {
			return nil, err
		}
		actual, err := bytes.HammingDistance(first, second)
		if err != nil {
			return nil, err
		}
		norm := float64(actual) / float64(i)
		candidates = append(candidates, keySizeResult{
			score: norm,
			size:  i,
		})
		input.Seek(0, 0)
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score < candidates[j].score
	})
	result := make([]int, candidateCount)
	for i := 0; i < candidateCount; i += 1 {
		result[i] = candidates[i].size
	}
	return result, nil
}

type keySizeResult struct {
	score float64
	size  int
}
