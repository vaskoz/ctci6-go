package five

import "testing"

var testcases = []struct {
	in1 string
	in2 string
	out bool
}{
	{"pale", "pale", true},
	{"too", "tookem", false},
	{"too", "too", true},
	{"too", "took", true},
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"bale", "bales", true},
	{"pale", "bake", false},
}

func TestOneAway(t *testing.T) {
	for _, c := range testcases {
		if out := OneAway(c.in1, c.in2); out != c.out {
			t.Errorf("OneAway given %s and %s returns %t but expected %t\n", c.in1, c.in2, out, c.out)
		}
	}
}

func BenchmarkOneAway(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			OneAway(c.in1, c.in2)
		}
	}
}
