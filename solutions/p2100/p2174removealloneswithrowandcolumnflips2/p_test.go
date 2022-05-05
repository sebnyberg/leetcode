package p2174removealloneswithrowandcolumnflips2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeOnes(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 2},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 0}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, removeOnes(tc.grid))
		})
	}
}

// func removeOnes(grid [][]int) int {
// 	minMoves := math.MaxInt64
// 	m, n := len(grid), len(grid[0])

// 	var findMin func(i, j, rows, cols int, moves int)
// 	findMin = func(i, j, rows, cols int, moves int) {
// 		// Return early if this DFS cannot improve the best result so far
// 		if moves >= minMoves {
// 			return
// 		}

// 		if j == n {
// 			j = 0
// 			i++
// 		}
// 		if i == m {
// 			minMoves = min(minMoves, moves)
// 			return
// 		}

// 		// If this cell is already covered, continue
// 		if grid[i][j] != 1 || rows&(1<<i) > 0 || cols&(1<<j) > 0 {
// 			findMin(i, j+1, rows, cols, moves)
// 			return
// 		}

// 		// Fix the current row and try all columns
// 		for jj := 0; jj < n; jj++ {
// 			if grid[i][jj] != 1 {
// 				continue
// 			}
// 			findMin(i, j+1, rows|(1<<i), cols|(1<<jj), moves+1)
// 		}

// 		// Fix the current column and try all rows
// 		for ii := 0; ii < m; ii++ {
// 			if grid[ii][j] != 1 {
// 				continue
// 			}
// 			findMin(i, j+1, rows|(1<<ii), cols|(1<<j), moves+1)
// 		}
// 	}

// 	findMin(0, 0, 0, 0, 0)
// 	return minMoves
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func removeOnes(grid [][]int) int {
	var findMin func(rows, cols int, moves int) int
	findMin = func(rows, cols int, moves int) int {
		min := math.MaxInt32
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == 0 || rows&(1<<i) > 0 || cols&(1<<j) > 0 {
					continue
				}
				res := findMin(rows|(1<<i), cols|(1<<j), moves+1)
				if res < min {
					min = res
				}
			}
		}
		if min == math.MaxInt32 {
			return moves
		}
		return min
	}

	return findMin(0, 0, 0)
}
