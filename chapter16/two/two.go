package two

import (
	"strings"
	"unicode"
)

func WordFrequency(book string) map[string]int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(book, f)
	results := make(map[string]int)
	for _, word := range words {
		results[strings.ToLower(word)]++
	}
	return results
}
