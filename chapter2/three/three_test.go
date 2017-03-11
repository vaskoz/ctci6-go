package three

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	in  []int
	pos int
	out []int
}{
	{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 5}},
}

func createList(data []int) *Elem {
	var head, current *Elem = nil, nil
	for i, v := range data {
		if i == 0 {
			current = &Elem{Value: v}
			head = current
			continue
		}
		current.next = &Elem{Value: v}
		current = current.next
	}
	return head
}

func getData(list *Elem) []int {
	result := []int{}
	for list.next != nil {
		result = append(result, list.Value.(int))
		list = list.next
	}
	return result
}

func TestDeleteMiddle(t *testing.T) {
	for _, c := range testcases {
		head := createList(c.in)
		current := head
		for i := 0; i < c.pos; i++ {
			current = current.next
		}
		DeleteMiddle(current)
		if out := getData(head); !reflect.DeepEqual(out, c.out) {
			t.Errorf("DeleteMiddle: given %v got %v expected %v\n", c.in, out, c.out)
		}
	}
}

func BenchmarkDeleteMiddle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			head := createList(c.in)
			current := head
			for i := 0; i < c.pos; i++ {
				current = current.next
			}
			DeleteMiddle(current)
		}
	}
}
