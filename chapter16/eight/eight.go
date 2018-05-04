package eight

import (
	"strings"
)

var teens = []string{"", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen",
	"fourteen", "fifteen", "sixteen", "seventeen",
	"eighteen", "nineteen"}

var tens = []string{"", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninety"}

// EnglishInt take an integer and returns a string representing it in English.
// Works up to billions.
func EnglishInt(n int) string {
	if n == 0 {
		return "zero"
	}
	var result []string

	if billions := n / 1000000000; billions > 0 {
		result = append(result, []string{convert(billions), "billion"}...)
	}
	if millions := (n / 1000000) % 1000; millions > 0 {
		result = append(result, []string{convert(millions), "million"}...)
	}
	if thousands := (n / 1000) % 1000; thousands > 0 {
		result = append(result, []string{convert(thousands), "thousand"}...)
	}
	if ones := n % 1000; ones > 0 {
		result = append(result, convert(ones))
	}

	return strings.Join(result, " ")
}

func convert(n int) string {
	var result []string
	if n%100 < 20 {
		result = append(result, teens[n%100])
		n /= 100
	} else {
		result = append(result, teens[n%10])
		n /= 10
		front := []string{tens[n%10]}
		result = append(front, result...)
		n /= 10
	}
	if n == 0 {
		return strings.Join(result, " ")
	}
	front := []string{teens[n]}
	front = append(front, "hundred")
	result = append(front, result...)
	return strings.Join(result, " ")
}
