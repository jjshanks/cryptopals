package base64

import "errors"

const (
	pad = "="
)

var (
	charMap = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9"}
	byteMap = map[string]byte{"A": 0x0, "B": 0x1, "C": 0x2, "D": 0x3, "E": 0x4, "F": 0x5, "G": 0x6, "H": 0x7, "I": 0x8,
		"J": 0x9, "K": 0xa, "L": 0xb, "M": 0xc, "N": 0xd, "O": 0xe, "P": 0xf, "Q": 0x10, "R": 0x11, "S": 0x12,
		"T": 0x13, "U": 0x14, "V": 0x15, "W": 0x16, "X": 0x17, "Y": 0x18, "Z": 0x19, "a": 0x1a, "b": 0x1b, "c": 0x1c,
		"d": 0x1d, "e": 0x1e, "f": 0x1f, "g": 0x20, "h": 0x21, "i": 0x22, "j": 0x23, "k": 0x24, "l": 0x25, "m": 0x26,
		"n": 0x27, "o": 0x28, "p": 0x29, "q": 0x2a, "r": 0x2b, "s": 0x2c, "t": 0x2d, "u": 0x2e, "v": 0x2f, "w": 0x30,
		"x": 0x31, "y": 0x32, "z": 0x33, "0": 0x34, "1": 0x35, "2": 0x36, "3": 0x37, "4": 0x38, "5": 0x39, "6": 0x3a,
		"7": 0x3b, "8": 0x3c, "9": 0x3d, "+": 0x3e, "/": 0x3f, "=": 0x0}
)

func Decode(s string) ([]byte, error) {
	if len(s)%4 != 0 {
		return nil, errors.New("input must have a length that is a multiple of 4")
	}
	result := make([]byte, 0)
	var intermediate int
	// each input character represents 6 bits
	// 4 characters can be 3 bytes unless there is padding (=)
	// base64 - 0b111111 222222 333333 444444
	// 111111 << 6*3
	// 222222 << 6*2
	// 333333 << 6
	// 444444
	// bytes  - 0b11111122 22223333 33444444
	for i := 0; i < len(s); i += 4 {
		// store all bits in an int for easier retrieval
		intermediate = 0
		// for each character
		// make sure it is valid
		// bitwise or with the intermediate after shifting to the correct position
		c1 := s[i : i+1]
		v1, ok := byteMap[c1]
		if !ok {
			return nil, errors.New("invalid character")
		}
		intermediate |= int(v1) << (6 * 3)

		c2 := s[i+1 : i+2]
		v2, ok := byteMap[c2]
		if !ok {
			return nil, errors.New("invalid character")
		}
		intermediate |= int(v2) << (6 * 2)

		c3 := s[i+2 : i+3]
		// if 3rd character is padding the end has been reached
		// and only the first byte is needed
		if string(c3) == pad {
			result = append(result, byte((intermediate>>16)&0xff))
			break
		}
		v3, ok := byteMap[c3]
		if !ok {
			return nil, errors.New("invalid character")
		}
		intermediate |= int(v3) << 6

		c4 := s[i+3 : i+4]
		v4, ok := byteMap[c4]
		if !ok {
			return nil, errors.New("invalid character")
		}
		intermediate |= int(v4)
		result = append(result, byte((intermediate>>16)&0xff), byte((intermediate>>8)&0xff))
		// only add the last byte if the last character isn't padding
		if c4 != pad {
			result = append(result, byte(intermediate&0xff))
		}
	}
	return result, nil
}

func Encode(ba []byte) string {
	result := ""
	for i := 0; i < len(ba); i += 3 {
		result += encodeChunk(ba, i)
	}
	return result
}

func encodeChunk(ba []byte, start int) string {
	// 3 bytes are turned into 4 base64 chars
	// 01234567 890abcde fghijklm
	// 012345 67890a bcdefg hijklm
	// first char is the upper 6 bits of byte 0
	first := (ba[start] & 0xfc) >> 2
	// second char starts out as the lower 2 bits of byte 0
	second := (ba[start] & 0x3) << 4
	// add padding if no more bytes to process
	if len(ba) <= start+1 {
		return charMap[first] + charMap[second] + pad + pad
	}
	// if byte 1 is given then add the upper 4 bits to second
	second |= ba[start+1] >> 4
	// third char starts out as the lower 4 bits of byte 1
	third := (ba[start+1] & 0xf) << 2
	// add padding if no more bytes to process
	if len(ba) <= start+2 {
		return charMap[first] + charMap[second] + charMap[third] + pad
	}
	// if byte 3 is given then add the upper 2 bits to third
	third |= (ba[start+2] & 0xc0) >> 6
	// fourth if the lower 6 bits of byte 3
	fourth := ba[start+2] & 0x3f

	return charMap[first] + charMap[second] + charMap[third] + charMap[fourth]
}
