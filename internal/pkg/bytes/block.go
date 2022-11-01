package bytes

import (
	"errors"
	"math"
)

func Chunk(input []byte, chunkSize int) ([][]byte, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunk size must be greater than zero")
	}
	chunks := int(math.Ceil(float64(len(input)) / float64(chunkSize)))
	result := make([][]byte, chunks)
	for i := 0; i < chunks; i += 1 {
		end := (i + 1) * chunkSize
		if end > len(input) {
			end = len(input)
		}
		result[i] = input[i*chunkSize : end]
	}
	return result, nil
}
