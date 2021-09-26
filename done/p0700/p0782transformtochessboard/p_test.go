package p0782transformtochessboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_movesToChessboard(t *testing.T) {
	for _, tc := range []struct {
		board [][]int
		want  int
	}{
		{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}, 2},
		{[][]int{{0, 1}, {1, 0}}, 0},
		{[][]int{{0, 1}, {0, 1}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, movesToChessboard(tc.board))
		})
	}
}

func movesToChessboard(board [][]int) int {
	// A valid chessboard exists in two exact configurations, and one is the
	// inverse of the other. When N is odd, there is only one possible solution
	// (the total number of ones or zeroes is also uneven). When N is even,
	// there are two possible solutions, so finding one solution is equal to
	// to finding the other.

	// For a given row, it is not possible to change the number of zeroes and
	// ones, only the order of its contents. The same constraint is valid for
	// columns.

	// So, if the first row cannot be re-configured to valid configuration using
	// col swaps, then there is no solution. Similarly, if the first column cannot
	// be re-configured to a valid configuration using row swaps then there is no
	// solution either.

	// To check whether all squares formed by a re-configuration are valid, we can
	// xor the corners of each possible square starting in the top-left corner.
	// Courtesy of lee215 magic.
	n := len(board)
	for i := range board {
		for j := range board[i] {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] >= 1 {
				return -1
			}
		}
	}

	// Reconfigure the board to (possibly) have a valid first row / column
	var rowSum, colSum, rowSwap, colSwap int
	for i := range board {
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}
	}
	// Row sum and col sum must be half/half (even case), or
	// half+1/half (odd case). When n is even, n/2 and (n+1)/2 evaluates to the
	// same number (half).
	if rowSum != n/2 && rowSum != (n+1)/2 ||
		colSum != n/2 && colSum != (n+1)/2 {
		return -1
	}
	if n%2 == 1 {
		// Switch to optimal swaps for odd
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
	} else {
		// Get optimal swaps for even
		colSwap = min(colSwap, n-colSwap)
		rowSwap = min(rowSwap, n-rowSwap)
	}
	return (colSwap + rowSwap) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
