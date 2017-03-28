package one

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	a   []int
	b   []int
	out []int
}{
	{a: []int{1, 3, 5}, b: []int{2, 4, 6}, out: []int{1, 2, 3, 4, 5, 6}},
	{a: []int{1, 3, 5}, b: []int{2}, out: []int{1, 2, 3, 5}},
}

func TestMerge(t *testing.T) {
	for _, c := range testcases {
		if out := Merge(c.a, c.b); !reflect.DeepEqual(c.out, out) {
			t.Errorf("Expected %v but got %v", c.out, out)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Merge(c.a, c.b)
		}
	}
}
