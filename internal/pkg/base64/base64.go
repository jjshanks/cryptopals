package base64

const (
	pad = "="
)

var (
	charMap = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9"}
)

func Encode(ba []byte) string {
	result := ""
	for i := 0; i < len(ba); i += 3 {
		result += processChunk(ba, i)
	}
	return result
}

func processChunk(ba []byte, start int) string {
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
