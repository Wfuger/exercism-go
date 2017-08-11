package bob

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const (
	testVersion = 3
	q           = iota
	y
	w
	other
)

type teenager struct {
	question   string
	yell       string
	whitespace string
	def        string
}

func getAlphas(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		fmt.Printf("ERROR With regex %v", err)
	}
	return reg.ReplaceAllString(s, "")
}

func (t teenager) typeOfInput(s string) int {
	alphas := getAlphas(s)
	allCaps := true
	for _, char := range alphas {
		if unicode.IsLower(char) {
			allCaps = false
		}
	}
	if allCaps && len(alphas) > 0 {
		return y
	}
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return w
	}
	if rune(trimmed[len(trimmed)-1]) == '?' {
		return q
	}
	return other
}

func Hey(input string) string {
	bob := teenager{
		"Sure.",
		"Whoa, chill out!",
		"Fine. Be that way!",
		"Whatever.",
	}
	typeOf := bob.typeOfInput(input)
	switch typeOf {
	case q:
		return bob.question
	case y:
		return bob.yell
	case w:
		return bob.whitespace
	default:
		return bob.def
	}
}
