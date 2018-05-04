package one

// NumberSwap swaps the values pointed to by x and y using XOR.
func NumberSwap(x, y *int) {
	*x ^= *y
	*y ^= *x
}

// NumberSwapGo swaps the values pointed to by x and y using parallel assignment.
// https://en.wikipedia.org/wiki/Assignment_(computer_science)#Parallel_assignment
func NumberSwapGo(x, y *int) {
	*x, *y = *y, *x
}

// NumberSwapTemp swaps the values pointed to by x and y using a temporary variable.
func NumberSwapTemp(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
