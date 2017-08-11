package acronym

import (
	"strings"
	"bytes"
)

const testVersion = 3

func Abbreviate(full string) string {
	var buffer bytes.Buffer
	xs := strings.FieldsFunc(full, func (r rune) bool {
		if r == '-' || r == ' ' {
			return true
		}
		return false
	})
	for _, s := range xs {
		buffer.WriteByte(s[0])
	}
	return strings.ToUpper(buffer.String())
}
