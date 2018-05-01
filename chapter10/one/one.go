package one

// Merge takes two sorted int slices and merges them into a new int slice.
func Merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for {
		if i < len(a) && j < len(b) {
			if a[i] < b[j] {
				result = append(result, a[i])
				i++
			} else {
				result = append(result, b[j])
				j++
			}
		} else if i < len(a) {
			result = append(result, a[i])
			i++
		} else if j < len(b) {
			result = append(result, b[j])
			j++
		} else {
			return result
		}
	}
}
