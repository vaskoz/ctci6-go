package eleven

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	short, long, k int
	expected       []int
}{
	{short: 1, long: 2, k: 3, expected: []int{3, 4, 5, 6}},
	{short: 1, long: 1, k: 3, expected: []int{3}},
}

func TestDivingBoards(t *testing.T) {
	for _, c := range testcases {
		if got := DivingBoards(c.short, c.long, c.k); !reflect.DeepEqual(got, c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, got)
		}
	}
}

func BenchmarkDivingBoards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			DivingBoards(c.short, c.long, c.k)
		}
	}
}
