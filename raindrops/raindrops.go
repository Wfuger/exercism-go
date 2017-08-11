package raindrops

import (
	"fmt"
	"bytes"
)

const testVersion = 3

func Convert(num int) string {
	var buffer bytes.Buffer

	pling := []rune{'P','l','i','n','g'}
	plang := []rune{'P','l','a','n','g'}
	plong := []rune{'P','l','o','n','g'}

	var converted string
	if num % 3 == 0 {
		for b, _ := range pling {
			buffer.WriteRune(pling[b])
		}
	}
	if num % 5 == 0 {
		for b, _ := range plang {
			buffer.WriteRune(plang[b])
		}
	}
	if num % 7 == 0 {
		for b, _ := range plong {
			buffer.WriteRune(plong[b])
		}
	}
	if len(buffer.String()) == 0 {
		converted = fmt.Sprintf("%v", num)
	} else {
		converted = fmt.Sprintf("%s", buffer.String())
	}
	return converted
}

