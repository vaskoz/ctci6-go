package nine

// Add returns the arithmetic sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers.
// Implemented using the Add function.
func Subtract(a, b int) int {
	return Add(a, -b)
}

// Multiply returns the multiplication of two integers.
// Implemented using the Add function.
func Multiply(a, b int) int {
	var neg = false
	if a < 0 {
		a = -a
		neg = true
	}
	if b < 0 {
		b = -b
		if neg {
			neg = false
		}
	}
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
	if neg {
		return -result
	}
	return result
}

// Divide returns the division of two integers.
// Implemented by using the Multiply function.
func Divide(a, b int) int {
	var neg = false
	if a < 0 {
		a = -a
		neg = true
	}
	if b < 0 {
		b = -b
		if neg {
			neg = false
		}
	}
	result := 0
	for ; Multiply(result, b) < a; result++ {
	}
	if Multiply(result, b) > a {
		if neg {
			return -result + 1
		}
		return result - 1
	}
	if neg {
		return -result
	}
	return result
}
