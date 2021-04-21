package fibo

// Value of the fibonacci number at the index
func Calc(n int) int {
	if n <= 2 {
		return 1
	}
	x, y := 1, 1
	result := 0
	for i := 2; i < n; i++ {
		result = x + y
		x = y
		y = result
	}
	return result
}
