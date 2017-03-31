package six

import (
	"math"
	"sort"
)

func SmallestDifference(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	smallest := math.MaxInt64
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			if b[j]-a[i] < smallest {
				smallest = b[j] - a[i]
			}
			i++
		} else {
			if a[i]-b[j] < smallest {
				smallest = a[i] - b[j]
			}
			j++
		}
	}
	return smallest
}
