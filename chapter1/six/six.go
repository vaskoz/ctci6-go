package six

import "fmt"

func StringCompress(orig string) string {
	if len(orig) < 3 {
		return orig
	}
	result := make([]rune, 0, len(orig))
	count := 0
	var last rune
	for i, r := range orig {
		if i == 0 {
			last = r
			count = 1
		} else if r != last {
			result = append(result, []rune(fmt.Sprintf("%s%d", string(last), count))...)
			count = 1
			last = r
		} else {
			count++
		}
	}
	result = append(result, []rune(fmt.Sprintf("%s%d", string(last), count))...)
	if len(result) >= len(orig) {
		return orig
	}
	return string(result)
}

func StringCompressList(orig string) string {
	if len(orig) < 3 {
		return orig
	}
	indicies := []int{0}
	runes := []rune(orig)
	for i, r := range runes {
		if r != runes[indicies[len(indicies)-1]] {
			indicies = append(indicies, i)
		}
	}
	indicies = append(indicies, len(runes))
	result := make([]rune, 0, len(orig))
	for i := 0; i < len(indicies)-1; i++ {
		result = append(result, []rune(fmt.Sprintf("%s%d", string(runes[indicies[i]]), indicies[i+1]-indicies[i]))...)
	}
	if len(result) >= len(runes) {
		return orig
	}
	return string(result)
}
