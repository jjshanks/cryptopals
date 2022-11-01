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

func Transpose(input [][]byte) ([][]byte, error) {
	if len(input) == 0 || len(input[0]) == 0 {
		return nil, errors.New("input can not be empty")
	}
	result := make([][]byte, len(input[0]))
	for _, chunk := range input {
		for idx, b := range chunk {
			result[idx] = append(result[idx], b)
		}
	}
	return result, nil
}
