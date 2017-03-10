package nine

import "testing"

var testcases = []struct {
	s1, s2 string
	out    bool
}{
	{"waterbottle", "erbottlewat", true},
	{"aa", "aa", true},
	{"a", "a", true},
	{"abc", "bac", false},
}

func TestStringRotation(t *testing.T) {
	for _, c := range testcases {
		if out := StringRotation(c.s1, c.s2); out != c.out {
			t.Errorf("StringRotation: %s, %s return %t expected %t\n", c.s1, c.s2, out, c.out)
		}
	}
}

func BenchmarkStringRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			StringRotation(c.s1, c.s2)
		}
	}
}
