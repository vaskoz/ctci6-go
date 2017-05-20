package four

import "testing"

var testcases = []struct {
	board [][]int
	won   bool
}{
	{
		board: [][]int{
			{X, O, X},
			{X, O, X},
			{X, E, E},
		},
		won: true,
	},
	{
		board: [][]int{
			{E, O, X},
			{E, O, X},
			{E, O, E},
		},
		won: true,
	},
	{
		board: [][]int{
			{E, E, X},
			{E, O, X},
			{E, O, E},
		},
		won: false,
	},
	{
		board: [][]int{
			{E, E, X},
			{X, X, X},
			{E, O, E},
		},
		won: true,
	},
	{
		board: [][]int{
			{E, E, X},
			{X, X, O},
			{X, O, E},
		},
		won: true,
	},
	{
		board: [][]int{
			{O, E, X},
			{X, O, O},
			{X, O, O},
		},
		won: true,
	},
	{
		board: [][]int{
			{O, E, E, E},
			{E, O, E, E},
			{E, E, O, E},
			{E, E, E, O},
		},
		won: true,
	},
	{
		board: [][]int{
			{X, E, X, E},
			{E, O, E, E},
			{E, E, O, E},
			{X, X, E, O},
		},
		won: false,
	},
	{
		board: [][]int{
			{E, E, X, O},
			{E, O, E, O},
			{E, E, O, O},
			{E, X, E, O},
		},
		won: true,
	},
	{
		board: [][]int{
			{E, E, E, E},
			{E, E, E, E},
			{E, E, O, E},
			{O, O, O, O},
		},
		won: true,
	},
}

func TestHasWon(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if won := HasWon(c.board); won != c.won {
			t.Errorf("Expected %t but got %t for board %v", c.won, won, c.board)
		}
	}
}

func BenchmarkHasWon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			HasWon(c.board)
		}
	}
}
