package four

import "strings"

func CouldBePalindrome(str string) bool {
	m := map[rune]int{}
	for _, r := range strings.ToLower(str) {
		m[r]++
	}
	delete(m, ' ')
	odds := 0
	for _, v := range m {
		if v%2 != 0 {
			odds++
		}
	}
	return odds <= 1
}
