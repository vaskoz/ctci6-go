package seven

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	in1, in2 []int
	out      int
}{
	{[]int{1, 4, 5}, []int{2, 4, 5}, 4},
	{[]int{1, 2, 4, 5}, []int{2, 4, 5}, 2},
	{[]int{1, 2, 4, 5}, []int{6, 7, 8}, 0},
	{[]int{1, 2, 4, 5}, []int{5}, 5},
	{[]int{1, 2, 4, 5}, []int{4, 5}, 4},
	{[]int{1, 2, 4, 5}, []int{0, 1, 2, 4, 5}, 1},
}

func createIntersectingLists(l1, l2 []int) (*Elem, *Elem) {
	var common *Elem
	index := 0
	for index <= len(l1)-1 && index <= len(l2)-1 && l1[len(l1)-index-1] == l2[len(l2)-index-1] {
		index++
	}
	for i := len(l1) - 1; i > len(l1)-1-index; i-- {
		common = &Elem{Value: l1[i], next: common}
	}
	list1 := &Elem{}
	result1 := list1
	for i := 0; i < len(l1)-index; i++ {
		list1.next = &Elem{Value: l1[i]}
		list1 = list1.next
	}
	list1.next = common

	list2 := &Elem{}
	result2 := list2
	for i := 0; i < len(l2)-index; i++ {
		list2.next = &Elem{Value: l2[i]}
		list2 = list2.next
	}
	list2.next = common

	return result1.next, result2.next
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

func TestIntersection(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		first, second := createIntersectingLists(c.in1, c.in2)
		if out := Intersection(first, second); out == nil && c.out != 0 {
			t.Errorf("Intersection: given %s and %s returned 0 expected %d\n", sprint(first), sprint(second), c.out)
		} else if out != nil && out.Value != c.out {
			t.Errorf("Intersection: given %s and %s returned %d expected %d\n", sprint(first), sprint(second), out.Value, c.out)
		}
	}
}

func BenchmarkIntersection(b *testing.B) {
	first, second := createIntersectingLists(testcases[0].in1, testcases[1].in1)
	for i := 0; i < b.N; i++ {
		Intersection(first, second)
	}
}
