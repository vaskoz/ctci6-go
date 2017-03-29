package three

import "testing"

var testcases = []struct {
	in     []int
	search int
	out    int
}{
	{in: []int{7, 1, 2}, search: 7, out: 0},
	{in: []int{4, 5, 6, 7, 1, 2}, search: 1, out: 4},
	{in: []int{7, 1, 2, 4, 5}, search: 5, out: 4},
	{in: []int{7, 1, 2, 3, 4, 5, 6}, search: 1, out: 1},
	{in: []int{7, 1, 2, 3, 4, 5, 6}, search: 10, out: -1},
}

func TestFindRotatedIndex(t *testing.T) {
	for _, c := range testcases {
		if out := FindRotatedIndex(c.in, c.search); out != c.out {
			t.Errorf("Expected %v but got %v", c.out, out)
		}
	}
}

func BenchmarkFindRotatedIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			FindRotatedIndex(c.in, c.search)
		}
	}
}
