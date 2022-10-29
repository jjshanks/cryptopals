package scoring

import (
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
		' ': 10, // not actual freq
	}
)

// EnglishScore returns a score representing how likely
// the input is to be english. Higher scores mean it is
// more likely
func EnglishScore(input string) float64 {
	input = strings.ToLower(input)
	total := 0.0
	for _, r := range input {
		// penalize controls and extended ascii
		if r < ' ' || r > '~' {
			total -= 100.0
		}
		// penalize unlikely characters
		if r >= '!' && r <= '&' || r >= '{' {
			total -= 20.0
		}
		total += freq[r]
	}
	return total
}
