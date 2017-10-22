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

func generateUniqueString(length int) string {
	var c rune
	var result []rune
	for i := 0; i < length; i++ {
		result = append(result, c)
		c++
	}
	return string(result)
}

func TestIsUnique(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if IsUnique(c.in) != c.out {
			t.Errorf("IsUnique given %s should output %t\n", c.in, c.out)
		}
	}
}

func TestIsUniqueSort(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if IsUniqueSort(c.in) != c.out {
			t.Errorf("IsUniqueSort given %s should output %t\n", c.in, c.out)
		}
	}
}

func TestIsUniqueNoDS(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if IsUniqueNoDS(c.in) != c.out {
			t.Errorf("IsUniqueNoDS given %s should output %t\n", c.in, c.out)
		}
	}
}

func BenchmarkIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			IsUnique(c.in)
		}
	}
}

func BenchmarkIsUniqueSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			IsUniqueSort(c.in)
		}
	}
}

func BenchmarkIsUniqueNoDS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			IsUniqueNoDS(c.in)
		}
	}
}

func BenchmarkIsUnique1kString(b *testing.B) {
	s := generateUniqueString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsUnique(s)
	}
}

func BenchmarkIsUniqueSort1kString(b *testing.B) {
	s := generateUniqueString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsUniqueSort(s)
	}
}

func BenchmarkIsUniqueNoDS1kString(b *testing.B) {
	s := generateUniqueString(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsUniqueNoDS(s)
	}
}
