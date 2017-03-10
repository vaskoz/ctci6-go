package eight

import "testing"

var testcases = []struct {
	in, out [][]int
}{
	{
		[][]int{
			{1, 2, 3, 4},
			{2, 0, 4, 5},
			{3, 4, 5, 6},
			{4, 5, 6, 0},
		},
		[][]int{
			{1, 0, 3, 0},
			{0, 0, 0, 0},
			{3, 0, 5, 0},
			{0, 0, 0, 0},
		},
	},
	{
		[][]int{
			{1, 2, 3, 4, 5},
			{2, 3, 4, 5, 0},
			{3, 0, 5, 6, 7},
			{4, 5, 6, 7, 8},
			{5, 6, 7, 8, 9},
		},
		[][]int{
			{1, 0, 3, 4, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{4, 0, 6, 7, 0},
			{5, 0, 7, 8, 0},
		},
	},
}

func TestZeroMatrix(t *testing.T) {
	for _, c := range testcases {
		if out := ZeroMatrix(c.in); !MatrixEqual(out, c.out) {
			t.Errorf("ZeroMatrix: given %v returned %v expected %v\n", c.in, out, c.out)
		}
	}
}

func BenchmarkZeroMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			ZeroMatrix(c.in)
		}
	}
}

func MatrixEqual(one, two [][]int) bool {
	if len(one) != len(two) {
		return false
	}
	for i := 0; i < len(one); i++ {
		for j := 0; j < len(one[i]); j++ {
			if one[i][j] != two[i][j] {
				return false
			}
		}
	}
	return true
}
