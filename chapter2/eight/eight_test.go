package eight

import (
	"testing"
)

var testcases = []struct {
	in  []int
	out int
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 3}, 3},
	{[]int{1, 2, 3, 4}, 0},
	{[]int{1, 2, 1}, 1},
}

func createList(data []int) *Elem {
	head := &Elem{}
	current := head
	used := map[int]*Elem{}
	for _, v := range data {
		if ptr, found := used[v]; found {
			current.next = ptr
		} else {
			current.next = &Elem{Value: v}
			current = current.next
			used[v] = current
		}
	}
	return head.next
}

func TestDetectLoop(t *testing.T) {
	for _, c := range testcases {
		if out := DetectLoop(createList(c.in)); out == nil || out.Value != c.out {
			if out == nil && c.out == 0 {
				continue
			}
			t.Errorf("DetectLoop: given %v got %v expected %v\n", c.in, out, c.out)
		}
	}
}

func BenchmarkDetectLoop(b *testing.B) {
	l := createList([]int{1, 2, 3, 4, 5, 2})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DetectLoop(l)
	}
}
