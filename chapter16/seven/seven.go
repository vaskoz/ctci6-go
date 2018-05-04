package seven

import (
	"strconv"
)

// MaxNoCompare returns the maximum of two integers using bit manipulation rather than comparison.
func MaxNoCompare(a, b int) int {
	var intBits uint = strconv.IntSize
	neg := ((a - b) >> (intBits - 1)) & 1
	return a*(neg^1) + neg*b
}

// MaxCompare returns the maximum of two integers using a comparison operation.
func MaxCompare(a, b int) int {
	if a < b {
		return b
	}
	return a
}
