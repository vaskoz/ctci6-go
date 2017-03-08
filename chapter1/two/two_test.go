package two

import "testing"

var testcases = []struct {
	in  []string
	out bool
}{
	{[]string{"a", "a"}, true},
	{[]string{"a", "b"}, false},
	{[]string{"of", "of"}, true},
	{[]string{"off", "fof"}, true},
	{[]string{"off", "oof"}, false},
}

func TestIsPermutation(t *testing.T) {
	for _, c := range testcases {
		if IsPermutation(c.in[0], c.in[1]) != c.out {
			t.Errorf("IsPermutation given %s and %s but got %t\n", c.in[0], c.in[1], c.out)
		}
	}
}

func TestIsPermutationDS(t *testing.T) {
	for _, c := range testcases {
		if IsPermutationDS(c.in[0], c.in[1]) != c.out {
			t.Errorf("IsPermutationDS given %s and %s but got %t\n", c.in[0], c.in[1], c.out)
		}
	}
}
