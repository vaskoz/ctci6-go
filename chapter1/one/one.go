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
