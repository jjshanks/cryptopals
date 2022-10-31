package bytes

import "errors"

func HammingDistance(left, right []byte) (total int, _ error) {
	if left == nil || right == nil {
		return 0, errors.New("inputs must not be nil")
	}
	if len(left) != len(right) {
		return 0, errors.New("inputs must be the same length")
	}
	for i := 0; i < len(left); i += 1 {
		total += calcHammingDistance(left[i], right[i])
	}
	return
}

func calcHammingDistance(left, right byte) (total int) {
	mask := byte(1)
	for i := 0; i < 8; i += 1 {
		if left&(mask<<i)^right&(mask<<i) > 0 {
			total += 1
		}
	}
	return
}
