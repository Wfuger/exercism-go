package house

import (
	"bytes"
	"strings"
)

const final = `This is the horse and the hound and the horn
that belonged to the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`

const testVersion = 1

func iterate(buf bytes.Buffer, verse string) string {
	if verse == "This is the house that Jack built." {
		return buf.String()
	}
	lines := verse.split("/n")
	lines = lines[1:]
	firstLine := lines[0]
	noun := firstLine.split("the")[1]
	nfl := "This is the" + noun
	lines[0] = nfl
	lines = append(lines, "\n\n")
	lineString := strings.Join(lines)
	buf.WriteString(lines)
	return buf.String() + iterate(buf, lines)
}

// Verse takes a verse number and returns the verse
func Verse(num int) string {

}

// Song returns the full song
func Song() string {

}
