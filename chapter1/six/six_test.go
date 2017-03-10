package six

import "testing"

var testcases = []struct {
	in, out string
}{
	{"aabcccccaa", "a2b1c5a2"},
	{"a", "a"},
	{"abcd", "abcd"},
	{"aaaad", "a4d1"},
}

func TestStringCompress(t *testing.T) {
	for _, c := range testcases {
		if out := StringCompress(c.in); out != c.out {
			t.Errorf("StringCompress given %s returns %s expects %s\n", c.in, out, c.out)
		}
	}
}

func BenchmarkStringCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			StringCompress(c.in)
		}
	}
}

func TestStringCompressList(t *testing.T) {
	for _, c := range testcases {
		if out := StringCompressList(c.in); out != c.out {
			t.Errorf("StringCompressList given %s returns %s expects %s\n", c.in, out, c.out)
		}
	}
}

func BenchmarkStringCompressList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			StringCompressList(c.in)
		}
	}
}
