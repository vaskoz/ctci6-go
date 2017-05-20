package three

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	in  string
	out string
}{
	{"a", "a"},
	{"a a  ", "a%20a"},
	{"Mr John Smith    ", "Mr%20John%20Smith"},
}

func TestInplaceURLify(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if out := InplaceURLify(c.in); !reflect.DeepEqual(out, c.out) {
			t.Errorf("InplaceURLify: given %v expected %v but got %v\n", c.in, c.out, out)
		}
	}
}

func TestSimpleURLify(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if out := SimpleURLify(c.in); !reflect.DeepEqual(out, c.out) {
			t.Errorf("SimpleURLify: given %v expected %v but got %v\n", c.in, c.out, out)
		}
	}
}

func BenchmarkInplaceURLify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			InplaceURLify(c.in)
		}
	}
}

func BenchmarkSimpleURLify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			SimpleURLify(c.in)
		}
	}
}
