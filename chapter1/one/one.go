package one

func IsUnique(str string) bool {
	m := map[rune]int{}
	runes := []rune(str)
	for _, r := range runes {
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
