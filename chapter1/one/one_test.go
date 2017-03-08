package one

import "testing"

var testcases = []struct {
	in  string
	out bool
}{
	{"", true},
	{"a", true},
	{"aa", false},
	{"aba", false},
	{"baa", false},
}

func TestIsUnique(t *testing.T) {
	for _, c := range testcases {
		if IsUnique(c.in) != c.out {
			t.Errorf("IsUnique given %s should output %t\n", c.in, c.out)
		}
	}
}

func TestIsUniqueNoDS(t *testing.T) {
	for _, c := range testcases {
		if IsUniqueNoDS(c.in) != c.out {
			t.Errorf("IsUniqueNoDS given %s should output %t\n", c.in, c.out)
		}
	}
}
