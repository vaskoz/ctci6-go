package seven

func MatrixRotation90Inplace(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
	return matrix
}
