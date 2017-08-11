package triangle

import "math"

const (
	testVersion = 3
	NaT         = iota // not a triangle
	Equ                // equilateral
	Iso                // isosceles
	Sca                // scalene
)

type Kind int

func checkNaT(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || (a < 0 || b < 0 || c < 0) || (b == 0 && c == 0 && a == 0) {
		return true
	}
	lengths := map[string]float64{
		"a": a,
		"b": b,
		"c": c,
	}
	max := a
	holder := map[string]string{"key": "a"}
	for key, val := range lengths {
		if val > max {
			max = val
			holder["key"] = key
		}
	}
	delete(lengths, holder["key"])
	var sum float64
	for _, val := range lengths {
		sum += val
	}
	if sum < max {
		return true
	}
	return false
}

func KindFromSides(a, b, c float64) Kind {
	if checkNaT(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}