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
	{[]string{"off", "oofo"}, false},
}

func generateUniqueString(length int) string {
	var c rune = 0
	var result []rune
	for i := 0; i < length; i++ {
		result = append(result, c)
		c++
	}
	return string(result)
}

func TestIsPermutation(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if IsPermutation(c.in[0], c.in[1]) != c.out {
			t.Errorf("IsPermutation given %s and %s but got %t\n", c.in[0], c.in[1], c.out)
		}
	}
}

func TestIsPermutationDS(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if IsPermutationDS(c.in[0], c.in[1]) != c.out {
			t.Errorf("IsPermutationDS given %s and %s but got %t\n", c.in[0], c.in[1], c.out)
		}
	}
}

func BenchmarkIsPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			IsPermutation(c.in[0], c.in[1])
		}
	}
}

func BenchmarkIsPermutationDS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			IsPermutationDS(c.in[0], c.in[1])
		}
	}
}

func BenchmarkIsPermutation1kString(b *testing.B) {
	s := generateUniqueString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPermutation(s, s)
	}
}

func BenchmarkIsPermutationDS1kString(b *testing.B) {
	s := generateUniqueString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPermutationDS(s, s)
	}
}
