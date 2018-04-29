package three

import "strings"

// InplaceURLify takes a string with extra trailing space to fit the encoded string.
func InplaceURLify(str string) string {
	input := []rune(str)
	endIndex := len(input) - 1
	for i := range input {
		if input[len(input)-1-i] != ' ' {
			input[endIndex] = input[len(input)-1-i]
			endIndex--
		} else if endIndex != len(input)-1 {
			input[endIndex] = '0'
			input[endIndex-1] = '2'
			input[endIndex-2] = '%'
			endIndex -= 3
		}
	}
	return string(input)
}

// SimpleURLify replaces all spaces with %20
func SimpleURLify(str string) string {
	str = strings.TrimSpace(str)
	return strings.Replace(str, " ", "%20", -1)
}
