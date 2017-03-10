package eight

func ZeroMatrix(matrix [][]int) [][]int {
	is, js := []int{}, []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				is = append(is, i)
				js = append(js, j)
			}
		}
	}
	for i := 0; i < len(is); i++ {
		zi, zj := is[i], js[i]
		for x := 0; x < len(matrix); x++ {
			matrix[x][zj] = 0
		}
		for x := 0; x < len(matrix[zi]); x++ {
			matrix[zi][x] = 0
		}
	}
	return matrix
}
