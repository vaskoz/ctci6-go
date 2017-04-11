package nine

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return Add(a, -b)
}

func Multiply(a, b int) int {
	var small, large int
	if a < b {
		small = a
		large = b
	} else {
		small = b
		large = a
	}
	result := 0
	for i := 1; i <= small; i++ {
		result = Add(result, large)
	}
	return result
}

func Divide(a, b int) int {
	result := 0
	for ; Multiply(result, b) < a; result++ {
	}
	if Multiply(result, b) > a {
		return result - 1
	} else {
		return result
	}
}
