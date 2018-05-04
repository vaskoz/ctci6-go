package five

// FactorialZeros returns the number of trailing zeros in the answer for n!.
func FactorialZeros(n int) int {
	result := 0
	for divisor := 5; divisor <= n; divisor *= 5 {
		result += (n / divisor)
	}
	return result
}
