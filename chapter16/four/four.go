package four

// Constants to represent an empty cell and also 'X' and 'O'.
const (
	E = iota // E is Empty
	X
	O
)

// FullRow answers true if an entire row is controlled by either 'X' or 'O'.
func FullRow(rows int, m map[int]int) bool {
	return len(m) == 1 && (m[X] == rows || m[O] == rows)
}

// HasWon checks if any player has won this game of tic-tac-toe.
func HasWon(board [][]int) bool {
	rows := len(board)
	cols := len(board[0])
	for col := 0; col < cols; col++ {
		m := map[int]int{}
		for row := 0; row < rows; row++ {
			m[board[row][col]]++
		}
		if FullRow(rows, m) {
			return true
		}
	}
	for row := 0; row < rows; row++ {
		m := map[int]int{}
		for col := 0; col < cols; col++ {
			m[board[row][col]]++
		}
		if FullRow(rows, m) {
			return true
		}
	}
	leftDiag := map[int]int{}
	rightDiag := map[int]int{}
	for i := 0; i < rows; i++ {
		leftDiag[board[i][i]]++
		rightDiag[board[i][cols-i-1]]++
	}
	return FullRow(rows, leftDiag) || FullRow(rows, rightDiag)
}
