package six

import (
	"container/list"
	"testing"
)

var testcases = []struct {
	in  []int
	out bool
}{
	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 1, 2}, false},
	{[]int{1}, true},
}

func createList(data []int) *list.List {
	l := list.New()
	for _, v := range data {
		l.PushBack(v)
	}
	return l
}

func TestIsPalindrome(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		l := createList(c.in)
		if out := IsPalindrome(l); out != c.out {
			t.Errorf("IsPalindrome: given %v got %t expected %t\n", c.in, out, c.out)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	l := createList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(l)
	}
}
