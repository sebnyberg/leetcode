package p0764largestplus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_orderOfLargestPlusSign(t *testing.T) {
	for _, tc := range []struct {
		N     int
		mines [][]int
		want  int
	}{
		{5, [][]int{{4, 2}}, 2},
		{2, [][]int{}, 1},
		{1, [][]int{{0, 0}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.N), func(t *testing.T) {
			require.Equal(t, tc.want, orderOfLargestPlusSign(tc.N, tc.mines))
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

func orderOfLargestPlusSign(N int, mines [][]int) int {
	// Since it is faster to initialize a board of zeroes rather than ones,
	// this solution marks mines as true on the board.
	board := make([][]bool, N)
	// For each position on the board, store the number of values found
	// on either side of that position.
	right := make([][]int, N)
	left := make([][]int, N)
	below := make([][]int, N)
	for i := range board {
		board[i] = make([]bool, N)
		right[i] = make([]int, N)
		left[i] = make([]int, N)
		below[i] = make([]int, N)
	}

	// Place mines
	for _, mine := range mines {
		board[mine[0]][mine[1]] = true
	}

	// Fill right matrix by going from right to left
	for col := N - 2; col >= 0; col-- {
		for row := 0; row < N; row++ {
			if board[row][col+1] {
				right[row][col] = 0
				continue
			}
			right[row][col] = right[row][col+1] + 1
		}
	}

	// Fill left matrix by going from left to right
	for col := 1; col < N; col++ {
		for row := 0; row < N; row++ {
			if board[row][col-1] {
				left[row][col] = 0
				continue
			}
			left[row][col] = left[row][col-1] + 1
		}
	}

	// Fill below matrix by going from bottom to the top
	for row := N - 2; row >= 0; row-- {
		for col := 0; col < N; col++ {
			if board[row+1][col] {
				below[row][col] = 0
				continue
			}
			below[row][col] = below[row+1][col] + 1
		}
	}

	// Finally, keep track of the values above
	// If the board position is zero (sorry about the inversion)
	// the order of the local plus-sign is the minimum of above/right/bottom/left
	// + 1
	above := make([]int, N)
	var maxOrder int
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] { // Remember! Inverse. True marks a mine
				above[col] = 0
				continue
			}
			maxOrder = max(maxOrder, 1+min(min(min(above[col], right[row][col]), below[row][col]), left[row][col]))
			above[col]++
		}
	}

	return maxOrder
}
