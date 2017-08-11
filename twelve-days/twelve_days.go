package twelve

import "bytes"

const (
	testVersion = 1
	basePref = "On the "
	baseSuf = " day of Christmas my true love gave to me,"
)

var keyMap = map[int]string{
	0: "first",
	1: "second",
	2: "third",
	3: "fourth",
	4: "fifth",
	5: "sixth",
	6: "seventh",
	7: "eighth",
	8: "ninth",
	9: "tenth",
	10: "eleventh",
	11: "twelfth",
}

var daysSlice = []string{
	" a Partridge in a Pear Tree.",
	" two Turtle Doves,",
	" three French Hens,",
	" four Calling Birds,",
	" five Gold Rings,",
	" six Geese-a-Laying,",
	" seven Swans-a-Swimming,",
	" eight Maids-a-Milking,",
	" nine Ladies Dancing,",
	" ten Lords-a-Leaping,",
	" eleven Pipers Piping,",
	" twelve Drummers Drumming,",
}
func writeVerse(i int, addLine bool) string {
	var buffer bytes.Buffer
	buffer.WriteString(basePref)
	buffer.WriteString(keyMap[i])
	buffer.WriteString(baseSuf)
	for j := i ; j >= 0; j-- {
		if j == 0 && i != 0 {
			buffer.WriteString(" and")
		}
		buffer.WriteString(daysSlice[j])
	}
	if addLine {
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func Song() string  {
	var buffer bytes.Buffer
	for i := range daysSlice {
		buffer.WriteString(writeVerse(i, true))
	}
	return buffer.String()
}

func Verse(i int) string {
	return writeVerse((i - 1), false)
}