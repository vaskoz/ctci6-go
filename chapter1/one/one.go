package one

func IsUnique(str string) bool {
	m := map[rune]int{}
	runes := []rune(str)
	for _, r := range runes {
		m[r]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
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
