package seven

import "testing"

var testcases = []struct {
	a, b, max int
}{
	{10, 1, 10},
	{1, 10, 10},
	{-1, 10, 10},
	{10, -1, 10},
	{-100, 10, 10},
	{10, -100, 10},
}

func TestMaxNoCompare(t *testing.T) {
	for _, c := range testcases {
		if max := MaxNoCompare(c.a, c.b); max != c.max {
			t.Errorf("Expected %v, but got %v", c.max, max)
		}
	}
}

func BenchmarkMaxNoCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			MaxNoCompare(c.a, c.b)
		}
	}
}

func TestMaxCompare(t *testing.T) {
	for _, c := range testcases {
		if max := MaxCompare(c.a, c.b); max != c.max {
			t.Errorf("Expected %v, but got %v", c.max, max)
		}
	}
}

func BenchmarkMaxCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			MaxCompare(c.a, c.b)
		}
	}
}
