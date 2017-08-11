package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA sequence samples are not of equal length")
	}
	var counter int
	for i, _ := range a {
		if b[i] != a[i] {
			counter++
		}
	}
	return counter, nil
}
