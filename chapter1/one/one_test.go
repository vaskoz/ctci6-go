package one

import "testing"

func TestIsUnique(t *testing.T) {
	testcases := []struct {
		in  string
		out bool
	}{
		{"", true},
		{"a", true},
		{"aa", false},
		{"aba", false},
		{"baa", false},
	}
	for _, c := range testcases {
		if IsUnique(c.in) != c.out {
			t.Errorf("IsUnique given %s should output %t\n", c.in, c.out)
		}
	}
}
