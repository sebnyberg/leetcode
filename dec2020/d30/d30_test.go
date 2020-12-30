package d30_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gameOfLife(t *testing.T) {
	for _, tc := range []struct {
		in   [][]int
		want [][]int
	}{
		{[][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		}, [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1},
			{0, 1, 0},
		}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			// incpy := make([][]int, len(tc.in))
			// for i, row := range tc.in {
			// 	incpy[i] = make([]int, len(row))
			// 	copy(incpy[i], row)
			// }
			gameOfLife(tc.in)
			require.Equal(t, tc.want, tc.in)
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gameOfLife(board [][]int) {
	// m == board.length
	// n == board[i].length
	// 1 <= m, n <= 25
	// board[i][j] is 0 or 1

	// Any live (1) cell with fewer than two live neighbours dies as if caused by under-population.
	// Any live (1) cell with two or three live neighbours lives on to the next generation.
	// Any live (1) cell with more than two live neighbours dies, as if by over-population.
	// Any dead (0) cell with exactly three live neighbours becomes a live cell, as if by reproduction
	m := len(board)
	n := len(board[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			aliveNeighbours := 0
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				for l := max(j-1, 0); l < min(j+2, n); l++ {
					if k == i && j == l {
						continue
					}
					if board[k][l] == 1 {
						aliveNeighbours++
					}
				}
			}
			if board[i][j] == 0 {
				if aliveNeighbours == 3 {
					res[i][j] = 1
				}
				continue
			}

			switch aliveNeighbours {
			case 0, 1: // Underpopulation
				res[i][j] = 0
			case 2, 3:
				res[i][j] = 1
			default: // Overpopulation
				res[i][j] = 0
			}
		}
	}

	for i := range res {
		copy(board[i], res[i])
	}

	// Follow-up:
	// Can you solve it in-place?
	// How to solve an infinite board? (store coordinates in a map)
}
