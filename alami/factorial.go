package alami

// Factorial function
// 5 = 120
func Factorial(n int) int {
	var factorial int

	for i := 0; i < n; i++ {
		if i == 0 {
			factorial = n
		} else {
			factorial *= n - i
		}
	}

	return factorial
}
