package five

import "testing"

var testcases = []struct {
	in     []string
	search string
	out    int
}{
	{in: []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, search: "ball", out: 4},
	{in: []string{"at", "", "ball", "", "car"}, search: "ball", out: 2},
	{in: []string{"at", "", "ball", "", "car"}, search: "bar", out: -1},
	{in: []string{"at", "bake", "ball", "cab", "car"}, search: "at", out: 0},
	{in: []string{"at", "bake", "ball", "cab", "car"}, search: "ate", out: -1},
	{in: []string{"at", "", "", "ball", "car"}, search: "car", out: 4},
}

func TestSparseStringSearch(t *testing.T) {
	for _, c := range testcases {
		if out := SparseStringSearch(c.in, c.search); out != c.out {
			t.Errorf("Expected %v got %v", c.out, out)
		}
	}
}

func BenchmarkSparseStringSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			SparseStringSearch(c.in, c.search)
		}
	}
}
