package two

import "sort"

func IsPermutation(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	if len(r1) != len(r2) {
		return false
	}
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
