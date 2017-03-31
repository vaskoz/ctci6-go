package seven

import (
	"strconv"
)

func MaxNoCompare(a, b int) int {
	var intBits uint = strconv.IntSize
	neg := ((a - b) >> (intBits - 1)) & 1
	return a*(neg^1) + neg*b
}

func MaxCompare(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
