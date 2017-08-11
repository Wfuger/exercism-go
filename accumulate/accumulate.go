package accumulate

const testVersion = 1

func Accumulate(xs []string, mapper func(string) string) []string {
	var accum []string
	for i, _ := range xs {
		accum = append(accum, mapper(xs[i]))
	}
	return accum
}
