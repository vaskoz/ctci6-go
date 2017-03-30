package one

func NumberSwap(x, y *int) {
	*x ^= *y
	*y ^= *x
}

func NumberSwapGo(x, y *int) {
	*x, *y = *y, *x
}

func NumberSwapTemp(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
