package secret

import (
	"strconv"
	//"reflect"
)

const testVersion = 2

func toBinary(num uint) string {
	return strconv.FormatUint(uint64(num), 2)
}

func reverse(ss []interface{}) []interface{} {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

func Handshake(num uint) []string {
	var hs []string
	bin := toBinary(num)
	binRunes := []rune(bin)
	t := make([]interface{}, len(bin))
	for i, v := range binRunes {
		t[i] = interface{}(v)
	}
	reverse(t)
	for i, v := range t {
		binRunes[i] = v.(rune)
	}
	for i, num := range binRunes {
		if i == 0 && num == '1' {
			hs = append(hs, "wink")
		}
		if i == 1 && num == '1' {
			hs = append(hs, "double blink")
		}
		if i == 2 && num == '1' {
			hs = append(hs, "close your eyes")
		}
		if i == 3 && num == '1' {
			hs = append(hs, "jump")
		}
	}
	if len(bin) == 5 {
		t = make([]interface{}, len(hs))
		for i, v := range hs {
			t[i] = interface{}(v)
		}
		reverse(t)
		for i, v := range t {
			hs[i] = v.(string)
		}
	}
	return hs
}
