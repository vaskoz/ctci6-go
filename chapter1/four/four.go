package four

import "strings"

func CouldBePalindrome(str string) bool {
	runes := []rune(strings.ToLower(str))
	m := map[rune]int{}
	for _, r := range runes {
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
