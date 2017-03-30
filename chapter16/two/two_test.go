package two

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	book string
	freq map[string]int
}{
	{
		book: "So humble, so humble. Watch out, watch out",
		freq: map[string]int{
			"so":     2,
			"humble": 2,
			"watch":  2,
			"out":    2,
		},
	},
}

func TestWordFrequency(t *testing.T) {
	for _, c := range testcases {
		if out := WordFrequency(c.book); !reflect.DeepEqual(c.freq, out) {
			t.Errorf("got %v expected %v", out, c.freq)
		}
	}
}

func BenchmarkWordFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			WordFrequency(c.book)
		}
	}
}
