package four

import "testing"

var testcases = []struct {
	in  string
	out bool
}{
	{"Tact coa", true},
	{"a a", true},
	{"abbb", false},
}

func TestCouldBePalindrome(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if out := CouldBePalindrome(c.in); out != c.out {
			t.Errorf("CouldBePalindrome given %s and got %t but expected %t\n", c.in, out, c.out)
		}
	}
}

func BenchmarkCouldBePalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			CouldBePalindrome(c.in)
		}
	}
}
