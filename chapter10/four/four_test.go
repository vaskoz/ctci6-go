package four

import "testing"

var testcases = []struct {
	in     *Listy
	search int
	out    int
}{
	{in: &Listy{1, 3, 4, 5, 6, 10, 12, 13}, search: 10, out: 6},
	{in: &Listy{1, 3, 4, 5, 6, 10, 12, 13}, search: 3, out: 2},
	{in: &Listy{1, 3, 4, 5, 6, 10, 12, 13}, search: 7, out: -1},
}

func TestSearchListy(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if out := SearchListy(c.in, c.search); out != c.out {
			t.Errorf("Expected %v got %v", c.out, out)
		}
	}
}

func BenchmarkSearchListy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			SearchListy(c.in, c.search)
		}
	}
}
