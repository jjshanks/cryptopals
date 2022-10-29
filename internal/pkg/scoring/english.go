package scoring

import (
	"math"
	"regexp"
	"strings"
)

var (
	// from https://en.wikipedia.org/wiki/Letter_frequency
	freq = map[rune]float64{
		'a': 8.2,
		'b': 1.5,
		'c': 2.8,
		'd': 4.3,
		'e': 13,
		'f': 2.2,
		'g': 2,
		'h': 6.1,
		'i': 7,
		'j': 0.15,
		'k': 0.77,
		'l': 4,
		'm': 2.4,
		'n': 6.7,
		'p': 1.9,
		'q': 0.095,
		'r': 6,
		's': 6.3,
		't': 9.1,
		'u': 2.8,
		'v': 0.98,
		'w': 2.4,
		'x': 0.15,
		'y': 2,
		'z': 0.074,
	}

	// uncommon but expected so ignore
	skips = regexp.MustCompile(`^[ .?,!'"0-9]$`)
	// rare but possible, treat as anomaly
	lows         = regexp.MustCompile(`^[@#$%^&*()_+=<>/\;-|]$`)
	anomalyScore = 100.0
)

// EnglishScore returns a score representing how close to
// the expected frequency of english characters was observed.
// The closer to zero the better frequency match.
func EnglishScore(input string) float64 {
	input = strings.ToLower(input)
	counts := make(map[rune]float64)
	for _, r := range input {
		counts[r] += 1
	}
	total := 0.0
	for k, v := range counts {
		if _, ok := freq[k]; !ok {
			if skips.Match([]byte{byte(k)}) {
				continue
			}
			if lows.Match([]byte{byte(k)}) {
				total += anomalyScore / 100
				continue
			}
			total += anomalyScore
			continue
		}
		ratio := (v / float64(len(input)) * 100)
		drift := math.Abs(freq[k] - ratio)
		total += drift
	}
	return total
}
