package two

import "testing"

var testcases = []struct {
	in  []int
	k   int
	out int
}{
	{[]int{1, 2, 3, 4, 5}, 2, 4},
	{[]int{1, 2, 3, 4, 5}, 4, 2},
}

func createList(data []int) *Elem {
	if len(data) == 0 {
		return nil
	}
	head := &Elem{}
	current := head
	for _, v := range data {
		current.Value = v
		current.next = &Elem{}
		current = current.next
	}
	return head
}

func TestFindKth(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		head := createList(c.in)
		if out := FindKth(head, c.k); out.Value != c.out {
			t.Errorf("FindKth %d returned %d expected %d\n", c.k, out, c.out)
		}
	}
}

func BenchmarkFindKth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			head := createList(c.in)
			FindKth(head, c.k)
		}
	}
}

func TestFindKth2Ptr(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		head := createList(c.in)
		if out := FindKth2Ptr(head, c.k); out.Value != c.out {
			t.Errorf("FindKth %d returned %d expected %d\n", c.k, out, c.out)
		}
	}
}

func BenchmarkFindKth2Ptr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			head := createList(c.in)
			FindKth2Ptr(head, c.k)
		}
	}
}

func TestFindKthRecursive(t *testing.T) {
}
