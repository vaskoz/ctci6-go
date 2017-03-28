package two

import (
	"sort"
)

func GroupAnagrams(words []string) {
	sort.SliceStable(words, func(i, j int) bool {
		a := []rune(words[i])
		b := []rune(words[j])
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		sa, sb := string(a), string(b)
		return sa < sb
	})
}
