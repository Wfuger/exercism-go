package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
