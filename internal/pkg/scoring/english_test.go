package scoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnglishScore(t *testing.T) {
	input := "it was the best of times it was the worst of times"
	assert.InDelta(t, 366.1, EnglishScore(input), 0.1)
	input = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	assert.InDelta(t, 2787.66, EnglishScore(input), 0.1)
	input = "!.? ,158"
	assert.Equal(t, -10.0, EnglishScore(input))
	input = "@#$%^&*()"
	assert.InDelta(t, -80.0, EnglishScore(input), 0.1)
	input = "Ûêö"
	assert.InDelta(t, -360.0, EnglishScore(input), 0.1)
}
