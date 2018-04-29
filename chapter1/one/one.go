package one

import "sort"

func IsUnique(str string) bool {
	m := map[rune]int{}
	for _, r := range str {
		if count := m[r]; count == 1 {
			return false
		}
		m[r]++
	}
	return true
}

func IsUniqueNoDS(str string) bool {
	runes := []rune(str)
	for i := range runes {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func IsUniqueSort(str string) bool {
	runes := []rune(str)
	less := func(i, j int) bool { return runes[i] < runes[j] }
	if !sort.SliceIsSorted(runes, less) {
		sort.Slice(runes, less)
	}
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == runes[i] {
			return false
		}
	}
	return true
}
