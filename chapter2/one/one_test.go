package one

import (
	"container/list"
	"reflect"
	"testing"
)

var testcases = []struct {
	in, out []int
}{
	{[]int{1, 1}, []int{1}},
	{[]int{1, 1, 2}, []int{1, 2}},
	{[]int{1, 1, 2, 2}, []int{1, 2}},
	{[]int{1, 1, 3, 2, 2, 3}, []int{1, 2, 3}},
}

func TestRemoveDups(t *testing.T) {
	for _, c := range testcases {
		l := createList(c.in)
		RemoveDups(l)
		data := getDataFromList(l)
		if !reflect.DeepEqual(c.out, data) {
			t.Errorf("RemoveDups: given %v returns %v expected %v\n", c.in, data, c.out)
		}
	}
}

func BenchmarkRemoveDups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			l := createList(c.in)
			RemoveDups(l)
			getDataFromList(l)
		}
	}
}

func getDataFromList(l *list.List) []int {
	result := []int{}
	for e := l.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}
	return result
}

func createList(data []int) *list.List {
	l := list.New()
	for _, v := range data {
		l.PushBack(v)
	}
	return l
}
