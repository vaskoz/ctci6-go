package seven

import "testing"

var testcases = []struct {
	in, out [][]int
}{
	{
		[][]int{
			{1, 2, 3, 4},
			{2, 3, 4, 5},
			{3, 4, 5, 6},
			{4, 5, 6, 7},
		},
		[][]int{
			{4, 3, 2, 1},
			{5, 4, 3, 2},
			{6, 5, 4, 3},
			{7, 6, 5, 4},
		},
	},
	{
		[][]int{
			{1, 2, 3, 4, 5},
			{2, 3, 4, 5, 6},
			{3, 4, 5, 6, 7},
			{4, 5, 6, 7, 8},
			{5, 6, 7, 8, 9},
		},
		[][]int{
			{5, 4, 3, 2, 1},
			{6, 5, 4, 3, 2},
			{7, 6, 5, 4, 3},
			{8, 7, 6, 5, 4},
			{9, 8, 7, 6, 5},
		},
	},
}

func TestMatrixRotation90Inplace(t *testing.T) {
	for _, c := range testcases {
		if out := MatrixRotation90Inplace(c.in); !MatrixEqual(out, c.out) {
			t.Errorf("MatrixRotation: given %v returned %v expected %v\n", c.in, out, c.out)
		}
	}
}

func BenchmarkMatrixRotation90Inplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			MatrixRotation90Inplace(c.in)
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
