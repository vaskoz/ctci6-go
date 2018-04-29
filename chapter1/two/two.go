package two

import (
	"reflect"
	"sort"
)

// IsPermutation returns true if strings are permutations of each other.
// It sorts each rune in the string and then compares them for equality.
// Since each string is sorted runtime is O(N log N) where N=len(s1)+len(s2)
func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	return reflect.DeepEqual(r1, r2)
}

// IsPermutationDS returns true if strings are permutations of each other.
// It uses a map of space O(N) to count occurrences of runes in one string.
// It then removes each rune occurrence that appears in the second string.
// Runtime is O(N) thanks to the map.
func IsPermutationDS(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := map[rune]int{}
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
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
