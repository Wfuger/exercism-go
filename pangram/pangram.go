package pangram

import "strings"

const testVersion = 1


func IsPangram(s string) bool {
	alphabet := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	for _, char := range s {
		lower := strings.ToLower(string(char))
		if _, ok := alphabet[lower]; ok {
			delete(alphabet, lower)
		}

	}
	if len(alphabet) == 0 {
		return true
	}
	return false
}
