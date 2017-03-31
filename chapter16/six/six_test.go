package six

import "testing"

var testcases = []struct {
	a, b []int
	diff int
}{
	{[]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}, 3},
}

func TestSmallestDifference(t *testing.T) {
	for _, c := range testcases {
		if diff := SmallestDifference(c.a, c.b); diff != c.diff {
			t.Errorf("Expected %v, but got %v", c.diff, diff)
		}
	}
}

func BenchmarkSmallestDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			SmallestDifference(c.a, c.b)
		}
	}
}
