package five

import "testing"

var testcases = []struct {
	in  int
	out int
}{
	{6, 1},
	{100, 24},
	{130, 32},
}

func TestFactorialZeros(t *testing.T) {
	for _, c := range testcases {
		if out := FactorialZeros(c.in); out != c.out {
			t.Errorf("Expected %v, but got %v", c.out, out)
		}
	}
}

func BenchmarkFactorialZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			FactorialZeros(c.in)
		}
	}
}
