package one

import "testing"

var testcases = []struct {
	data       DiGraph
	start, end int
	out        bool
}{
	{
		data: DiGraph{
			1: map[int]struct{}{
				3: struct{}{},
				4: struct{}{},
			},
			2: map[int]struct{}{
				3: struct{}{},
			},
			3: map[int]struct{}{
				4: struct{}{},
			},
		},
		start: 1,
		end:   4,
		out:   true,
	},
	{
		data: DiGraph{
			1: map[int]struct{}{
				2: struct{}{},
			},
			2: map[int]struct{}{
				4: struct{}{},
			},
			3: map[int]struct{}{
				1: struct{}{},
			},
		},
		start: 1,
		end:   3,
		out:   false,
	},
}

func TestRouteExists(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if out := RouteExists(c.data, c.start, c.end); out != c.out {
			t.Errorf("start %v end %v expected %t got %t", c.start, c.end, c.out, out)
		}
	}
}

func BenchmarkRouteExists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			RouteExists(c.data, c.start, c.end)
		}
	}
}
