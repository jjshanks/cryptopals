package hex

import "errors"

// Normally we'd use the built in hex package but since the goal of this
// project is learning crypto fundamentals wanted to write my own
func Decode(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("input can not be blank")
	}
	if len(s)%2 != 0 {
		return nil, errors.New("input must be even length")
	}
	ba := make([]byte, len(s)/2)
	for i, r := range s {
		b, err := runeToByte(r)
		if err != nil {
			return nil, err
		}
		if i%2 == 0 {
			b = b << 4
		}
		ba[i/2] |= b
	}
	return ba, nil
}

func runeToByte(r rune) (byte, error) {
	if r >= '0' && r <= '9' {
		return byte(r - '0'), nil
	}
	if r >= 'a' && r <= 'f' {
		return byte(r-'a') + 10, nil
	}
	return 0, errors.New("rune must be one of 0123456789abcdef")
}
