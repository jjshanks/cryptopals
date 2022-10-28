package base64

const (
	pad = "="
)

var (
	charMap = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5",
		"6", "7", "8", "9", pad}
)

func Encode(ba []byte) string {
	result := ""
	for i := 0; i < len(ba); i += 3 {
		result += processChunk(ba, i)
	}
	return result
}

func processChunk(ba []byte, start int) string {
	first := (ba[start] & 0xfc) >> 2
	second := (ba[start] & 0x3) << 4
	if len(ba) <= start+1 {
		return charMap[first] + charMap[second] + pad + pad
	}
	second |= ba[start+1] >> 4
	third := (ba[start+1] & 0xf) << 2
	if len(ba) <= start+2 {
		return charMap[first] + charMap[second] + charMap[third] + pad
	}
	third |= (ba[start+2] & 0xc0) >> 6
	fourth := ba[start+2] & 0x3f

	return charMap[first] + charMap[second] + charMap[third] + charMap[fourth]
}
