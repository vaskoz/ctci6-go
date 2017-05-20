package five

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	in1, in2 []int
	out      int
}{
	{[]int{1, 2, 3}, []int{3, 2, 1}, 444},
	{[]int{1}, []int{3}, 4},
	{[]int{1}, []int{1, 3}, 32},
	{[]int{9, 9}, []int{1, 1}, 110},
	{[]int{9, 9}, []int{1}, 100},
}

func TestSumReverseList(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		first, second := createList(c.in1), createList(c.in2)
		if out := SumReverseList(first, second); out != c.out {
			t.Errorf("SumReverseList: given %s and %s returned %d expected %d\n", sprint(first), sprint(second), out, c.out)
		}
	}
}

func BenchmarkSumReverseList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		first, second := createList(testcases[0].in1), createList(testcases[0].in2)
		SumReverseList(first, second)
	}
}

func createList(data []int) *Elem {
	result := &Elem{}
	current := result
	for _, v := range data {
		current.next = &Elem{Value: v}
		current = current.next
	}
	return result.next
}

func sprint(head *Elem) string {
	var result string = "["
	for head != nil {
		result += fmt.Sprintf("%d->", head.Value)
		head = head.next
	}
	result += "]"
	return result
}
