package two

import (
	"reflect"
	"sort"
)

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	return reflect.DeepEqual(r1, r2)
}

func IsPermutationDS(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	if len(r1) != len(r2) {
		return false
	}
	m := map[rune]int{}
	for _, r := range r1 {
		m[r]++
	}
	for _, r := range r2 {
		if count := m[r]; count > 1 {
			m[r]--
		} else if count == 1 {
			delete(m, r)
		} else {
			return false
		}
	}
	return len(m) == 0
}
