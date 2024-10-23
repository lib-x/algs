package base62

import (
	"errors"
)

const base = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const b = 62

var decodeMap [256]byte

func init() {
	// init map
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(base); i++ {
		decodeMap[base[i]] = byte(i)
	}
}

// Encode uint64 to string
func Encode(n uint64) string {
	if n == 0 {
		return string(base[0])
	}

	chars := make([]byte, 0, 11)
	for n > 0 {
		chars = append(chars, base[n%b])
		n = n / b
	}

	// reverse
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// Decode base62 string to uint64
func Decode(s string) (uint64, error) {
	var n uint64
	for i := 0; i < len(s); i++ {
		dec := decodeMap[s[i]]
		if dec == 0xFF {
			return 0, errors.New("invalid character in Base62 string")
		}
		n = n*b + uint64(dec)
	}
	return n, nil
}
