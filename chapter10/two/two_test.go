package two

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	in  []string
	out []string
}{
	{in: []string{"aba", "c", "baa"}, out: []string{"aba", "baa", "c"}},
	{in: []string{"cdc", "c", "ccd"}, out: []string{"c", "cdc", "ccd"}},
	{in: []string{"cdc", "c", "ccd"}, out: []string{"c", "cdc", "ccd"}},
	{in: []string{"bbaa", "aba", "abba"}, out: []string{"aba", "bbaa", "abba"}},
}

func TestGroupAnagrams(t *testing.T) {
	for _, c := range testcases {
		if GroupAnagrams(c.in); !reflect.DeepEqual(c.in, c.out) {
			t.Errorf("Expected %v but got %v", c.out, c.in)
		}
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			GroupAnagrams(c.in)
		}
	}
}
